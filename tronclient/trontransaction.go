package tronclient

type TronTransactionRawData struct {
	Contract      []*TronTransactionInfo `json:"contract"`
	RefBlockBytes string                 `json:"ref_block_bytes"`
	RefBlockHash  string                 `json:"ref_block_hash"`
	Expiration    int64                  `json:"expiration"`
	FeeLimit      int                    `json:"fee_limit"`
	Timestamp     int64                  `json:"timestamp"`
}

type TronTransactionInfo struct {
	Parameter *TransactionInfoParam `json:"parameter"`
	Type      string                `json:"type"`
}

type TransactionInfoParam struct {
	Value   *TransactionInfoParamValue `json:"value"`
	TypeUrl string                     `json:"type_url"`
}

type TransactionInfoParamValue struct {
	Data            string `json:"data"`
	OwnerAddress    string `json:"owner_address"`
	AssetName       string `json:"asset_name,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"`
	Amount          int64  `json:"amount,omitempty"`
	ToAddress       string `json:"to_address,omitempty"`
}

type TronTransaction struct {
	Ret []struct {
		ContractRet string `json:"contractRet"`
	} `json:"ret"`
	Signature  []string                `json:"signature"`
	TxID       string                  `json:"txID"`
	RawData    *TronTransactionRawData `json:"raw_data"`
	RawDataHex string                  `json:"raw_data_hex"`
}

type rawTransactionSignRequest struct {
	Transaction struct {
		TxID    string                  `json:"txID"`
		RawData *TronTransactionRawData `json:"raw_data"`
	} `json:"transaction"`
	PrivateKey string `json:"privateKey"`
}
