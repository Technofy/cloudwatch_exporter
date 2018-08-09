package config

import (
	"io/ioutil"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

var DefaultInteval = 120

type Instance struct {
	Instance     string `yaml:"instance"`
	Region       string `yaml:"region"`
	Interval     int    `yaml:"interval"`
	AwsAccessKey string `yaml:"aws_access_key"`
	AwsSecretKey string `yaml:"aws_secret_key"`
}

func (i *Instance) setInverval(interval int) {
	i.Interval = interval
}

type Config struct {
	Instances []*Instance `yaml:"instances"`
}

type Settings struct {
	config Config
	sync.RWMutex
	// AfterLoad is run after every Load request but before releasing Mutex
	AfterLoad func(config Config) error
}

func (s *Settings) Load(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()
	if err := yaml.Unmarshal(content, &s.config); err != nil {
		return err
	}
	for _, c := range s.config.Instances {
		if c.Interval == 0 {
			c.setInverval(DefaultInteval)
		}
	}

	if s.AfterLoad != nil {
		return s.AfterLoad(s.config)
	}

	return nil
}

func (s *Settings) Config() Config {
	s.RLock()
	defer s.RUnlock()
	return s.config
}
