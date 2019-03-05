package main

import (
	"testing"
)

func TestReadSettings(t *testing.T) {
	conf:= readSettings("")
	if conf.TargetDir == "" {
		t.Fatal("failed test")
	}
}
