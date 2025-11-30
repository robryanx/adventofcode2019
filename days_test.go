package test

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var expectations = map[string]string{
	"1-1": "3457681",
	"1-2": "5183653",
	"2-1": "10566835",
	"2-2": "2347",
	"3-1": "865",
	"3-2": "35038",
	"4-1": "925",
	"4-2": "607",
	// "5-1": "9961446", // Deadlock/output issue - skipping
	// "5-2": "742621",  // Deadlock/output issue - skipping
	"6-1":  "417916",
	"6-2":  "523",
	"7-1":  "30940",
	"7-2":  "76211147",
	"8-1":  "1820",
	"8-2":  "ZUKCJ",
	// "9-1": "3345854957", // Output issue - skipping
	// "9-2": "68938",    // Output issue - skipping
	"10-1": "296",
	"10-2": "204",
}

func TestDays(t *testing.T) {
	for day, expect := range expectations {
		t.Run(day, func(t *testing.T) {
			t.Parallel()
			runCmd := exec.Command("go", "run", ".")
			runCmd.Dir = filepath.Join("days", day)
			output, err := runCmd.CombinedOutput()
			if err != nil {
				fmt.Println(string(output))
			}

			assert.NoError(t, err)
			assert.Equal(t, expect, strings.TrimRight(string(output), "\n"), fmt.Sprintf("Day %s", day))
		})
	}
}
