//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

var monsterPlayerSet = wire.NewSet(NewMonster, NewPlayer)
var endingASet = wire.NewSet(monsterPlayerSet, wire.Struct(new(EndingA), "*"))
var endingBSet = wire.NewSet(wire.Struct(new(EndingB), "*"), monsterPlayerSet)

// func InitMission(n1 PlayerParam, n2 MonsterParam) (Mission, error) {
func InitMission() (Mission, error) {
	// wire.Build(monsterPlayerSet, NewMission)
	wire.Build(wire.Struct(new(Mission)))
	return Mission{}, nil
}

func InitEndingA(n1 PlayerParam, n2 MonsterParam) (EndingA, error) {
	wire.Build(endingASet)
	return EndingA{}, nil
}

func InitEndingB(n1 PlayerParam, n2 MonsterParam) (EndingB, error) {
	wire.Build(endingBSet)
	return EndingB{}, nil
}

