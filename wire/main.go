package main

import (
	"errors"
	"fmt"
	"time"
)

type PlayerParam string
type MonsterParam string
type Monster struct {
	Name string
}

func NewMonster(name MonsterParam) Monster {
	return Monster{Name: string(name)}
}

type Player struct {
	Name string
}

func NewPlayer(name PlayerParam) (Player, error) {
	//return Player{Name: string(name)}
	if time.Now().Unix()%2 == 0 {
		return Player{},  errors.New("player dead")
	}
	return Player{Name: string(name)},nil
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

type EndingA struct {
	Player  Player
	Monster Monster
}


func (e EndingA) appear() {
	fmt.Printf("%s defeats %s, world peace!\n", e.Player.Name, e.Monster.Name)
}

type EndingB struct {
	Player  Player
	Monster Monster
}


func (e EndingB) appear() {
	fmt.Printf("%s defeats %s, World peace!\n", e.Player.Name, e.Monster.Name)
}

func main() {
	//mission, _ := InitMission("Dj", "jake")
	//
	//mission.Start()

	//endingA ,_ := InitEndingA("A", "a")
	//
	//endingA.appear()

	endingB, _ := InitEndingB("B", "b")
	endingB.appear()
}
