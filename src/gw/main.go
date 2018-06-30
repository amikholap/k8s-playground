package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := getEnvInt("PLAYGROUND_PORT", 8080)

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
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}
	w.Write(data)
}
