package main

import (
	"testing"
)

func TestFightResult_PlayerWins(t *testing.T) {
	player := Player{
		Name:    "TestPlayer",
		Level:   1,
		MaxHP:   20,
		HP:      20,
		Attack:  10,
		Defense: 5,
	}

	enemy := Enemy{
		Name:    "TestEnemy",
		Level:   1,
		HP:      10,
		Attack:  5,
		Defense: 3,
	}

	result := fightResult(player, enemy)

	if !result {
		t.Errorf("fightResult() returned false, expected true")
	}
}

func TestBossFight(t *testing.T) {
	player := Player{
		Name:    "TestPlayer",
		Level:   1,
		MaxHP:   20,
		HP:      20,
		Attack:  10,
		Defense: 5,
	}

	boss := Enemy{
		Name:    "TestBoss",
		Level:   5,
		HP:      50,
		Attack:  15,
		Defense: 10,
	}

	result := bossFight(player, boss)

	if !result {
		t.Errorf("bossFight() returned false, expected true")
	}
}
