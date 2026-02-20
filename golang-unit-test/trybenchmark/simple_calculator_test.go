package trybenchmark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	res := Add(1, 2)
	assert.Equal(t, 3, res, "Failed bro")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(99999999999999999, 99999999999)
	}
}
func BenchmarkAddSub(b *testing.B) {
	b.Run("3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Add(1, 2)
		}
	})
	b.Run("2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Add(1, 1)
		}
	})
}
func BenchmarkAddTable(b *testing.B) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "Add(Percbaan1)",
			a:        1,
			b:        2,
			expected: 3,
		},
		{
			name:     "Add(Percbaan2)",
			a:        1,
			b:        3,
			expected: 5,
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(t *testing.B) {
			result := Add(test.a, test.b)
			assert.Equal(t, test.expected, result)
		})
	}
}
