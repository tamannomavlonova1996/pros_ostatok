package bank

import "pros_ostatok/types"

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !card.Active {
		return
	}
	if card.MinBalance < 0 {
		return
	}
	bonus := int(card.MinBalance) * percent * daysInMonth / daysInYear
	if bonus > 5000 {
		bonus = 5000
	}
	card.Balance = types.Money(bonus) + card.Balance

}
