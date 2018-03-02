package ripple

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

	return queryServer(rippledURL, p)
}

// SignXaction sign a payment transaction using a rippled server.
func SignXaction(secret string, account string, dest string, amount int) (*Response, error) {
	r := request{Method: "sign"}
	t := transaction{
		Account:         account,
		Amount:          amount,
		Destination:     dest,
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

	return queryServer(rippledURL, payload)
}

// SubmitSignedXaction submits a signed transaction blob for entry into the ledger.
func SubmitSignedXaction(txBlob string) (*Response, error) {
	r := request{Method: "submit"}
	p := param{
		TxBlob: txBlob,
	}
	r.Params = make([]param, 0)
	r.Params = append(r.Params, p)

	payload, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return queryServer(rippledURL, payload)
}

// OpenPaymentChannel opens a payment channel from account -> dest.
//
// pubkey is the HEX representation of the sender's public key.
func OpenPaymentChannel(secret string, account string, amt int, dest string,
	pubkey string) (*Response, error) {
	r := request{Method: "submit"}
	t := transaction{
		Account:         account,
		Amount:          amt,
		Destination:     dest,
		TransactionType: "PaymentChannelCreate",
	}
	p := param{
		Secret: secret,
		TxJSON: t,
	}

	r.Params = make([]param, 0)
	r.Params = append(r.Params, p)

	payload, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return queryServer(rippledURL, payload)
}

func queryServer(url string, payload []byte) (*Response, error) {
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
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var resp *Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
