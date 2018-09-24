package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type InstanceType string

const (
	Unknown     InstanceType = "unknown"
	AuroraMySQL InstanceType = "aurora_mysql"
	MySQL       InstanceType = "mysql"
)

type Instance struct {
	Region   string `yaml:"region"`
	Instance string `yaml:"instance"`
	// DBIResourceID string       `yaml:"dbi_resource_id"` // may be empty
	// Type          InstanceType `yaml:"type"`            // may be empty
	AWSAccessKey string `yaml:"aws_access_key"`
	AWSSecretKey string `yaml:"aws_secret_key"`
}

type Config struct {
	Instances []Instance `yaml:"instances"`
}

func Load(filename string) (*Config, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	if err = yaml.Unmarshal(b, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
