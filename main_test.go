package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFailing(t *testing.T) {
	assert.Equal(t, 1, 2, "we are really testing this test, so should fail")
}
