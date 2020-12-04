package feet_test

import (
	"testing"

	"github.com/hsmtkk/udemy_go_bootcamp/ch112/feet"
	"github.com/stretchr/testify/assert"
)

func TestFeetMeter(t *testing.T) {
	want := 3048
	got := int(feet.New().FeetMeter(10000))
	assert.Equal(t, want, got)
}

func TestFeetYard(t *testing.T) {
	want := 3333
	got := int(feet.New().FeetYard(10000))
	assert.Equal(t, want, got)
}
