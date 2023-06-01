package randomstring

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/glopezep/arithmetic-calculator/internal/infrastructure/config"
	"github.com/stretchr/testify/require"
)

func TestRandomStringGenerate(t *testing.T) {
	conf := config.NewConfig()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`
			{
				"jsonrpc": "2.0",
				"result": {
					"random": {
								"data": [
										"grvhglvahj", "hjrmosjwed", "nivjyqptyy", "lhogeshsmi",
										"syilbgsytb", "birvcmgdrz", "wgclyynpcq", "eujwnhgonh"
								],
								"completionTime": "2011-10-10 13:19:12Z"
						},
						"bitsUsed": 376,
						"bitsLeft": 199624,
						"requestsLeft": 9999,
						"advisoryDelay": 0
				},
				"id": 42
			}
		`))
	}))
	defer ts.Close()

	conf.RandomServiceURL = ts.URL
	randomStringService := NewRandomStringService(conf)

	str, err := randomStringService.Generate()

	require.NoError(t, err)
	require.NotEmpty(t, str)
}

func TestRandomStringGenerateFail(t *testing.T) {
	conf := config.NewConfig()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
		w.Write([]byte{})
	}))
	defer ts.Close()

	conf.RandomServiceURL = ts.URL
	randomStringService := NewRandomStringService(conf)

	_, err := randomStringService.Generate()

	require.Error(t, err)
}
