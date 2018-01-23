package main

import (
	//"time"
	"net/http"
	"fmt"
	"flag"
	"encoding/json"

	"log"
	//"math/rand"
	//"io"
	"io"
	"github.com/tj/go-gracefully"
	"time"
)


type Ivent struct {
	Lat float32 `json: lat`
	Lon float32 `json: lon`
}



var WorkQueue = make(chan Ivent, 10)

func Collector(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var content = &Ivent{}
	err := json.NewDecoder(io.LimitReader(r.Body, 10000)).Decode(&content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
		work := Ivent{content.Lat, content.Lon}
		WorkQueue <- work
		fmt.Println(len(WorkQueue))

		}


func NewWorker(id int, workerQueue chan chan Ivent) Worker {
	worker := Worker{
		ID:          id,
		Work:        make(chan Ivent),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

type Worker struct {
	ID          int
	Work        chan Ivent
	WorkerQueue chan chan Ivent
	QuitChan    chan bool
}


func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:

				log.Printf("%f \n",work)
				//time.Sleep(time.Millisecond*10)

			case <-w.QuitChan:
				fmt.Printf("worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}


func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

var WorkerQueue chan chan Ivent

func StartDispatcher(nworkers int) {
	WorkerQueue = make(chan chan Ivent, nworkers)

	for i := 0; i<nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}
	go func() {
		for {
			select {
			case work := <-WorkQueue:
				go func() {
					worker := <-WorkerQueue
					worker <- work
				}()
			}
		}
	}()
}


var (
	NWorkers = flag.Int("n", 5, "The number of workers to start")
	HTTPAddr = flag.String("http", ":8000", "Address to listen for HTTP requests on")
)



func server() {
	StartDispatcher(*NWorkers)
	http.HandleFunc("/", Collector)
	fmt.Println("HTTP server listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
	fmt.Println(err.Error())
	}
}

func main() {

	go server()
	gracefully.Timeout = 10 * time.Second
	gracefully.Shutdown()


}