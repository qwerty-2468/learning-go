package clear

import (
	"fmt"

	"splitdim/pkg/api"
)

// Clear returns the set of transfers required to zero-out all balances.
// It mutates the supplied map, so callers should pass a copy if they need to retain the original values.
func Clear(accounts map[string]int) ([]api.Transfer, error) {
	// Consistency check: total balance must be zero.
	total := 0
	for _, balance := range accounts {
		total += balance
	}
	if total != 0 {
		return nil, fmt.Errorf("accounts inconsistent: total balance is not zero")
	}

	transfers := make([]api.Transfer, 0)

	// Find creditors (negative balance) and debtors (positive balance).
	for sender, balance := range accounts {
		if balance >= 0 {
			continue
		}
		for receiver, receiverBalance := range accounts {
			if receiverBalance <= 0 {
				continue
			}
			transferAmount := min(-balance, receiverBalance)
			if transferAmount == 0 {
				continue
			}
			transfers = append(transfers, api.Transfer{
				Sender:   sender,
				Receiver: receiver,
				Amount:   transferAmount,
			})
			accounts[sender] += transferAmount
			accounts[receiver] -= transferAmount
			if accounts[sender] == 0 {
				break
			}
		}
	}
	return transfers, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
