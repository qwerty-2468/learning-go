package kvstore

import (
	"fmt"
	"log"
	"strconv"

	clientapi "kvstore/pkg/api"
	"kvstore/pkg/client"

	"splitdim/pkg/api"
	"splitdim/pkg/clear"
)

type kvstore struct {
	client.Client
}

// NewDataLayer creates a new database of scores.
func NewDataLayer(kvStoreAddr string) api.DataLayer {
	return &kvstore{Client: client.NewClient(kvStoreAddr)}
}

func (db *kvstore) setBalance(user string, amount int) error {
	vv, err := db.Get(user)
	if err != nil {
		return err
	}

	balance, _ := strconv.Atoi(vv.Value)
	vv.Value = fmt.Sprintf("%d", balance+amount)
	vkv := clientapi.VersionedKeyValue{Key: user, VersionedValue: vv}

	if err := db.Put(vkv); err != nil {
		return err
	}
	return nil
}

func (db *kvstore) Transfer(t api.Transfer) error {
	if t.Sender == "" || t.Receiver == "" || t.Sender == t.Receiver {
		return fmt.Errorf("invalid transfer")
	}

	// Credit the sender.
	for {
		err := db.setBalance(t.Sender, t.Amount)
		if err == nil {
			break
		}
		log.Printf("Retrying setBalance for sender: %v", err)
	}

	// Debit the receiver.
	for {
		err := db.setBalance(t.Receiver, -t.Amount)
		if err == nil {
			break
		}
		log.Printf("Retrying setBalance for receiver: %v", err)
	}

	return nil
}

func (db *kvstore) AccountList() ([]api.Account, error) {
	accounts, err := db.List()
	if err != nil {
		return nil, err
	}

	ret := []api.Account{}
	for _, kv := range accounts {
		balance, err := strconv.Atoi(kv.Value)
		if err != nil {
			return nil, fmt.Errorf("invalid balance for user %s: %v", kv.Key, err)
		}
		ret = append(ret, api.Account{Holder: kv.Key, Balance: balance})
	}
	return ret, nil
}

func (db *kvstore) Clear() ([]api.Transfer, error) {
	accounts, err := db.List()
	if err != nil {
		return nil, err
	}

	balances := make(map[string]int)
	for _, kv := range accounts {
		balance, err := strconv.Atoi(kv.Value)
		if err != nil {
			return nil, fmt.Errorf("invalid balance for user %s: %v", kv.Key, err)
		}
		balances[kv.Key] = balance
	}

	// Make a copy to avoid mutating the original map
	balancesCopy := make(map[string]int)
	for k, v := range balances {
		balancesCopy[k] = v
	}

	transfers, err := clear.Clear(balancesCopy)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}

// Reset sets all balances to zero.
func (db *kvstore) Reset() error {
	return db.Client.Reset()
}
