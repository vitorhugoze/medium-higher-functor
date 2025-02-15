package main

import "fmt"

type Player struct {
	Name   string
	Health int
	Power  func(Player) Player
}

func NewPlayer(name string, power func(Player) Player) Player {
	return Player{
		Name:   name,
		Health: 100,
		Power:  power,
	}
}

func (p1 Player) Apply(p2 Player) Player {
	return p1.Power(p2)
}

func (p Player) Info() string {
	if p.Health > 0 {
		return fmt.Sprintf("Player %s has %d health", p.Name, p.Health)
	} else {
		return fmt.Sprintf("Player %s was taken in combat", p.Name)
	}
}

func Heal(healing int) func(Player) Player {
	return func(p Player) Player {
		p.Health += healing
		return p
	}
}

func DealDamage(damage int) func(Player) Player {
	return func(p Player) Player {
		p.Health -= damage
		return p
	}
}

func main() {
	knight := NewPlayer("Sir Lancelot", DealDamage(90))
	wizard := NewPlayer("Gandalf", Heal(50))
	goblin := NewPlayer("Mozair", DealDamage(30))

	//Sneaky goblin attacks Sir Lancelot twice
	knight = goblin.Apply(knight).Apply(knight)
	fmt.Println(knight.Info())
	//Prints: Player Sir Lancelot has 40 health

	//The mighty Gandalf heals the knight's wounds
	knight = wizard.Apply(knight)
	fmt.Println(knight.Info())
	//Prints: Player Sir Lancelot has 90 health

	//Goblin stumbles and falls face down
	goblin = goblin.Apply(goblin)
	fmt.Println(goblin.Info())
	//Prints: Player Mozair has 70 health

	//The final blow has been delivered
	goblin = knight.Apply(goblin)
	fmt.Println(goblin.Info())
	//Prints: Player Mozair was taken in combat
}
