package main

import (
	//"encoding/json"

	"time"
	"math/rand"
	"github.com/tj/go-gracefully"

	"encoding/json"

	"fmt"
	"net/http"
	"bytes"
)

type Ivent struct {
	Lat float32 `json: lat`
	Lon float32 `json: lon`
}

var k=0
func client(max_iv int) {
	for i:=0; i<max_iv; i++ {

		locJson, _ := json.Marshal(Ivent{Lat: rand.Float32(), Lon: rand.Float32()})
		req, err := http.NewRequest("POST", "http://localhost:8000", bytes.NewBuffer(locJson))
		////req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		_ = resp
		_ = err
		k++
		fmt.Println(locJson, k)
	}
}



func main() {

		go client(10000)
		gracefully.Timeout = 5 * time.Second
		gracefully.Shutdown()

}