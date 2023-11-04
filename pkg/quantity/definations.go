package quantity

/*Memory*/

type byteSize float64

const (
	_             = iota //ignore first value by assigning to blank identifier
	byte byteSize = 1 << (10 * (iota - 1))
	kiloByte
	megaByte
	gigaByte
	TeraByte
	PetaByte
)

var Bytes = unit{fullName: "Bytes", shortName: "B", conversionFactor: float64(byte)}
var KiloBytes = unit{fullName: "Kilobytes", shortName: "KB", conversionFactor: float64(kiloByte)}
var MegaBytes = unit{fullName: "Megabytes", shortName: "MB", conversionFactor: float64(megaByte)}
var GigaBytes = unit{fullName: "Gigabytes", shortName: "GB", conversionFactor: float64(gigaByte)}
