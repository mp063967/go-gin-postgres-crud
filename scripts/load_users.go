package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	apiURL    = "http://localhost:8082/users"
	total     = 50
	batchSize = 5
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	var wg sync.WaitGroup
	client := &http.Client{}

	batches := total / batchSize

	for b := 0; b < batches; b++ {
		wg.Add(1)
		fmt.Printf("== Start Batch: %d ==\n", b+1)
		go func(batch int) {
			defer wg.Done()
			fmt.Printf("== Go Batch: %d ==\n", batch)
			for i := 0; i < batchSize; i++ {
				id := batch*batchSize + i + 1

				user := User{
					Name:  fmt.Sprintf("User-%d", id),
					Email: fmt.Sprintf("user%d@example.com", id),
				}

				body, _ := json.Marshal(user)

				req, err := http.NewRequest(
					http.MethodPost,
					apiURL,
					bytes.NewBuffer(body),
				)
				if err != nil {
					fmt.Println("request error:", err)
					continue
				}

				req.Header.Set("Content-Type", "application/json")

				resp, err := client.Do(req)
				if err != nil {
					fmt.Println("api error:", err)
					continue
				}

				resp.Body.Close()
				// fmt.Printf("Created user %d (batch %d) at time %d\n", id, batch+1, time.Now().UnixMilli())
				// Lets print real date time for clarity
				fmt.Printf("Created user %d (batch %d) at time %s\n", id, batch+1, time.Now().Format(time.RFC3339))
			}
		}(b)
	}

	wg.Wait()
	fmt.Println("✅ All users created")
}
