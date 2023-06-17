package quantity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuantity(t *testing.T) {

	t.Run("ConvertTo", func(t *testing.T) {
		t.Run("1 GB should be equal to 1024 MB", func(t *testing.T) {
			oneGb := New(1, GigaBytes)
			val := oneGb.ConvertTo(MegaBytes)
			assert.Equal(t, float64(1024), val.value)
			assert.Equal(t, "MB", val.unit.shortName)
			assert.Equal(t, "Megabytes", val.unit.fullName)
		})

		t.Run("1 GB should be equal to 10,48,576 KB", func(t *testing.T) {
			oneGb := New(1, GigaBytes)
			val := oneGb.ConvertTo(KiloBytes)
			assert.Equal(t, float64(1048576), val.value)
			assert.Equal(t, "KB", val.unit.shortName)
			assert.Equal(t, "Kilobytes", val.unit.fullName)
		})

		t.Run("900 MB should be equal to 0.88 GB if precision is 2", func(t *testing.T) {
			oneMB := New(900, MegaBytes)
			val := oneMB.ConvertTo(GigaBytes, WithPrecision(2))
			assert.Equal(t, 0.88, val.value)
			assert.Equal(t, "GB", val.unit.shortName)
			assert.Equal(t, "Gigabytes", val.unit.fullName)
		})

		t.Run("900 MB should be equal to 0.87891 GB if precision is 5", func(t *testing.T) {
			oneMB := New(900, MegaBytes)
			val := oneMB.ConvertTo(GigaBytes, WithPrecision(5))
			assert.Equal(t, 0.87891, val.value)
			assert.Equal(t, "GB", val.unit.shortName)
			assert.Equal(t, "Gigabytes", val.unit.fullName)
		})
	})
}
