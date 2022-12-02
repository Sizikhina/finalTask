package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const url = "https://kuvaev-ituniversity.vps.elewise.com/tasks"

func GetSolutionData(taskName string) [10][]json.RawMessage {
	resp, err := http.Get(fmt.Sprintf("%s/%s", url, taskName))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var result [10][]json.RawMessage

	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	return result
}

type SolutionCheck struct {
	UserName string `json:"user_name"`
	Task     string `json:"task"`
	Results  struct {
		Payload interface{} `json:"payload"`
		Results interface{} `json:"results"`
	} `json:"results"`
}

func CheckSolution(check SolutionCheck) json.RawMessage {
	b, err := json.Marshal(check)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(fmt.Sprintf("%s/solution", url), "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("status code, %d", resp.StatusCode))
	}
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return b
}
