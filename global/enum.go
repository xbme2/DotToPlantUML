package global

type Direction int

const (
	LR Direction = iota
	RL
	BT
	TB
)

type ClassType int

const (
	DEF ClassType = iota
	EXTEND
	IMPLEMENT
	ASSOCIATE
	POLY
	COMBANITION
)
