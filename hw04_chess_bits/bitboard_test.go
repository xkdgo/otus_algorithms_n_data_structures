package bitboard

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var testdirectory = "testdata"

func TestPopcntCicle(t *testing.T) {
	testBitBoardDirectory := path.Join(testdirectory, "1_Bitboard_King")
	for i := 0; i <= 9; i++ {
		t.Run(fmt.Sprintf("test popcnt with cicle %d", i), func(t *testing.T) {
			outData, err := os.ReadFile(path.Join(testBitBoardDirectory, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			output := strings.TrimSpace(string(outData))
			parts := strings.SplitN(output, "\r\n", 2)
			wantCount, err := strconv.Atoi(parts[0])
			log.Println("wantCount", wantCount)
			require.NoError(t, err)
			wantBitboard, err := strconv.ParseUint(parts[1], 0, 64)
			require.NoError(t, err)
			log.Printf("wantBitboard %064b", wantBitboard)
			gotCount := PopcntCicle(wantBitboard)
			require.Equal(t, wantCount, gotCount)
		})
	}
}

func TestPopcntSparseFast(t *testing.T) {
	testBitBoardDirectory := path.Join(testdirectory, "1_Bitboard_King")
	for i := 0; i <= 9; i++ {
		t.Run(fmt.Sprintf("test popcnt with fast sparce cnt %d", i), func(t *testing.T) {
			outData, err := os.ReadFile(path.Join(testBitBoardDirectory, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			output := strings.TrimSpace(string(outData))
			parts := strings.SplitN(output, "\r\n", 2)
			wantCount, err := strconv.Atoi(parts[0])
			log.Println("wantCount", wantCount)
			require.NoError(t, err)
			wantBitboard, err := strconv.ParseUint(parts[1], 0, 64)
			require.NoError(t, err)
			log.Printf("wantBitboard %064b", wantBitboard)
			gotCount := PopcntSparseFast(wantBitboard)
			require.Equal(t, wantCount, gotCount)
		})
	}
}

func TestPopcnt(t *testing.T) {
	testBitBoardDirectory := path.Join(testdirectory, "1_Bitboard_King")
	for i := 0; i <= 9; i++ {
		t.Run(fmt.Sprintf("test popcnt with fast sparce cnt %d", i), func(t *testing.T) {
			outData, err := os.ReadFile(path.Join(testBitBoardDirectory, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			output := strings.TrimSpace(string(outData))
			parts := strings.SplitN(output, "\r\n", 2)
			wantCount, err := strconv.Atoi(parts[0])
			log.Println("wantCount", wantCount)
			require.NoError(t, err)
			wantBitboard, err := strconv.ParseUint(parts[1], 0, 64)
			require.NoError(t, err)
			log.Printf("wantBitboard %064b", wantBitboard)
			gotCount := Popcnt(wantBitboard)
			require.Equal(t, wantCount, gotCount)
		})
	}
}

func TestKing(t *testing.T) {
	testBitBoardDirectory := path.Join(testdirectory, "1_Bitboard_King")
	for i := 0; i <= 9; i++ {
		t.Run(fmt.Sprintf("test king %d", i), func(t *testing.T) {
			inData, err := os.ReadFile(path.Join(testBitBoardDirectory, fmt.Sprintf("test.%d.in", i)))
			require.NoError(t, err)
			input := strings.TrimSpace(string(inData))
			inDigit, err := strconv.ParseUint(input, 0, 64)
			require.NoError(t, err)
			log.Println("indata", inDigit)

			outData, err := os.ReadFile(path.Join(testBitBoardDirectory, fmt.Sprintf("test.%d.out", i)))
			require.NoError(t, err)
			output := strings.TrimSpace(string(outData))
			parts := strings.SplitN(output, "\r\n", 2)
			wantCount, err := strconv.Atoi(parts[0])
			log.Println("wantCount", wantCount)
			require.NoError(t, err)
			wantBitboard, err := strconv.ParseUint(parts[1], 0, 64)
			require.NoError(t, err)
			log.Printf("wantBitboard %064b", wantBitboard)
		})
	}
}
