package global

import (
	"fmt"
	"log"
)

type ConFig struct {
	InputFile  string
	OutputFile string
	Params     []byte
}

var Config = &ConFig{
	InputFile:  "",
	OutputFile: "",
	Params:     []byte{},
}

func InitConfig(commandArgs []string) {
	if len(commandArgs) <= 1 {
		log.Fatalf("you should add input file")
	}
	for _, v := range commandArgs[1:] {
		fmt.Println(v)
		if v[0] == '-' {
			if len(v) > 2 {
				log.Fatalf("wrong param: you can input -h for help")
			}
			Config.Params = append(Config.Params, v[1])

		} else if v != "." {
			if Config.InputFile == "" {
				Config.InputFile = v
			} else {
				Config.OutputFile = v
			}

		}
	}
	if Config.OutputFile == "" {
		Config.OutputFile = Config.InputFile[:len(Config.InputFile)-4] + ".puml"
	}
	fmt.Println(Config.InputFile)
	fmt.Println(Config.OutputFile)

}
