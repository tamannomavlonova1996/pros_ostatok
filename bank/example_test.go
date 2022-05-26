package bank

import (
	"fmt"
	"pros_ostatok/types"
)

func ExampleAddBonus() {
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
	AddBonus(&toma, 3, 30, 365)
	fmt.Println(toma.Balance)

	// Output:
	// 62
}
