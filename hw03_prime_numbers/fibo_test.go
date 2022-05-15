package calculate

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecursionFibo(t *testing.T) {
	testFiboDirectory := path.Join(testdirectory, "fibo")
	for i := 0; i < 7; i++ {
		t.Run(fmt.Sprintf("recursion test %d", i), func(t *testing.T) {
			inData, err := os.ReadFile(path.Join(testFiboDirectory, fmt.Sprintf("test.%d.in", i)))
			require.NoError(t, err)
			digitsN, err := strconv.ParseUint(strings.TrimSpace(string(inData)), 0, 64)
			require.NoError(t, err)

			outData, err := os.ReadFile(path.Join(testFiboDirectory, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			wantResult := strings.TrimSpace(string(outData))
			require.NoError(t, err)
			got := FiboRecursion(digitsN)
			require.Equal(t, wantResult, got.String())
		})
	}
}

func TestCicleFibo(t *testing.T) {
	testFiboDirectory := path.Join(testdirectory, "fibo")
	for i := 0; i < 12; i++ {
		t.Run(fmt.Sprintf("cicle test %d", i), func(t *testing.T) {
			inData, err := os.ReadFile(path.Join(testFiboDirectory, fmt.Sprintf("test.%d.in", i)))
			require.NoError(t, err)
			digitsN, err := strconv.ParseUint(strings.TrimSpace(string(inData)), 0, 64)
			require.NoError(t, err)

			outData, err := os.ReadFile(path.Join(testFiboDirectory, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			wantResult := strings.TrimSpace(string(outData))
			require.NoError(t, err)
			got := FiboCicle(digitsN)
			require.Equal(t, wantResult, got.String())
		})
	}
}

func TestMatrixFibo(t *testing.T) {
	testFiboDirectory := path.Join(testdirectory, "fibo")
	for i := 0; i < 13; i++ {
		t.Run(fmt.Sprintf("matrix fibo test %d", i), func(t *testing.T) {
			inData, err := os.ReadFile(path.Join(testFiboDirectory, fmt.Sprintf("test.%d.in", i)))
			require.NoError(t, err)
			digitsN, err := strconv.ParseUint(strings.TrimSpace(string(inData)), 0, 64)
			require.NoError(t, err)
			outData, err := os.ReadFile(path.Join(testFiboDirectory, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			wantResult := strings.TrimSpace(string(outData))
			require.NoError(t, err)
			got := FiboMatrix(digitsN)
			require.Equal(t, wantResult, got.String())
		})
	}
}
