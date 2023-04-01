package asciiartfs

import (

	"os"
	"strings"
)


func WithFlag(ouput string,color string, filename string, argsRune []rune, lettersToColor []string){


		strTab := strings.Split(string(argsRune), `\n`)
		asciiTable := GetAsciiValues(filename)
	
		for _, v := range strTab {
			word := WordFormation2([]rune(v), asciiTable,lettersToColor,color)
			PrintWord(word)
			WriteWordToFile(word,&ouput)
		}
		os.Exit(0)
}