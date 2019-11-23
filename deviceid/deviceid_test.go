package deviceid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateDeviceID(t *testing.T) {
	table := []struct {
		steamID string
		valid   string
	}{
		{"76561198261386394", "962d9818-8b83-81d6-afa4-e09118f1295b"},
		{"76561198261386395", "fbc879ac-572e-750e-001e-25da190475ba"},
		{"76561198261386396", "c054e045-b749-0699-f1dc-254211cb2e0d"},
		{"76561198261386397", "34fb601c-3513-37bb-9052-ba8e0a400329"},
		{"76561198261386398", "8e7087df-4a6a-643e-7a21-20d875e05ec4"},
		{"76561198261386399", "8771ffff-6714-569d-6dde-b7be85e8a23b"},
		{"76561198261386400", "c70cca8e-9359-086e-82d6-3911e4e734f3"},
		{"76561198261386401", "4ec451d0-6fc2-9972-c03c-fab5cce1db96"},
		{"76561198261386402", "0dc3bd6b-bf6c-0e6f-75df-0da9985899d3"},
		{"76561198261386403", "d2305bb9-a690-e199-9383-017f6709cc8e"},
	}

	for _, row := range table {
		assert.Equal(t, row.valid, GenerateDeviceID(row.steamID))
	}
}

func BenchmarkGenerateDeviceID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateDeviceID("76561198261386403")
	}
}
