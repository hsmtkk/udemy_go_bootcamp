package feet

type Converter interface {
	FeetMeter(feet float64) float64
	FeetYard(feet float64) float64
}

type converterImpl struct{}

func New() Converter {
	return &converterImpl{}
}

func (c *converterImpl) FeetMeter(feet float64) float64 {
	return feet * 0.3048
}

func (c *converterImpl) FeetYard(feet float64) float64 {
	return feet * 0.3333
}
