package ripple

type transaction struct {
	Account         string `json:",omitempty"`
	Amount          int    `json:",omitempty"`
	Destination     string `json:",omitempty"`
	TransactionType string `json:",omitempty"`
	SettleDelay     int    `json:",omitempty"`
	PublicKey       string `json:",omitempty"`
}

type param struct {
	Offline     bool         `json:"offline,omitempty"`
	Secret      string       `json:"secret,omitempty"`
	TxJSON      *transaction `json:"tx_json,omitempty"`
	TxBlob      string       `json:"tx_blob,omitempty"`
	FeeMultMax  int          `json:"fee_mult_max,omitempty"`
	Account     string       `json:"account,omitemtpy"`
	DestAccount string       `json:"destination_account,omitempty"`
}

type request struct {
	Method string  `json:"method"`
	Params []param `json:"params,omitempty"`
}

type channel struct {
}

// Response a response from the rippled server
type Response struct {
	Result *struct {
		Role             string `json:"role"`
		Status           string `json:"status"`
		EngineResult     string `json:"engine_result"`
		EngineResultCode *int   `json:"engine_result_code"` // pointer to int allows nil (instead of 0 as default value)
		EngineResultMsg  string `json:"engine_result_message"`
		TxBlob           string `json:"tx_blob"`
		ErrorMsg         string `json:"error_message"`
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
		Account  string `json:"account"`
		Channels *[]struct {
			Account            string `json:"account"`
			Amount             string `json:"amount"`
			Balance            string `json:"balance"`
			ChannelID          string `json:"channel_id"`
			DestinationAccount string `json:"destination_account"`
			PublicKey          string `json:"public_key"`
			PublicKeyHex       string `json:"public_key_hex"`
			SettleDelay        int    `json:"settle_delay"`
		} `json:"channels"`
	}
}
