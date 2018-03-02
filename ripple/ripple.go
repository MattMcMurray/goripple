package ripple

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type transaction struct {
	Account         string `json:"Account"`
	Amount          int    `json:"Amount"`
	Destination     string `json:"Destination"`
	TransactionType string `json:"TransactionType"`
}

type param struct {
	Offline bool        `json:"offline"`
	Secret  string      `json:"secret"`
	TxJSON  transaction `json:"tx_json"`
}

type request struct {
	Method string  `json:"method"`
	Params []param `json:"params"`
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
	r := request{
		Method: "ping",
		Params: nil,
	}

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

// SignXaction sign a payment transaction using a rippled server
func SignXaction(secret string, account string, destination string, amount int) (*Response, error) {
	r := request{Method: "sign"}
	t := transaction{
		Account:         account,
		Amount:          amount,
		Destination:     destination,
		TransactionType: "Payment",
	}
	p := param{
		Offline: false,
		Secret:  secret,
		TxJSON:  t,
	}
	r.Params = make([]param, 0)
	r.Params = append(r.Params, p)

	payload, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	body, err := makeHTTPRequest(rippledURL, payload)
	if err != nil {
		return nil, err
	}

	var res *Response
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
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
