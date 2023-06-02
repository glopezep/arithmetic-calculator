package randomstring

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/config"
)

type RandomStringService interface {
	Generate() (string, error)
}

type randomStringService struct {
	config *config.Config
}

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
			"apiKey":      s.config.RandomServiceApiKey,
			"n":           2,
			"length":      10,
			"characters":  "abcdefghijklmnopqrstuvwxyz",
			"replacement": true,
		},
		ID: 42,
	}

	postBody, _ := json.Marshal(req)

	resp, err := http.Post(
		s.config.RandomServiceURL,
		"application/json",
		bytes.NewBuffer(postBody),
	)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", errors.New("failed to generate random string")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var res response

	json.Unmarshal([]byte(body), &res)

	return res.Result.Random.Data[0], nil
}

func NewRandomStringService(conf *config.Config) RandomStringService {
	return &randomStringService{
		config: conf,
	}
}
