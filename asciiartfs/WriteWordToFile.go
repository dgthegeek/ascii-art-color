package asciiartfs

import (
"fmt"
"os"
)

func WriteWordToFile(words [][]string, outputfile *string){

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
	
		fileWriter(line1, outputfile)
		fileWriter(line2, outputfile)
		fileWriter(line3, outputfile)
		fileWriter(line4, outputfile)
		fileWriter(line5, outputfile)
		fileWriter(line6, outputfile)
		fileWriter(line7, outputfile)
		fileWriter(line8, outputfile)
}


func fileWriter(content string, outputfile *string) {

	file, err := os.OpenFile(*outputfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content + "\n")
	if err != nil {
		fmt.Println(err)
		return
	}
}
