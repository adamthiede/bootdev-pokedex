package pokeapi

import (
    "fmt"
    "net/http"
    //"encoding/json"
    "io"
    "log"
    "time"
    )

    const baseURL="https://pokeapi.co/api/v2"


type Client struct {
    httpClient http.Client
}

func NewClient() Client {
    return Client {
	httpClient: http.Client {
	    Timeout: time.Minute,
	},
    }
}

func getJsonData() error {
    // example get https://pkg.go.dev/net/http#example-Get

    fmt.Println("command map")
    res, err := http.Get("https://pokeapi.co/api/v2/location-area/")

    if err != nil {
	log.Fatal(err)
    }
    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if res.StatusCode > 299 {
	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
    }
    if err != nil {
	log.Fatal(err)
    }
    strBody:=string(body)

    // parse json body

    fmt.Println("Got response:", strBody)

    return nil
}
