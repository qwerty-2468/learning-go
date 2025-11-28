package local

import (
	"fmt"
	"sort"
	"sync"

	"splitdim/pkg/api"
)

// localDB is a simple implementation of the DataLayer API.
type localDB struct {
    // accounts maintains the balance for each user name
    accounts map[string]int
    // The read-write mutex makes sure concurrent access is safe.
    mu sync.RWMutex
}

// NewDataLayer creates a new database of accounts.
func NewDataLayer() api.DataLayer {
    return &localDB{accounts: make(map[string]int)}
}

func (db *localDB) Transfer(t api.Transfer) error {
    if t.Sender == t.Receiver {
        return fmt.Errorf("sender and receiver must be different")
    }
    if t.Amount <= 0 {
        return fmt.Errorf("amount must be positive")
    }
    db.mu.Lock()
    defer db.mu.Unlock()
    // Initialize sender and receiver to zero if not present
    if _, ok := db.accounts[t.Sender]; !ok {
        db.accounts[t.Sender] = 0
    }
    if _, ok := db.accounts[t.Receiver]; !ok {
        db.accounts[t.Receiver] = 0
    }
    db.accounts[t.Sender] += t.Amount
    db.accounts[t.Receiver] -= t.Amount
    return nil

}
func (db *localDB) AccountList() ([]api.Account, error) {
    db.mu.RLock()
    defer db.mu.RUnlock()
    ret := []api.Account{}
    for name, balance := range db.accounts {
        ret = append(ret, api.Account{Holder: name, Balance: balance})
    }
    // Sort by account holder name
    sort.Slice(ret, func(i, j int) bool {
        return ret[i].Holder < ret[j].Holder
    })
    return ret, nil
}
func (db *localDB) Clear() ([]api.Transfer, error) {
    db.mu.RLock()

    // Consistency check: total balance must be zero
    total := 0
    for _, balance := range db.accounts {
        total += balance
    }
    if total != 0 {
        return nil, fmt.Errorf("database inconsistent: total balance is not zero")
    }

    // Copy accounts to tempAcc
    tempAcc := make(map[string]int, len(db.accounts))
    for k, v := range db.accounts {
        tempAcc[k] = v
    }

    db.mu.RUnlock()

    transfers := make([]api.Transfer, 0)

    // Find creditors (negative balance) and debtors (positive balance)
    for sender, balance := range tempAcc {
        if balance >= 0 {
            continue
        }
        for receiver, receiverBalance := range tempAcc {
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
            tempAcc[sender] += transferAmount
            tempAcc[receiver] -= transferAmount
            if tempAcc[sender] == 0 {
                break
            }
        }
    }
    return transfers, nil
}

// min returns the smaller of two ints.
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func (db *localDB) Reset() error { 
    db.mu.Lock()
    db.accounts = make(map[string]int)
    defer db.mu.Unlock()
    return nil
}