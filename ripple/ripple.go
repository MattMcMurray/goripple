package ripple

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var RippledURL = ""

// Ping Pings the ripple server
func Ping() (*Response, error) {
	r := request{
		Method: "ping",
		Params: nil,
	}

	return queryServer(RippledURL, &r)
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
		TxJSON:  &t,
	}
	r.Params = make([]param, 0)
	r.Params = append(r.Params, p)

	return queryServer(RippledURL, &r)
}

// SubmitSignedXaction submits a signed transaction blob for entry into the ledger.
func SubmitSignedXaction(txBlob string) (*Response, error) {
	r := request{Method: "submit"}
	p := param{
		TxBlob: txBlob,
	}
	r.Params = make([]param, 0)
	r.Params = append(r.Params, p)

	return queryServer(RippledURL, &r)
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
		SettleDelay:     86400,
		PublicKey:       pubkey,
	}
	p := param{
		Secret:     secret,
		TxJSON:     &t,
		FeeMultMax: 1000,
	}

	r.Params = make([]param, 0)
	r.Params = append(r.Params, p)

	return queryServer(RippledURL, &r)
}

// GetChannels get the ID and other important information about a channel open between two accounts.
//
// Since there's no great way to specify which channel you want, this function takes an
// index and returns that transaction from the returned list.
func GetChannels(account string, dest string) (*Response, error) {
	r := request{Method: "account_channels"}
	p := param{
		Account:     account,
		DestAccount: dest,
	}
	r.Params = make([]param, 0)
	r.Params = append(r.Params, p)

	return queryServer(RippledURL, &r)
}

func ChannelAuthorize(channelID string, secret string, amt int) (*Response, error) {
	r := request{Method: "channel_authorize"}
	p := param{
		ChannelID: channelID,
		Secret:    secret,
		Amount:    amt,
	}
	r.Params = make([]param, 0)
	r.Params = append(r.Params, p)

	return queryServer(RippledURL, &r)
}

func queryServer(url string, r *request) (*Response, error) {
	errLog := log.New(os.Stderr, " [HTTP] ", log.Ldate|log.Ltime|log.Lshortfile)
	if url == "" {
		return nil, errors.New("URL is empty")
	}
	payload, err := json.Marshal(r)
	if err != nil {
		errLog.Println("Error marshalling request")
		errLog.Println(r)
		return nil, err
	}

	p := bytes.NewBuffer(payload)

	req, err := http.NewRequest("POST", url, p)
	if err != nil {
		errLog.Println("Error creating HTTP request")
		errLog.Println(p.String())
		errLog.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		errLog.Println("Error in response to HTTP POST")
		errLog.Println(req)
		errLog.Println(err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		errLog.Println("Error reading response body")
		errLog.Println(res.Body)
		return nil, err
	}

	var resp *Response
	err = json.Unmarshal(body, &resp)
	if err != nil {
		errLog.Println("Error unmarshalling JSON response body")
		errLog.Println(string(body))
		return nil, errors.New(string(body))
	}

	return resp, nil
}
