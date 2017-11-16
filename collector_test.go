package main

import (
	"testing"

	"github.com/Percona-Lab/rds_exporter/config"
)

func TestNew(t *testing.T) {
	settings := &config.Settings{}
	c := New(settings)

	if c == nil {
		t.Fatal("collector should not be nil")
	}
}
