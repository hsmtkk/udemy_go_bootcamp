package bang

import "strings"

type Banger interface {
	Bang(orig string) string
}

type bangerImpl struct{}

func (b *bangerImpl) Bang(orig string) string {
	return strings.ToUpper(orig) + strings.Repeat("!", len(orig))
}

func New() Banger {
	return &bangerImpl{}
}
