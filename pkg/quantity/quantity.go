package quantity

import (
	"math"
)

type quantity struct {
	value float64
	unit  unit

	//optional
	precision int
}

func New(value float64, u unit) quantity {
	return quantity{
		value: value,
		unit:  u,
	}
}

type option func(f *quantity)

func (q quantity) GetValue(opts ...option) float64 {
	for _, opt := range opts {
		opt(&q)
	}
	return q.value
}

func (q quantity) GetUnitName() string {
	return q.unit.fullName
}

func (q quantity) GetUnitShortName() string {
	return q.unit.shortName
}

func WithPrecision(precision int) option {
	return func(q *quantity) {
		q.precision = precision
		q.value = toFixed(q.value, precision)
	}
}

func (q quantity) ConvertTo(u unit, opts ...option) *quantity {
	value := q.value * (q.unit.conversionFactor / u.conversionFactor)
	quant := &quantity{value: value, unit: u}
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
