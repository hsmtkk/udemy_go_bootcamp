package valid_test

import (
	"github.com/hsmtkk/udemy_go_bootcamp/ch132/valid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	v := valid.New()
	got := v.Validate("alice", "bob")
	assert.True(t, got)
	got = v.Validate("alice", "charlie")
	assert.False(t, got)
	got = v.Validate("delta", "bob")
	assert.False(t, got)
}
