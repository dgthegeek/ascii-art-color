package main

import (
	"asciiart/asciiartfs"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Watch your inputs...")
		}
	}()
	var ouput string
	flag.StringVar(&ouput, "output","output","string")
	var color string
	flag.StringVar(&color,"color","white","string")
	flag.Parse()
	args:= flag.Args()
		
	
		if len(args) == 1 {
			asciiartfs.DefaultWork("standard.txt",[]rune(args[0]))
			os.Exit(0)
		}
		if len(args) < 1  {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			os.Exit(0)
		}
		if args[0] == "\\n" {
			fmt.Println()
			os.Exit(0)
		}
		if args[0] == "" {
			os.Exit(0)
		}
		lettersToColor:= strings.Split(args[0], "")
		filename := args[len(args)-1] 
		argsRune := []rune(args[1]) // 
		asciiartfs.WithFlag(ouput,color,filename,argsRune,lettersToColor)
	
}






