package utils

import (
	"encoding/json"
	"github.com/crc/zlzk/chaincode/define"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)

var Logger = shim.NewLogger("chainCode")

func ResponseFailed(format string, a ...interface{}) ([]byte, error) {
	err := errors.Errorf(format, a...)
	Logger.Error("Response failed: ", err.Error())
	return nil, err
}

func GetData(stub shim.ChaincodeStubInterface, getKey string) ([]byte, error) {
	Logger.Debugf("getKey: %v\n", getKey)
	res, err := stub.GetState(getKey)
	if err != nil {
		return ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func PutData(stub shim.ChaincodeStubInterface, data interface{}, putKey string) error {
	putData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	Logger.Debugf("putKey: %v\n", putKey)
	if err := stub.PutState(putKey, putData); err != nil {
		return err
	}
	return nil
}

func GetHistoryData(stub shim.ChaincodeStubInterface, getKey string) ([]byte, error) {
	Logger.Info("getKey: %v\n", getKey)
	historyIter, err := stub.GetHistoryForKey(getKey)

	if err != nil {
		return ResponseFailed("Form query failed to GetState err: %v", err)
	}
	var dat1 []interface{}
	for i := 0; historyIter.HasNext(); i++ {

		modification, err := historyIter.Next()
		if err != nil {
			return ResponseFailed("Form query failed to GetState err: %v", err)
		}

		data := &define.ThirdPartUploadRecord{}
		err = json.Unmarshal(modification.Value, data)
		if err != nil {
			return ResponseFailed("format history data err: %v", err)
		}

		dat1 = append(dat1, data)
	}

	return json.Marshal(dat1)
}
