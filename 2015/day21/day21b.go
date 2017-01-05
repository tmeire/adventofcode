package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Cost   int
	Damage int
	Armor  int
}

type ItemSet struct {
	Weapon *Item
	Armor  *Item
	Ring1  *Item
	Ring2  *Item
}

func (is *ItemSet) CostPoints() int {
	c := is.Weapon.Cost
	if is.Armor != nil {
		c += is.Armor.Cost
	}
	if is.Ring1 != nil {
		c += is.Ring1.Cost
	}
	if is.Ring2 != nil {
		c += is.Ring2.Cost
	}
	return c
}

func (is *ItemSet) DamagePoints() int {
	c := is.Weapon.Damage
	if is.Armor != nil {
		c += is.Armor.Damage
	}
	if is.Ring1 != nil {
		c += is.Ring1.Damage
	}
	if is.Ring2 != nil {
		c += is.Ring2.Damage
	}
	return c
}

func (is *ItemSet) ArmorPoints() int {
	c := is.Weapon.Armor
	if is.Armor != nil {
		c += is.Armor.Armor
	}
	if is.Ring1 != nil {
		c += is.Ring1.Armor
	}
	if is.Ring2 != nil {
		c += is.Ring2.Armor
	}
	return c
}

type ItemSetList []*ItemSet

func (s ItemSetList) Len() int {
	return len(s)
}
func (s ItemSetList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ItemSetList) Less(i, j int) bool {
	return s[i].CostPoints() > s[j].CostPoints()
}

func variations(weapons []*Item, armors []*Item, rings []*Item) []*ItemSet {
	vars := make([]*ItemSet, 0)

	// only the weapons
	for _, w := range weapons {
		vars = append(vars, &ItemSet{w, nil, nil, nil})
	}

	// only the weapons and armor
	for _, w := range weapons {
		for _, a := range armors {
			vars = append(vars, &ItemSet{w, a, nil, nil})
		}
	}

	// only the weapons and rings
	for _, w := range weapons {
		for i, r1 := range rings {
			vars = append(vars, &ItemSet{w, nil, r1, nil})

			for j := i + 1; j < len(rings); j += 1 {
				vars = append(vars, &ItemSet{w, nil, r1, rings[j]})
			}
		}
	}

	// the weapons, armor and rings
	for _, w := range weapons {
		for _, a := range armors {
			for i, r1 := range rings {
				vars = append(vars, &ItemSet{w, a, r1, nil})

				for j := i + 1; j < len(rings); j += 1 {
					vars = append(vars, &ItemSet{w, a, r1, rings[j]})
				}
			}
		}
	}
	return vars
}

type Fighter struct {
	health int
	damage int
	armor  int
}

func fight(f1 Fighter, f2 Fighter) bool {
	for true {
		f2.health -= f1.damage - f2.armor
		if f2.health <= 0 {
			return true
		}
		f1.health -= f2.damage - f1.armor
		if f1.health <= 0 {
			return false
		}
	}
	return false
}

func main() {
	weapons := []*Item{
		&Item{8, 4, 0},
		&Item{10, 5, 0},
		&Item{25, 6, 0},
		&Item{40, 7, 0},
		&Item{74, 8, 0},
	}
	armors := []*Item{
		&Item{13, 0, 1},
		&Item{31, 0, 2},
		&Item{53, 0, 3},
		&Item{75, 0, 4},
		&Item{102, 0, 5},
	}
	rings := []*Item{
		&Item{25, 1, 0},
		&Item{50, 2, 0},
		&Item{100, 3, 0},
		&Item{20, 0, 1},
		&Item{40, 0, 2},
		&Item{80, 0, 3},
	}

	vars := variations(weapons, armors, rings)
	sort.Sort(ItemSetList(vars))
	for _, varx := range vars {
		if !fight(Fighter{100, varx.DamagePoints(), varx.ArmorPoints()}, Fighter{109, 8, 2}) {
			fmt.Println(varx.CostPoints())
			return
		}
	}
}
