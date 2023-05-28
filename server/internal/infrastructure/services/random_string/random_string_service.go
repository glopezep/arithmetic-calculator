package randomstring

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RandomStringService interface {
	Generate() (string, error)
}

type randomStringService struct{}

type request struct {
	JsonRPC string         `json:"jsonrpc"`
	Method  string         `json:"method"`
	Params  map[string]any `json:"params"`
	ID      int64          `json:"id"`
}

type random struct {
	Data           []string `json:"data"`
	CompletionTime string   `json:"completionTime"`
}

type result struct {
	Random        random
	BitsUsed      int64 `json:"bitsUsed"`
	BitsLeft      int64 `json:"bitsLeft"`
	RequestsLeft  int64 `json:"requestsLeft"`
	AdvisoryDelay int64 `json:"advisoryDelay"`
}

type response struct {
	JsonRPC string `json:"jsonrpc"`
	Result  result `json:"result"`
	ID      int64  `json:"id"`
}

func (s *randomStringService) Generate() (string, error) {

	req := request{
		JsonRPC: "2.0",
		Method:  "generateStrings",
		Params: map[string]any{
			"apiKey":      "ffcfa8ec-f661-44cf-86fc-1f08925f5880",
			"n":           2,
			"length":      10,
			"characters":  "abcdefghijklmnopqrstuvwxyz",
			"replacement": true,
		},
		ID: 42,
	}

	postBody, _ := json.Marshal(req)

	resp, err := http.Post(
		"https://api.random.org/json-rpc/4/invoke",
		"application/json",
		bytes.NewBuffer(postBody),
	)
	if err != nil {
		return "", nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}

	var res response

	json.Unmarshal([]byte(body), &res)

	return res.Result.Random.Data[0], nil
}

func NewRandomStringService() RandomStringService {
	return &randomStringService{}
}
