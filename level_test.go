package main

import (
	"testing"
)

func TestCreateNewLevel(t *testing.T) {
	expectedLevel := Level{
		KillCondition: 3,
		isBossLevel:   false,
	}

	level := createNewLevel(false, 3)

	if level.KillCondition != expectedLevel.KillCondition ||
		level.isBossLevel != expectedLevel.isBossLevel {
		t.Errorf("createNewLevel() = %+v, esperado %+v", level, expectedLevel)
	}
}
