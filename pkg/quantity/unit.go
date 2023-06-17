package quantity

type UnitType string

type unit struct {
	fullName         string
	shortName        string
	conversionFactor float64
}
