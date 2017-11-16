package config

import (
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

type Instance struct {
	Instance string `yaml:"instance"`
	Region   string `yaml:"region"`
}

type Config struct {
	Instances []Instance `yaml:"instances"`
}

type Settings struct {
	config Config
	sync.RWMutex
}

func (s *Settings) Load(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()
	return yaml.Unmarshal(content, &s.config)
}

func (s *Settings) Config() Config {
	s.RLock()
	defer s.RUnlock()
	return s.config
}
