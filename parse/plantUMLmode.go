package parse

import (
	"log"
	"os"

	"main.go/global"
)

var plantClassRelation = []string{
	" {\n\tfield : type\n\tmethod(type) : type\n}\n",
	" --|> ", //继承关系
	" ..|> ", //实现关系
	" -- ",   //关联关系
	" --o ",  //聚合关系
	" --* ",  //组合关系
}

type UMLS struct {
	PackageName string
	Dir         global.Direction
	ArrUml      []*Uml
	//
}

type Uml struct {
	Type       global.ClassType
	Label      string
	LeftClass  string
	RightClass string
}

func WriteUml(U *UMLS, outFile *os.File) {

	outFile.WriteString("@startUml\n")
	switch U.Dir {
	case global.LR:
		if _, err := outFile.WriteString("left to right direction\n"); err != nil {
			log.Fatalf("write dir wrong!")
		}
	case global.BT:
		if _, err := outFile.WriteString("bottom to top direction\n"); err != nil {
			log.Fatalf("write dir wrong!")
		}
	case global.TB:
		if _, err := outFile.WriteString("top to bottom direction\n"); err != nil {
			log.Fatalf("write dir wrong!")
		}
	case global.RL:
		if _, err := outFile.WriteString("right to left direction\n"); err != nil {
			log.Fatalf("write dir wrong!")
		}
	default:
		log.Fatalf("no such direction!")
	}

	for _, u := range U.ArrUml {
		writeUml(u, outFile)
	}

	outFile.WriteString("@endUml\n")
}

func writeUml(u *Uml, outFile *os.File) {
	if u.Type == global.DEF {
		if _, err := outFile.WriteString("class " + u.LeftClass + plantClassRelation[u.Type]); err != nil {
			log.Fatalf("write wrong!")
		}
	} else {
		if _, err := outFile.WriteString(u.LeftClass + plantClassRelation[u.Type] + u.RightClass); err != nil {
			log.Fatalf("write classrelationship wrong!")
		}
		if u.Label != "" {
			if _, err := outFile.WriteString(" : " + u.Label + "\n"); err != nil {
				log.Fatalf("write label wrong!")
			}
		} else {
			outFile.WriteString("\n")
		}

	}
}
