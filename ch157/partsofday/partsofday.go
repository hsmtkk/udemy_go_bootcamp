package partsofday

import (
	"fmt"
	"time"
)

type Getter interface {
	GetCurrent() string
	Get(hour int) (string, error)
}

type getterImpl struct{}

func New() Getter {
	return &getterImpl{}
}

const (
	Night     = "Good night"
	Morning   = "Good morning"
	Afternoon = "Good afternoon"
	Evening   = "Good evening"
)

func (g *getterImpl) Get(hour int) (string, error) {
	switch {
	case 0 <= hour && hour < 6:
		return Night, nil
	case 6 <= hour && hour < 12:
		return Morning, nil
	case 12 <= hour && hour < 18:
		return Afternoon, nil
	case 18 <= hour && hour < 24:
		return Evening, nil
	}
	return "", fmt.Errorf("invalid hour %d", hour)
}

func (g *getterImpl) GetCurrent() string {
	h := time.Now().Hour()
	s, _ := g.Get(h)
	return s
}
