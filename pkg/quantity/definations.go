package quantity

/*Memory*/

type byteSize int

const (
	_ = iota //ignore first value by assigning to blank identifier
	byte
	kiloByte byteSize = 1 << (10 * iota)
	megaByte
	gigaByte
	TeraByte
	PetaByte
)

var Bytes = unit{fullName: "Bytes", shortName: "B", conversionFactor: byte}
var KiloBytes = unit{fullName: "Kilobytes", shortName: "KB", conversionFactor: float64(kiloByte)}
var MegaBytes = unit{fullName: "Megabytes", shortName: "MB", conversionFactor: float64(megaByte)}
var GigaBytes = unit{fullName: "Gigabytes", shortName: "GB", conversionFactor: float64(gigaByte)}
