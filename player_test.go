package main

import (
	"testing"
)

func TestCreateNewPlayer(t *testing.T) {
	expectedPlayer := Player{
		Name:    "TestPlayer",
		Level:   1,
		HP:      20,
		MaxHP:   20,
		Attack:  5,
		Defense: 5,
	}

	player := createNewPlayer("TestPlayer")

	if player.Name != expectedPlayer.Name ||
		player.Level != expectedPlayer.Level ||
		player.HP != expectedPlayer.HP ||
		player.MaxHP != expectedPlayer.MaxHP ||
		player.Attack != expectedPlayer.Attack ||
		player.Defense != expectedPlayer.Defense {
		t.Errorf("createNewPlayer() = %+v, esperado %+v", player, expectedPlayer)
	}
}

func TestPlayerLevelUp(t *testing.T) {
	player := Player{
		Name:    "TestPlayer",
		Level:   1,
		MaxHP:   20,
		HP:      20,
		Attack:  5,
		Defense: 5,
	}

	player.LevelUp()

	expectedLevel := 2
	expectedMaxHP := 22
	expectedAttack := 7
	expectedDefense := 7

	if player.Level != expectedLevel ||
		player.MaxHP != expectedMaxHP ||
		player.Attack != expectedAttack ||
		player.Defense != expectedDefense {
		t.Errorf("LevelUp() = %+v, esperado Level=%d, MaxHP=%d, Attack=%d, Defense=%d",
			player, expectedLevel, expectedMaxHP, expectedAttack, expectedDefense)
	}
}
