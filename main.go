package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	ticker := time.NewTicker(15 * time.Second)

	for {
		select {
		case <-ticker.C:
			makePostRequest()
		}
	}
}

func makePostRequest() {
	rand.Seed(time.Now().UnixNano())
	data := map[string]interface{}{
		"title":  "Albert Einstein Quote",
		"body":   "Two things are infinite: the universe and human stupidity; and I'm not sure about the universe",
		"userId": rand.Intn(100) + 1,
		"wind":   rand.Intn(100) + 1,
		"water":  rand.Intn(100) + 1,
	}

	requestJSON, err := json.Marshal(data)
	fmt.Println()
	fmt.Println("Request : ")
	fmt.Println(string(requestJSON))
	fmt.Println()
	if err != nil {
		fmt.Println("Failed to marshal request data: ", err)
		return
	}

	res, err := http.Post("https://jsonplaceholder.typicode.com/posts", "Application/json", bytes.NewBuffer(requestJSON))
	defer res.Body.Close()

	bodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Failed to read response body: ", err)
		return
	}
	err = json.Unmarshal(bodyByte, &data)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response : ")
	fmt.Println(string(bodyByte))
	fmt.Println()
	if data["water"].(float64) <= 5 {
		fmt.Println("status water : aman")
	} else if data["water"].(float64) > 5 && data["water"].(float64) <= 8 {
		fmt.Println("status water : siaga")
	} else {
		fmt.Println("status water : bahaya")
	}

	if data["wind"].(float64) <= 6 {
		fmt.Println("status wind : aman")
	} else if data["wind"].(float64) > 6 && data["wind"].(float64) <= 15 {
		fmt.Println("status wind : siaga")
	} else {
		fmt.Println("status wind : bahaya")
	}
}

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	ticker := time.NewTicker(15 * time.Second)

// 	for {
// 		select {
// 		case <-ticker.C:
// 			makePostRequest()
// 		}
// 	}
// }

// func makePostRequest() {
// 	url := "https://example.com/api/endpoint"
// 	jsonStr := []byte(`{"key1":"value1","key2":"value2"}`)

// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("POST request sent successfully")
// }
