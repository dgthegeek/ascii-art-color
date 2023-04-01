package asciiartfs

import (
"fmt"
)
// This function print the ouput of your treatment
func PrintWord(words [][]string) {
	if len(words) == 0 {
		fmt.Println()
		return
	}
	line1 := ""
	line2 := ""
	line3 := ""
	line4 := ""
	line5 := ""
	line6 := ""
	line7 := ""
	line8 := ""

	for i := 0; i < len(words); i++ {
		line1 += words[i][0]
		line2 += words[i][1]
		line3 += words[i][2]
		line4 += words[i][3]
		line5 += words[i][4]
		line6 += words[i][5]
		line7 += words[i][6]
		line8 += words[i][7]
	}

	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println(line3)
	fmt.Println(line4)
	fmt.Println(line5)
	fmt.Println(line6)
	fmt.Println(line7)
	fmt.Println(line8)
}
