package asciiartfs

import (
	"bufio"
	"os"
	"strings"
)

func GetAsciiValues(filename string) [][]string {
	// read the file and split it by line
	filename = selectFile(filename)
	asciiData, _ := os.ReadFile(filename)
	asciiDataStr := string(asciiData)
	scanner := bufio.NewScanner(strings.NewReader(asciiDataStr))
	scannerTable := []string{}
	for scanner.Scan() {
		scannerTable = append(scannerTable, scanner.Text())
	}

	asciiTable := strings.Split(strings.Join(scannerTable, "\n"), "\n")
	// parse the data and get each 8 lines as a letter
	bigTable := [][]string{}
	for i := 0; i < len(asciiTable)-9; i += 9 {
		bigTable = append(bigTable, asciiTable[i+1:i+9])
	}
	return bigTable
}
