package quantity

import (
	"math"
)

type Quantity struct {
	value float64
	unit  unit

	//optional
	precision int
}

func New(value float64, u unit) Quantity {
	return Quantity{
		value: value,
		unit:  u,
	}
}

type option func(f *Quantity)

func (q Quantity) Value(opts ...option) float64 {
	for _, opt := range opts {
		opt(&q)
	}
	return q.value
}

func (q Quantity) UnitName() string {
	return q.unit.fullName
}

func (q Quantity) ShortName() string {
	return q.unit.shortName
}

func WithPrecision(precision int) option {
	return func(q *Quantity) {
		q.precision = precision
		q.value = toFixed(q.value, precision)
	}
}

func (q Quantity) ConvertTo(u unit, opts ...option) *Quantity {
	value := q.value * (q.unit.conversionFactor / u.conversionFactor)
	quant := &Quantity{value: value, unit: u}
	for _, opt := range opts {
		opt(quant)
	}
	return quant
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
