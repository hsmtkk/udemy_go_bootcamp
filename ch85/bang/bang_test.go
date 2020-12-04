package bang_test

import (
	"testing"

	"github.com/hsmtkk/udemy_go_bootcamp/ch85/bang"
	"github.com/stretchr/testify/assert"
)

func TestBang(t *testing.T) {
	banger := bang.New()

	want := "HEY!!!"
	got := banger.Bang("hey")
	assert.Equal(t, want, got)

	want = "HELLO!!!!!"
	got = banger.Bang("hello")
	assert.Equal(t, want, got)
}
