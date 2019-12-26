package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(t *testing.T) {
	want := "dummy"
	assert.Equal(t, want, DummyFunc())
}
