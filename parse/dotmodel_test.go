package parse

import (
	"reflect"
	"strings"
	"testing"

	"main.go/global"
)

func TestGetRelationShipType(t *testing.T) {
	str := "arrowhead=\"empty\", arrowtail=\"none\", fontcolor=\"green\", label=\"style\", style=\"solid\"];]"
	got := getRelationShipType(0, strings.Index(str, "arrowtail"), str)
	want := global.EXTEND
	if got != want {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

func TestGetDefClassName(t *testing.T) {
	str := "pydantic.main.BaseModel\""
	got := getDefClassName(str)
	want := "BaseModel"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}
}

func TestGetLabel(t *testing.T) {
	line := "\"ragas.embeddings.base.BaseRagasEmbeddings\" -> \"ragas.prompt.few_shot_pydantic_prompt.InMemoryExampleStore\" [arrowhead=\"diamond\", arrowtail=\"none\", fontcolor=\"green\", label=\"embeddings\", style=\"solid\"];"
	s := strings.Split(line, "[")
	_, relationType := s[0], s[1]
	got := getLabel(relationType)
	want := "embeddings"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("excepted:%v, got:%v", want, got)
	}

}
