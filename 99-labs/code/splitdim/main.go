package main

import (
	"encoding/json"
	"log"
	"net/http"
	"splitdim/pkg/api"
	"splitdim/pkg/db/local"
	"splitdim/pkg/db/kvstore"
	"os"
)

// KVStoreMode defines the data layer mode (local/redis/kvstore).
var KVStoreMode = "local"

// KVStoreAddr stores the key-value store address as a DNS domain name or IP address.
var KVStoreAddr = "localhost:8001"

var db api.DataLayer

// TransferHandler is a HTTP handler that implements the money transfer API.
func TransferHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Return HTTP 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Printf("%s %s", r.Method, r.RequestURI)
	var t api.Transfer
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		// Return HTTP 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Transfer request: %+v", t)
	err = db.Transfer(t)
	if err != nil {
		// Return HTTP 400
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

// AccountListHandler is a HTTP handler that returns the current balance of each registered user.
func AccountListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Return HTTP 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Printf("%s %s", r.Method, r.RequestURI)
	accountList, err := db.AccountList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonData, err := json.Marshal(accountList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// ClearHandler is a HTTP handler that returns a list of transfers to clear the balance of each user.
func ClearHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Return HTTP 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Printf("%s %s", r.Method, r.RequestURI)
	transfers, err := db.Clear()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonData, err := json.Marshal(transfers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// ResetHandler is a HTTP handler that allows to zero out all balances.
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Printf("%s %s", r.Method, r.RequestURI)
	err := db.Reset()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	if os.Getenv("KVSTORE_MODE") != "" {
        KVStoreMode = os.Getenv("KVSTORE_MODE")
    }
    if os.Getenv("KVSTORE_ADDR") != "" {
        KVStoreAddr = os.Getenv("KVSTORE_ADDR")
    }

    switch KVStoreMode {
    case "kvstore":
        log.Printf("Using the kvstore datalayer using at %q", KVStoreAddr)
        db = kvstore.NewDataLayer(KVStoreAddr)
    case "local":
        fallthrough
    default:
        log.Println("Using the local datalayer")
        db = local.NewDataLayer()
    }

    // Set the default logger to a fancier log format.
    log.SetFlags(log.LstdFlags | log.Lshortfile)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/api/transfer", TransferHandler)
	http.HandleFunc("/api/accounts", AccountListHandler)
	http.HandleFunc("/api/clear", ClearHandler)
	http.HandleFunc("/api/reset", ResetHandler)

	log.Println("Server listening on http://:8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) 
}