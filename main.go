package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/cancel", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Main handler job has started... You have 3 seconds to " +
			"cancel the request to preserve downstream request from being executed.")

		time.Sleep(time.Second * 3)
		err := handleRequestWithCancelledContext(w, r)

		if err != nil {
			fmt.Printf("Received an error from the downstream job: %s\n", err)
		}
		fmt.Println("Main handler job has finished...")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handleRequestWithCancelledContext(w http.ResponseWriter, r *http.Request) error {
	var pow = []int{1, 2, 4, 8}

	select {
	case <-r.Context().Done():
		fmt.Println("Channel was closed")
		return r.Context().Err()

	default:
		fmt.Println("Downstream job has started...")

		for i, v := range pow {
			time.Sleep(time.Second * 1)
			fmt.Printf("2**%d = %d\n", i, v)
		}

		fmt.Fprintf(w, "OK")

		fmt.Println("Downstream job has finished...")
		return nil
	}
}
