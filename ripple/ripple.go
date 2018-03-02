package ripple

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type request struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params,omitempty"`
}

// Response a response from the rippled server
type Response struct {
	Result *struct {
		Role             string `json:"role"`
		Status           string `json:"status"`
		EngineResult     string `json:"engine_result"`
		EngineResultCode string `json:"engine_result_code"`
		EngineResultMsg  string `json:"engine_result_message"`
		TxBlob           string `json:"tx_blob"`
		Tx               *struct {
			Account         string
			Amount          string
			Destination     string
			Fee             string
			Flags           int
			Sequence        int
			SigningPubKey   string
			TransactionType string
			TxnSignature    string
			Hash            string `json:"hash"`
		} `json:"tx_json"`
	}
}

// TODO: move to config (?)
var rippledURL = "http://0.0.0.0:5005" // 5005 is the admin port

// Ping Pings the ripple server
func Ping() (*Response, error) {
	r := request{"ping", nil}

	p, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	body, err := makeHTTPRequest(rippledURL, p)

	var pr Response
	err = json.Unmarshal(body, &pr)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}

func makeHTTPRequest(url string, payload []byte) ([]byte, error) {
	p := bytes.NewBuffer(payload)

	req, err := http.NewRequest("POST", url, p)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
