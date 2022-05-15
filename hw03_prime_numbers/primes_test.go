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

var testdirectory = "testdata"

func TestCountPrimesLinear(t *testing.T) {
	testPrimesDirectory := path.Join(testdirectory, "primes")
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("primes count linear %d", i), func(t *testing.T) {
			inData, err := os.ReadFile(path.Join(testPrimesDirectory, fmt.Sprintf("test.%d.in", i)))
			require.NoError(t, err)
			input := strings.TrimSpace(string(inData))
			inDigit, err := strconv.Atoi(input)
			require.NoError(t, err)

			outData, err := os.ReadFile(path.Join(testPrimesDirectory, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			output := strings.TrimSpace(string(outData))
			wantResult, err := strconv.Atoi(output)
			require.NoError(t, err)
			got := CountPrimesLinear(inDigit)
			require.Equal(t, wantResult, got)
		})
	}
}

func TestPrimesCountEratosphen(t *testing.T) {
	testPrimesDirectory := path.Join(testdirectory, "primes")
	for i := 0; i < 15; i++ {
		t.Run(fmt.Sprintf("primes count Eratosphen %d", i), func(t *testing.T) {
			inData, err := os.ReadFile(path.Join(testPrimesDirectory, fmt.Sprintf("test.%d.in", i)))
			require.NoError(t, err)
			input := strings.TrimSpace(string(inData))
			inDigit, err := strconv.Atoi(input)
			require.NoError(t, err)

			outData, err := os.ReadFile(path.Join(testPrimesDirectory, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			output := strings.TrimSpace(string(outData))
			wantResult, err := strconv.Atoi(output)
			require.NoError(t, err)
			got := EratosphenClassic(inDigit)
			require.Equal(t, wantResult, got)
		})
	}
}
