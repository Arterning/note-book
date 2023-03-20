package main

import (
	"fmt"
	"github.com/crc/zlzk/chaincode/define"
	"github.com/crc/zlzk/chaincode/handler"
	"github.com/crc/zlzk/chaincode/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type handlerFunc func(stub shim.ChaincodeStubInterface, args []string) ([]byte, error)

var funcHandler = map[string]handlerFunc{
	define.Query:                    handler.Query,
	define.QueryUploadRecord:        handler.QueryUploadRecord,
	define.QueryUploadRecords:       handler.QueryUploadRecords,
	define.QueryFarm:                handler.QueryFarm,
	define.QueryLand:                handler.QueryLand,
	define.QueryPlant:               handler.QueryPlant,
	define.QueryFarmRecord:          handler.QueryFarmRecord,
	define.QueryHarvest:             handler.QueryHarvest,
	define.QueryDryInfo:             handler.QueryDryInfo,
	define.QueryWarehouseInfo:       handler.QueryWarehouseInfo,
	define.QueryProcessInfo:         handler.QueryProcessInfo,
	define.QueryFeedSectionCrop:     handler.QueryFeedSectionCrop,
	define.QueryDict:                handler.QueryDict,
	define.QueryPesticideFertilizer: handler.QueryPesticideFertilizer,
	define.PutUploadRecord:          handler.PutUploadRecord,
	define.PutFarm:                  handler.PutFarm,
	define.PutLand:                  handler.PutLand,
	define.PutPlant:                 handler.PutPlant,
	define.PutFarmRecord:            handler.PutFarmRecord,
	define.PutHarvest:               handler.PutHarvest,
	define.PutDryInfo:               handler.PutDryInfo,
	define.PutWarehouseInfo:         handler.PutWarehouseInfo,
	define.PutProcessInfo:           handler.PutProcessInfo,
	define.PutFeedSectionCrop:       handler.PutFeedSectionCrop,
	define.PutDict:                  handler.PutDict,
	define.PutPesticideFertilizer:   handler.PutPesticideFertilizer,
}

type ChainCode struct {
}

func (t *ChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	utils.Logger.Debug("Init ChainCode...")
	return shim.Success([]byte("SUCCESS"))
}

func (t *ChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	utils.Logger.Debugf("Invoke function=%v,args=%v\n", function, args)

	if len(args) < 1 || len(args[0]) == 0 {
		err := fmt.Sprintf("the invoke args length < 1 or arg[0] is empty")
		utils.Logger.Error(err)
		return shim.Error(err)
	}

	currentFunc, ok := funcHandler[function]
	if !ok {
		err := fmt.Sprintf("the function name not exist")
		utils.Logger.Error(err)
		return shim.Error(err)
	}

	payload, err := currentFunc(stub, args)
	if err != nil {
		utils.Logger.Error(err)
		return shim.Error(err.Error())
	}
	return shim.Success(payload)
}

func main() {
	err := shim.Start(new(ChainCode))
	if err != nil {
		utils.Logger.Errorf("Error starting chainCode: %s", err)
	}
}
