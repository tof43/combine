package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/atotto/clipboard" // Clipboard library
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: combine <file1> <file2> ...")
		os.Exit(1)
	}

	var combinedText strings.Builder

	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", file, err)
			os.Exit(1)
		}
		combinedText.WriteString(fmt.Sprintf("===== START OF %s =====\n", file))
		combinedText.Write(data)
		combinedText.WriteString(fmt.Sprintf("\n===== END OF %s =====\n", file))
	}

	// Copy the combined text to the clipboard
	err := clipboard.WriteAll(combinedText.String())
	if err != nil {
		fmt.Printf("Error copying to clipboard: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Combined text copied to clipboard!")
}

