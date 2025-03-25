package main

import (
	"log"
	"os"

	"main.go/global"
	"main.go/parse"
)

func main() {
	global.InitConfig(os.Args)
	outFile, err := os.OpenFile(global.Config.OutputFile, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开写文件错误:")
	}
	defer outFile.Close()

	U := &parse.UMLS{
		PackageName: "llms",
		Dir:         global.LR,
		ArrUml:      make([]*parse.Uml, 0),
	}
	parse.ReadInputFile(U)
	parse.WriteUml(U, outFile)
}
