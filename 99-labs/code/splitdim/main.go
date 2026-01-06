package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"splitdim/pkg/api"
	"splitdim/pkg/db/kvstore"
	"splitdim/pkg/db/local"
)

// KVStoreMode defines the data layer mode (local/redis/kvstore).
var KVStoreMode = "local"

// KVStoreAddr stores the key-value store address as a DNS domain name or IP address.
// Default fallback used when neither environment variable nor CLI arg is provided.
var KVStoreAddr = "localhost:8081"

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
	// CLI flags with sensible defaults from environment variables.
	defMode := os.Getenv("KVSTORE_MODE")
	if defMode == "" {
		defMode = "local"
	}
	defAddr := os.Getenv("KVSTORE_ADDR")
	if defAddr == "" {
		defAddr = "localhost:8081"
	}

	modeFlag := flag.String("mode", defMode, "data-layer mode (local/kvstore)")
	addrFlag := flag.String("addr", defAddr, "key-value store address (host:port)")
	flag.Parse()

	// Final settings: CLI flag takes precedence over environment variable.
	KVStoreMode = *modeFlag
	KVStoreAddr = *addrFlag

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

	server := &http.Server{Addr: ":8080"}

	go func() {
		log.Println("Server listening on http://:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutdown signal received, waiting for in-flight requests")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
}
