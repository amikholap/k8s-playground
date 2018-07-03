package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	port := getEnvInt("PLAYGROUND_PORT", 8080)

	log.Printf("starting server on port %d\n", port)

	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func getEnvInt(name string, fallback int) int {
	var value int
	var err error

	valueStr := os.Getenv(name)
	if valueStr == "" {
		value = fallback
	} else {
		value, err = strconv.Atoi(valueStr)
		if err != nil {
			log.Panic(err)
		}
	}

	return value
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s%s\n", r.RemoteAddr, r.Method, r.Host, r.URL)

	x := rand.Int() % 1024
	fibX := fib(x)
	data := map[int]*big.Int{x: fibX}

	json.NewEncoder(w).Encode(data)
}

func fib(n int) *big.Int {
	a := big.NewInt(1)
	b := big.NewInt(0)
	tmp := big.NewInt(0)

	for i := 0; i < n; i++ {
		tmp = tmp.Add(a, b)
		a, b = tmp, a
	}

	return b
}
