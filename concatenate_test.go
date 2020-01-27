package ditto_yaml_test

import (
	ditto_yaml "github.com/payfazz/ditto-yaml"
	"testing"
)

func TestGet(t *testing.T) {
	_, err := ditto_yaml.Get("test")
	if err == nil {
		t.Fatal("no such directory error expected")
	}

	res, err := ditto_yaml.Get("v0.1")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}