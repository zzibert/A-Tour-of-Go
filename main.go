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

type weapon interface {
	Wield() bool
}

func wielder(w weapon) bool {
	fmt.Println("Wielding...")
	return w.Wield()
}

func main() {
	sword1 := sword{attacker: attacker{attackPower: 1, damageBonus: 5}, twoHanded: true}
	gun1 := gun{attacker: attacker{attackPower: 10, damageBonus: 20}, bullletsRemaining: 11}
	wielder(sword1)
	wielder(gun1)
}
