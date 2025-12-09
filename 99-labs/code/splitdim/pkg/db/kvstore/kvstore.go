package kvstore

import (
	"fmt"
	"log"
	"strconv"
	"time"

	clientapi "kvstore/pkg/api"
	"kvstore/pkg/client"

	"resilient"
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

var defaultBackoff = resilient.Backoff{
	Base:      150 * time.Millisecond,
	Cap:       2 * time.Second,
	Jitter:    1,
	NumTrials: 6,
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

func (db *kvstore) setBalanceForUser(user string, amount int) resilient.Closure {
	return func() error { return db.setBalance(user, amount) }
}

func (db *kvstore) Transfer(t api.Transfer) error {
	if t.Sender == "" || t.Receiver == "" || t.Sender == t.Receiver {
		return fmt.Errorf("invalid transfer")
	}

	// Credit the sender.
	sender := resilient.WithRetry(db.setBalanceForUser(t.Sender, t.Amount), defaultBackoff)
	if err := sender(); err != nil {
		return fmt.Errorf("transfer: could not set balance for user %q: %w", t.Sender, err)
	}

	// Debit the receiver.
	receiver := resilient.WithRetry(db.setBalanceForUser(t.Receiver, -t.Amount), defaultBackoff)
	if err := receiver(); err != nil {
		log.Printf("transfer: could not set balance for receiver %q after retries, attempting undo: %v", t.Receiver, err)

		undoBackoff := resilient.Backoff{Base: 300 * time.Millisecond, Cap: 4 * time.Second, Jitter: 1, NumTrials: 8}
		undoSender := resilient.WithRetry(db.setBalanceForUser(t.Sender, -t.Amount), undoBackoff)
		if undoErr := undoSender(); undoErr != nil {
			return fmt.Errorf("transfer: receiver update failed (%v) and undo for sender %q also failed: %w", err, t.Sender, undoErr)
		}

		return fmt.Errorf("transfer: receiver update failed after retries: %w", err)
	}

	return nil
}

func (db *kvstore) AccountList() ([]api.Account, error) {
	list := resilient.WithRetry(func() error {
		_, err := db.List()
		return err
	}, defaultBackoff)

	if err := list(); err != nil {
		return nil, err
	}

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
	list := resilient.WithRetry(func() error {
		_, err := db.List()
		return err
	}, defaultBackoff)

	if err := list(); err != nil {
		return nil, err
	}

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
	reset := resilient.WithRetry(func() error {
		return db.Client.Reset()
	}, defaultBackoff)
	return reset()
}
