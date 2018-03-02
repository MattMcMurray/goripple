package ripple

type transaction struct {
	Account         string `json:"Account"`
	Amount          int    `json:"Amount"`
	Destination     string `json:"Destination"`
	TransactionType string `json:"TransactionType"`
}

type param struct {
	Offline bool        `json:"offline,omitempty"`
	Secret  string      `json:"secret,omitempty"`
	TxJSON  transaction `json:"tx_json,omitempty"`
	TxBlob  string      `json:"tx_blob,omitempty"`
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
		EngineResultCode *int   `json:"engine_result_code"`
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
