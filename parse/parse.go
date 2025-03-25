package parse

import (
	"bufio"
	"log"
	"os"
	"strings"

	"main.go/global"
)

func ReadInputFile(U *UMLS) {
	f, err := os.Open(global.Config.InputFile)
	if err != nil {
		log.Fatalf("读文件错误:")
	}
	defer f.Close()

	outFile, err := os.OpenFile(global.Config.OutputFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开写文件错误:")
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(f)
	lineCount := 0
	var parseLine string
	for scanner.Scan() {
		line := scanner.Text()
		if lineCount < 3 {
			DotParse(line, U)
		} else {
			if strings.Contains(line, ";") {
				parseLine += line
				DotParse(parseLine, U)
			} else {
				parseLine = line
			}
		}
	}
}

// func ScanSentences(data []byte, atEOF bool) (advance int, token []byte, err error) {

// }
