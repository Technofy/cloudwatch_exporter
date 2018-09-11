package config

import (
	"io/ioutil"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	// DefaultInterval defines default interval in case it is missing in yaml.
	DefaultInterval = time.Minute
)

type Instance struct {
	Instance     string        `yaml:"instance"`
	Region       string        `yaml:"region"`
	Interval     time.Duration `yaml:"interval"`
	AwsAccessKey string        `yaml:"aws_access_key"`
	AwsSecretKey string        `yaml:"aws_secret_key"`
}

// Labels returns slice of labels.
func (i Instance) Labels() []string {
	return []string{
		"instance",
		"region",
	}
}

// LabelsValues returns slice of labels values.
func (i Instance) LabelsValues() []string {
	return []string{
		i.Instance,
		i.Region,
	}
}

type Config struct {
	Instances []Instance `yaml:"instances"`
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
	for i := range s.config.Instances {
		if s.config.Instances[i].Interval.Nanoseconds() == 0 {
			s.config.Instances[i].Interval = DefaultInterval
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
