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

func TestLowSpeed(t *testing.T) {
	for i := 0; i < 4; i++ {
		t.Run(fmt.Sprintf("low speed test %d", i), func(t *testing.T) {
			testroot := path.Join(".", testdirectory)
			inData, err := os.ReadFile(path.Join(testroot, fmt.Sprintf("test.%d.in", i)))
			require.NoError(t, err)
			digitsN, err := strconv.Atoi(strings.TrimSpace(string(inData)))
			require.NoError(t, err)

			outData, err := os.ReadFile(path.Join(testroot, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			wantResult, err := strconv.Atoi(strings.TrimSpace(string(outData)))
			require.NoError(t, err)

			got := LuckyNumbersLowSpeed(digitsN)
			require.Equal(t, wantResult, got)
		})
	}
}

func TestHighSpeed(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("high speed test %d", i), func(t *testing.T) {
			testroot := path.Join(".", testdirectory)
			inData, err := os.ReadFile(path.Join(testroot, fmt.Sprintf("test.%d.in", i)))
			require.NoError(t, err)
			digitsN, err := strconv.Atoi(strings.TrimSpace(string(inData)))
			require.NoError(t, err)

			outData, err := os.ReadFile(path.Join(testroot, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			wantResult, err := strconv.Atoi(strings.TrimSpace(string(outData)))
			require.NoError(t, err)

			got := LuckyNumbersHighSpeed(digitsN)
			require.Equal(t, wantResult, got)
		})
	}
}
