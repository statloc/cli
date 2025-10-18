package internal_test

import (
	"bufio"
	"cli/internal"
	"os"
	"strings"
	"testing"

	core "github.com/statloc/core"
	"github.com/stretchr/testify/assert"
)

func TestRespond(t *testing.T) {
    file, _ := os.Open("../testdata/results.txt")
	defer file.Close() // nolint:errcheck

	response := internal.Respond("../testdata")

	// go line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        println(response, scanner.Text())
		assert.True(t, strings.Contains(response, scanner.Text()))
	}
}

func TestGetTable(t *testing.T) {
    file, _ := os.Open("../testdata/languages.txt")
	defer file.Close() // nolint:errcheck

	statistics, _ := core.GetStatistics("../testdata")
	table := internal.GetTable(statistics.Languages, 5, 3, 5)

	// go line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        println(table, scanner.Text())
		assert.True(t, strings.Contains(table, scanner.Text()))
	}
}
