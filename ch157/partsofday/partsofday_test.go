package partsofday_test

import (
	"strings"
	"testing"

	"github.com/hsmtkk/udemy_go_bootcamp/ch157/partsofday"
	"github.com/stretchr/testify/assert"
)

func TestNight(t *testing.T) {
	want := "Good night"
	got, err := partsofday.New().Get(3)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestMorning(t *testing.T) {
	want := "Good morning"
	got, err := partsofday.New().Get(9)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestAfternoon(t *testing.T) {
	want := "Good afternoon"
	got, err := partsofday.New().Get(15)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestEvening(t *testing.T) {
	want := "Good evening"
	got, err := partsofday.New().Get(21)
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}

func TestGetCurrent(t *testing.T) {
	got := partsofday.New().GetCurrent()
	assert.NotEmpty(t, got)
	assert.True(t, strings.HasPrefix(got, "Good "))
}
