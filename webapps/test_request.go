package webapps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
)

var ans = [4]string{"Sangat Bermanfaat", "Cukup Bermanfaat", "Bermanfaat", "Kurang Bermanfaat"}

func SendRandomPoll(url string, count int64) {
	counter := [4]int{}
	for i := 0; i <= int(count); i++ {
		randomInt := rand.IntN(4-0) + 0
		answer := ans[randomInt]
		counter[randomInt] = counter[randomInt] + 1

		fmt.Println("Send Respons ", answer)

		jsonBytes, err := json.Marshal(map[string]string{"jawaban": answer})
		if err != nil {
			log.Fatalf("Error marshaling map to JSON: %v", err)
		}
		jsonString := string(jsonBytes)

		response := SendPoll(url, jsonString)
		fmt.Println(string(response))
	}

	for i, v := range ans {
		fmt.Println(v, " : ", counter[i])
	}

}

func SendPollShouldBalanced(url string) {
	counter := [4]int{}
	specificInt := 0
	stop := false
	for !stop {
		indexInt := specificInt
		answer := ans[indexInt]
		counter[indexInt] = counter[indexInt] + 1
		fmt.Println("Send Respons ", answer)

		jsonBytes, err := json.Marshal(map[string]string{"jawaban": answer})
		if err != nil {
			log.Fatalf("Error marshaling map to JSON: %v", err)
		}
		jsonString := string(jsonBytes)

		response := SendPoll(url, jsonString)

		var bodyMap = map[string]any{}

		_ = json.Unmarshal(response, &bodyMap)
		fmt.Println(string(response))
		specificInt, stop = findMinValue(bodyMap)
		fmt.Println("specificInt:", specificInt)
	}

	for i, v := range ans {
		fmt.Println(v, " : ", counter[i])
	}
}

func SendPoll(url string, dataJsonString string) []byte {
	contentType := "application/json"
	data := []byte(dataJsonString)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return []byte("Error")
	}

	req.Header.Add("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte("Error")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return []byte("Error")
	}

	return body
}

func findMinValue(param map[string]any) (int, bool) {
	var temp, val float64 = 100, 0
	tempStop := true
	var index int
	for key, value := range param {
		if key == "data" {
			for k, v := range value.(map[string]any) {
				switch c := v.(type) {
				case float64:
					val = float64(c)
				}
				if temp > val {
					temp = val
					num, _ := strconv.Atoi(k)
					index = num
				}

				if val != 25.00 {
					tempStop = false
				}
			}
		}
	}

	return index - 1, tempStop
}
