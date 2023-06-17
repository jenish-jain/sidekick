package quantity_test

import (
	. "github.com/jenish-jain/sidekick/pkg/quantity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuantity(t *testing.T) {

	t.Run("ConvertTo", func(t *testing.T) {
		t.Run("1 GB should be equal to 1024 MB", func(t *testing.T) {
			oneGb := New(1, GigaBytes)
			val := oneGb.ConvertTo(MegaBytes)
			assert.Equal(t, float64(1024), val.Value())
			assert.Equal(t, "MB", val.ShortName())
			assert.Equal(t, "Megabytes", val.UnitName())
		})

		t.Run("1 GB should be equal to 10,48,576 KB", func(t *testing.T) {
			oneGb := New(1, GigaBytes)
			val := oneGb.ConvertTo(KiloBytes)
			assert.Equal(t, float64(1048576), val.Value())
			assert.Equal(t, "KB", val.ShortName())
			assert.Equal(t, "Kilobytes", val.UnitName())
		})

		t.Run("900 MB should be equal to 0.88 GB if precision is 2", func(t *testing.T) {
			oneMB := New(900, MegaBytes)
			val := oneMB.ConvertTo(GigaBytes, WithPrecision(2))
			assert.Equal(t, 0.88, val.Value())
			assert.Equal(t, "GB", val.ShortName())
			assert.Equal(t, "Gigabytes", val.UnitName())
		})

		t.Run("900 MB should be equal to 0.87891 GB if precision is 5", func(t *testing.T) {
			oneMB := New(900, MegaBytes)
			val := oneMB.ConvertTo(GigaBytes, WithPrecision(5))
			assert.Equal(t, 0.87891, val.Value())
			assert.Equal(t, "GB", val.ShortName())
			assert.Equal(t, "Gigabytes", val.UnitName())
		})
	})

	t.Run("Value", func(t *testing.T) {
		t.Run("return value without precision for a quantity if not specified", func(t *testing.T) {
			q := New(6.123456789, GigaBytes)
			assert.Equal(t, 6.123456789, q.Value())
		})

		t.Run("return value with specified precision for a quantity", func(t *testing.T) {
			q := New(6.123456789, GigaBytes)
			assert.Equal(t, 6.12, q.Value(WithPrecision(2)))
		})
	})
}
