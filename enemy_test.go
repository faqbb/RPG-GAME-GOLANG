package main

import (
	"testing"
)

func TestCreateNewEnemy(t *testing.T) {
	expectedEnemy := Enemy{
		Name:    "TestEnemy",
		Level:   1,
		HP:      20,
		Attack:  5,
		Defense: 5,
	}

	enemy := createNewEnemy("TestEnemy", 1, 20, 5, 5)

	if enemy.Name != expectedEnemy.Name ||
		enemy.Level != expectedEnemy.Level ||
		enemy.HP != expectedEnemy.HP ||
		enemy.Attack != expectedEnemy.Attack ||
		enemy.Defense != expectedEnemy.Defense {
		t.Errorf("createNewEnemy() = %+v, esperado %+v", enemy, expectedEnemy)
	}
}
