package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Player struct {
	Name    string
	Level   int
	MaxHP   int
	HP      int
	Attack  int
	Defense int
}

type Enemy struct {
	Name    string
	Level   int
	HP      int
	Attack  int
	Defense int
}

type Level struct {
	KillCondition int
	isBossLevel   bool
}

type Levels struct {
	levels []Level
}

func createNewPlayer(name string) Player {
	newPLayer := Player{
		Name:    name,
		Level:   1,
		HP:      20,
		MaxHP:   20,
		Attack:  5,
		Defense: 5,
	}
	return newPLayer

}

func createNewEnemy(name string, level int, hp int, damage int, defense int) Enemy {
	newEnemy := Enemy{
		Name:    name,
		Level:   level,
		HP:      hp,
		Attack:  damage,
		Defense: defense,
	}
	return newEnemy
}

func createNewLevel(isBoss bool, killCondition int) Level {
	newLevel := Level{
		KillCondition: killCondition,
		isBossLevel:   isBoss,
	}
	return newLevel
}

func fightResult(player Player, enemy Enemy) bool {
	var isEnemyDead bool = false
	for player.HP > 0 && enemy.HP > 0 {
		var playerAttack int = player.Attack + randomNum(1, 10)
		var enemyAttack int = enemy.Attack + randomNum(1, 10)
		fmt.Printf("Atacas con %d puntos de ataque! %s se defiende con %d puntos de defensa \n", playerAttack, enemy.Name, enemy.Defense)
		enemy.HP -= playerAttack - enemy.Defense
		if enemy.HP <= 0 {
			isEnemyDead = true
			fmt.Printf("El %s murio!!! ", enemy.Name)
			break
		} else {
			fmt.Printf("Al enemigo le quedan %d puntos de vida \n", enemy.HP)
		}
		fmt.Printf("Te quedan %d puntos de vida \n", player.HP)
		fmt.Printf("El enemigo ataca con %d de poder! %s se defiende con %d puntos de defensa \n", enemyAttack, player.Name, player.Defense)
		player.HP -= enemyAttack - player.Defense
		if player.HP <= 0 {
			isEnemyDead = false
			break
		} else {
			fmt.Printf("te quedan %d puntos de vida \n", player.HP)
		}
	}
	return isEnemyDead
}

func bossFight(player Player, boss Enemy) bool {
	fmt.Printf("El jefe %s te desafia! Que vas a hacer? \n", boss.Name)
	fmt.Printf("Presiona 1 para pelear o 2 para huir \n")
	input := getInput()
	result := false
	if input == 1 {
		result = fightResult(player, boss)
	} else {
		fmt.Printf("No pudiste huir, %s te apuñaló por la espalda \n", boss.Name)
	}
	return result
}

func (player *Player) LevelUp() {
	fmt.Printf("Felicidades %s! Subiste de nivel! ", player.Name)
	player.Level++
	player.MaxHP += 2
	player.Attack += 2
	player.Defense += 2
	fmt.Printf("Estas son tus nuevas estadísticas:\n")
	fmt.Printf("Nivel: %d\n", player.Level)
	fmt.Printf("Max HP: %d\n", player.MaxHP)
	fmt.Printf("Ataque: %d\n", player.Attack)
	fmt.Printf("Defensa: %d\n\n", player.Defense)
	continueInput()
}

func (player *Player) GetLoot(quality int) {
	fmt.Printf("Has encontrado un tesoro! Elige uno \n\n")
	fmt.Printf("Estas son tus estadisticas \n")
	fmt.Printf("HP: %d/%d \n", player.HP, player.MaxHP)
	fmt.Printf("Ataque: %d\n", player.Attack)
	fmt.Printf("Defensa: %d\n\n", player.Defense)
	fmt.Printf("Presiona 1 para tomar la pocion que te recupera toda la vida \n")
	fmt.Printf("Presiona 2 para tomar una espada que mejora tu ataque \n")
	fmt.Printf("Presiona 3 para tomar una armadura que mejora tu defensa \n")

	input := getInput()
	switch input {
	case 1:
		fmt.Printf("Haz recuperado toda la vida! \n\n")
		player.HP = player.MaxHP
	case 2:
		fmt.Printf("Haz tomado la espada, tu ataque mejora en %d puntos! \n\n", quality)
		player.Attack += 1 * quality
	case 3:
		fmt.Printf("Haz tomado la armadura, tu defensa mejora en %d puntos! \n\n", quality)
		player.Defense += 1 * quality
	}

}

func getPlayerName() string {
	var data string
	fmt.Scanln(&data)
	return data
}

func continueInput() int {
	fmt.Println("Presione un numero para continuar")
	input := getInput()
	return input
}

func getEnemyName() string {
	names := []string{"hongo", "perro", "lobo", "oso", "dragon", "mapache", "goblin", "fantasma", "demonio"}
	return names[randomNum(0, len(names)-1)]
}

func getInput() int {
	var input string
	fmt.Scanln(&input)
	command, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Ingrese un numero valido")
		return 0
	}
	return command
}

func randomNum(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {
	fmt.Println("Ingrese 1 para iniciar el juego")
	var input int = getInput()
	if input == 1 {
		fmt.Println("Ingrese su nombre")
		boss := Enemy{
			Name:    "Damian de RRHH",
			Level:   5,
			HP:      5,
			Attack:  10,
			Defense: 5,
		}
		var isBossDead bool = false
		var isPlayerDead bool = false
		var cantLevels int = 1
		var name string = getPlayerName()
		var player Player = createNewPlayer(name)
		for !isPlayerDead && !isBossDead {
			var min int = 1 * cantLevels
			var max int = 2 * cantLevels
			var isBossLevel bool = (cantLevels == 4)
			var level Level = createNewLevel(isBossLevel, randomNum(1, 4))
			if !isBossLevel {
				var expGained int = 0
				for killsInLevel := 0; killsInLevel <= level.KillCondition; {
					var enemy Enemy = createNewEnemy(getEnemyName(), randomNum(min, max), randomNum(min, max), randomNum(min, max), randomNum(min, max))
					fmt.Printf("Ha aparecido el enemigo %s que tiene %d cantidad de vida\n", enemy.Name, enemy.HP)
					continueInput()
					if fightResult(player, enemy) {
						fmt.Printf("Has vencido al %s \n\n", enemy.Name)
						expGained++
						if randomNum(0, 1) == 1 {
							player.GetLoot(cantLevels)
						}
						if expGained == 2 {
							player.LevelUp()
						}
						killsInLevel++
					} else {
						isPlayerDead = true
						fmt.Println("Has muerto. Fin del juego.")
						break // Salimos del bucle si el jugador muere
					}

				}
			} else {
				fmt.Println("Te encontraste al jefe!!")
				isBossDead = bossFight(player, boss)
				if !isBossDead {
					isPlayerDead = true
					fmt.Println("Has muerto. Fin del juego.")
					break
				}
			}
			if isBossDead {
				fmt.Printf("Felicidades %s, ganaste el juego!! \n\n", player.Name)
				break
			}
			if !isPlayerDead {
				fmt.Printf("Pasas al siguiente area. \n\n")
				cantLevels++
				continueInput()
			}
		}
	}

}
