package gover_test

import (
	"strings"
	"testing"

	"github.com/hsmtkk/udemy_go_bootcamp/ch36/gover"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	v := gover.Version()
	assert.NotEmpty(t, v)
	assert.True(t, strings.HasPrefix(v, "go1"))
}
