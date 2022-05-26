package main

import (
	"fmt"
	"pros_ostatok/bank"
	"pros_ostatok/types"
)

func main() {

	toma := types.Card{
		ID:         1,
		PAN:        "20250297",
		Balance:    60,
		MinBalance: 10,
		Currency:   "TJS",
		Color:      "black",
		Name:       "Tamanno",
		Active:     true,
	}
	bank.AddBonus(&toma, 3, 30, 365)
	fmt.Println(toma)
}
