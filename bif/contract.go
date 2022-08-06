package bif

import (
	"encoding/hex"
	"encoding/json"
	"go-zero-demo/internal/svc"

	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/utils/http"

	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/proto"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/utils/hash"

	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/common"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/exception"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/blockchain"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/encryption/key"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/request"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/types/response"
	protobuf "github.com/golang/protobuf/proto"
)

func CreateAccount(svCtx *svc.ServiceContext, r CreateAccountRequest) response.BIFCreateAccountResponse {
	if r.SenderAddress == "" || r.DestAddress == "" {
		return response.BIFCreateAccountResponse{
			BIFBaseResponse: exception.REQUEST_NULL_ERROR,
		}
	}
	if !key.IsAddressValid(r.SenderAddress) {
		return response.BIFCreateAccountResponse{
			BIFBaseResponse: exception.INVALID_ADDRESS_ERROR,
		}
	}
	if !key.IsAddressValid(r.DestAddress) {
		return response.BIFCreateAccountResponse{
			BIFBaseResponse: exception.INVALID_DESTADDRESS_ERROR,
		}
	}
	if r.InitBalance == 0 || r.InitBalance <= common.INIT_ZERO {
		return response.BIFCreateAccountResponse{
			BIFBaseResponse: exception.INVALID_INITBALANCE_ERROR,
		}
	}

	if r.PrivateKey == "" {
		return response.BIFCreateAccountResponse{
			BIFBaseResponse: exception.PRIVATEKEY_NULL_ERROR,
		}
	}

	if r.FeeLimit == common.INIT_ZERO {
		r.FeeLimit = common.FEE_LIMIT
	}
	if r.GasPrice == common.INIT_ZERO {
		r.GasPrice = common.GAS_PRICE
	}

	// 广播交易
	bifAccountActivateOperation := request.BIFAccountActivateOperation{
		BIFBaseOperation: request.BIFBaseOperation{
			OperationType: common.ACCOUNT_ACTIVATE,
		},
		DestAddress: r.DestAddress,
		InitBalance: r.InitBalance,
	}

	bifRadioTransactionRequest := RadioTransactionRequest{
		BIFRadioTransactionRequest: request.BIFRadioTransactionRequest{
			SenderAddress:    r.SenderAddress,
			FeeLimit:         r.FeeLimit,
			GasPrice:         r.GasPrice,
			Operation:        bifAccountActivateOperation,
			CeilLedgerSeq:    r.CeilLedgerSeq,
			Remarks:          r.Remarks,
			SenderPrivateKey: r.PrivateKey,
		},
		Nonce: r.Nonce,
	}

	radioTransactionResponse := RadioTransaction(svCtx, bifRadioTransactionRequest)
	if radioTransactionResponse.ErrorCode != common.SUCCESS {
		return response.BIFCreateAccountResponse{
			BIFBaseResponse: radioTransactionResponse.BIFBaseResponse,
		}
	}

	return response.BIFCreateAccountResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFAccountCreateAccountResult{
			Hash: radioTransactionResponse.Result.Hash,
		},
	}
}

// RadioTransaction 广播交易
func RadioTransaction(svCtx *svc.ServiceContext, r RadioTransactionRequest) response.BIFRadioTransactionResponse {
	// 一、构建操作、序列化交易
	// 初始化请求参数 BIFTransactionSerializeRequest
	serializeRequest := request.BIFTransactionSerializeRequest{
		SourceAddress: r.SenderAddress,
		Nonce:         r.Nonce,
		GasPrice:      r.GasPrice,
		FeeLimit:      r.FeeLimit,
		Operation:     r.Operation,
		CeilLedgerSeq: r.CeilLedgerSeq,
		Metadata:      r.Remarks,
	}
	// BIFTransactionSerializeResponse
	serializeResponse := Serializable(svCtx, serializeRequest)
	if serializeResponse.ErrorCode != common.SUCCESS {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: serializeResponse.BIFBaseResponse,
		}
	}
	transactionBlob := serializeResponse.Result.TransactionBlob
	blob, err := hex.DecodeString(transactionBlob)
	if err != nil {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	// 二、签名
	signData, err := key.Sign([]byte(r.SenderPrivateKey), blob)
	if err != nil {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	publicKey, err := key.GetEncPublicKey([]byte(r.SenderPrivateKey))
	if err != nil {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	// 三、提交交易 BIFTransactionSubmitRequest
	submitRequest := request.BIFTransactionSubmitRequest{
		Serialization: transactionBlob,
		SignData:      hex.EncodeToString(signData),
		PublicKey:     publicKey,
	}

	// 调用bifSubmit接口 BIFTransactionSubmitResponse
	transactionSubmitResponse := Submit(svCtx, submitRequest)
	if transactionSubmitResponse.ErrorCode != common.SUCCESS {
		return response.BIFRadioTransactionResponse{
			BIFBaseResponse: transactionSubmitResponse.BIFBaseResponse,
		}
	}

	return response.BIFRadioTransactionResponse{
		BIFBaseResponse: transactionSubmitResponse.BIFBaseResponse,
		Result: response.BIFRadioTransactionResult{
			Hash: transactionSubmitResponse.Result.Hash,
		},
	}
}

// Submit 交易提交
func Submit(svCtx *svc.ServiceContext, r request.BIFTransactionSubmitRequest) response.BIFTransactionSubmitResponse {
	if svCtx.Config.Chain.Url == "" {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.URL_EMPTY_ERROR,
		}
	}
	if r.SignData == "" {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.SIGNDATA_NULL_ERROR,
		}
	}
	if r.PublicKey == "" {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.PUBLICKEY_NULL_ERROR,
		}
	}
	if r.Serialization == "" {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.INVALID_SERIALIZATION_ERROR,
		}
	}

	var transactionSubmit request.TransactionSubmit
	transactionSubmit.TransactionBlob = r.Serialization
	var signature request.Signature
	signature.SignData = r.SignData
	signature.PublicKey = r.PublicKey
	transactionSubmit.Signatures = append(transactionSubmit.Signatures, signature)

	var transactionSubmitRequest request.TransactionSubmitRequest
	transactionSubmitRequest.Items = append(transactionSubmitRequest.Items, transactionSubmit)
	transactionRequest, err := json.Marshal(transactionSubmitRequest)
	if err != nil {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	submitURL := common.TransactionSubmitURL(svCtx.Config.Chain.Url)
	dataByte, err := http.HttpPost(submitURL, transactionRequest)
	if err != nil {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.CONNECTNETWORK_ERROR,
		}
	}

	var res response.TransactionSubmitResponse
	err = json.Unmarshal(dataByte, &res)
	if err != nil {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}

	if res.Results[0].ErrorCode != common.SUCCESS {
		return response.BIFTransactionSubmitResponse{
			BIFBaseResponse: response.BIFBaseResponse{
				ErrorCode: res.Results[0].ErrorCode,
				ErrorDesc: res.Results[0].ErrorDesc,
			},
		}
	}

	return response.BIFTransactionSubmitResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFTransactionSubmitResult{
			Hash: res.Results[0].Hash,
		},
	}
}

// Serializable 交易序列化
func Serializable(svCtx *svc.ServiceContext, r request.BIFTransactionSerializeRequest) response.BIFTransactionSerializeResponse {

	if !key.IsAddressValid(r.SourceAddress) {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.INVALID_SOURCEADDRESS_ERROR,
		}
	}
	if r.Nonce <= 0 {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.INVALID_NONCE_ERROR,
		}
	}
	if r.CeilLedgerSeq < 0 {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.INVALID_CEILLEDGERSEQ_ERROR,
		}
	}
	if r.GasPrice < 0 {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.INVALID_GASPRICE_ERROR,
		}
	}
	if r.FeeLimit < 0 {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.INVALID_FEELIMIT_ERROR,
		}
	}

	if r.Operation == nil {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.OPERATIONS_EMPTY_ERROR,
		}
	}

	operationService := blockchain.GetOperationInstance(svCtx.Config.Chain.Url)
	operations, bifBaseResponse := operationService.GetOperations(r.Operation, r.SourceAddress)
	if bifBaseResponse.ErrorCode != common.SUCCESS {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: bifBaseResponse,
		}
	}
	if r.CeilLedgerSeq < 0 {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.INVALID_CEILLEDGERSEQ_ERROR,
		}
	}
	var seq int64 = 0
	if r.CeilLedgerSeq > 0 {
		blockService := blockchain.GetBlockInstance(svCtx.Config.Chain.Url)
		blockGetNumberResponse := blockService.GetBlockNumber()
		if blockGetNumberResponse.ErrorCode != common.SUCCESS {
			return response.BIFTransactionSerializeResponse{
				BIFBaseResponse: blockGetNumberResponse.BIFBaseResponse,
			}
		}

		seq = r.CeilLedgerSeq + blockGetNumberResponse.Result.Header.BlockNumber
	}
	transaction := proto.Transaction{
		SourceAddress: r.SourceAddress,
		Nonce:         r.Nonce,
		CeilLedgerSeq: seq,
		FeeLimit:      r.FeeLimit,
		GasPrice:      r.GasPrice,
		Metadata:      []byte(r.Metadata),
		Operations:    operations,
	}
	blobByte, err := protobuf.Marshal(&transaction)
	if err != nil {
		return response.BIFTransactionSerializeResponse{
			BIFBaseResponse: exception.SYSTEM_ERROR,
		}
	}
	blob := hex.EncodeToString(blobByte)

	return response.BIFTransactionSerializeResponse{
		BIFBaseResponse: exception.SUCCESS,
		Result: response.BIFTransactionSerializeResult{
			TransactionBlob: blob,
			Hash:            string(hash.GenerateHashHex(blobByte, hash.SHA256)),
		},
	}
}
