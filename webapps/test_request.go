package webapps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
)

func SendRandomPoll(url string, count int64) {
	ans := [4]string{"Sangat Bermanfaat", "Bermanfaat", "Cukup Bermanfaat", "Kurang Bermanfaat"}
	counter := [4]int{}
	for i := 0; i <= int(count); i++ {
		contentType := "application/json"
		randomInt := rand.IntN(4-0) + 0
		answer := ans[randomInt]
		counter[randomInt] = counter[randomInt] + 1

		fmt.Println("Respons ", answer)

		jsonBytes, err := json.Marshal(map[string]string{"jawaban": answer})
		if err != nil {
			log.Fatalf("Error marshaling map to JSON: %v", err)
		}
		jsonString := string(jsonBytes)

		data := []byte(jsonString)
		client := &http.Client{}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			fmt.Println(err)
			return
		}

		req.Header.Add("Content-Type", contentType)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
	}

	for i, v := range ans {
		fmt.Println(v, " : ", counter[i])
	}

}
