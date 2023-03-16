package main

import (
	"bufio"
	"fmt"
	"github.com/adrg/strutil/metrics"
	"os"
)

func main() {
	j := metrics.NewJaccard()
	j.CaseSensitive = false

	for {
		inputText := readUserInput()
		existingTexts := readFile()

		for _, line := range existingTexts {
			jaccardResult := j.Compare(inputText, line)
			fmt.Printf("%s - Jaccaard Score: %v\n", line, jaccardResult)
		}
		writeFile(inputText)
	}
}

func readUserInput() string {
	fmt.Println("Type new message from C1...")
	inputReader := bufio.NewReader(os.Stdin)
	inputText, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		panic("Unable to read input")
	}
	return inputText
}

func readFile() []string {
	var resultSlice []string
	readFile, err := os.Open("existing_messages.txt")

	if err != nil {
		fmt.Println(err)
		panic("Unable to open existing messages file")
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		resultSlice = append(resultSlice, fileScanner.Text())
	}
	return resultSlice
}

func writeFile(line string) {
	f, err := os.OpenFile("existing_messages.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		panic("Unable to open existing messages file")
	}
	f.WriteString(line)
}
