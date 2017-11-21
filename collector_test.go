package main

import (
	"testing"

	"github.com/percona/rds_exporter/config"
	"github.com/percona/rds_exporter/sessions"
)

func TestNew(t *testing.T) {
	settings := &config.Settings{}
	sessions := &sessions.Sessions{}
	c := New(settings, sessions)

	if c == nil {
		t.Fatal("collector should not be nil")
	}
}
