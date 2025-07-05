package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func printLine(AsciiShapes []string, char int, i int) {
	str := strings.Split(AsciiShapes[char-32], "\n")
	fmt.Print(str[i])
}

func CheckNewLines(arr []string) bool {
	for _, v := range arr {
		if v != "" {
			return false
		}
	}
	return true
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Invalid number of areguments...")
		return
	}
	input := args[0]
	fileName := "standard.txt"
	//fileName := "shadow.txt"
	//fileName := "thinkertoy.txt"

	info, err := os.Stat(fileName)
	if err == nil {
		if info.Mode() != 0o400 {
			erRemoveFile := exec.Command("rm", "-f", fileName).Run()
			if erRemoveFile != nil {
				fmt.Println(erRemoveFile)
			}
		}
	}

	_, errOpen := os.Open(fileName)
	if errOpen != nil {
		if os.IsNotExist(errOpen) {
			url := "https://learn.zone01oujda.ma/api/content/root/public/subjects/ascii-art/" + fileName

			// download the file
			errDownload := exec.Command("wget", "-q", url).Run()
			if errDownload != nil {
				fmt.Println(errDownload)
			}

			// change the permission
			erPermission := exec.Command("chmod", "400", fileName).Run()
			if erPermission != nil {
				fmt.Println("Error changing the permission:", erPermission)
				fmt.Println(erPermission)
			}
		}
	}

	file, er := os.Open(fileName)
	if er != nil {
		fmt.Println(er)
	}
	defer file.Close()

	fileData, erR := io.ReadAll(file)
	if erR != nil {
		fmt.Println(erR)
	}

	CleanFileData := strings.ReplaceAll(string(fileData), "\r", "")
	AsciiShapes := strings.Split(CleanFileData, "\n\n")

	// validate user input
	for _, char := range input {
		if char < 32 || char > 126 {
			fmt.Println("The string includes characters outside the ASCII range...")
			return
		}
	}

	inputLines := strings.Split(input, "\\n")

	// if the input contain only new lines
	if CheckNewLines(inputLines) {
		inputLines = inputLines[1:]
	}
	for _, Inpline := range inputLines {
		if Inpline == "" {
			fmt.Println()
			continue
		}
		for i := 0; i < 8; i++ {
			for _, char := range Inpline {
				if char != ' ' {
					printLine(AsciiShapes, int(char), i)
				} else {
					printLine(AsciiShapes, int(char), i+1)
				}
			}
			fmt.Println()
		}
	}
}
