package calculate

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestStandard(t *testing.T) {
	testPowerDirectory := path.Join(testdirectory, "power")
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("power via multiply test %d", i), func(t *testing.T) {
			inData, err := os.ReadFile(path.Join(testPowerDirectory, fmt.Sprintf("test.%d.in", i)))
			require.NoError(t, err)
			input := strings.TrimSpace(string(inData))
			parts := strings.Fields(input)
			require.Equal(t, 2, len(parts))
			x, err := strconv.ParseFloat(parts[0], 64)
			require.NoError(t, err)
			y, err := strconv.ParseFloat(parts[1], 64)
			require.NoError(t, err)

			outData, err := os.ReadFile(path.Join(testPowerDirectory, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			output := strings.TrimSpace(string(outData))
			wantResult, err := strconv.ParseFloat(output, 64)
			require.NoError(t, err)
			start := time.Now()
			got := PowerStandard(x, int(y))
			log.Printf("%s %.10f^%d time=%d msecs", "PowerStandard", x, int(y), time.Since(start).Milliseconds())
			require.Equal(t, fmt.Sprintf("%.6f", wantResult), fmt.Sprintf("%.6f", got))
		})
	}
}

func TestPowerViaMultiply(t *testing.T) {
	testPowerDirectory := path.Join(testdirectory, "power")
	for i := 0; i < 11; i++ {
		t.Run(fmt.Sprintf("power via multiply test %d", i), func(t *testing.T) {
			inData, err := os.ReadFile(path.Join(testPowerDirectory, fmt.Sprintf("test.%d.in", i)))
			require.NoError(t, err)
			input := strings.TrimSpace(string(inData))
			parts := strings.Fields(input)
			require.Equal(t, 2, len(parts))
			x, err := strconv.ParseFloat(parts[0], 64)
			require.NoError(t, err)
			y, err := strconv.ParseFloat(parts[1], 64)
			require.NoError(t, err)

			outData, err := os.ReadFile(path.Join(testPowerDirectory, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			output := strings.TrimSpace(string(outData))
			wantResult, err := strconv.ParseFloat(output, 64)
			require.NoError(t, err)
			start := time.Now()
			got := PowerViaMultiply(x, int(y))
			log.Printf("%s %.10f^%d time=%d msecs", "PowerViaMultiply", x, int(y), time.Since(start).Milliseconds())
			require.Equal(t, fmt.Sprintf("%.11f", wantResult), fmt.Sprintf("%.11f", got))
		})
	}
}

func TestPowerViaPowOfTwo(t *testing.T) {
	testPowerDirectory := path.Join(testdirectory, "power")
	for i := 0; i < 11; i++ {
		t.Run(fmt.Sprintf("power via multiply test %d", i), func(t *testing.T) {
			inData, err := os.ReadFile(path.Join(testPowerDirectory, fmt.Sprintf("test.%d.in", i)))
			require.NoError(t, err)
			input := strings.TrimSpace(string(inData))
			parts := strings.Fields(input)
			require.Equal(t, 2, len(parts))
			x, err := strconv.ParseFloat(parts[0], 64)
			require.NoError(t, err)
			y, err := strconv.ParseFloat(parts[1], 64)
			require.NoError(t, err)

			outData, err := os.ReadFile(path.Join(testPowerDirectory, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			output := strings.TrimSpace(string(outData))
			wantResult, err := strconv.ParseFloat(output, 64)
			require.NoError(t, err)
			start := time.Now()
			got := PowerViaPowOfTwo(x, int(y))
			log.Printf("%s %.10f^%d time=%d msecs", "PowerViaPowOfTwo", x, int(y), time.Since(start).Milliseconds())
			require.Equal(t, fmt.Sprintf("%.6f", wantResult), fmt.Sprintf("%.6f", got))
		})
	}
}
