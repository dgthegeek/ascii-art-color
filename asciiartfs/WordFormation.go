package asciiartfs

import "strings"

// import (
// 	"fmt"
// 	"strings"
// )



func WordFormation(argsRune []rune, asciiTable [][]string ) [][]string{
	word := [][]string{}
	for _, v := range argsRune {
		if v >= 26 && v <= 126 {
			word = append(word, asciiTable[v-32])
		}
	}
	return word
}

func WordFormation2(argsRune []rune, asciiTable [][]string, lettersToColor []string,color string) [][]string {
	word := [][]string{}
	for _, v := range argsRune {
		if v >= 26 && v <= 126 {
			if IsColored(string(v),lettersToColor){
				colored:=Colorize(asciiTable[v-32],color)
				word = append(word, colored)
			}else{
				word = append(word, asciiTable[v-32])
			}	
		}
	}
	return word
}

func IsColored(v string, lettersToColor []string)bool{
	for _,val:= range lettersToColor{
		if string(val)==v{
			return true
		}
	}
	return false
}

func Colorize(char[]string,color string)[]string{
	reset:= "\033[0m"
	colorized:=[]string{}
	for _,v:= range char{
		coloredstr:= SelectColor(color) +string(v) + reset 
		colorized = append(colorized,coloredstr )
	}
	return colorized
}


func SelectColor(color string) string {
	color = strings.ToLower(color)
	switch color {
	case "white" :
		color = "\u001b[37m"
	case "black" :
		color = "\u001b[30m"
	case "red" :
		color = "\u001b[31m"
	case "green":
		color= "\u001b[32m"
	case "blue":
		color = "\u001b[34m"
	case "yellow" :
		color= "\u001b[33m"
	case "magenta":
		color= "\u001b[35m"
	case "cyan":
		color= "\u001b[36m"
	default : color = "\u001b[37m"
	}
return color
}