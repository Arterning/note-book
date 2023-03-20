package handler

import (
	"encoding/json"
	"github.com/crc/zlzk/chaincode/define"
	"github.com/crc/zlzk/chaincode/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func Query(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

// QueryUploadRecord 查询上传记录/**
func QueryUploadRecord(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, "ThirdPartUploadRecord"+args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

// QueryUploadRecords 查询历史记录/**
func QueryUploadRecords(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetHistoryData(stub, "ThirdPartUploadRecord"+args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QueryFarm(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, "Farm"+args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QueryLand(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, "Land"+args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QueryPlant(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, "Plant"+args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QueryFarmRecord(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, "FarmRecord"+args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QueryHarvest(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, "Harvest"+args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QueryDryInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, "Dry"+args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QueryWarehouseInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, "Warehouse"+args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QueryProcessInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, "Process"+args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QueryFeedSectionCrop(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, "FeedSectionCrop"+args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QueryDict(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QueryPesticideFertilizer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	res, err := utils.GetData(stub, args[0])
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

// PutUploadRecord 上传第三方记录/**
func PutUploadRecord(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	data := &define.ThirdPartUploadRecord{}
	if err := json.Unmarshal([]byte(args[0]), data); err != nil {
		return utils.ResponseFailed("land handler unmarshal data error %s", err)
	}
	putKey := "ThirdPartUploadRecord" + data.UploadRecordId
	if err := utils.PutData(stub, data, putKey); err != nil {
		return utils.ResponseFailed("land handler put data status error %s", err)
	}
	return nil, nil
}

func PutFarm(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	data := &define.Farm{}
	if err := json.Unmarshal([]byte(args[0]), data); err != nil {
		return utils.ResponseFailed("land handler unmarshal data error %s", err)
	}
	putKey := "Farm" + data.FarmId
	if err := utils.PutData(stub, data, putKey); err != nil {
		return utils.ResponseFailed("land handler put data status error %s", err)
	}
	return nil, nil
}

func PutLand(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	data := &define.Land{}
	if err := json.Unmarshal([]byte(args[0]), data); err != nil {
		return utils.ResponseFailed("land handler unmarshal data error %s", err)
	}
	putKey := "Land" + data.SectionId
	if err := utils.PutData(stub, data, putKey); err != nil {
		return utils.ResponseFailed("land handler put data status error %s", err)
	}
	return nil, nil
}

func PutPlant(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	data := &define.Plant{}
	if err := json.Unmarshal([]byte(args[0]), data); err != nil {
		return utils.ResponseFailed("plant handler unmarshal data error %s", err)
	}
	putKey := "Plant" + data.SectionCropsId
	if err := utils.PutData(stub, data, putKey); err != nil {
		return utils.ResponseFailed("plant handler put data status error %s", err)
	}
	return nil, nil
}

func PutFarmRecord(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	data := &define.FarmRecord{}
	if err := json.Unmarshal([]byte(args[0]), data); err != nil {
		return utils.ResponseFailed("farm record handler unmarshal data error %s", err)
	}
	putKey := "FarmRecord" + data.OperateId
	if err := utils.PutData(stub, data, putKey); err != nil {
		return utils.ResponseFailed("farm record handler put data status error %s", err)
	}
	return nil, nil
}

func PutHarvest(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	data := &define.Harvest{}
	if err := json.Unmarshal([]byte(args[0]), data); err != nil {
		return utils.ResponseFailed("harvest handler unmarshal data error %s", err)
	}
	data.TxId = stub.GetTxID()
	putKey := "Harvest" + data.HarvestId
	if err := utils.PutData(stub, data, putKey); err != nil {
		return utils.ResponseFailed("harvest handler put data status error %s", err)
	}
	return nil, nil
}

func PutDryInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	data := &define.DryInfo{}
	if err := json.Unmarshal([]byte(args[0]), data); err != nil {
		return utils.ResponseFailed("dry handler unmarshal data error %s", err)
	}
	putKey := "Dry" + data.DryId
	if err := utils.PutData(stub, data, putKey); err != nil {
		return utils.ResponseFailed("dry handler put data status error %s", err)
	}
	return nil, nil
}

func PutWarehouseInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	data := &define.WarehouseInfo{}
	if err := json.Unmarshal([]byte(args[0]), data); err != nil {
		return utils.ResponseFailed("warehouse handler unmarshal data error %s", err)
	}
	putKey := "Warehouse" + data.WarehouseId
	if err := utils.PutData(stub, data, putKey); err != nil {
		return utils.ResponseFailed("warehouse handler put data status error %s", err)
	}
	return nil, nil
}

func PutProcessInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	data := &define.ProcessInfo{}
	if err := json.Unmarshal([]byte(args[0]), data); err != nil {
		return utils.ResponseFailed("process handler unmarshal data error %s", err)
	}
	data.TxId = stub.GetTxID()
	putKey := "Process" + data.ProcessId
	if err := utils.PutData(stub, data, putKey); err != nil {
		return utils.ResponseFailed("process handler put data status error %s", err)
	}
	return nil, nil
}

func PutFeedSectionCrop(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	data := &define.FeedSectionCrop{}
	if err := json.Unmarshal([]byte(args[0]), data); err != nil {
		return utils.ResponseFailed("process handler unmarshal data error %s", err)
	}
	data.TxId = stub.GetTxID()
	putKey := "FeedSectionCrop" + data.FeedListId
	if err := utils.PutData(stub, data, putKey); err != nil {
		return utils.ResponseFailed("process handler put data status error %s", err)
	}
	return nil, nil
}

func PutDict(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	data := &define.AgDictInfo{}
	if err := json.Unmarshal([]byte(args[0]), data); err != nil {
		return utils.ResponseFailed("dict handler unmarshal data error %s", err)
	}
	putKey := "AgDict"
	if err := utils.PutData(stub, data, putKey); err != nil {
		return utils.ResponseFailed("dict handler put data status error %s", err)
	}
	return nil, nil
}

func PutPesticideFertilizer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	data := &define.AgPesticideFertilizerInfo{}
	if err := json.Unmarshal([]byte(args[0]), data); err != nil {
		return utils.ResponseFailed("reap handler unmarshal data error %s", err)
	}
	putKey := "AgPesticideFertilizer"
	if err := utils.PutData(stub, data, putKey); err != nil {
		return utils.ResponseFailed("tenant handler put data status error %s", err)
	}
	return nil, nil
}

/*
func PutTenant(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	tenant := &define.Tenant{}
	if err := json.Unmarshal([]byte(args[0]), tenant); err != nil {
		return utils.ResponseFailed("tenant handler unmarshal data error %s", err)
	}
	// 构造复合key
	putKey, err := stub.CreateCompositeKey("Tenant", []string{tenant.TenantId})
	if err !=nil {
		return utils.ResponseFailed("PutFarm create composite key error: %s", err)
	}
	//putKey := "Tenant"+ tenant.TenantId
	if err := utils.PutData(stub, tenant, putKey); err != nil {
		return utils.ResponseFailed("tenant handler put data status error %s", err)
	}
	return nil, nil
}

func PutFarm(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	farm := &define.Farm{}
	if err := json.Unmarshal([]byte(args[0]), farm); err != nil {
		return utils.ResponseFailed("farm handler unmarshal data error %s", err)
	}
	// 构造复合key
	putKey, err := stub.CreateCompositeKey("Farm", []string{farm.TenantId, farm.FarmId})
	fmt.Printf("putkey: %v", putKey)
	if err !=nil {
		return utils.ResponseFailed("PutFarm create composite key error: %s", err)
	}
	if err := utils.PutData(stub, farm, putKey); err != nil {
		return utils.ResponseFailed("tenant handler put data status error %s", err)
	}
	return nil, nil
}

func PutSection(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	section := &define.Section{}
	if err := json.Unmarshal([]byte(args[0]), section); err != nil {
		return utils.ResponseFailed("section handler unmarshal data error %s", err)
	}
	// 构造复合key
	putKey, err := stub.CreateCompositeKey("Section", []string{"Section", section.TenantId, section.FarmId, section.SectionId})
	if err !=nil {
		return utils.ResponseFailed("PutFarm create composite key error: %s", err)
	}
	if err := utils.PutData(stub, section, putKey); err != nil {
		return utils.ResponseFailed("tenant handler put data status error %s", err)
	}
	return nil, nil
}

func PutReap(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	reap := &define.Reap{}
	if err := json.Unmarshal([]byte(args[0]), reap); err != nil {
		return utils.ResponseFailed("reap handler unmarshal data error %s", err)
	}
	// 构造复合key
	putKey, err := stub.CreateCompositeKey("Reap", []string{reap.TenantId, reap.FarmId, reap.ReapBatchId})
	if err !=nil {
		return utils.ResponseFailed("PutFarm create composite key error: %s", err)
	}
	if err := utils.PutData(stub, reap, putKey); err != nil {
		return utils.ResponseFailed("tenant handler put data status error %s", err)
	}
	return nil, nil
}

func PutVarieties(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	varieties := &define.Varieties{}
	if err := json.Unmarshal([]byte(args[0]), varieties); err != nil {
		return utils.ResponseFailed("reap handler unmarshal data error %s", err)
	}
	putKey := "agVarieties"
	if err := utils.PutData(stub, varieties, putKey); err != nil {
		return utils.ResponseFailed("tenant handler put data status error %s", err)
	}
	return nil, nil
}

func QueryTenant(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	utils.Logger.Errorf("query key: %v\n", args[0])
	// 构造复合key
	queryKey, err := stub.CreateCompositeKey("Tenant", []string{args[0]})
	if err !=nil {
		return nil, err
	}
	res, err := stub.GetState(queryKey)
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QueryFarm(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	utils.Logger.Errorf("query key: %v\n", args[0])
	// 构造复合key
	queryKey, err := stub.CreateCompositeKey("Farm", strings.Split(args[0], "_"))
	if err !=nil {
		return nil, err
	}
	res, err := stub.GetState(queryKey)
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}

func QuerySection(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	// 构造复合key
	key := strings.Split(args[0], "_")
	// 查询请求参数需要满足，公司ID_农场ID_地块ID。
	if len(key) > 3 {
		// 如果查询参数齐全，就直接查询，否则进行复合查询
		if len(key[1])>0 && len(key[2])>0 && len(key[3])>0{
			queryKey, err := stub.CreateCompositeKey("Section", strings.Split(args[0], "_"))
			if err !=nil {
				return nil, err
			}
			utils.Logger.Errorf("query key: %v\n", queryKey)
			res, err := stub.GetState(queryKey)
			if err != nil {
				return utils.ResponseFailed("Form query failed to GetState err: %v", err)
			}
			return res, nil
		}
		if len(key[1])==0 && len(key[2])>0 && len(key[3])>0{ // 没有公司ID，复合查询所有匹配的键值
			sectionIter, err:= stub.GetStateByPartialCompositeKey("Section", []string{"Section"})
			if err != nil {
				return utils.ResponseFailed("Form query failed to GetStateByPartialCompositeKey err: %v", err)
			}
			defer sectionIter.Close()
			for sectionIter.HasNext() {
				item, _ := sectionIter.Next()
				_, keyParts, err := stub.SplitCompositeKey(item.Key)
				if err != nil {
					return utils.ResponseFailed("Form query failed to SplitCompositeKey err: %v", err)
				}
				if keyParts[2] == key[2] && keyParts[3] == key[3] {
					key[1] = keyParts[1]
				}
			}
			if len(key[1]) > 0 {
				queryKey, err := stub.CreateCompositeKey("Section", key)
				if err !=nil {
					return nil, err
				}
				res, err := stub.GetState(queryKey)
				if err != nil {
					return utils.ResponseFailed("Form query failed to GetState err: %v", err)
				}
				return res, nil
			}
		}
		if len(key[1])>0 && len(key[2])==0 && len(key[3])>0{ // 没有农场ID，复合查询所有匹配的键值
			sectionIter, err:= stub.GetStateByPartialCompositeKey("Section", []string{"Section", key[1]})
			if err != nil {
				return utils.ResponseFailed("Form query failed to GetStateByPartialCompositeKey err: %v", err)
			}
			defer sectionIter.Close()
			for sectionIter.HasNext() {
				item, _ := sectionIter.Next()
				_, keyParts, err := stub.SplitCompositeKey(item.Key)
				if err != nil {
					return utils.ResponseFailed("Form query failed to SplitCompositeKey err: %v", err)
				}
				if keyParts[3] == key[3] {
					key[2] = keyParts[2]
				}
			}
			if len(key[2]) > 0 {
				queryKey, err := stub.CreateCompositeKey("Section", key)
				if err !=nil {
					return nil, err
				}
				utils.Logger.Errorf("query key: %v\n", queryKey)
				res, err := stub.GetState(queryKey)
				if err != nil {
					return utils.ResponseFailed("Form query failed to GetState err: %v", err)
				}
				return res, nil
			}
		}
		if key[0] == "SoilExt" {
			sectionIter, err:= stub.GetStateByPartialCompositeKey("Section", []string{"Section"})
			if err != nil {
				return utils.ResponseFailed("Form query failed to GetStateByPartialCompositeKey err: %v", err)
			}
			defer sectionIter.Close()
			for sectionIter.HasNext() {
				item, _ := sectionIter.Next()
				_, keyParts, err := stub.SplitCompositeKey(item.Key)
				if err != nil {
					return utils.ResponseFailed("Form query failed to SplitCompositeKey err: %v", err)
				}
				queryKey, err := stub.CreateCompositeKey("Section", keyParts)
				if err !=nil {
					return nil, err
				}
				utils.Logger.Errorf("query key: %v\n", queryKey)
				res, err := stub.GetState(queryKey)
				if err != nil {
					return utils.ResponseFailed("Form query failed to GetState err: %v", err)
				}
				section := &define.Section{}
				if err := json.Unmarshal(res, &section); err != nil {
					return nil, fmt.Errorf("unmarshal file tx info error %v", err)
				}
				for _, v := range section.AgSoil {
					if string(v.AgSoilExt.Id) == keyParts[3] {
						return res, nil
					}
				}
				return nil, nil
			}
		}
		if key[0] == "FarmingRecordDetail" {
			sectionIter, err:= stub.GetStateByPartialCompositeKey("Section", []string{"Section", key[1]})
			if err != nil {
				return utils.ResponseFailed("Form query failed to GetStateByPartialCompositeKey err: %v", err)
			}
			defer sectionIter.Close()
			for sectionIter.HasNext() {
				item, _ := sectionIter.Next()
				_, keyParts, err := stub.SplitCompositeKey(item.Key)
				if err != nil {
					return utils.ResponseFailed("Form query failed to SplitCompositeKey err: %v", err)
				}
				queryKey, err := stub.CreateCompositeKey("Section", keyParts)
				if err !=nil {
					return nil, err
				}
				utils.Logger.Errorf("query key: %v\n", queryKey)
				res, err := stub.GetState(queryKey)
				if err != nil {
					return utils.ResponseFailed("Form query failed to GetState err: %v", err)
				}
				section := &define.Section{}
				if err := json.Unmarshal(res, &section); err != nil {
					return nil, fmt.Errorf("unmarshal file tx info error %v", err)
				}
				for _, v := range section.FarmingRecord {
					if v.Id == keyParts[3] {
						return res, nil
					}
				}
				return nil, nil
			}
		}
	}
	return nil, nil
}

func QueryReap(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	utils.Logger.Errorf("query key: %v\n", args[0])
	// 构造复合key
	key := strings.Split(args[0], "_")
	// 查询请求参数需要满足，公司ID_农场ID_地块ID。
	if len(key) > 3 {
		// 如果查询参数齐全，就直接查询，否则进行复合查询
		if len(key[1]) > 0 && len(key[2]) > 0 && len(key[3]) > 0 {
			queryKey, err := stub.CreateCompositeKey("Reap", strings.Split(args[0], "_"))
			if err != nil {
				return nil, err
			}
			res, err := stub.GetState(queryKey)
			if err != nil {
				return utils.ResponseFailed("Form query failed to GetState err: %v", err)
			}
			return res, nil
		}
		if len(key[1])>0 && len(key[2])==0 && len(key[3])>0{ // 没有农场ID，复合查询所有匹配的键值
			sectionIter, err:= stub.GetStateByPartialCompositeKey("Reap", []string{"Reap", key[1]})
			if err != nil {
				return utils.ResponseFailed("Form query failed to GetStateByPartialCompositeKey err: %v", err)
			}
			defer sectionIter.Close()
			for sectionIter.HasNext() {
				item, _ := sectionIter.Next()
				_, keyParts, err := stub.SplitCompositeKey(item.Key)
				if err != nil {
					return utils.ResponseFailed("Form query failed to SplitCompositeKey err: %v", err)
				}
				if keyParts[3] == key[3] {
					key[2] = keyParts[2]
				}
			}
			if len(key[2]) > 0 {
				queryKey, err := stub.CreateCompositeKey("Section", key)
				if err !=nil {
					return nil, err
				}
				res, err := stub.GetState(queryKey)
				if err != nil {
					return utils.ResponseFailed("Form query failed to GetState err: %v", err)
				}
				return res, nil
			}
		}
	}

	return nil, nil
}

func QueryVarieties(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	utils.Logger.Errorf("query key: %v\n", args[0])
	queryKey := "agVarieties"
	res, err := stub.GetState(queryKey)
	if err != nil {
		return utils.ResponseFailed("Form query failed to GetState err: %v", err)
	}
	return res, nil
}
*/
