package bif

import "github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/request"

type CreateAccountRequest struct {
	request.BIFCreateAccountRequest
	Nonce int64 `json:"nonce"`
}

type RadioTransactionRequest struct {
	request.BIFRadioTransactionRequest
	Nonce int64 `json:"nonce"`
}
