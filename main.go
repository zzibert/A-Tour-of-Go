package main

import "fmt"

type attacker struct {
	attackPower int
	damageBonus int
}

type sword struct {
	attacker
	twoHanded bool
}

type gun struct {
	attacker
	bullletsRemaining int
}

func (s sword) Wield() bool {
	fmt.Println("You've wielded a sword!")
	return true
}

func (g gun) Wield() bool {
	fmt.Println("You've wielded a gun!")
	return true
}

func main() {
	sword1 := sword{attacker: attacker{attackPower: 1, damageBonus: 5}, twoHanded: true}
	gun1 := gun{attacker: attacker{attackPower: 10, damageBonus: 20}, bullletsRemaining: 11}
	fmt.Printf("Weapons: sword: %v, gun: %v\n", sword1, gun1)
	sword1.Wield()
	gun1.Wield()
}
