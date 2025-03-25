package parse

import (
	"fmt"
	"log"
	"strings"

	"main.go/global"
)

func DotParse(line string, U *UMLS) {
	fmt.Println(line)
	s := strings.Split(line, "[")
	// 先判断是全局信息还是类关系定义
	if len(s) == 1 {
		// 存储全局信息
		// digraph "classes_prompt" {
		// 	rankdir=LR
		// 	charset="utf-8"
		if s[0] == "}" {
			return
		}
		target := s[0]
		switch target {
		case "digraph":
			className := strings.Split(line, "_")
			for _, char := range className[1] {
				if char != rune('"') {
					U.PackageName += string(char)
				} else {
					break
				}

			}
		case "rankdir":
			dst := line[8:]
			switch dst {
			case "LR":
				U.Dir = global.LR
			case "BT":
				U.Dir = global.BT
			case "RL":
				U.Dir = global.RL
			case "TB":
				U.Dir = global.TB
			default:
				log.Fatalf("Direction wrong!")
			}
		case "charset":
			break
		}
	} else {
		//"pydantic.main.BaseModel" [color="black", fontcolor="black", label=<BaseModel>, shape="record", style="solid"];
		// "ragas.embeddings.base.BaseRagasEmbeddings" -> "ragas.prompt.few_shot_pydantic_prompt.InMemoryExampleStore" [arrowhead="diamond", arrowtail="none", fontcolor="green", label="embeddings", style="solid"];
		relationShip, relationType := s[0], s[1]
		var u *Uml = new(Uml)
		midIndex := strings.Index(relationShip, "->")
		if midIndex != -1 {
			// 是类与类之间关系的定义
			u.LeftClass = getDefClassName(relationShip[:midIndex])
			u.RightClass = getDefClassName(relationShip[midIndex:])
			u.Label = getLabel(relationType)
			u.Type = getRelationShipType(0, strings.Index(relationType, "arrowtail"), relationType)
		} else {
			// 类定义
			u.Type = global.DEF
			u.LeftClass = getDefClassName(relationShip)
		}
		U.ArrUml = append(U.ArrUml, u)
	}
}

func getDefClassName(relationShip string) string {
	//"pydantic.main.BaseModel" [
	left, right := 0, 0
	for i := len(relationShip) - 1; i >= 0; i-- {
		if relationShip[i] == '"' && right == 0 {
			right = i
		}
		if relationShip[i] == '.' {
			left = i + 1
			break
		}
	}
	return relationShip[left:right]
}

func getLabel(relationType string) string {
	labelIndex := strings.Index(relationType, "label")
	if labelIndex == -1 {
		return ""
	}
	left, right := 0, 0
	for i := labelIndex; i < len(relationType); i++ {
		if relationType[i] == '"' && left == 0 {
			left = i + 1
		} else if relationType[i] == '"' {
			right = i
			break
		}
	}
	return relationType[left:right]
}

func getRelationShipType(arrowheadIdx int, arrowtailIdx int, relationType string) global.ClassType {
	//arrowhead="diamond", arrowtail="none", fontcolor="green", label="embeddings", style="solid"];

	distance := 11
	// if arrowheadIdx+distance >= len(relationType) || arrowtailIdx+distance >= len(relationType) {
	// 	return global.ASSOCIATE // 默认返回关联关系
	// }
	headByte, tailByte := relationType[arrowheadIdx+distance], relationType[arrowtailIdx+distance]
	switch headByte {
	case 'e':
		switch tailByte {
		case 'v':
			return global.ASSOCIATE
		case 'n':
			return global.EXTEND
		}
	case 'd':
		return global.POLY
	case 'o':
		return global.COMBANITION
	default:
		fmt.Println("all not match")
		fmt.Println("headByte is ", string(headByte))
		fmt.Println("tailByte is ", string(tailByte))
		log.Fatalf("never return here")

	}
	// never return	here
	log.Fatalf("never return here")
	return global.ASSOCIATE
}
