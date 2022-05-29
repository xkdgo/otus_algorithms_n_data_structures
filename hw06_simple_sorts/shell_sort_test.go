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

func TestShell(t *testing.T) {
	for _, sort := range []sortFunc{ShellSortClassical, ShellSort2Kminus1, ShellSort3Kminus1} {
		for _, dirname := range xtypeData {
			testNameDirectory := path.Join(testdirectory, dirname)
			for i := 0; i <= 7; i++ {
				t.Run(fmt.Sprintf("%s test on %s data %d", GetFunctionName(sort), GetDirectoryName(dirname), i),
					func(t *testing.T) {
						inData, err := os.ReadFile(path.Join(testNameDirectory, fmt.Sprintf("test.%d.in", i)))
						require.NoError(t, err)
						input := strings.TrimSpace(string(inData))
						parts := strings.SplitN(input, "\n", 2)
						require.Equal(t, 2, len(parts))
						arrLen, err := strconv.Atoi(strings.TrimSpace(parts[0]))

						require.NoError(t, err)
						arr := make([]int, 0, arrLen)
						arrStrings := strings.Fields(strings.TrimSpace(parts[1]))
						require.Equal(t, arrLen, len(arrStrings))
						for _, strItem := range arrStrings {
							intItem, err := strconv.Atoi(strings.TrimSpace(strItem))
							require.NoError(t, err)
							arr = append(arr, intItem)
						}

						outData, err := os.ReadFile(path.Join(testNameDirectory, fmt.Sprintf("test.%d.out", i)))
						require.NoError(t, err)
						output := strings.TrimSpace(string(outData))
						wantResult := make([]int, 0, arrLen)
						arrStrings = strings.Fields(strings.TrimSpace(output))
						for _, strItem := range arrStrings {
							intItem, err := strconv.Atoi(strings.TrimSpace(strItem))
							require.NoError(t, err)
							wantResult = append(wantResult, intItem)
						}
						start := time.Now()
						sort(arr)
						log.Printf("%s test on %s data %d: N=%d time=%d msecs",
							GetFunctionName(sort), GetDirectoryName(dirname), i, arrLen, time.Since(start).Milliseconds())
						require.Equal(t, wantResult, arr)
					})
			}
		}
	}
}