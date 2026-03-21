package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Anggiat")
	}
}
func BenchmarkHelloWorldCoki(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Coki")
	}
}
func TestHelloWorldAnggiat(t *testing.T) {
	result := HelloWorld("Anggiat")
	if result != "Hello Joko" {
		t.Error("Result must be Hello Anggiat")
	}
	// fmt.Println("TestHelloWorldAnggiat is done ")
}

func TestHelloWorldFretty(t *testing.T) {
	result := HelloWorld("Fretty")
	if result != "Hello Coki" {
		t.Fatal("Error must be Hello Fretty")
	}
	// fmt.Println("TestHelloWorldFretty is done your")
}
func TestHelloWorldAnggiatAssertion(t *testing.T) {
	res := HelloWorld("Anggiat")
	assert.Equal(t, "Hello Coki", res, "Result must be Coki")
	fmt.Println("Test Case telah dieksekusi")
}
func TestHelloWorldAnggiatRequire(t *testing.T) {
	res := HelloWorld("Anggiat")
	require.Equal(t, "Hello Anggi", res, "res must be Anggiat")
	fmt.Println("Telah dieksekusi")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit Test Tidak dijalankan di OS MAC")
	}
	res := HelloWorld("Anggiat")
	require.Equal(t, "Hello Coki", res, "Result must be coki")
}

func TestMain(m *testing.M) {
	fmt.Println("Sebelum unit test")
	m.Run()
	fmt.Println("Setelah unit test")
}
func TestSubTest(t *testing.T) {
	t.Run("Anggiat", func(t *testing.T) {
		res := HelloWorld("Anggiat")
		require.Equal(t, "Hello Anggiat", res, "Result must be Hello Anggiat")
	})
	t.Run("Pretty", func(t *testing.T) {
		res := HelloWorld("Pretty")
		fmt.Println("Ini sub test Pretty")
		require.Equal(t, "Hello Pretty", res, "Result must be Hello Pretty")
	})
}
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Anggiat)",
			request:  "Anggiat",
			expected: "Hello Anggiat",
		},
		{
			name:     "HelloWorld(Pretty)",
			request:  "Pretty",
			expected: "Hello Pretty",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
