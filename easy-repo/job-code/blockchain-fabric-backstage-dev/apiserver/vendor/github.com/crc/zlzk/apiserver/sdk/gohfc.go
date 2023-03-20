package sdk

import (
	"fmt"
	"github.com/crc/gohfc"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
	"time"
)

func InitSDKs(path, name string) error {
	viper.SetEnvPrefix("core")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	configFilePath := filepath.Join(path, name+".yaml")
	viper.SetConfigFile(configFilePath)

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Fatal error when initializing %s config : %v\n", "SDK", err)
	}
	err = gohfc.InitSDK(configFilePath)
	if err != nil {
		return err
	}

	if err := gohfc.GetHandler().HandleTxStatus(""); err != nil {
		return err
	}

	//utils.Logger.Debugf("config file is %s", configFilePath)
	return nil
}

func SetLogLevel(level, name string) error {
	return gohfc.SetLogLevel(level, name)
}

func Query(key, opType string) ([]byte, error) {
	var info []string
	info = append(info, opType, key)
	Handler := gohfc.GetHandler()
	res, err := Handler.Query(info, "", "")
	if err != nil {
		return nil, err
	} else {
		if res[0].Error != nil {
			return nil, err
		}
	}
	return res[0].Response.Response.Payload, nil
}

func GetTransactionByID(txId, myChannel string) (string, error) {
	txData, err := gohfc.GetHandler().GetTransactionById(txId, myChannel)
	if err != nil {
		return "", fmt.Errorf("get transaction by id error :%v", err)
	}

	if txData.ValidationCode != 0 {
		return "", fmt.Errorf("tx is invalidation")
	}
	return txData.ChaincodeSpec.Input.Args[1], nil
}

func GetBlockInfoByTxID(txid string) (uint64, string, error) {
	chainId := viper.GetString("channel.channelId")
	args := []string{"GetBlockByTxID", chainId, txid}
	resps, err := gohfc.GetHandler().QueryByQscc(args, chainId)
	if err != nil {
		return 0, "", fmt.Errorf("can not get installed chaincodes :%s", err.Error())
	} else if len(resps) == 0 {
		return 0, "", fmt.Errorf("GetBlockByTxID empty response from peer")
	}
	if resps[0].Error != nil {
		return 0, "", resps[0].Error
	}
	data := resps[0].Response.Response.Payload
	var block = new(common.Block)
	err = proto.Unmarshal(data, block)
	if err != nil {
		return 0, "", fmt.Errorf("GetBlockByTxID Unmarshal from payload failed: %s", err.Error())
	}
	blockNumber := block.Header.Number
	txData, err := gohfc.GetHandler().GetTransactionById(txid, chainId)
	if err != nil {
		return 0, "", fmt.Errorf("get transaction by id error :%v", err)
	}

	if txData.ValidationCode != 0 {
		return 0, "", fmt.Errorf("tx is invalidation")
	}
	channelTimestamp := txData.ChannelHeader.Timestamp
	channelDate := time.Unix(channelTimestamp.GetSeconds(), int64(channelTimestamp.GetNanos())).Format("2006-01-02 15:04:05")

	return blockNumber,channelDate, nil
}
