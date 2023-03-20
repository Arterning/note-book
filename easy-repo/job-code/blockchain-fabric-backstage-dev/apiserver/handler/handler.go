package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/crc/gohfc"
	apiDefine "github.com/crc/zlzk/apiserver/define"
	"github.com/crc/zlzk/apiserver/sdk"
	"github.com/crc/zlzk/apiserver/utils"
	"github.com/crc/zlzk/chaincode/define"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func PutUploadRecord(c *gin.Context) {
	//farm := &define.Farm{}
	uploadRecords := &define.UploadRecordBatch{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, uploadRecords); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if len(uploadRecords.UploadRecordList) == 0 {
		return
	}
	txId := ""
	for _, farm := range uploadRecords.UploadRecordList {
		// 查询历史数据并校验，暂时不做
		data, err := json.Marshal(farm)
		if err != nil {
			errStr := fmt.Sprintf("json marshal error %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}

		// 更新链上数据
		info := []string{"PutUploadRecord", string(data)}
		resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		txId = resp.TxID
	}
	utils.Response(txId, c, http.StatusOK)
	return
}

func PutFarm(c *gin.Context) {
	//farm := &define.Farm{}
	farmBatch := &define.FarmBatch{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, farmBatch); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if len(farmBatch.FarmList) == 0 {
		return
	}
	txId := ""
	for _, farm := range farmBatch.FarmList {
		// 查询历史数据并校验，暂时不做
		data, err := json.Marshal(farm)
		if err != nil {
			errStr := fmt.Sprintf("json marshal error %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		// 更新链上数据
		info := []string{"PutFarm", string(data)}
		resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		txId = resp.TxID
	}
	utils.Response(txId, c, http.StatusOK)
	return
}

func PutLand(c *gin.Context) {
	//land := &define.Land{}
	landBatch := &define.LandBatch{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, landBatch); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if len(landBatch.LandList) == 0 {
		return
	}
	txId := ""
	for _, land := range landBatch.LandList {
		// 查询历史数据并校验，暂时不做
		data, err := json.Marshal(land)
		if err != nil {
			errStr := fmt.Sprintf("json marshal error %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		// 更新链上数据
		info := []string{"PutLand", string(data)}
		resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		txId = resp.TxID
	}
	utils.Response(txId, c, http.StatusOK)
	return
}

func PutPlant(c *gin.Context) {
	//plant := &define.Plant{}
	plantBatch := &define.PlantBatch{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, plantBatch); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if len(plantBatch.PlantList) == 0 {
		return
	}
	txId := ""
	for _, plant := range plantBatch.PlantList {
		// 查询历史数据并校验，暂时不做
		data, err := json.Marshal(plant)
		if err != nil {
			errStr := fmt.Sprintf("json marshal error %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		// 更新链上数据
		info := []string{"PutPlant", string(data)}
		resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		txId = resp.TxID
	}
	utils.Response(txId, c, http.StatusOK)
	return
}

func PutFarmRecord(c *gin.Context) {
	//farmRecord := &define.FarmRecord{}
	farmRecordBatch := &define.FarmRecordBatch{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, farmRecordBatch); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if len(farmRecordBatch.FarmRecordList) == 0 {
		return
	}
	txId := ""
	for _, farmRecord := range farmRecordBatch.FarmRecordList {
		// 查询历史数据并校验，暂时不做
		data, err := json.Marshal(farmRecord)
		if err != nil {
			errStr := fmt.Sprintf("json marshal error %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		// 更新链上数据
		info := []string{"PutFarmRecord", string(data)}
		resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		txId = resp.TxID
	}
	utils.Response(txId, c, http.StatusOK)
	return
}

func PutHarvest(c *gin.Context) {
	//harvest := &define.Harvest{}
	harvestBatch := &define.HarvestBatch{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, harvestBatch); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if len(harvestBatch.HarvestList) == 0 {
		return
	}
	txId := ""
	for _, harvest := range harvestBatch.HarvestList {
		// 查询历史数据并校验，暂时不做
		data, err := json.Marshal(harvest)
		if err != nil {
			errStr := fmt.Sprintf("json marshal error %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		// 更新链上数据
		info := []string{"PutHarvest", string(data)}
		resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		txId = resp.TxID
	}
	utils.Response(txId, c, http.StatusOK)
	return
}

func PutDryInfo(c *gin.Context) {
	//dryInfo := &define.DryInfo{}
	dryInfoBatch := &define.DryInfoBatch{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, dryInfoBatch); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if len(dryInfoBatch.DryInfoList) == 0 {
		return
	}
	txId := ""
	for _, dryInfo := range dryInfoBatch.DryInfoList {
		// 查询历史数据并校验，暂时不做
		data, err := json.Marshal(dryInfo)
		if err != nil {
			errStr := fmt.Sprintf("json marshal error %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		// 更新链上数据
		info := []string{"PutDryInfo", string(data)}
		resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		txId = resp.TxID
	}
	utils.Response(txId, c, http.StatusOK)
	return
}

func PutWarehouseInfo(c *gin.Context) {
	//warehouseInfo := &define.WarehouseInfo{}
	warehouseInfoBatch := &define.WarehouseInfoBatch{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, warehouseInfoBatch); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if len(warehouseInfoBatch.WarehouseInfoList) == 0 {
		return
	}
	txId := ""
	for _, warehouseInfo := range warehouseInfoBatch.WarehouseInfoList {
		// 查询历史数据并校验，暂时不做
		data, err := json.Marshal(warehouseInfo)
		if err != nil {
			errStr := fmt.Sprintf("json marshal error %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		// 更新链上数据
		info := []string{"PutWarehouseInfo", string(data)}
		resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		txId = resp.TxID
	}
	utils.Response(txId, c, http.StatusOK)
	return
}

func PutProcessInfo(c *gin.Context) {
	//processInfo := &define.ProcessInfo{}
	processInfoBatch := &define.ProcessInfoBatch{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, processInfoBatch); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if len(processInfoBatch.ProcessInfoList) == 0 {
		return
	}
	txId := ""
	for _, processInfo := range processInfoBatch.ProcessInfoList {
		// 查询历史数据并校验，暂时不做
		data, err := json.Marshal(processInfo)
		if err != nil {
			errStr := fmt.Sprintf("json marshal error %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		// 更新链上数据
		info := []string{"PutProcessInfo", string(data)}
		resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		txId = resp.TxID
	}

	utils.Response(txId, c, http.StatusOK)
	return
}

func PutFeedSectionMap(c *gin.Context) {
	feedSectionMap := &define.FeedSectionCrop{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, feedSectionMap); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if len(feedSectionMap.FeedListId) == 0 {
		errStr := fmt.Sprintf("input params error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	txId := ""
	// 查询历史数据并校验，暂时不做
	data, err := json.Marshal(feedSectionMap)
	if err != nil {
		errStr := fmt.Sprintf("json marshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	// 更新链上数据
	info := []string{"PutFeedSectionCrop", string(data)}
	resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
	if err != nil {
		errStr := fmt.Sprintf("数据上链失败 %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	txId = resp.TxID
	utils.Response(txId, c, http.StatusOK)
	return
}

func PutDict(c *gin.Context) {
	agDictInfo := &define.AgDictInfo{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, agDictInfo); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	// 查询历史数据并校验，暂时不做
	data, err := json.Marshal(agDictInfo)
	if err != nil {
		errStr := fmt.Sprintf("json marshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	// 更新链上数据
	info := []string{"PutDict", string(data)}
	resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
	if err != nil {
		errStr := fmt.Sprintf("数据上链失败 %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	utils.Response(resp.TxID, c, http.StatusOK)
	return
}

func PutPesticideFertilizer(c *gin.Context) {
	agPesticideFertilizerInfo := &define.AgPesticideFertilizerInfo{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, agPesticideFertilizerInfo); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	// 查询历史数据并校验，暂时不做
	data, err := json.Marshal(agPesticideFertilizerInfo)
	if err != nil {
		errStr := fmt.Sprintf("json marshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	// 更新链上数据
	info := []string{"PutPesticideFertilizer", string(data)}
	resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
	if err != nil {
		errStr := fmt.Sprintf("数据上链失败 %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	utils.Response(resp.TxID, c, http.StatusOK)
	return
}

/* KafkaData 接受kafka订阅数据
* 每次数据库有新增（insert）更新（update）删除（delete）操作时
* 会将变更的数据表及数据发送到kafka。
* 解析对应表数据，分别更新公司、农场、地块、收割对应的上链数据
* 先查询获取原有链上数据，将对应部分数据更新后，再次上链新数据
* 除上链数据采用复合key上链，需要采用复合key查询*/
func KafkaData(c *gin.Context) {
	var err error
	kafkaMessage := &apiDefine.KafkaMessage{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, kafkaMessage); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}

	fmt.Printf("kafkaMessage: %v\n", kafkaMessage)
	switch kafkaMessage.Table {
	case "ag_farm":
		// 农场信息
		res, err := HandleFarm(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "ag_section":
		// 地块信息
		// 更新农场信息，中的地块ID列表
		res, err := HandleFarm(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		res, err = HandleLand(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "ag_test_report":
		// 土壤、水质检测报告
		res, err := HandleFarm(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "ag_farming_record":
		// 农事记录
		// 更新地块信息，中的农事记录
		res, err := HandlePlant(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "ag_farming_record_detail":
		// 农事记录详情
		// 更新地块信息，中的农事记录详情
		res, err := HandlePlant(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "sys_attachment":
		// 农事记录照片
		// 更新农事记录中的照片信息，及种植信息中的照片信息
		res, err := HandlePlant(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "ag_reap_batch":
		// 收割批次
		// 更新收割信息，中的收割批次信息
		res, err := HandleHarvest(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "ag_section_crops":
		utils.Response("", c, http.StatusOK)
		return
	case "ag_section_crops_reap":
		// 地块农作物收割关联
		// 更新地块信息，中的收割列表
		res, err := HandleHarvest(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "srv_dry_info":
		// 烘干批次
		// 更新收割信息，中的烘干批次
		res, err := HandleDry(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "srv_storage_info":
		// 仓储
		res, err := HandleWarehouse(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "srv_mach_packing":
		// 加工记录
		// 更新收割信息，中的加工记录
		res, err := HandleProcess(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "ag_dict":
		// 字典表，type code + value=code唯一
		res, err := HandleDict(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "ag_dict_type":
		utils.Response("", c, http.StatusOK)
		return
	case "ag_pesticide_fertilizer":
		// 农药化肥表
		res, err := HandlePesticideFertilizer(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	case "srv_feed_section_map":
		// 智农云筛选的地块列表
		res, err := HandleFeedSectionMap(*kafkaMessage)
		if err != nil {
			errStr := fmt.Sprintf("数据上链失败 %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
		}
		utils.Response(res, c, http.StatusOK)
		return
	default:
		// 未知表，请联系开发人员添加
		errStr := fmt.Sprintf("未知表，请联系开发人员添加")
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err != nil {
		errStr := fmt.Sprintf("数据上链失败 %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	utils.Response("", c, http.StatusOK)
	return
}

// 农场上链信息
// key: 农场ID
// value：地块ID列表
func HandleFarm(farm apiDefine.KafkaMessage) (string, error) {
	response := ""
	var err error
	switch farm.Table {
	case "ag_farm":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup

		for gIndex, land := range farm.Data {
			wg.Add(1)
			f := &define.Farm{}
			f.TenantId = reflect.ValueOf(land["tenant_id"]).String()
			f.FarmId = reflect.ValueOf(land["id"]).String()
			isActive := reflect.ValueOf(land["is_active"]).String()

			switch farm.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(f.FarmId, define.QueryFarm)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.Farm{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					// 删除对应数据
					old.Status = 1
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutFarm", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				go func() {
					d, err := sdk.Query(f.FarmId, define.QueryFarm)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.Farm{}
					if len(d) > 0 {
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						old.TenantId = f.TenantId
					} else {
						old = f
					}
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutFarm", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				_, okTenantId := farm.Old[gIndex]["tenant_id"]
				_, okFarmId := farm.Old[gIndex]["id"]
				_, okIsActive := farm.Old[gIndex]["is_active"]
				go func() {
					farmId := ""
					if okFarmId {
						farmId = reflect.ValueOf(farm.Old[gIndex]["id"]).String()
					} else {
						farmId = f.FarmId
					}
					d, err := sdk.Query(farmId, define.QueryFarm)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.Farm{}
					if len(d) > 0 {
						// 存在历史值就更新
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						if okIsActive && isActive == "N" {
							// 删除对应数据
							old.Status = 1
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutFarm", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
							wg.Done()
							return
						}
						if okFarmId {
							old.FarmId = f.FarmId
						}
						if okTenantId {
							old.TenantId = f.TenantId
						}
						if okFarmId || okTenantId {
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutFarm", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
						}
					} else if len(d) == 0 {
						// 不存在历史值，就插入
						old.FarmId = f.FarmId
						old.TenantId = f.TenantId
						data, err := json.Marshal(old)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						info := []string{"PutFarm", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
					}
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			}
		}
		wg.Wait()
		cancel()
	case "ag_section":
		for gIndex, land := range farm.Data {
			f := &define.Farm{}
			f.TenantId = reflect.ValueOf(land["tenant_id"]).String()
			f.FarmId = reflect.ValueOf(land["farm_id"]).String()
			sectionId := reflect.ValueOf(land["id"]).String()
			isActive := reflect.ValueOf(land["is_active"]).String()

			switch farm.Type {
			case "DELETE":
				d, err := sdk.Query(f.FarmId, define.QueryFarm)
				if err != nil || len(d) == 0 {
					//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				old := &define.Farm{}
				err = json.Unmarshal(d, old)
				if err != nil {
					return "", err
				}
				// 删除对应数据
				for i := 0; i < len(old.SectionIdList); i++ {
					if old.SectionIdList[i] == sectionId {
						old.SectionIdList = append(old.SectionIdList[:i], old.SectionIdList[i+1:]...)
						i--
					}
				}
				data, err := json.Marshal(old)
				if err != nil {
					return "", err
				}
				// 更新链上数据
				info := []string{"PutFarm", string(data)}
				res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
				if err != nil {
					return "", err
				}
				response = res.TxID

			case "INSERT":
				d, err := sdk.Query(f.FarmId, define.QueryFarm)
				if err != nil {
					//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				old := &define.Farm{}
				if len(d) > 0 {
					err = json.Unmarshal(d, old)
					if err != nil {
						return "", err
					}
					if old.Status == 0 && !in(sectionId, old.SectionIdList) {
						fmt.Printf("section id: %v\n", sectionId)
						old.SectionIdList = append(old.SectionIdList, sectionId)
					}
				} else {
					old.FarmId = f.FarmId
					old.TenantId = f.TenantId
					old.SectionIdList = append(old.SectionIdList, sectionId)
				}
				data, err := json.Marshal(old)
				if err != nil {
					return "", err
				}
				// 更新链上数据
				info := []string{"PutFarm", string(data)}
				res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
				if err != nil {
					return "", err
				}
				response = res.TxID

			case "UPDATE":
				_, okSectionId := farm.Old[gIndex]["id"]
				_, okTenantId := farm.Old[gIndex]["tenant_id"]
				_, okFarmId := farm.Old[gIndex]["farm_id"]
				_, okIsActive := farm.Old[gIndex]["is_active"]
				farmId := ""
				if okFarmId {
					farmId = reflect.ValueOf(farm.Old[gIndex]["farm_id"]).String()
				} else {
					farmId = f.FarmId
				}
				d, err := sdk.Query(farmId, define.QueryFarm)
				if err != nil {
					//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				old := &define.Farm{}
				if len(d) > 0 {
					// 存在历史值就更新
					err = json.Unmarshal(d, old)
					if err != nil {
						return "", err
					}
					if okIsActive && isActive == "N" {
						// 删除对应数据
						for i := 0; i < len(old.SectionIdList); i++ {
							if old.SectionIdList[i] == sectionId {
								old.SectionIdList = append(old.SectionIdList[:i], old.SectionIdList[i+1:]...)
								i--
							}
						}
						data, err := json.Marshal(old)
						if err != nil {
							return "", err
						}
						// 更新链上数据
						info := []string{"PutFarm", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							return "", err
						}
						response = res.TxID
						return response, err
					}
					oldSectionId := ""
					if okSectionId {
						oldSectionId = reflect.ValueOf(farm.Old[gIndex]["id"]).String()
					}
					for i, v := range old.SectionIdList {
						if v == oldSectionId {
							old.SectionIdList[i] = sectionId
						}
					}
					if okFarmId {
						old.FarmId = f.FarmId
					}
					if okTenantId {
						old.TenantId = f.TenantId
					}
					if okFarmId || okTenantId || okSectionId {
						data, err := json.Marshal(old)
						if err != nil {
							return "", err
						}
						// 更新链上数据
						info := []string{"PutFarm", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							return "", err
						}
						response = res.TxID
					}
				} else if len(d) == 0 {
					//不存在历史值就插入
					if !in(sectionId, old.SectionIdList) {
						old.SectionIdList = append(old.SectionIdList, sectionId)
					}
					old.FarmId = f.FarmId
					old.TenantId = f.TenantId
					data, err := json.Marshal(old)
					if err != nil {
						return "", err
					}
					// 更新链上数据
					info := []string{"PutFarm", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						return "", err
					}
					response = res.TxID
				}
			}
		}
	case "ag_test_report":
		for gIndex, land := range farm.Data {
			farmId := reflect.ValueOf(land["farm_id"]).String()
			testReportType := reflect.ValueOf(land["type"]).String() // 检测报告类型：1土壤，2水质
			beginDate, _ := time.Parse("2006-01-02 15:04:05", reflect.ValueOf(land["begin_date"]).String())
			endDate, _ := time.Parse("2006-01-02 15:04:05", reflect.ValueOf(land["end_date"]).String())
			isActive := reflect.ValueOf(land["is_active"]).String()
			imageUrl := ""
			if _, okImg := land["attachment_id"]; okImg {
				imageUrl = reflect.ValueOf(land["attachment_id"]).String()
			}

			switch farm.Type {
			case "DELETE":
				f, err := sdk.Query(farmId, define.QueryFarm)
				if err != nil || len(f) == 0 {
					return "", err
				}
				oldFarm := &define.Farm{}
				err = json.Unmarshal(f, &oldFarm)
				if err != nil {
					return "", err
				}
				if oldFarm.Status == 0 {
					if testReportType == "1" {
						for i := 0; i < len(oldFarm.SoilTestReportList); i++ {
							if oldFarm.SoilTestReportList[i].ImageUrl == imageUrl {
								oldFarm.SoilTestReportList = append(oldFarm.SoilTestReportList[:i], oldFarm.SoilTestReportList[i+1:]...)
								i--
							}
						}
					} else if testReportType == "2" {
						for i := 0; i < len(oldFarm.WaterTestReportList); i++ {
							if oldFarm.WaterTestReportList[i].ImageUrl == imageUrl {
								oldFarm.WaterTestReportList = append(oldFarm.WaterTestReportList[:i], oldFarm.WaterTestReportList[i+1:]...)
								i--
							}
						}
					}
					data, err := json.Marshal(oldFarm)
					if err != nil {
						continue
					}
					// 更新链上数据
					info := []string{"PutFarm", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						continue
					}
					response = res.TxID
				}
			case "INSERT":
				f, err := sdk.Query(farmId, define.QueryFarm)
				if err != nil || len(f) == 0 {
					return "", err
				}
				oldFarm := &define.Farm{}
				err = json.Unmarshal(f, &oldFarm)
				if err != nil {
					return "", err
				}
				if oldFarm.Status == 0 {
					testReportList := &define.TestReport{}
					testReportList.InspectionDate = beginDate
					testReportList.InspectionEndDate = endDate
					testReportList.ImageUrl = imageUrl
					if testReportType == "1" {
						existFlag := false
						for _, v := range oldFarm.SoilTestReportList {
							if v.ImageUrl == imageUrl {
								existFlag = true
							}
						}
						if !existFlag {
							oldFarm.SoilTestReportList = append(oldFarm.SoilTestReportList, *testReportList)
						}
					} else if testReportType == "2" {
						existFlag := false
						for _, v := range oldFarm.WaterTestReportList {
							if v.ImageUrl == imageUrl {
								existFlag = true
							}
						}
						if !existFlag {
							oldFarm.WaterTestReportList = append(oldFarm.WaterTestReportList, *testReportList)
						}
					}
					data, err := json.Marshal(oldFarm)
					if err != nil {
						continue
					}
					// 更新链上数据
					info := []string{"PutFarm", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						continue
					}
					response = res.TxID
				}
			case "UPDATE":
				_, okFarmId := farm.Old[gIndex]["farm_id"]
				_, okTestReportType := farm.Old[gIndex]["type"]
				_, okBeginDate := farm.Old[gIndex]["begin_date"]
				_, okEndDate := farm.Old[gIndex]["end_date"]
				_, okImageUrl := farm.Old[gIndex]["attachment_id"]
				_, okIsActive := farm.Old[gIndex]["is_active"]
				oldFarmId := ""
				if okFarmId {
					oldFarmId = reflect.ValueOf(farm.Old[gIndex]["farm_id"]).String()
				} else {
					oldFarmId = farmId
				}
				f, err := sdk.Query(oldFarmId, define.QueryFarm)
				if err != nil {
					return "", err
				}
				oldFarm := &define.Farm{}
				if len(f) > 0 {
					err = json.Unmarshal(f, &oldFarm)
					if err != nil {
						return "", err
					}
					if okIsActive && isActive == "N" {
						if oldFarm.Status == 0 {
							if testReportType == "1" {
								for i := 0; i < len(oldFarm.SoilTestReportList); i++ {
									if oldFarm.SoilTestReportList[i].ImageUrl == imageUrl {
										oldFarm.SoilTestReportList = append(oldFarm.SoilTestReportList[:i], oldFarm.SoilTestReportList[i+1:]...)
										i--
									}
								}
							} else if testReportType == "2" {
								for i := 0; i < len(oldFarm.WaterTestReportList); i++ {
									if oldFarm.WaterTestReportList[i].ImageUrl == imageUrl {
										oldFarm.WaterTestReportList = append(oldFarm.WaterTestReportList[:i], oldFarm.WaterTestReportList[i+1:]...)
										i--
									}
								}
							}
							data, err := json.Marshal(oldFarm)
							if err != nil {
								continue
							}
							// 更新链上数据
							info := []string{"PutFarm", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								continue
							}
							response = res.TxID
							continue
						}
					}
					if oldFarm.Status == 0 {
						oldImageUrl := ""
						if okImageUrl {
							oldImageUrl = reflect.ValueOf(farm.Old[gIndex]["attachment_id"]).String()
						}
						if testReportType == "1" {
							if len(oldFarm.SoilTestReportList) > 0 {
								existFlag := false
								for i, v := range oldFarm.SoilTestReportList {
									if v.ImageUrl == oldImageUrl || v.ImageUrl == imageUrl {
										oldFarm.SoilTestReportList[i].ImageUrl = imageUrl
										if okBeginDate {
											oldFarm.SoilTestReportList[i].InspectionDate = beginDate
										}
										if okEndDate {
											oldFarm.SoilTestReportList[i].InspectionEndDate = endDate
										}
										existFlag = true
									}
								}
								if !existFlag {
									testReport := define.TestReport{
										ImageUrl:          imageUrl,
										InspectionDate:    beginDate,
										InspectionEndDate: endDate,
									}
									oldFarm.SoilTestReportList = append(oldFarm.SoilTestReportList, testReport)
								}
							} else {
								var testReport define.TestReport
								testReport.ImageUrl = imageUrl
								testReport.InspectionDate = beginDate
								testReport.InspectionEndDate = endDate
								oldFarm.SoilTestReportList = append(oldFarm.SoilTestReportList, testReport)
							}
						} else if testReportType == "2" {
							if len(oldFarm.WaterTestReportList) > 0 {
								existFlag := false
								for i, v := range oldFarm.WaterTestReportList {
									if v.ImageUrl == oldImageUrl || v.ImageUrl == imageUrl {
										oldFarm.WaterTestReportList[i].ImageUrl = imageUrl
										if okBeginDate {
											oldFarm.WaterTestReportList[i].InspectionDate = beginDate
										}
										if okEndDate {
											oldFarm.WaterTestReportList[i].InspectionEndDate = endDate
										}
										existFlag = true
									}
								}
								if !existFlag {
									testReport := define.TestReport{
										ImageUrl:          imageUrl,
										InspectionDate:    beginDate,
										InspectionEndDate: endDate,
									}
									oldFarm.WaterTestReportList = append(oldFarm.WaterTestReportList, testReport)
								}
							} else {
								var testReport define.TestReport
								testReport.ImageUrl = imageUrl
								testReport.InspectionDate = beginDate
								testReport.InspectionEndDate = endDate
								oldFarm.WaterTestReportList = append(oldFarm.WaterTestReportList, testReport)
							}
						}
						if okFarmId || okTestReportType || okBeginDate || okEndDate || okImageUrl {
							data, err := json.Marshal(farm)
							if err != nil {
								continue
							}
							// 更新链上数据
							info := []string{"PutFarm", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								continue
							}
							response = res.TxID
						}
					}
				} else if len(f) == 0 {
					oldFarm.FarmId = farmId
					if testReportType == "1" {
						var testReport define.TestReport
						testReport.ImageUrl = imageUrl
						testReport.InspectionDate = beginDate
						testReport.InspectionEndDate = endDate
						oldFarm.SoilTestReportList = append(oldFarm.SoilTestReportList, testReport)
					} else if testReportType == "2" {
						var testReport define.TestReport
						testReport.ImageUrl = imageUrl
						testReport.InspectionDate = beginDate
						testReport.InspectionEndDate = endDate
						oldFarm.WaterTestReportList = append(oldFarm.WaterTestReportList, testReport)
					}
					data, err := json.Marshal(oldFarm)
					if err != nil {
						continue
					}
					// 更新链上数据
					info := []string{"PutFarm", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						continue
					}
					response = res.TxID
				}
				return response, nil
			}
		}
	}
	return response, err
}

// 地块上链信息
// key: 地块ID
// value：地块ID、地块名称、地块位置、地块图片URL
func HandleLand(landSection apiDefine.KafkaMessage) (string, error) {
	response := ""
	var err error
	switch landSection.Table {
	case "ag_section":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup

		for gIndex, land := range landSection.Data {
			wg.Add(1)
			l := &define.Land{}
			l.SectionId = reflect.ValueOf(land["id"]).String()
			l.FarmId = reflect.ValueOf(land["farm_id"]).String()
			l.Name = reflect.ValueOf(land["name"]).String()
			l.Province = reflect.ValueOf(land["province"]).String()
			l.City = reflect.ValueOf(land["city"]).String()
			l.District = reflect.ValueOf(land["district"]).String()
			l.Address = reflect.ValueOf(land["address"]).String()
			l.DetailedAddress = reflect.ValueOf(land["detailed_address"]).String()
			l.ImgUrl = reflect.ValueOf(land["img_url"]).String()
			l.Area, _ = strconv.ParseFloat(reflect.ValueOf(land["area"]).String(), 64)
			isActive := reflect.ValueOf(land["is_active"]).String()
			switch landSection.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(l.SectionId, define.QueryLand)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.Land{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					// 删除对应数据
					old.Status = 1
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutLand", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				go func() {
					d, err := sdk.Query(l.SectionId, define.QueryLand)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					if len(d) > 0 {
						old := &define.Land{}
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						if old.Status == 0 {
							fmt.Println("数据数据已存在")
							wg.Done()
							return
						}
					}
					data, err := json.Marshal(l)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutLand", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				go func() {
					_, ok := landSection.Old[gIndex]["id"]
					_, okFarmId := landSection.Old[gIndex]["farm_id"]
					_, okName := landSection.Old[gIndex]["name"]
					_, okProvince := landSection.Old[gIndex]["province"]
					_, okCity := landSection.Old[gIndex]["city"]
					_, okDistrict := landSection.Old[gIndex]["district"]
					_, okAddress := landSection.Old[gIndex]["address"]
					_, okDetailedAddress := landSection.Old[gIndex]["detailed_address"]
					_, okImgUrl := landSection.Old[gIndex]["img_url"]
					_, okArea := landSection.Old[gIndex]["area"]
					_, okIsActive := landSection.Old[gIndex]["is_active"]
					sectionId := ""
					if ok {
						sectionId = reflect.ValueOf(landSection.Old[gIndex]["id"]).String()
					} else {
						sectionId = l.SectionId
					}
					d, err := sdk.Query(sectionId, define.QueryLand)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.Land{}
					if len(d) > 0 {
						// 存在历史值就更新
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						if okIsActive && isActive == "N" {
							// 删除对应数据
							old.Status = 1
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutLand", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
							wg.Done()
							return
						}
						if ok {
							old.SectionId = l.SectionId
						}
						if okFarmId {
							old.FarmId = l.FarmId
						}
						if okName {
							old.Name = l.Name
						}
						if okProvince {
							old.Province = l.Province
						}
						if okCity {
							old.City = l.City
						}
						if okDistrict {
							old.District = l.District
						}
						if okAddress {
							old.Address = l.Address
						}
						if okDetailedAddress {
							old.DetailedAddress = l.DetailedAddress
						}
						if okImgUrl {
							old.ImgUrl = l.ImgUrl
						}
						if okArea {
							old.Area = l.Area
						}
						if ok || okFarmId || okName || okProvince || okCity || okDistrict || okAddress ||
							okDetailedAddress || okImgUrl || okArea {
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutLand", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
						}
					} else if len(d) == 0 {
						// 不存在历史值就插入
						old.SectionId = l.SectionId
						old.FarmId = l.FarmId
						old.Name = l.Name
						old.Province = l.Province
						old.City = l.City
						old.District = l.District
						old.Address = l.Address
						old.DetailedAddress = l.DetailedAddress
						old.ImgUrl = l.ImgUrl
						old.Area = l.Area
						data, err := json.Marshal(old)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						info := []string{"PutLand", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
					}
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()
			}
		}
		wg.Wait()
		cancel()
	}
	return response, err
}

// 种植上链信息
// key: 种植ID
// value: 种植ID、地块ID、农事记录ID列表、农事记录及详情
// 农事记录按农事八周期分类（大田准备、移栽、返青期、分蘖期、孕穗期、抽穗扬花、灌浆期、成熟期收获）
func HandlePlant(plantData apiDefine.KafkaMessage) (string, error) {
	response := ""
	switch plantData.Table {
	case "ag_farming_record":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup

		// 农事记录id、农事操作人、处理时间
		for gIndex, plant := range plantData.Data {
			wg.Add(1)
			l := &define.Plant{}
			// 农事记录Id
			farmRecordId := reflect.ValueOf(plant["id"]).String()
			// 地块Id
			l.SectionId = reflect.ValueOf(plant["section_id"]).String()
			// 种植Id
			l.SectionCropsId = reflect.ValueOf(plant["section_crops_id"]).String()
			handleDate, _ := time.Parse("2006-01-02 15:04:05", reflect.ValueOf(plant["handle_date"]).String())
			operatorName := reflect.ValueOf(plant["operator_name"]).String()
			stageCode := reflect.ValueOf(plant["stage_code"]).String()
			isActive := reflect.ValueOf(plant["is_active"]).String()

			switch plantData.Type {
			// 删除农事记录
			case "DELETE":
				go func() {
					d, err := sdk.Query(l.SectionCropsId, define.QueryPlant)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.Plant{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					// 删除种植信息中对应农事记录，这里只做记录，
					// 删除在ag_farming_record_detail表时处理
					old.FarmRecordIdDeleteList = append(old.FarmRecordIdDeleteList, farmRecordId)
					for i := 0; i < len(old.FarmRecordIdList); i++ {
						if old.FarmRecordIdList[i] == farmRecordId {
							old.FarmRecordIdList = append(old.FarmRecordIdList[:i], old.FarmRecordIdList[i+1:]...)
							i--
						}
					}
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutPlant", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					// 删除农事记录
					f, err := sdk.Query(farmRecordId, define.QueryFarmRecord)
					if err != nil || len(f) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					ff := &define.FarmRecord{}
					err = json.Unmarshal(f, ff)
					if err != nil {
						wg.Done()
						return
					}
					ff.Status = 1
					fData, err := json.Marshal(ff)
					if err != nil {
						wg.Done()
						return
					}
					// 更新plant上链数据
					fInfo := []string{"PutFarmRecord", string(fData)}
					_, err = gohfc.GetHandler().SyncInvoke(fInfo, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				go func() {
					d, err := sdk.Query(farmRecordId, define.QueryFarmRecord)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					// 如果已经存在此农事操作则不进行上链更新
					if len(d) > 0 {
						ff := &define.FarmRecord{}
						err = json.Unmarshal(d, ff)
						if err != nil {
							wg.Done()
							return
						}
						if ff.Status == 0 {
							ff.SectionCropsId = l.SectionCropsId
							ff.HandleDate = handleDate
							ff.OperatorName = operatorName
							ff.StageCode = stageCode
							fData, err := json.Marshal(ff)
							if err != nil {
								wg.Done()
								return
							}
							// 更新plant上链数据
							fInfo := []string{"PutFarmRecord", string(fData)}
							res, err := gohfc.GetHandler().SyncInvoke(fInfo, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
							wg.Done()
							return
						}
					}
					// 更新farmRecord上链数据
					farmRecord := &define.FarmRecord{}
					farmRecord.SectionCropsId = l.SectionCropsId
					farmRecord.OperateId = farmRecordId
					farmRecord.HandleDate = handleDate
					farmRecord.OperatorName = operatorName
					farmRecord.StageCode = stageCode
					fData, err := json.Marshal(farmRecord)
					if err != nil {
						wg.Done()
						return
					}
					// 更新plant上链数据
					fInfo := []string{"PutFarmRecord", string(fData)}
					res, err := gohfc.GetHandler().SyncInvoke(fInfo, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				_, okFarmRecordId := plantData.Old[gIndex]["id"]
				_, okSectionId := plantData.Old[gIndex]["section_id"]
				_, okSectionCropsId := plantData.Old[gIndex]["section_crops_id"]
				_, okHandleDate := plantData.Old[gIndex]["handle_date"]
				_, okOperatorName := plantData.Old[gIndex]["operator_name"]
				_, okStageCode := plantData.Old[gIndex]["stage_code"]
				_, okIsActive := plantData.Old[gIndex]["is_active"]
				go func() {
					oldFarmRecordId := ""
					if okFarmRecordId {
						oldFarmRecordId = reflect.ValueOf(plantData.Old[gIndex]["id"]).String()
					} else {
						oldFarmRecordId = farmRecordId
					}
					d, err := sdk.Query(oldFarmRecordId, define.QueryFarmRecord)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.FarmRecord{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					if okIsActive && isActive == "N" {
						old.Status = 1
						fData, err := json.Marshal(old)
						if err != nil {
							wg.Done()
							return
						}
						// 更新plant上链数据
						fInfo := []string{"PutFarmRecord", string(fData)}
						res, err := gohfc.GetHandler().SyncInvoke(fInfo, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
						wg.Done()
					} else {
						// 更新数据
						if okFarmRecordId {
							old.OperateId = farmRecordId
						}
						if okSectionCropsId {
							old.SectionCropsId = l.SectionCropsId
						}
						if okHandleDate {
							old.HandleDate = handleDate
						}
						if okOperatorName {
							old.OperatorName = operatorName
						}
						if okStageCode {
							old.StageCode = stageCode
						}
						if okFarmRecordId || okSectionCropsId || okHandleDate || okOperatorName || okStageCode {
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutFarmRecord", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
						}
					}
					sectionCropsId := ""
					if okSectionCropsId {
						sectionCropsId = reflect.ValueOf(plant["section_crops_id"]).String()
					} else {
						sectionCropsId = l.SectionCropsId
					}
					p, err := sdk.Query(sectionCropsId, define.QueryPlant)
					if err != nil || len(p) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					oldPlant := &define.Plant{}
					err = json.Unmarshal(p, oldPlant)
					if err != nil {
						wg.Done()
						return
					}
					if okIsActive && isActive == "N" {
						// 删除种植信息中对应农事记录，这里只做记录，
						// 删除在ag_farming_record_detail表时处理
						oldPlant.FarmRecordIdDeleteList = append(oldPlant.FarmRecordIdDeleteList, farmRecordId)
						for i := 0; i < len(oldPlant.FarmRecordIdList); i++ {
							if oldPlant.FarmRecordIdList[i] == farmRecordId {
								oldPlant.FarmRecordIdList = append(oldPlant.FarmRecordIdList[:i], oldPlant.FarmRecordIdList[i+1:]...)
								i--
							}
						}
						data, err := json.Marshal(oldPlant)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						info := []string{"PutPlant", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
						return
					}
					if okSectionId {
						oldPlant.SectionId = reflect.ValueOf(plantData.Old[gIndex]["section_id"]).String()
					}
					for i, v := range oldPlant.FarmRecordIdList {
						if v == reflect.ValueOf(plantData.Old[gIndex]["id"]).String() {
							oldPlant.FarmRecordIdList[i] = farmRecordId
						}
					}
					if okSectionId || okFarmRecordId {
						data, err := json.Marshal(oldPlant)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						info := []string{"PutPlant", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
					}
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()
			}
		}
		if plantData.Type == "INSERT" {
			type SectionPlant struct {
				SectionId        string
				SectionCropsList []string
			}
			sectionCropsMap := make(map[string]SectionPlant)
			for _, plant := range plantData.Data {
				// 农事记录Id
				farmRecordId := reflect.ValueOf(plant["id"]).String()
				// 地块Id
				SecId := reflect.ValueOf(plant["section_id"]).String()
				// 种植Id
				sectionCropsId := reflect.ValueOf(plant["section_crops_id"]).String()
				if v, ok := sectionCropsMap[sectionCropsId]; ok {
					if !in(farmRecordId, v.SectionCropsList) {
						v.SectionCropsList = append(v.SectionCropsList, farmRecordId)
						sectionCropsMap[sectionCropsId] = SectionPlant{SectionId: SecId, SectionCropsList: v.SectionCropsList}
					}
				} else {
					sectionCropsMap[sectionCropsId] = SectionPlant{SectionId: SecId, SectionCropsList: []string{farmRecordId}}
				}
			}
			for k, v := range sectionCropsMap {
				wg.Add(1)
				go func() {
					// 查询种植信息
					dd, err := sdk.Query(k, define.QueryPlant)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					// 没有种植信息就插入种植信息
					if len(dd) == 0 {
						l := &define.Plant{}
						l.SectionCropsId = k
						l.SectionId = v.SectionId
						l.FarmRecordIdList = append(l.FarmRecordIdList, v.SectionCropsList...)
						data, err := json.Marshal(l)
						if err != nil {
							wg.Done()
							return
						}
						// 更新plant上链数据
						info := []string{"PutPlant", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
					} else {
						// 存在种植信息，就更新农事记录列表
						old := &define.Plant{}
						err = json.Unmarshal(dd, old)
						if err != nil {
							wg.Done()
							return
						}
						newFlag := false
						for _, r := range v.SectionCropsList {
							if !in(r, old.FarmRecordIdList) {
								old.FarmRecordIdList = append(old.FarmRecordIdList, r)
								newFlag = true
							}
						}
						if newFlag {
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新plant上链数据
							info := []string{"PutPlant", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
						}
					}
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()
			}
		}
		wg.Wait()
		cancel()
	case "ag_farming_record_detail":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		// 更新农事操作码、使用物品和用量
		for gIndex, plant := range plantData.Data {
			wg.Add(1)
			l := &define.FarmRecord{}
			// 农事记录Id
			l.OperateId = reflect.ValueOf(plant["record_id"]).String()
			// 农事操作码
			/*l.OperationCode = reflect.ValueOf(plant["operation_code"]).String()
			if len(l.OperationCode) == 0 {
				err := fmt.Errorf("农事记录详情操作码不能为空")
				return "", err
			}*/
			r := &define.RecordDetail{}
			// 使用物品，从ag_dict中找映射数据
			r.Id = reflect.ValueOf(plant["id"]).String()
			r.Type = reflect.ValueOf(plant["type"]).String()
			r.UseType = reflect.ValueOf(plant["use_type"]).String()
			r.UseGood = reflect.ValueOf(plant["use_good"]).String()
			r.TypeCode = reflect.ValueOf(plant["type_code"]).String()
			r.OrderNo = reflect.ValueOf(plant["order_no"]).String()
			r.StartValue, _ = strconv.ParseFloat(reflect.ValueOf(plant["start_value"]).String(), 64)
			r.EndValue, _ = strconv.ParseFloat(reflect.ValueOf(plant["end_value"]).String(), 64)
			isActive := reflect.ValueOf(plant["is_active"]).String()

			switch plantData.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(l.OperateId, define.QueryFarmRecord)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.FarmRecord{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					// 删除对应数据
					for i := 0; i < len(old.RecordDetailList); i++ {
						if old.RecordDetailList[i].Id == r.Id {
							old.RecordDetailList = append(old.RecordDetailList[:i], old.RecordDetailList[i+1:]...)
							i--
						}
					}
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutFarmRecord", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					/*p, err := sdk.Query(old.SectionCropsId, define.QueryPlant)
					if err != nil || len(p) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					pp := &define.Plant{}
					err = json.Unmarshal(p, pp)
					if err != nil {
						wg.Done()
						return
					}
					// 删除对应数据
					for i := 0; i < len(pp.FarmRecordIdDeleteList); i++ {
						if pp.FarmRecordIdDeleteList[i] == old.OperateId {
							stageCode := strings.Split(old.StageCode, "_")
							stageCodeLen := len(stageCode)
							if stageCodeLen < 5 {
								// stage_rice_16_34_prepare
								continue
								//wg.Done()
								//return
							}
							switch stageCode[stageCodeLen-1] {
							case "prepare": // 大田准备
								for _, n := range pp.FieldPreparation {
									if n.OperateId == old.OperateId {
										for i := 0; i < len(n.RecordDetailList); i++ {
											if n.RecordDetailList[i].Id == r.Id {
												n.RecordDetailList = append(n.RecordDetailList[:i], n.RecordDetailList[i+1:]...)
												i--
											}
										}
									}
								}
							case "sowing": // 移栽
								for _, n := range pp.Transplant {
									if n.OperateId == old.OperateId {
										for i := 0; i < len(n.RecordDetailList); i++ {
											if n.RecordDetailList[i].Id == r.Id {
												n.RecordDetailList = append(n.RecordDetailList[:i], n.RecordDetailList[i+1:]...)
												i--
											}
										}
									}
								}
							case "rejuvenation": // 返青期
								for _, n := range pp.ReGreening {
									if n.OperateId == old.OperateId {
										for i := 0; i < len(n.RecordDetailList); i++ {
											if n.RecordDetailList[i].Id == r.Id {
												n.RecordDetailList = append(n.RecordDetailList[:i], n.RecordDetailList[i+1:]...)
												i--
											}
										}
									}
								}
							case "tillering": // 分蘖期
								for _, n := range pp.Tillering {
									if n.OperateId == old.OperateId {
										for i := 0; i < len(n.RecordDetailList); i++ {
											if n.RecordDetailList[i].Id == r.Id {
												n.RecordDetailList = append(n.RecordDetailList[:i], n.RecordDetailList[i+1:]...)
												i--
											}
										}
									}
								}
							case "booting": // 孕穗期
								for _, n := range pp.Booting {
									if n.OperateId == old.OperateId {
										for i := 0; i < len(n.RecordDetailList); i++ {
											if n.RecordDetailList[i].Id == r.Id {
												n.RecordDetailList = append(n.RecordDetailList[:i], n.RecordDetailList[i+1:]...)
												i--
											}
										}
									}
								}
							case "heading": // 抽穗扬花
								for _, n := range pp.HeadingAndBlooming {
									if n.OperateId == old.OperateId {
										for i := 0; i < len(n.RecordDetailList); i++ {
											if n.RecordDetailList[i].Id == r.Id {
												n.RecordDetailList = append(n.RecordDetailList[:i], n.RecordDetailList[i+1:]...)
												i--
											}
										}
									}
								}
							case "filling": // 灌浆期
								for _, n := range pp.GrainFillingStage {
									if n.OperateId == old.OperateId {
										for i := 0; i < len(n.RecordDetailList); i++ {
											if n.RecordDetailList[i].Id == r.Id {
												n.RecordDetailList = append(n.RecordDetailList[:i], n.RecordDetailList[i+1:]...)
												i--
											}
										}
									}
								}
							case "harvest": // 成熟期收获
								for _, n := range pp.Maturity {
									if n.OperateId == old.OperateId {
										for i := 0; i < len(n.RecordDetailList); i++ {
											if n.RecordDetailList[i].Id == r.Id {
												n.RecordDetailList = append(n.RecordDetailList[:i], n.RecordDetailList[i+1:]...)
												i--
											}
										}
									}
								}
							}
							pp.FarmRecordIdDeleteList = append(pp.FarmRecordIdDeleteList[:i], pp.FarmRecordIdDeleteList[i+1:]...)
							i--
						}
					}
					pData, err := json.Marshal(pp)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					pInfo := []string{"PutPlant", string(pData)}
					res, err := gohfc.GetHandler().SyncInvoke(pInfo, "", "")
					if err != nil {
						wg.Done()
						return
					}*/
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				f := &define.FarmRecord{}
				d, err := sdk.Query(l.OperateId, define.QueryFarmRecord)
				if err != nil || len(d) == 0 {
					err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				err = json.Unmarshal(d, f)
				if err != nil {
					return "", err
				}

				// 更新farmRecord上链数据
				//f.OperationCode = l.OperationCode

				// 使用物品需要从字典中获取映射值
				switch r.TypeCode {
				case "pesticideType", "fertilizerName": // 农药或化肥
					pesticideFertilizer, err := sdk.Query("AgPesticideFertilizer", define.QueryPesticideFertilizer)
					if err != nil || len(pesticideFertilizer) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						return "", err
					}
					agPesticideFertilizerInfo := &define.AgPesticideFertilizerInfo{}
					err = json.Unmarshal(pesticideFertilizer, agPesticideFertilizerInfo)
					if err != nil {
						return "", err
					}
					for _, v := range agPesticideFertilizerInfo.AgPesticideFertilizerList {
						if v.TypeCode == r.TypeCode && v.Value == r.UseGood {
							r.UseType = v.Name
							r.Unit = v.Unit
						}
					}
				default:
					dict, err := sdk.Query("AgDict", define.QueryDict)
					if err != nil || len(dict) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						return "", err
					}
					agDictInfo := &define.AgDictInfo{}
					err = json.Unmarshal(dict, agDictInfo)
					if err != nil {
						return "", err
					}
					for _, v := range agDictInfo.AgDictList {
						useGood, _ := strconv.Atoi(r.UseGood)
						if v.TypeCode == r.TypeCode && v.Value == useGood {
							r.UseType = v.Name
						}
					}
				}
				existFlag := false
				for _, n := range f.RecordDetailList {
					if n.Id == r.Id {
						existFlag = true
					}
				}
				if !existFlag {
					f.RecordDetailList = append(f.RecordDetailList, *r)
				}
				fData, err := json.Marshal(f)
				if err != nil {
					return "", err
				}
				// 更新plant上链数据
				fmt.Printf("")
				fInfo := []string{"PutFarmRecord", string(fData)}
				res, err := gohfc.GetHandler().SyncInvoke(fInfo, "", "")
				if err != nil {
					return "", err
				}
				response = res.TxID
				wg.Done()

				// 更新种植数据
				// 查询种植信息
				/*dd, err := sdk.Query(f.SectionCropsId, define.QueryPlant)
				if err != nil || len(dd) == 0 {
					err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				// 更新农事记录列表
				old := &define.Plant{}
				err = json.Unmarshal(dd, old)
				if err != nil {
					return "", err
				}
				// 根据操作码分类农事记录到8个过程中
				stageCode := strings.Split(f.StageCode, "_")
				stageCodeLen := len(stageCode)
				if stageCodeLen < 5 {
					// stage_rice_16_34_prepare
					return "", err
				}
				// 使用物品需要从字典中获取映射值
				switch r.TypeCode {
				case "pesticideType", "fertilizerName": // 农药或化肥
					pesticideFertilizer, err := sdk.Query("AgPesticideFertilizer", define.QueryPesticideFertilizer)
					if err != nil || len(pesticideFertilizer) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						return "", err
					}
					agPesticideFertilizerInfo := &define.AgPesticideFertilizerInfo{}
					err = json.Unmarshal(pesticideFertilizer, agPesticideFertilizerInfo)
					if err != nil {
						return "", err
					}
					for _, v := range agPesticideFertilizerInfo.AgPesticideFertilizerList {
						if v.TypeCode == r.TypeCode && v.Value == r.UseGood {
							r.UseType = v.Name
							r.Unit = v.Unit
						}
					}
				default:
					dict, err := sdk.Query("AgDict", define.QueryDict)
					if err != nil || len(dict) == 0 {
						err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						return "", err
					}
					agDictInfo := &define.AgDictInfo{}
					err = json.Unmarshal(dict, agDictInfo)
					if err != nil {
						return "", err
					}
					for _, v := range agDictInfo.AgDictList {
						useGood, _ := strconv.Atoi(r.UseGood)
						if v.TypeCode == r.TypeCode && v.Value == useGood {
							r.UseType = v.Name
						}
					}
				}

				existFlag := false
				for _, n := range f.RecordDetailList {
					if n.Id == r.Id {
						existFlag = true
					}
				}
				if !existFlag {
					f.RecordDetailList = append(f.RecordDetailList, *r)
				}
				//fmt.Printf("f.RecordDetailList=%v\n", f.RecordDetailList)
				op := define.OpDetail{
					OperateId:     f.OperateId,
					StageCode: 	   f.StageCode,
					OperatorName:  f.OperatorName,
					HandleDate:    f.HandleDate,
					RecordDetailList:       f.RecordDetailList,
					ImageUrl:f.ImageUrl,
					SubImageUrl:f.SubImageUrl,
				}
				switch stageCode[stageCodeLen-1] {
				case "prepare": // 大田准备
					existFlag := false
					for i, v := range old.FieldPreparation {
						if v.OperateId == f.OperateId {
							old.FieldPreparation[i] = op
							existFlag = true
						}
					}
					if !existFlag {
						old.FieldPreparation = append(old.FieldPreparation, op)
					}
				case "sowing": // 移栽
					existFlag := false
					for i, v := range old.Transplant {
						if v.OperateId == f.OperateId {
							old.Transplant[i] = op
							existFlag = true
						}
					}
					if !existFlag {
						old.Transplant = append(old.Transplant, op)
					}
				case "rejuvenation": // 返青期
					existFlag := false
					for i, v := range old.ReGreening {
						if v.OperateId == f.OperateId {
							old.ReGreening[i] = op
							existFlag = true
						}
					}
					if !existFlag {
						old.ReGreening = append(old.ReGreening, op)
					}
				case "tillering": // 分蘖期
					existFlag := false
					for i, v := range old.Tillering {
						if v.OperateId == f.OperateId {
							old.Tillering[i] = op
							existFlag = true
						}
					}
					if !existFlag {
						old.Tillering = append(old.Tillering, op)
					}
				case "booting": // 孕穗期
					existFlag := false
					for i, v := range old.Booting {
						if v.OperateId == f.OperateId {
							old.Booting[i] = op
							existFlag = true
						}
					}
					if !existFlag {
						old.Booting = append(old.Booting, op)
					}
				case "heading": // 抽穗扬花
					existFlag := false
					for i, v := range old.HeadingAndBlooming {
						if v.OperateId == f.OperateId {
							old.HeadingAndBlooming[i] = op
							existFlag = true
						}
					}
					if !existFlag {
						old.HeadingAndBlooming = append(old.HeadingAndBlooming, op)
					}
				case "filling": // 灌浆期
					existFlag := false
					for i, v := range old.GrainFillingStage {
						if v.OperateId == f.OperateId {
							old.GrainFillingStage[i] = op
							existFlag = true
						}
					}
					if !existFlag {
						old.GrainFillingStage = append(old.GrainFillingStage, op)
					}
				case "harvest": // 成熟期收获
					existFlag := false
					for i, v := range old.Maturity {
						if v.OperateId == f.OperateId {
							old.Maturity[i] = op
							existFlag = true
						}
					}
					if !existFlag {
						old.Maturity = append(old.Maturity, op)
					}
				}
				for i := 0; i < len(old.FarmRecordIdList); i++ {
					if old.FarmRecordIdList[i] == f.OperateId {
						old.FarmRecordIdList = append(old.FarmRecordIdList[:i], old.FarmRecordIdList[i+1:]...)
						i--
					}
				}
				data, err := json.Marshal(old)
				if err != nil {
					return "", err
				}
				// 更新plant上链数据
				info := []string{"PutPlant", string(data)}
				res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
				if err != nil {
					return "", err
				}
				response = res.TxID*/

			case "UPDATE":
				_, okOperateId := plantData.Old[gIndex]["record_id"]
				//_, okOperationCode := plantData.Old[gIndex]["operation_code"]
				_, okRecordDetailId := plantData.Old[gIndex]["id"]
				_, okType := plantData.Old[gIndex]["type"]
				_, okUseType := plantData.Old[gIndex]["use_type"]
				_, okUseGood := plantData.Old[gIndex]["use_good"]
				_, okTypeCode := plantData.Old[gIndex]["type_code"]
				_, okStartValue := plantData.Old[gIndex]["start_value"]
				_, okEndValue := plantData.Old[gIndex]["end_value"]
				_, okIsActive := plantData.Old[gIndex]["is_active"]
				go func() {
					operateId := ""
					if okOperateId {
						operateId = reflect.ValueOf(plantData.Old[gIndex]["record_id"]).String()
					} else {
						operateId = l.OperateId
					}
					d, err := sdk.Query(operateId, define.QueryFarmRecord)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.FarmRecord{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					if okIsActive && isActive == "N" {
						// 删除对应数据
						for i := 0; i < len(old.RecordDetailList); i++ {
							if old.RecordDetailList[i].Id == r.Id {
								old.RecordDetailList = append(old.RecordDetailList[:i], old.RecordDetailList[i+1:]...)
								i--
							}
						}
						data, err := json.Marshal(old)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						info := []string{"PutFarmRecord", string(data)}
						_, err = gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
					} else {
						// 更新数据
						if okOperateId {
							old.OperateId = l.OperateId
						}
						/*if okOperationCode {
							old.OperationCode = l.OperationCode
						}*/
						for i, recordDetailId := range old.RecordDetailList {
							if recordDetailId.Id == r.Id {
								if okRecordDetailId {
									old.RecordDetailList[i].Id = r.Id
								}
								if okUseType {
									old.RecordDetailList[i].UseType = r.UseType
								}
								if okUseGood {
									old.RecordDetailList[i].UseGood = r.UseGood
									dict, err := sdk.Query("AgDict", define.QueryDict)
									if err != nil || len(dict) == 0 {
										//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
										wg.Done()
										return
									}
									agDict := &define.AgDictInfo{}
									err = json.Unmarshal(dict, agDict)
									if err != nil {
										wg.Done()
										return
									}
									for _, v := range agDict.AgDictList {
										for j, q := range old.RecordDetailList {
											useGood, _ := strconv.Atoi(r.UseGood)
											if v.TypeCode == q.TypeCode && v.Value == useGood {
												old.RecordDetailList[j].UseType = v.Name
											}
										}
									}
								}
								if okType {
									old.RecordDetailList[i].Type = r.Type
								}
								if okTypeCode {
									old.RecordDetailList[i].TypeCode = r.TypeCode
								}
								if okStartValue {
									old.RecordDetailList[i].StartValue = r.StartValue
								}
								if okEndValue {
									old.RecordDetailList[i].EndValue = r.EndValue
								}
							}
						}

						/*if okOperateId || okOperationCode || okUseType || okUseGood || okTypeCode || okStartValue ||
						okEndValue {*/
						if okRecordDetailId || okOperateId || okType || okUseType || okUseGood || okTypeCode || okStartValue ||
							okEndValue {
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutFarmRecord", string(data)}
							_, err = gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
						}
					}

					/*p, err := sdk.Query(old.SectionCropsId, define.QueryPlant)
					if err != nil || len(p) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					pp := &define.Plant{}
					err = json.Unmarshal(p, pp)
					if err != nil {
						wg.Done()
						return
					}
					if okIsActive && isActive == "N" {
						// 删除对应数据
						for i := 0; i < len(pp.FarmRecordIdDeleteList); i++ {
						//for i, v := range pp.FarmRecordIdDeleteList {
							if pp.FarmRecordIdDeleteList[i] == old.OperateId {
								stageCode := strings.Split(old.StageCode, "_")
								stageCodeLen := len(stageCode)
								if stageCodeLen < 5 {
									// stage_rice_16_34_prepare
									wg.Done()
									return
								}
								switch stageCode[stageCodeLen-1] {
								case "prepare": // 大田准备
									for _, n := range pp.FieldPreparation {
										if n.OperateId == old.OperateId {
											for i := 0; i < len(n.RecordDetailList); i++ {
												if n.RecordDetailList[i].Id == r.Id  {
													n.RecordDetailList = append(n.RecordDetailList[:i],n.RecordDetailList[i+1:]...)
													i--
												}
											}
										}
									}
								case "sowing": // 移栽
									for _, n := range pp.Transplant {
										if n.OperateId == old.OperateId {
											for i := 0; i < len(n.RecordDetailList); i++ {
												if n.RecordDetailList[i].Id == r.Id  {
													n.RecordDetailList = append(n.RecordDetailList[:i],n.RecordDetailList[i+1:]...)
													i--
												}
											}
										}
									}
								case "rejuvenation": // 返青期
									for _, n := range pp.ReGreening {
										if n.OperateId == old.OperateId {
											for i := 0; i < len(n.RecordDetailList); i++ {
												if n.RecordDetailList[i].Id == r.Id  {
													n.RecordDetailList = append(n.RecordDetailList[:i],n.RecordDetailList[i+1:]...)
													i--
												}
											}
										}
									}
								case "tillering": // 分蘖期
									for _, n := range pp.Tillering {
										if n.OperateId == old.OperateId {
											for i := 0; i < len(n.RecordDetailList); i++ {
												if n.RecordDetailList[i].Id == r.Id  {
													n.RecordDetailList = append(n.RecordDetailList[:i],n.RecordDetailList[i+1:]...)
													i--
												}
											}
										}
									}
								case "booting": // 孕穗期
									for _, n := range pp.Booting {
										if n.OperateId == old.OperateId {
											for i := 0; i < len(n.RecordDetailList); i++ {
												if n.RecordDetailList[i].Id == r.Id  {
													n.RecordDetailList = append(n.RecordDetailList[:i],n.RecordDetailList[i+1:]...)
													i--
												}
											}
										}
									}
								case "heading": // 抽穗扬花
									for _, n := range pp.HeadingAndBlooming {
										if n.OperateId == old.OperateId {
											for i := 0; i < len(n.RecordDetailList); i++ {
												if n.RecordDetailList[i].Id == r.Id  {
													n.RecordDetailList = append(n.RecordDetailList[:i],n.RecordDetailList[i+1:]...)
													i--
												}
											}
										}
									}
								case "filling": // 灌浆期
									for _, n := range pp.GrainFillingStage {
										if n.OperateId == old.OperateId {
											for i := 0; i < len(n.RecordDetailList); i++ {
												if n.RecordDetailList[i].Id == r.Id  {
													n.RecordDetailList = append(n.RecordDetailList[:i],n.RecordDetailList[i+1:]...)
													i--
												}
											}
										}
									}
								case "harvest": // 成熟期收获
									for _, n := range pp.Maturity {
										if n.OperateId == old.OperateId {
											for i := 0; i < len(n.RecordDetailList); i++ {
												if n.RecordDetailList[i].Id == r.Id  {
													n.RecordDetailList = append(n.RecordDetailList[:i],n.RecordDetailList[i+1:]...)
													i--
												}
											}
										}
									}
								}
								pp.FarmRecordIdDeleteList = append(pp.FarmRecordIdDeleteList[:i], pp.FarmRecordIdDeleteList[i+1:]...)
								i--
							}
						}
						pData, err := json.Marshal(pp)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						pInfo := []string{"PutPlant", string(pData)}
						res, err := gohfc.GetHandler().SyncInvoke(pInfo, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
						wg.Done()
						return
					}
					// 更新对应数据
					// 使用物品需要从字典中获取映射值
					switch r.TypeCode {
					case "pesticideType", "fertilizerName": // 农药或化肥
						pesticideFertilizer, err := sdk.Query("AgPesticideFertilizer", define.QueryPesticideFertilizer)
						if err != nil || len(pesticideFertilizer) == 0 {
							//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
							wg.Done()
							return
						}
						agPesticideFertilizer := &define.AgPesticideFertilizerInfo{}
						err = json.Unmarshal(pesticideFertilizer, agPesticideFertilizer)
						if err != nil {
							fmt.Println("json Unmarshal pesticideFertilizer error")
							wg.Done()
							return
						}
						for _, v := range agPesticideFertilizer.AgPesticideFertilizerList {
							for j, q := range old.RecordDetailList {
								if v.TypeCode == q.TypeCode && v.Value == q.UseGood {
									old.RecordDetailList[j].UseType = v.Name
									old.RecordDetailList[j].Unit = v.Unit
								}
							}
						}
					default:
						dict, err := sdk.Query("AgDict", define.QueryDict)
						if err != nil || len(dict) == 0 {
							//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
							wg.Done()
							return
						}
						agDict := &define.AgDictInfo{}
						err = json.Unmarshal(dict, agDict)
						if err != nil {
							wg.Done()
							return
						}
						for _, v := range agDict.AgDictList {
							for j, q := range old.RecordDetailList {
								useGood, _ := strconv.Atoi(r.UseGood)
								if v.TypeCode == q.TypeCode && v.Value == useGood {
									old.RecordDetailList[j].UseType = v.Name
								}
							}
						}
					}
					op := define.OpDetail{
						OperateId:     old.OperateId,
						StageCode:     old.StageCode,
						OperatorName:  old.OperatorName,
						HandleDate:    old.HandleDate,
						RecordDetailList: old.RecordDetailList,
						ImageUrl:old.ImageUrl,
						SubImageUrl:old.SubImageUrl,
					}
					stageCode := strings.Split(old.StageCode, "_")
					stageCodeLen := len(stageCode)
					//fmt.Printf("old.OperationCode=%v\n", old.OperationCode)
					if stageCodeLen < 5 {
						// stage_rice_16_34_prepare
						wg.Done()
						return
					}
					//fmt.Printf("len(stateCode)=%d, stateCode=%v\n", len(stateCode), stateCode)
					switch stageCode[stageCodeLen-1] {
					case "prepare": // 大田准备
						for m, n := range pp.FieldPreparation {
							if n.OperateId == old.OperateId {
								pp.FieldPreparation[m] = op
							}
						}
					case "sowing": // 移栽
						for m, n := range pp.Transplant {
							if n.OperateId == old.OperateId {
								pp.Transplant[m] = op
							}
						}
					case "rejuvenation": // 返青期
						for m, n := range pp.ReGreening {
							if n.OperateId == old.OperateId {
								pp.ReGreening[m] = op
							}
						}
					case "tillering": // 分蘖期
						for m, n := range pp.Tillering {
							if n.OperateId == old.OperateId {
								pp.Tillering[m] = op
							}
						}
					case "booting": // 孕穗期
						for m, n := range pp.Booting {
							if n.OperateId == old.OperateId {
								pp.Booting[m] = op
							}
						}
					case "heading": // 抽穗扬花
						for m, n := range pp.HeadingAndBlooming {
							if n.OperateId == old.OperateId {
								pp.HeadingAndBlooming[m] = op
							}
						}
					case "filling": // 灌浆期
						for m, n := range pp.GrainFillingStage {
							if n.OperateId == old.OperateId {
								pp.GrainFillingStage[m] = op
							}
						}
					case "harvest": // 成熟期收获
						for m, n := range pp.Maturity {
							if n.OperateId == old.OperateId {
								pp.Maturity[m] = op
							}
						}
					}
					pData, err := json.Marshal(pp)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					pInfo := []string{"PutPlant", string(pData)}
					res, err := gohfc.GetHandler().SyncInvoke(pInfo, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID*/
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()
			}
		}
		wg.Wait()
		cancel()
	case "sys_attachment":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		// 更新图片，subpath为压缩图、path为原始图
		for gIndex, plant := range plantData.Data {
			wg.Add(1)
			l := &define.FarmRecord{}
			// 农事记录Id
			l.OperateId = reflect.ValueOf(plant["pid"]).String()
			// 图片类型,农事记录file_class=farmingRecord,收割file_class=reapBatchRecord
			fileClass := reflect.ValueOf(plant["file_class"]).String()
			// 压缩图片路径
			subPath := reflect.ValueOf(plant["sub_path"]).String()
			// 正常图片路径
			//path := reflect.ValueOf(plant["path"]).String()
			path := reflect.ValueOf(plant["id"]).String()
			isActive := reflect.ValueOf(plant["is_active"]).String()

			switch plantData.Type {
			// 删除农事记录TODO
			case "DELETE":
				go func() {
					d, err := sdk.Query(l.OperateId, define.QueryFarmRecord)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					f := &define.FarmRecord{}
					err = json.Unmarshal(d, f)
					if err != nil {
						wg.Done()
						return
					}
					// 更新farmRecord上链数据
					if fileClass == "farmingRecord" {
						for i := 0; i < len(f.SubImageUrl); i++ {
							if f.SubImageUrl[i] == subPath {
								f.SubImageUrl = append(f.SubImageUrl[:i], f.SubImageUrl[i+1:]...)
								i--
							}
						}
						for i := 0; i < len(f.ImageUrl); i++ {
							if f.ImageUrl[i] == path {
								f.ImageUrl = append(f.ImageUrl[:i], f.ImageUrl[i+1:]...)
								i--
							}
						}
					}
					fData, err := json.Marshal(f)
					if err != nil {
						wg.Done()
						return
					}
					// 更新plant上链数据
					fInfo := []string{"PutFarmRecord", string(fData)}
					res, err := gohfc.GetHandler().SyncInvoke(fInfo, "", "")
					if err != nil {
						wg.Done()
						return
					}
					// 更新种植数据
					// 查询种植信息
					/*dd, err := sdk.Query(f.SectionCropsId, define.QueryPlant)
					if err != nil || len(dd) == 0{
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					// 更新农事记录列表
					old := &define.Plant{}
					err = json.Unmarshal(dd, old)
					if err != nil {
						wg.Done()
						return
					}
					// 根据操作码分类农事记录到8个过程中
					stageCode := strings.Split(f.StageCode,"_")
					stageCodeLen := len(stageCode)
					if stageCodeLen < 5 {
						// stage_rice_16_34_prepare
						wg.Done()
						return
					}
					op := define.OpDetail{
						OperateId:f.OperateId,
						StageCode:f.StageCode,
						OperatorName:f.OperatorName,
						HandleDate:f.HandleDate,
						RecordDetailList:       f.RecordDetailList,
						ImageUrl:f.ImageUrl,
						SubImageUrl:f.SubImageUrl,
					}
					switch stageCode[stageCodeLen-1] {
					case "prepare": // 大田准备
						for m, n := range old.FieldPreparation {
							if n.OperateId== f.OperateId {
								old.FieldPreparation[m] = op
							}
						}
					case "sowing": // 移栽
						for m, n := range old.Transplant {
							if n.OperateId== f.OperateId {
								old.Transplant[m] = op
							}
						}
					case "rejuvenation": // 返青期
						for m, n := range old.ReGreening {
							if n.OperateId== f.OperateId {
								old.ReGreening[m] = op
							}
						}
					case "tillering": // 分蘖期
						for m, n := range old.Tillering {
							if n.OperateId== f.OperateId {
								old.Tillering[m] = op
							}
						}
					case "booting": // 孕穗期
						for m, n := range old.Booting {
							if n.OperateId== f.OperateId {
								old.Booting[m] = op
							}
						}
					case "heading": // 抽穗扬花
						for m, n := range old.HeadingAndBlooming {
							if n.OperateId== f.OperateId {
								old.HeadingAndBlooming[m] = op
							}
						}
					case "filling": // 灌浆期
						for m, n := range old.GrainFillingStage {
							if n.OperateId== f.OperateId {
								old.GrainFillingStage[m] = op
							}
						}
					case "harvest": // 成熟期收获
						for m, n := range old.Maturity {
							if n.OperateId== f.OperateId {
								old.Maturity[m] = op
							}
						}
					}
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新plant上链数据
					info := []string{"PutPlant", string(data)}
					res, err = gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}*/
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				d, err := sdk.Query(l.OperateId, define.QueryFarmRecord)
				if err != nil || len(d) == 0 {
					err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				f := &define.FarmRecord{}
				err = json.Unmarshal(d, f)
				if err != nil {
					return "", err
				}
				go func() {
					// 更新farmRecord上链数据
					if fileClass == "farmingRecord" {
						if !in(subPath, f.SubImageUrl) {
							f.SubImageUrl = append(f.SubImageUrl, subPath)
						}
						if !in(path, f.ImageUrl) {
							f.ImageUrl = append(f.ImageUrl, path)
						}
					}
					fData, err := json.Marshal(f)
					if err != nil {
						wg.Done()
						return
					}
					// 更新plant上链数据
					fInfo := []string{"PutFarmRecord", string(fData)}
					_, err = gohfc.GetHandler().SyncInvoke(fInfo, "", "")
					if err != nil {
						wg.Done()
						return
					}

					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()
				// 更新种植数据
				// 查询种植信息
				/*dd, err := sdk.Query(f.SectionCropsId, define.QueryPlant)
				if err != nil || len(dd) == 0{
					err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				// 更新农事记录列表
				old := &define.Plant{}
				err = json.Unmarshal(dd, old)
				if err != nil {
					return "", err
				}
				// 根据操作码分类农事记录到8个过程中
				stageCode := strings.Split(f.StageCode,"_")
				stageCodeLen := len(stageCode)
				if stageCodeLen < 5 {
					// stage_rice_16_34_prepare
					return "", err
				}
				op := define.OpDetail{
					OperateId:f.OperateId,
					StageCode:f.StageCode,
					OperatorName:f.OperatorName,
					HandleDate:f.HandleDate,
					RecordDetailList:       f.RecordDetailList,
					ImageUrl:f.ImageUrl,
					SubImageUrl:f.SubImageUrl,
				}
				switch stageCode[stageCodeLen-1] {
				case "prepare": // 大田准备
					for m, n := range old.FieldPreparation {
						if n.OperateId== f.OperateId {
							old.FieldPreparation[m] = op
						}
					}
				case "sowing": // 移栽
					for m, n := range old.Transplant {
						if n.OperateId== f.OperateId {
							old.Transplant[m] = op
						}
					}
				case "rejuvenation": // 返青期
					for m, n := range old.ReGreening {
						if n.OperateId== f.OperateId {
							old.ReGreening[m] = op
						}
					}
				case "tillering": // 分蘖期
					for m, n := range old.Tillering {
						if n.OperateId== f.OperateId {
							old.Tillering[m] = op
						}
					}
				case "booting": // 孕穗期
					for m, n := range old.Booting {
						if n.OperateId== f.OperateId {
							old.Booting[m] = op
						}
					}
				case "heading": // 抽穗扬花
					for m, n := range old.HeadingAndBlooming {
						if n.OperateId== f.OperateId {
							old.HeadingAndBlooming[m] = op
						}
					}
				case "filling": // 灌浆期
					for m, n := range old.GrainFillingStage {
						if n.OperateId== f.OperateId {
							old.GrainFillingStage[m] = op
						}
					}
				case "harvest": // 成熟期收获
					for m, n := range old.Maturity {
						if n.OperateId== f.OperateId {
							old.Maturity[m] = op
						}
					}
				}
				data, err := json.Marshal(old)
				if err != nil {
					return "", err
				}
				// 更新plant上链数据
				info := []string{"PutPlant", string(data)}
				res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
				if err != nil {
					return "", err
				}
				response = res.TxID*/

			// 更新path
			case "UPDATE":
				_, okOperateId := plantData.Old[gIndex]["pid"]
				//_, okFileClass := oldPlantData["file_class"]
				_, okSubPath := plantData.Old[gIndex]["sub_path"]
				_, okPath := plantData.Old[gIndex]["id"]
				_, okIsActive := plantData.Old[gIndex]["is_active"]
				go func() {
					operateId := ""
					if okOperateId {
						operateId = reflect.ValueOf(plantData.Old[gIndex]["pid"]).String()
					} else {
						operateId = l.OperateId
					}
					d, err := sdk.Query(operateId, define.QueryFarmRecord)
					if err != nil {
						//err := fmt.Errorf("查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					if len(d) == 0 {
						r, err := sdk.Query(l.OperateId, define.QueryFarmRecord)
						if err != nil {
							//err := fmt.Errorf("查询链上数据失败：%v", err)
							wg.Done()
							return
						}
						if len(r) == 0 && isActive == "Y" {
							ff := &define.FarmRecord{}
							ff.OperateId = l.OperateId
							if fileClass == "farmingRecord" {
								ff.SubImageUrl = append(ff.SubImageUrl, subPath)
								ff.ImageUrl = append(ff.ImageUrl, path)
							}
							fData, err := json.Marshal(ff)
							if err != nil {
								wg.Done()
								return
							}
							// 更新plant上链数据
							fInfo := []string{"PutFarmRecord", string(fData)}
							res, err := gohfc.GetHandler().SyncInvoke(fInfo, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
							wg.Done()
							return
						}
						d = r
					}
					f := &define.FarmRecord{}
					err = json.Unmarshal(d, f)
					if err != nil {
						wg.Done()
						return
					}
					if okIsActive && isActive == "N" {
						// 更新farmRecord上链数据
						if fileClass == "farmingRecord" {
							for i := 0; i < len(f.SubImageUrl); i++ {
								if f.SubImageUrl[i] == subPath {
									f.SubImageUrl = append(f.SubImageUrl[:i], f.SubImageUrl[i+1:]...)
									i--
								}
							}
							for i := 0; i < len(f.ImageUrl); i++ {
								if f.ImageUrl[i] == path {
									f.ImageUrl = append(f.ImageUrl[:i], f.ImageUrl[i+1:]...)
									i--
								}
							}
						}
						fData, err := json.Marshal(f)
						if err != nil {
							wg.Done()
							return
						}
						// 更新plant上链数据
						fInfo := []string{"PutFarmRecord", string(fData)}
						res, err := gohfc.GetHandler().SyncInvoke(fInfo, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
					} else {
						// 更新farmRecord上链数据
						if okOperateId {
							f.OperateId = l.OperateId
						}
						if fileClass == "farmingRecord" {
							oldSubPath := ""
							if okSubPath {
								oldSubPath = reflect.ValueOf(plantData.Old[gIndex]["sub_path"]).String()
							}
							oldPath := ""
							if okPath {
								oldPath = reflect.ValueOf(plantData.Old[gIndex]["id"]).String()
							}
							if in(oldSubPath, f.SubImageUrl) {
								for i, v := range f.SubImageUrl {
									if v == oldSubPath {
										f.SubImageUrl[i] = subPath
									}
								}
							} else {
								if !in(subPath, f.SubImageUrl) {
									f.SubImageUrl = append(f.SubImageUrl, subPath)
								}
							}
							if in(oldPath, f.ImageUrl) {
								for i, v := range f.ImageUrl {
									if v == oldPath {
										f.ImageUrl[i] = path
									}
								}
							} else {
								if !in(path, f.ImageUrl) {
									f.ImageUrl = append(f.ImageUrl, path)
								}
							}
						}
						if okOperateId || okSubPath || okPath {
							fData, err := json.Marshal(f)
							if err != nil {
								wg.Done()
								return
							}
							// 更新plant上链数据
							fInfo := []string{"PutFarmRecord", string(fData)}
							res, err := gohfc.GetHandler().SyncInvoke(fInfo, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
						}
					}

					// 更新种植数据
					// 查询种植信息
					/*dd, err := sdk.Query(f.SectionCropsId, define.QueryPlant)
					if err != nil || len(dd) == 0{
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					// 更新农事记录列表
					old := &define.Plant{}
					err = json.Unmarshal(dd, old)
					if err != nil {
						wg.Done()
						return
					}
					if okIsActive && isActive == "N" {
						// 根据操作码分类农事记录到8个过程中
						stageCode := strings.Split(f.StageCode,"_")
						stageCodeLen := len(stageCode)
						if stageCodeLen < 5 {
							// stage_rice_16_34_prepare
							wg.Done()
							return
						}
						op := define.OpDetail{
							OperateId:f.OperateId,
							StageCode:f.StageCode,
							OperatorName:f.OperatorName,
							HandleDate:f.HandleDate,
							RecordDetailList:       f.RecordDetailList,
							ImageUrl:f.ImageUrl,
							SubImageUrl:f.SubImageUrl,
						}
						switch stageCode[stageCodeLen-1] {
						case "prepare": // 大田准备
							for m, n := range old.FieldPreparation {
								if n.OperateId== f.OperateId {
									old.FieldPreparation[m] = op
								}
							}
						case "sowing": // 移栽
							for m, n := range old.Transplant {
								if n.OperateId== f.OperateId {
									old.Transplant[m] = op
								}
							}
						case "rejuvenation": // 返青期
							for m, n := range old.ReGreening {
								if n.OperateId== f.OperateId {
									old.ReGreening[m] = op
								}
							}
						case "tillering": // 分蘖期
							for m, n := range old.Tillering {
								if n.OperateId== f.OperateId {
									old.Tillering[m] = op
								}
							}
						case "booting": // 孕穗期
							for m, n := range old.Booting {
								if n.OperateId== f.OperateId {
									old.Booting[m] = op
								}
							}
						case "heading": // 抽穗扬花
							for m, n := range old.HeadingAndBlooming {
								if n.OperateId== f.OperateId {
									old.HeadingAndBlooming[m] = op
								}
							}
						case "filling": // 灌浆期
							for m, n := range old.GrainFillingStage {
								if n.OperateId== f.OperateId {
									old.GrainFillingStage[m] = op
								}
							}
						case "harvest": // 成熟期收获
							for m, n := range old.Maturity {
								if n.OperateId== f.OperateId {
									old.Maturity[m] = op
								}
							}
						}
						data, err := json.Marshal(old)
						if err != nil {
							wg.Done()
							return
						}
						// 更新plant上链数据
						info := []string{"PutPlant", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
						wg.Done()
						return
					}
					// 根据操作码分类农事记录到8个过程中
					stageCode := strings.Split(f.StageCode,"_")
					stageCodeLen := len(stageCode)
					if stageCodeLen < 5 {
						// stage_rice_16_34_prepare
						wg.Done()
						return
					}
					op := define.OpDetail{
						OperateId:f.OperateId,
						StageCode:f.StageCode,
						OperatorName:f.OperatorName,
						HandleDate:f.HandleDate,
						RecordDetailList:       f.RecordDetailList,
						ImageUrl:f.ImageUrl,
						SubImageUrl:f.SubImageUrl,
					}
					switch stageCode[stageCodeLen-1] {
					case "prepare": // 大田准备
						for m, n := range old.FieldPreparation {
							if n.OperateId== f.OperateId {
								old.FieldPreparation[m] = op
							}
						}
					case "sowing": // 移栽
						for m, n := range old.Transplant {
							if n.OperateId== f.OperateId {
								old.Transplant[m] = op
							}
						}
					case "rejuvenation": // 返青期
						for m, n := range old.ReGreening {
							if n.OperateId== f.OperateId {
								old.ReGreening[m] = op
							}
						}
					case "tillering": // 分蘖期
						for m, n := range old.Tillering {
							if n.OperateId== f.OperateId {
								old.Tillering[m] = op
							}
						}
					case "booting": // 孕穗期
						for m, n := range old.Booting {
							if n.OperateId== f.OperateId {
								old.Booting[m] = op
							}
						}
					case "heading": // 抽穗扬花
						for m, n := range old.HeadingAndBlooming {
							if n.OperateId== f.OperateId {
								old.HeadingAndBlooming[m] = op
							}
						}
					case "filling": // 灌浆期
						for m, n := range old.GrainFillingStage {
							if n.OperateId== f.OperateId {
								old.GrainFillingStage[m] = op
							}
						}
					case "harvest": // 成熟期收获
						for m, n := range old.Maturity {
							if n.OperateId== f.OperateId {
								old.Maturity[m] = op
							}
						}
					}
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新plant上链数据
					info := []string{"PutPlant", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID*/
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()
			}
		}
		wg.Wait()
		cancel()
	}
	return response, nil
}

// 收割上链信息
// key: 收割批次ID
// value：收割批次ID、种植ID列表（按地块面积排序）、交易ID（作为农事相关区块链证书）
func HandleHarvest(harvestData apiDefine.KafkaMessage) (string, error) {
	response := ""
	switch harvestData.Table {
	case "ag_reap_batch":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		// 收割批次id（ag_reap_batch.id）、收割时间、Txid
		for gIndex, harvest := range harvestData.Data {
			wg.Add(1)
			h := &define.Harvest{}
			h.HarvestId = reflect.ValueOf(harvest["id"]).String()
			h.ReapCreateDate, _ = time.Parse("2006-01-02 15:04:05", reflect.ValueOf(harvest["handle_date"]).String())
			isActive := reflect.ValueOf(harvest["is_active"]).String()
			switch harvestData.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(h.HarvestId, define.QueryHarvest)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.Harvest{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					// 删除对应数据
					old.Status = 1
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutHarvest", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				go func() {
					d, err := sdk.Query(h.HarvestId, define.QueryHarvest)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					if len(d) > 0 {
						old := &define.Harvest{}
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						if old.Status == 0 {
							fmt.Println("数据数据已存在")
							wg.Done()
							return
						}
					}
					data, err := json.Marshal(h)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutHarvest", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				_, okHarvestId := harvestData.Old[gIndex]["id"]
				_, okReapCreateDate := harvestData.Old[gIndex]["create_date"]
				_, okIsActive := harvestData.Old[gIndex]["is_active"]
				go func() {
					d, err := sdk.Query(h.HarvestId, define.QueryHarvest)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.Harvest{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					if okIsActive && isActive == "N" {
						// 删除对应数据
						old.Status = 1
						data, err := json.Marshal(old)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						info := []string{"PutHarvest", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
						wg.Done()
						return
					}
					if okHarvestId {
						old.HarvestId = h.HarvestId
					}
					if okReapCreateDate {
						old.ReapCreateDate = h.ReapCreateDate
					}
					if okHarvestId || okReapCreateDate {
						data, err := json.Marshal(old)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						info := []string{"PutHarvest", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
					}
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()
			}
		}
		wg.Wait()
		cancel()
	case "ag_section_crops_reap":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		// 种植id
		for gIndex, harvest := range harvestData.Data {
			wg.Add(1)
			h := &define.Harvest{}
			h.HarvestId = reflect.ValueOf(harvest["reap_batch_id"]).String()
			sectionCropsId := reflect.ValueOf(harvest["section_crops_id"]).String()
			isActive := reflect.ValueOf(harvest["is_active"]).String()
			switch harvestData.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(h.HarvestId, define.QueryHarvest)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.Harvest{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					// 删除对应数据
					for i := 0; i < len(old.PlantIdList); i++ {
						if old.PlantIdList[i] == sectionCropsId {
							old.PlantIdList = append(old.PlantIdList[:i], old.PlantIdList[i+1:]...)
							i--
						}
					}
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutHarvest", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				d, err := sdk.Query(h.HarvestId, define.QueryHarvest)
				if err != nil || len(d) == 0 {
					//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					wg.Done()
					return "", nil
				}
				old := &define.Harvest{}
				err = json.Unmarshal(d, old)
				if err != nil {
					wg.Done()
					return "", nil
				}
				if !in(sectionCropsId, old.PlantIdList) {
					old.PlantIdList = append(old.PlantIdList, sectionCropsId)
				}
				data, err := json.Marshal(old)
				if err != nil {
					wg.Done()
					return "", nil
				}
				// 更新链上数据
				info := []string{"PutHarvest", string(data)}
				res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
				if err != nil {
					wg.Done()
					return "", nil
				}
				response = res.TxID
				wg.Done()

			case "UPDATE":
				_, okHarvestId := harvestData.Old[gIndex]["reap_batch_id"]
				_, okSectionCropsId := harvestData.Old[gIndex]["section_crops_id"]
				_, okIsActive := harvestData.Old[gIndex]["isActive"]
				go func() {
					harvestId := ""
					if okHarvestId {
						harvestId = reflect.ValueOf(harvestData.Old[gIndex]["reap_batch_id"]).String()
					} else {
						harvestId = h.HarvestId
					}
					d, err := sdk.Query(harvestId, define.QueryHarvest)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.Harvest{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					if okIsActive && isActive == "N" {
						// 删除对应数据
						for i := 0; i < len(old.PlantIdList); i++ {
							if old.PlantIdList[i] == sectionCropsId {
								old.PlantIdList = append(old.PlantIdList[:i], old.PlantIdList[i+1:]...)
								i--
							}
						}
						data, err := json.Marshal(old)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						info := []string{"PutHarvest", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
						wg.Done()
						return
					}
					// 删除对应数据
					if okSectionCropsId {
						oldSectionId := reflect.ValueOf(harvestData.Old[gIndex]["section_crops_id"]).String()
						for i, v := range old.PlantIdList {
							if v == oldSectionId {
								old.PlantIdList[i] = sectionCropsId
							}
						}
					}
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutHarvest", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()
			}
		}
		wg.Wait()
		cancel()
	}
	return response, nil
}

// 烘干上链信息
// key: 烘干批次ID
// value： 烘干批次ID、收割批次ID、烘干信息
func HandleDry(dryData apiDefine.KafkaMessage) (string, error) {
	response := ""
	switch dryData.Table {
	case "srv_dry_info":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		for gIndex, dry := range dryData.Data {
			wg.Add(1)
			l := &define.DryInfo{}
			l.DryId = reflect.ValueOf(dry["id"]).String()
			l.HarvestId = reflect.ValueOf(dry["sg_id"]).String() // 收割id
			l.HgDate, _ = time.Parse("2006-01-02 15:04:05", reflect.ValueOf(dry["hg_date"]).String())
			l.HgFactoryName = reflect.ValueOf(dry["hg_factory_name"]).String()
			l.ImageUrl = reflect.ValueOf(dry["attachment"]).String() // json 字符串
			_, okFeedListId := dry["feed_list_id"]
			if okFeedListId {
				l.FeedListId, _ = strconv.ParseInt(reflect.ValueOf(dry["feed_list_id"]).String(), 10, 64)
			}
			_, okHarvestYear := dry["harvest_year"]
			if okHarvestYear {
				l.HarvestYear = reflect.ValueOf(dry["harvest_year"]).String()
			}
			_, okReachingDate := dry["reaching_date"]
			if okReachingDate {
				l.ReachingDate, _ = time.Parse("2006-01-02 15:04:05", reflect.ValueOf(dry["reaching_date"]).String())
			}
			isActive := reflect.ValueOf(dry["is_active"]).String()
			switch dryData.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(l.DryId, define.QueryDryInfo)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.DryInfo{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					// 删除对应数据
					old.Status = 1
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutDryInfo", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				go func() {
					d, err := sdk.Query(l.DryId, define.QueryDryInfo)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					if len(d) > 0 {
						old := &define.DryInfo{}
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						if old.Status == 0 {
							fmt.Println("数据数据已存在")
							wg.Done()
							return
						}
					}
					data, err := json.Marshal(l)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutDryInfo", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				_, okDryId := dryData.Old[gIndex]["id"]
				_, okHarvestId := dryData.Old[gIndex]["sg_id"]
				_, okHgDate := dryData.Old[gIndex]["hg_date"]
				_, okHgFactoryName := dryData.Old[gIndex]["hg_factory_name"]
				_, okImageUrl := dryData.Old[gIndex]["attachment"]
				_, okIsActive := dryData.Old[gIndex]["is_active"]
				_, okOldFeedListId := dryData.Old[gIndex]["feed_list_id"]
				_, okOldHarvestYear := dryData.Old[gIndex]["harvest_year"]
				_, okOldReachingDate := dryData.Old[gIndex]["reaching_date"]
				go func() {
					dryId := ""
					if okDryId {
						dryId = reflect.ValueOf(dryData.Old[gIndex]["id"]).String()
					} else {
						dryId = l.DryId
					}
					d, err := sdk.Query(dryId, define.QueryDryInfo)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					if len(d) > 0 {
						old := &define.DryInfo{}
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						if okIsActive && isActive == "N" {
							// 删除对应数据
							old.Status = 1
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutDryInfo", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
							wg.Done()
							return
						}
						if okDryId {
							old.DryId = l.DryId
						}
						if okHarvestId {
							old.HarvestId = l.HarvestId
						}
						if okHgDate {
							old.HgDate = l.HgDate
						}
						if okHgFactoryName {
							old.HgFactoryName = l.HgFactoryName
						}
						if okImageUrl {
							old.ImageUrl = l.ImageUrl
						}
						if okOldFeedListId {
							old.FeedListId = l.FeedListId
						}
						if okOldHarvestYear {
							old.HarvestYear = l.HarvestYear
						}
						if okOldReachingDate {
							old.ReachingDate = l.ReachingDate
						}
						if okDryId || okHarvestId || okHgDate || okHgFactoryName || okImageUrl || okOldFeedListId || okOldHarvestYear || okOldReachingDate {
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutDryInfo", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
						}
					} else {
						data, err := json.Marshal(l)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						info := []string{"PutDryInfo", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
					}
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()
			}
		}
		wg.Wait()
		cancel()
	}
	return response, nil
}

// 仓储上链信息
// key: 仓储批次ID
// value：仓储批次ID、烘干批次ID、收割批次ID、仓储信息
func HandleWarehouse(warehouseData apiDefine.KafkaMessage) (string, error) {
	response := ""
	switch warehouseData.Table {
	case "srv_storage_info":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		for gIndex, warehouse := range warehouseData.Data {
			wg.Add(1)
			l := &define.WarehouseInfo{}
			l.WarehouseId = reflect.ValueOf(warehouse["id"]).String()
			l.DryId = reflect.ValueOf(warehouse["hg_id"]).String()
			l.StartDate, _ = time.Parse("2006-01-02 15:04:05", reflect.ValueOf(warehouse["start_date"]).String())
			l.DurationDate, _ = time.Parse("2006-01-02 15:04:05", reflect.ValueOf(warehouse["end_date"]).String())
			l.RefTemperature = reflect.ValueOf(warehouse["ref_temperature"]).String()
			l.CCFactoryName = reflect.ValueOf(warehouse["cc_factory_name"]).String()
			l.ImageUrl = reflect.ValueOf(warehouse["start_attachment"]).String() // json 字符串
			isActive := reflect.ValueOf(warehouse["is_active"]).String()
			switch warehouseData.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(l.WarehouseId, define.QueryWarehouseInfo)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.WarehouseInfo{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					// 删除对应数据
					old.Status = 1
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutWarehouseInfo", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				go func() {
					d, err := sdk.Query(l.WarehouseId, define.QueryWarehouseInfo)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					if len(d) > 0 {
						old := &define.WarehouseInfo{}
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						if old.Status == 0 {
							fmt.Println("数据数据已存在")
							wg.Done()
							return
						}
					}
					data, err := json.Marshal(l)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutWarehouseInfo", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				_, okWarehouseId := warehouseData.Old[gIndex]["id"]
				_, okDryId := warehouseData.Old[gIndex]["hg_id"]
				_, okStartDate := warehouseData.Old[gIndex]["start_date"]
				_, okDurationDate := warehouseData.Old[gIndex]["end_date"]
				_, okRefTemperature := warehouseData.Old[gIndex]["ref_temperature"]
				_, okCCFactoryName := warehouseData.Old[gIndex]["cc_factory_name"]
				_, okImageUrl := warehouseData.Old[gIndex]["start_attachment"]
				_, okIsActive := warehouseData.Old[gIndex]["is_active"]
				go func() {
					warehouseId := ""
					if okWarehouseId {
						warehouseId = reflect.ValueOf(warehouseData.Old[gIndex]["id"]).String()
					} else {
						warehouseId = l.WarehouseId
					}
					d, err := sdk.Query(warehouseId, define.QueryWarehouseInfo)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					if len(d) > 0 {
						old := &define.WarehouseInfo{}
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						if okIsActive && isActive == "N" {
							// 删除对应数据
							old.Status = 1
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutWarehouseInfo", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
							wg.Done()
							return
						}
						if okWarehouseId {
							old.WarehouseId = l.WarehouseId
						}
						if okDryId {
							old.DryId = l.DryId
						}
						if okStartDate {
							old.StartDate = l.StartDate
						}
						if okDurationDate {
							old.DurationDate = l.DurationDate
						}
						if okRefTemperature {
							old.RefTemperature = l.RefTemperature
						}
						if okCCFactoryName {
							old.CCFactoryName = l.CCFactoryName
						}
						if okImageUrl {
							old.ImageUrl = l.ImageUrl
						}
						if okWarehouseId || okDryId || okStartDate || okDurationDate || okRefTemperature ||
							okCCFactoryName || okImageUrl {
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutWarehouseInfo", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
						}
					} else {
						data, err := json.Marshal(l)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						info := []string{"PutWarehouseInfo", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
					}

					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()
			}
		}
		wg.Wait()
		cancel()
	}
	return response, nil
}

// 加工过程上链数据
// key: 加工批次ID
// value: 加工批次ID、上一环节类型（烘干或仓储）、上一环节批次ID、加工信息、交易ID（作为加工、检测报告相关区块链证书）
func HandleProcess(processData apiDefine.KafkaMessage) (string, error) {
	response := ""
	switch processData.Table {
	case "srv_mach_packing":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		for gIndex, process := range processData.Data {
			wg.Add(1)
			l := &define.ProcessInfo{}
			l.ProcessId = reflect.ValueOf(process["id"]).String()
			//l.PreType = reflect.ValueOf(process["up_flow_type"]).String() // 上一环节类型仓储CC，烘干HG
			l.PreType = "hg"
			hgId := strings.Split(reflect.ValueOf(process["sg_id"]).String(), ",") // sg_id
			//l.PreId = reflect.ValueOf(process["before_business_id"]).String() // 上一环节id
			if len(hgId) > 0 {
				l.PreId = hgId[0]
			}
			l.ProcessDate, _ = time.Parse("2006-01-02 15:04:05", reflect.ValueOf(process["jg_date"]).String())
			l.ProcessEnterpriseName = reflect.ValueOf(process["jgc_enterprise_name"]).String() // 加工厂名称
			l.ProcessImageUrl = reflect.ValueOf(process["jg_attachment"]).String()             // json字符串
			l.QualityReportImageUrl = reflect.ValueOf(process["pzjc_attachment"]).String()     // json字符串
			l.QualityInspectionDate, _ = time.Parse("2006-01-02 15:04:05", reflect.ValueOf(process["pzjc_date"]).String())
			l.SafeReportImageUrl = reflect.ValueOf(process["aqjc_attachment"]).String() // json字符串
			l.SafeInspectionDate, _ = time.Parse("2006-01-02 15:04:05", reflect.ValueOf(process["aqjc_date"]).String())
			l.PackDate, _ = time.Parse("2006-01-02 15:04:05", reflect.ValueOf(process["pack_date"]).String())
			l.PackageImageUrl = reflect.ValueOf(process["pack_attachment"]).String() // json字符串
			isActive := reflect.ValueOf(process["is_active"]).String()
			switch processData.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(l.ProcessId, define.QueryProcessInfo)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.ProcessInfo{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					// 删除对应数据
					old.Status = 1
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutProcessInfo", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				go func() {
					d, err := sdk.Query(l.ProcessId, define.QueryProcessInfo)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					if len(d) > 0 {
						old := &define.ProcessInfo{}
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						if old.Status == 0 {
							fmt.Println("数据数据已存在")
							wg.Done()
							return
						}
					}
					data, err := json.Marshal(l)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutProcessInfo", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				_, okProcessId := processData.Old[gIndex]["id"]
				_, okPreType := processData.Old[gIndex]["up_flow_type"]
				_, okPreId := processData.Old[gIndex]["before_business_id"]
				_, okSgId := processData.Old[gIndex]["sg_id"]
				_, okProcessDate := processData.Old[gIndex]["jg_date"]
				_, okProcessEnterpriseName := processData.Old[gIndex]["jgc_enterprise_name"]
				_, okProcessImageUrl := processData.Old[gIndex]["jg_attachment"]
				_, okQualityReportImageUrl := processData.Old[gIndex]["pzjc_attachment"]
				_, okQualityInspectionDate := processData.Old[gIndex]["pzjc_date"]
				_, okSafeReportImageUrl := processData.Old[gIndex]["aqjc_attachment"]
				_, okSafeInspectionDate := processData.Old[gIndex]["aqjc_date"]
				_, okPackDate := processData.Old[gIndex]["pack_date"]
				_, okPackageImageUrl := processData.Old[gIndex]["pack_attachment"]
				_, okIsActive := processData.Old[gIndex]["is_active"]
				go func() {
					processId := ""
					if okProcessId {
						processId = reflect.ValueOf(processData.Old[gIndex]["id"]).String()
					} else {
						processId = l.ProcessId
					}
					d, err := sdk.Query(processId, define.QueryProcessInfo)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					if len(d) > 0 {
						old := &define.ProcessInfo{}
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						if okIsActive && isActive == "N" {
							// 删除对应数据
							old.Status = 1
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutProcessInfo", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
							wg.Done()
							return
						}
						if okProcessId {
							old.ProcessId = l.ProcessId
						}
						/*if okPreType {
							old.PreType = l.PreType
						}
						if okPreId {
							old.PreId = l.PreId
						}*/
						if okSgId {
							old.PreId = l.PreId
						}
						if okProcessDate {
							old.ProcessDate = l.ProcessDate
						}
						if okProcessEnterpriseName {
							old.ProcessEnterpriseName = l.ProcessEnterpriseName
						}
						if okProcessImageUrl {
							old.ProcessImageUrl = l.ProcessImageUrl
						}
						if okQualityReportImageUrl {
							old.QualityReportImageUrl = l.QualityReportImageUrl
						}
						if okQualityInspectionDate {
							old.QualityInspectionDate = l.QualityInspectionDate
						}
						if okSafeReportImageUrl {
							old.SafeReportImageUrl = l.SafeReportImageUrl
						}
						if okSafeInspectionDate {
							old.SafeInspectionDate = l.SafeInspectionDate
						}
						if okPackDate {
							old.PackDate = l.PackDate
						}
						if okPackageImageUrl {
							old.PackageImageUrl = l.PackageImageUrl
						}
						if okProcessId || okPreType || okPreId || okSgId || okProcessDate || okProcessEnterpriseName ||
							okProcessImageUrl || okQualityReportImageUrl || okQualityInspectionDate ||
							okSafeReportImageUrl || okSafeInspectionDate || okPackDate || okPackageImageUrl {
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutProcessInfo", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
						}
					} else {
						data, err := json.Marshal(l)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						info := []string{"PutProcessInfo", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID

					}

					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()
			}
		}
		wg.Wait()
		cancel()
	case "srv_mach_packing_norms":
	}
	return response, nil
}

// 收粮ID与种植ID关联
// key: 收粮ID
// value：种植ID
func HandleFeedSectionMap(feedSectionMapData apiDefine.KafkaMessage) (string, error) {
	response := ""
	switch feedSectionMapData.Table {
	case "srv_feed_section_map":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		for gIndex, feedSectionMap := range feedSectionMapData.Data {
			wg.Add(1)
			l := &define.FeedSectionCrop{}
			l.FeedListId = reflect.ValueOf(feedSectionMap["feed_list_id"]).String()
			l.SectionCropsId = reflect.ValueOf(feedSectionMap["section_crops_id"]).String()
			l.IsActive = reflect.ValueOf(feedSectionMap["is_active"]).String()
			switch feedSectionMapData.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(l.FeedListId, define.QueryFeedSectionCrop)
					if err != nil || len(d) == 0 {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					old := &define.FeedSectionCrop{}
					err = json.Unmarshal(d, old)
					if err != nil {
						wg.Done()
						return
					}
					// 删除对应数据
					old.Status = 1
					data, err := json.Marshal(old)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutFeedSectionCrop", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				go func() {
					d, err := sdk.Query(l.FeedListId, define.QueryFeedSectionCrop)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					if len(d) > 0 {
						old := &define.FeedSectionCrop{}
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						if old.Status == 0 {
							fmt.Println("数据数据已存在")
							wg.Done()
							return
						}
					}
					data, err := json.Marshal(l)
					if err != nil {
						wg.Done()
						return
					}
					// 更新链上数据
					info := []string{"PutFeedSectionCrop", string(data)}
					res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				_, okFeedListId := feedSectionMapData.Old[gIndex]["feed_list_id"]
				_, okSectionCropsId := feedSectionMapData.Old[gIndex]["section_crops_id"]
				_, okIsActive := feedSectionMapData.Old[gIndex]["is_active"]
				go func() {
					feedListId := ""
					if okFeedListId {
						feedListId = reflect.ValueOf(feedSectionMapData.Old[gIndex]["feed_list_id"]).String()
					} else {
						feedListId = l.FeedListId
					}
					d, err := sdk.Query(feedListId, define.QueryFeedSectionCrop)
					if err != nil {
						//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
						wg.Done()
						return
					}
					if len(d) > 0 {
						old := &define.FeedSectionCrop{}
						err = json.Unmarshal(d, old)
						if err != nil {
							wg.Done()
							return
						}
						if okIsActive && l.IsActive == "N" {
							// 删除对应数据
							old.Status = 1
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutFeedSectionCrop", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
							wg.Done()
							return
						}
						if okSectionCropsId {
							old.SectionCropsId = l.SectionCropsId
						}
						if okFeedListId || okSectionCropsId || okIsActive {
							data, err := json.Marshal(old)
							if err != nil {
								wg.Done()
								return
							}
							// 更新链上数据
							info := []string{"PutFeedSectionCrop", string(data)}
							res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
							if err != nil {
								wg.Done()
								return
							}
							response = res.TxID
						}
					} else {
						data, err := json.Marshal(l)
						if err != nil {
							wg.Done()
							return
						}
						// 更新链上数据
						info := []string{"PutFeedSectionCrop", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							wg.Done()
							return
						}
						response = res.TxID
					}
					wg.Done()
					for {
						select {
						case <-ctx.Done():
							return
						}
					}
				}()
			}
		}
		wg.Wait()
		cancel()
	}
	return response, nil
}

// 数据字典
// key：AgDict
// value：id name value type_code
// 农事操作具体名字，比如灌溉，是喷灌、滴灌、还是满灌。
// 通过ag_farming_record_detail.type_code==ag_dict.type_code &&
// ag_farming_record_detail.use_good==ag_dict.value => ag_farming_record_detail.type_use=ag_dict.name
func HandleDict(dictData apiDefine.KafkaMessage) (string, error) {
	response := ""
	switch dictData.Table {
	case "ag_dict":
		// 收割批次id（ag_reap_batch.id）、收割时间、Txid
		if dictData.Type == "INSERT" {
			newDict := &define.AgDictInfo{}
			newFlag := false
			for _, dict := range dictData.Data {
				dictId, _ := strconv.Atoi(reflect.ValueOf(dict["id"]).String())
				name := reflect.ValueOf(dict["name"]).String()
				value, _ := strconv.Atoi(reflect.ValueOf(dict["value"]).String())
				typeCode := reflect.ValueOf(dict["type_code"]).String()
				fmt.Printf("insert dict id :%v\n", dictId)
				d, err := sdk.Query("AgDict", define.QueryDict)
				if err != nil {
					err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				dic := define.AgDict{}
				dic.Id = int64(dictId)
				dic.Name = name
				dic.Value = value
				dic.TypeCode = typeCode
				old := &define.AgDictInfo{}
				if len(d) > 0 {
					existFlag := false
					err = json.Unmarshal(d, old)
					if err != nil {
						return "", err
					}
					fmt.Printf("old is %v\n", *old)
					for _, v := range old.AgDictList {
						if v.Id == dic.Id {
							fmt.Printf("old id is %v\n", v.Id)
							existFlag = true
						}
					}
					if !existFlag {
						newDict.AgDictList = append(old.AgDictList, dic)
						newFlag = true
					}
				} else {
					newDict.AgDictList = append(newDict.AgDictList, dic)
					newFlag = true
				}
			}
			if newFlag {
				data, err := json.Marshal(newDict)
				if err != nil {
					return "", err
				}
				// 更新链上数据
				info := []string{"PutDict", string(data)}
				res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
				if err != nil {
					return "", err
				}
				response = res.TxID
			}
			return response, nil
		}

		for gIndex, dict := range dictData.Data {
			dictId, _ := strconv.Atoi(reflect.ValueOf(dict["id"]).String())
			name := reflect.ValueOf(dict["name"]).String()
			value, _ := strconv.Atoi(reflect.ValueOf(dict["value"]).String())
			typeCode := reflect.ValueOf(dict["type_code"]).String()
			isActive := reflect.ValueOf(dict["is_active"]).String()
			switch dictData.Type {
			case "DELETE":
				d, err := sdk.Query("AgDict", define.QueryDict)
				if err != nil || len(d) == 0 {
					err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				old := &define.AgDictInfo{}
				err = json.Unmarshal(d, old)
				if err != nil {
					return "", err
				}
				// 删除对应数据
				for i := 0; i < len(old.AgDictList); i++ {
					if old.AgDictList[i].Id == int64(dictId) {
						old.AgDictList = append(old.AgDictList[:i], old.AgDictList[i+1:]...)
						i--
					}
				}
				data, err := json.Marshal(old)
				if err != nil {
					return "", err
				}
				// 更新链上数据
				info := []string{"PutDict", string(data)}
				res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
				if err != nil {
					return "", err
				}
				response = res.TxID
			case "UPDATE":
				_, okIsActive := dictData.Old[gIndex]["is_active"]
				d, err := sdk.Query("AgDict", define.QueryDict)
				if err != nil {
					err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				old := &define.AgDictInfo{}
				if len(d) > 0 {
					// 链上有数据，要查询是否有此Id，有就更新，没有就插入
					err = json.Unmarshal(d, old)
					if err != nil {
						return "", err
					}
					if okIsActive && isActive == "N" {
						// 删除对应数据
						for i := 0; i < len(old.AgDictList); i++ {
							if old.AgDictList[i].Id == int64(dictId) {
								old.AgDictList = append(old.AgDictList[:i], old.AgDictList[i+1:]...)
								i--
							}
						}
						data, err := json.Marshal(old)
						if err != nil {
							return "", err
						}
						// 更新链上数据
						info := []string{"PutDict", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							return "", err
						}
						response = res.TxID
						return response, nil
					}
					existFlag := false
					for i, v := range old.AgDictList {
						if v.Id == int64(dictId) {
							old.AgDictList[i].Name = name
							old.AgDictList[i].Value = value
							old.AgDictList[i].TypeCode = typeCode
							existFlag = true
						}
					}
					if !existFlag {
						var agDict define.AgDict
						agDict.Id = int64(dictId)
						agDict.Name = name
						agDict.Value = value
						agDict.TypeCode = typeCode
						old.AgDictList = append(old.AgDictList, agDict)
					}
				} else {
					// 链上没有数据，直接插入
					var agDict define.AgDict
					agDict.Id = int64(dictId)
					agDict.Name = name
					agDict.Value = value
					agDict.TypeCode = typeCode
					old.AgDictList = append(old.AgDictList, agDict)
				}
				data, err := json.Marshal(old)
				if err != nil {
					return "", err
				}
				// 更新链上数据
				info := []string{"PutDict", string(data)}
				res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
				if err != nil {
					return "", err
				}
				response = res.TxID
			}
		}
	}
	return response, nil
}

// 数据字典
// key：AgPesticideFertilizer
// value：id name type type_code value unit
// 化肥、农药的具体名字，计量单位。
// 通过ag_farming_record_detail.type_code==ag_pesticide_fertilizer.type_code &&
// ag_farming_record_detail.use_good==ag_pesticide_fertilizer.value => ag_farming_record_detail.type_use=ag_pesticide_fertilizer.name
func HandlePesticideFertilizer(pesticideFertilizerData apiDefine.KafkaMessage) (string, error) {
	response := ""
	switch pesticideFertilizerData.Table {
	case "ag_pesticide_fertilizer":
		// 收割批次id（ag_reap_batch.id）、收割时间、Txid
		if pesticideFertilizerData.Type == "INSERT" {
			newPesticideFertilizer := &define.AgPesticideFertilizerInfo{}
			newFlag := false
			for _, pfData := range pesticideFertilizerData.Data {
				pfId := reflect.ValueOf(pfData["id"]).String()
				name := reflect.ValueOf(pfData["name"]).String()
				pfType := reflect.ValueOf(pfData["type"]).String()
				typeCode := reflect.ValueOf(pfData["type_code"]).String()
				value := reflect.ValueOf(pfData["value"]).String()
				unit := reflect.ValueOf(pfData["unit"]).String()
				d, err := sdk.Query("AgPesticideFertilizer", define.QueryPesticideFertilizer)
				if err != nil {
					err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				pf := &define.AgPesticideFertilizer{}
				pf.Id = pfId
				pf.Name = name
				pf.Type = pfType
				pf.TypeCode = typeCode
				pf.Value = value
				pf.Unit = unit
				old := &define.AgPesticideFertilizerInfo{}
				if len(d) > 0 {
					existFlag := false
					err = json.Unmarshal(d, old)
					if err != nil {
						return "", err
					}
					for _, v := range old.AgPesticideFertilizerList {
						if v.Id == pf.Id {
							fmt.Printf("old id is %v\n", v.Id)
							existFlag = true
						}
					}
					if !existFlag {
						newPesticideFertilizer.AgPesticideFertilizerList = append(old.AgPesticideFertilizerList, *pf)
						newFlag = true
					}
				} else {
					newPesticideFertilizer.AgPesticideFertilizerList = append(newPesticideFertilizer.AgPesticideFertilizerList, *pf)
					newFlag = true
				}
			}
			if newFlag {
				data, err := json.Marshal(newPesticideFertilizer)
				if err != nil {
					return "", err
				}
				// 更新链上数据
				info := []string{"PutPesticideFertilizer", string(data)}
				res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
				if err != nil {
					return "", err
				}
				response = res.TxID
			}
			return response, nil
		}
		for gIndex, pfData := range pesticideFertilizerData.Data {
			pfId := reflect.ValueOf(pfData["id"]).String()
			name := reflect.ValueOf(pfData["name"]).String()
			pfType := reflect.ValueOf(pfData["type"]).String()
			typeCode := reflect.ValueOf(pfData["type_code"]).String()
			value := reflect.ValueOf(pfData["value"]).String()
			unit := reflect.ValueOf(pfData["unit"]).String()
			isActive := reflect.ValueOf(pfData["is_active"]).String()
			switch pesticideFertilizerData.Type {
			case "DELETE":
				d, err := sdk.Query("AgPesticideFertilizer", define.QueryPesticideFertilizer)
				if err != nil || len(d) == 0 {
					err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				old := &define.AgPesticideFertilizerInfo{}
				err = json.Unmarshal(d, old)
				if err != nil {
					return "", err
				}
				// 删除对应数据
				for i := 0; i < len(old.AgPesticideFertilizerList); i++ {
					if old.AgPesticideFertilizerList[i].Id == pfId {
						old.AgPesticideFertilizerList = append(old.AgPesticideFertilizerList[:i], old.AgPesticideFertilizerList[i+1:]...)
						i--
					}
				}
				data, err := json.Marshal(old)
				if err != nil {
					return "", err
				}
				// 更新链上数据
				info := []string{"PutPesticideFertilizer", string(data)}
				res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
				if err != nil {
					return "", err
				}
				response = res.TxID
			case "UPDATE":
				_, okIsActive := pesticideFertilizerData.Old[gIndex]["is_active"]
				d, err := sdk.Query("AgPesticideFertilizer", define.QueryPesticideFertilizer)
				if err != nil {
					err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
					return "", err
				}
				old := &define.AgPesticideFertilizerInfo{}
				if len(d) > 0 {
					// 链上有数据，判断是否存在此ID，存在则更新，不存在则插入
					err = json.Unmarshal(d, old)
					if err != nil {
						return "", err
					}
					if okIsActive && isActive == "N" {
						// 删除对应数据
						for i := 0; i < len(old.AgPesticideFertilizerList); i++ {
							if old.AgPesticideFertilizerList[i].Id == pfId {
								old.AgPesticideFertilizerList = append(old.AgPesticideFertilizerList[:i], old.AgPesticideFertilizerList[i+1:]...)
								i--
							}
						}
						data, err := json.Marshal(old)
						if err != nil {
							return "", err
						}
						// 更新链上数据
						info := []string{"PutPesticideFertilizer", string(data)}
						res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
						if err != nil {
							return "", err
						}
						response = res.TxID
						return response, nil
					}
					existFlag := false
					for i, v := range old.AgPesticideFertilizerList {
						if v.Id == pfId {
							old.AgPesticideFertilizerList[i].Name = name
							old.AgPesticideFertilizerList[i].Type = pfType
							old.AgPesticideFertilizerList[i].TypeCode = typeCode
							old.AgPesticideFertilizerList[i].Value = value
							old.AgPesticideFertilizerList[i].Unit = unit
							existFlag = true
						}
					}
					if !existFlag {
						var agPesticideFertilizer define.AgPesticideFertilizer
						agPesticideFertilizer.Id = pfId
						agPesticideFertilizer.Name = name
						agPesticideFertilizer.Type = pfType
						agPesticideFertilizer.TypeCode = typeCode
						agPesticideFertilizer.Value = value
						agPesticideFertilizer.Unit = unit
						old.AgPesticideFertilizerList = append(old.AgPesticideFertilizerList, agPesticideFertilizer)
					}
				} else {
					// 链上不存在，直接插入
					var agPesticideFertilizer define.AgPesticideFertilizer
					agPesticideFertilizer.Id = pfId
					agPesticideFertilizer.Name = name
					agPesticideFertilizer.Type = pfType
					agPesticideFertilizer.TypeCode = typeCode
					agPesticideFertilizer.Value = value
					agPesticideFertilizer.Unit = unit
					old.AgPesticideFertilizerList = append(old.AgPesticideFertilizerList, agPesticideFertilizer)
				}
				data, err := json.Marshal(old)
				if err != nil {
					return "", err
				}
				// 更新链上数据
				info := []string{"PutPesticideFertilizer", string(data)}
				res, err := gohfc.GetHandler().SyncInvoke(info, "", "")
				if err != nil {
					return "", err
				}
				response = res.TxID
			}
		}
	}
	return response, nil
}

func Query(c *gin.Context) {
	key := c.Param("uuid")
	utils.Logger.Debugf("query key: %v\n", key)
	qRCodeResponse := &define.QRCodeResponse{}
	processInfo := &define.ProcessInfo{}
	fmt.Println("-----------查询加工批次--------------")
	processData, err := sdk.Query(key, define.QueryProcessInfo)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	fmt.Println("-----------查询加工批次---unmarshal-----------")
	err = json.Unmarshal(processData, processInfo)
	if err != nil {
		errStr := fmt.Sprintf("json unmarshal  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	fmt.Println("-----------查询加工批次---unmarshal--success---------")
	qRCodeResponse.ProcessRecordDetail.Process = processInfo.Process
	qRCodeResponse.ProcessRecordDetail.Package = processInfo.Package
	qRCodeResponse.ProcessRecordDetail.TxId = processInfo.TxId
	qRCodeResponse.ProcessRecordDetail.BlockNumber, qRCodeResponse.ProcessRecordDetail.TxCreateTime, err = sdk.GetBlockInfoByTxID(qRCodeResponse.ProcessRecordDetail.TxId)
	if err != nil {
		errStr := fmt.Sprintf("get block by txid error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	dryInfo := &define.DryInfo{}
	warehouseInfo := &define.WarehouseInfo{}
	switch processInfo.PreType {
	case "cc", "CC": // 上一环节为仓储
		fmt.Println("-----------查询仓储批次--------------")
		warehouseData, err := sdk.Query(processInfo.PreId, define.QueryWarehouseInfo)
		if err != nil {
			errStr := fmt.Sprintf("FormQuery  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询仓储批次---unmarshal-----------")
		err = json.Unmarshal(warehouseData, warehouseInfo)
		if err != nil {
			errStr := fmt.Sprintf("json unmarshal  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询仓储批次---unmarshal--success---------")
		qRCodeResponse.ProcessRecordDetail.Warehouse = warehouseInfo.Warehouse
		fmt.Println("-----------查询烘干信息--------------", warehouseInfo.DryId)
		dryData, err := sdk.Query(warehouseInfo.DryId, define.QueryDryInfo)
		if err != nil {
			errStr := fmt.Sprintf("FormQuery  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询烘干信息---unmarshal-----------")
		err = json.Unmarshal(dryData, dryInfo)
		if err != nil {
			errStr := fmt.Sprintf("json unmarshal  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询烘干信息---unmarshal--success---------")
		qRCodeResponse.ProcessRecordDetail.Dry = dryInfo.Dry
	case "hg", "HG": // 上一环节为烘干
		fmt.Println("-----------查询烘干信息--------------")
		dryData, err := sdk.Query(processInfo.PreId, define.QueryDryInfo)
		if err != nil {
			errStr := fmt.Sprintf("FormQuery  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询烘干信息---unmarshal-----------")
		err = json.Unmarshal(dryData, dryInfo)
		if err != nil {
			errStr := fmt.Sprintf("json unmarshal  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询烘干信息---unmarshal--success---------")
		qRCodeResponse.ProcessRecordDetail.Dry = dryInfo.Dry
	default:
		errStr := fmt.Sprintf("加工上一环节信息错误，请确认是否为仓储或烘干")
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	// 如果有收粮数据就走收粮溯源，没有收粮数据，走正常收割溯源
	plant := &define.Plant{}
	if dryInfo.FeedListId == 0 {
		// 收割溯源流程
		// 收割数据，通过收割获取种植信息
		harvest := &define.Harvest{}
		fmt.Println("-----------查询收割批次--------------")
		harvestData, err := sdk.Query(dryInfo.HarvestId, define.QueryHarvest)
		if err != nil {
			errStr := fmt.Sprintf("FormQuery  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询收割批次----Unmarshal json----------")
		err = json.Unmarshal(harvestData, harvest)
		if err != nil {
			errStr := fmt.Sprintf("json unmarshal  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询收割批次----Unmarshal json---success-------")
		qRCodeResponse.FarmingRecordDetail.TxId = harvest.TxId
		qRCodeResponse.FarmingRecordDetail.BlockNumber, qRCodeResponse.FarmingRecordDetail.TxCreateTime, err = sdk.GetBlockInfoByTxID(qRCodeResponse.FarmingRecordDetail.TxId)
		if err != nil {
			errStr := fmt.Sprintf("get block by txid error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		// 收割批次对应多个种植批次，种植批次有唯一地块，
		// 农事记录只显示面积最大地块的相关信息
		var landArea float64
		for i, v := range harvest.PlantIdList {
			// 查询种植id对应的地块id
			plant := &define.Plant{}
			fmt.Printf("-----------查询种植-----------v:%v\n", v)
			plantData, err := sdk.Query(v, define.QueryPlant)
			if err != nil {
				errStr := fmt.Sprintf("FormQuery  error : %v", err)
				utils.Response(errStr, c, http.StatusBadRequest)
				return
			}
			fmt.Println("-----------查询种植---unmarshal--------")
			err = json.Unmarshal(plantData, plant)
			if err != nil {
				errStr := fmt.Sprintf("json unmarshal  error : %v", err)
				utils.Response(errStr, c, http.StatusBadRequest)
				return
			}
			fmt.Println("-----------查询种植---unmarshal---success-----")
			// 查询地块面积
			land := &define.Land{}
			fmt.Println("-----------查询地块-----------")
			landData, err := sdk.Query(plant.SectionId, define.QueryLand)
			if err != nil {
				errStr := fmt.Sprintf("FormQuery  error : %v", err)
				utils.Response(errStr, c, http.StatusBadRequest)
				return
			}
			fmt.Println("-----------查询地块-----unmarshal------")
			err = json.Unmarshal(landData, land)
			if err != nil {
				errStr := fmt.Sprintf("json unmarshal  error : %v", err)
				utils.Response(errStr, c, http.StatusBadRequest)
				return
			}
			fmt.Println("-----------查询地块-----unmarshal--success----")
			if land.Area > landArea {
				harvest.PlantIdList[0] = harvest.PlantIdList[i]
				landArea = land.Area
			}
		}
		// 查询种植id对应的地块id
		fmt.Println("-----------查询种植-----------")
		plantData, err := sdk.Query(harvest.PlantIdList[0], define.QueryPlant)
		if err != nil {
			errStr := fmt.Sprintf("FormQuery  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询种植---unmarshal--------")
		err = json.Unmarshal(plantData, plant)
		if err != nil {
			errStr := fmt.Sprintf("json unmarshal  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		qRCodeResponse.ReapCreateDate = harvest.ReapCreateDate
	} else {
		// 收粮溯源流程
		feedSectionCrop := &define.FeedSectionCrop{}
		fmt.Println("-----------查询进粮清单--------------")
		feedListId := strconv.FormatInt(dryInfo.FeedListId, 10)
		feedSectionMapData, err := sdk.Query(feedListId, define.QueryFeedSectionCrop)
		if err != nil {
			errStr := fmt.Sprintf("FormQuery  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询进粮清单----Unmarshal json----------")
		err = json.Unmarshal(feedSectionMapData, feedSectionCrop)
		if err != nil {
			errStr := fmt.Sprintf("json unmarshal  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询进粮清单----Unmarshal json---success-------")
		qRCodeResponse.FarmingRecordDetail.TxId = feedSectionCrop.TxId
		qRCodeResponse.FarmingRecordDetail.BlockNumber, qRCodeResponse.FarmingRecordDetail.TxCreateTime, err = sdk.GetBlockInfoByTxID(qRCodeResponse.FarmingRecordDetail.TxId)
		if err != nil {
			errStr := fmt.Sprintf("get block by txid error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		// 查询种植id对应的地块id
		fmt.Println("-----------查询种植-----------")
		plantData, err := sdk.Query(feedSectionCrop.SectionCropsId, define.QueryPlant)
		if err != nil {
			errStr := fmt.Sprintf("FormQuery  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询种植---unmarshal--------")
		err = json.Unmarshal(plantData, plant)
		if err != nil {
			errStr := fmt.Sprintf("json unmarshal  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		qRCodeResponse.ReapCreateDate = dryInfo.ReachingDate
	}

	fmt.Println("-----------查询种植---unmarshal--success------")
	//qRCodeResponse.FarmingRecordDetail.PlantFarm = plant.PlantFarm
	for _, operateId := range plant.FarmRecordIdList {
		// 收粮溯源流程
		farmRecord := &define.FarmRecord{}
		fmt.Println("-----------查询农事操作--------------")

		farmRecordData, err := sdk.Query(operateId, define.QueryFarmRecord)
		if err != nil {
			errStr := fmt.Sprintf("FormQuery  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询农事操作----Unmarshal json----------")
		err = json.Unmarshal(farmRecordData, farmRecord)
		if err != nil {
			errStr := fmt.Sprintf("json unmarshal  error : %v", err)
			utils.Response(errStr, c, http.StatusBadRequest)
			return
		}
		fmt.Println("-----------查询农事操作---unmarshal--success------")
		qRCodeResponse.FarmingRecordDetail.OpDetailList = append(qRCodeResponse.FarmingRecordDetail.OpDetailList, farmRecord.OpDetail)
	}

	// 按白名单筛选农事记录
	// TODO

	// 查询地块面积
	land := &define.Land{}
	fmt.Println("-----------查询地块-----------")
	landData, err := sdk.Query(plant.SectionId, define.QueryLand)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	fmt.Println("-----------查询地块---unmarshal--------")
	err = json.Unmarshal(landData, land)
	if err != nil {
		errStr := fmt.Sprintf("json unmarshal  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	fmt.Println("-----------查询地块---unmarshal--success------")
	// 查询农场
	farm := &define.Farm{}
	fmt.Println("-----------查询农场-----------")
	farmData, err := sdk.Query(land.FarmId, define.QueryFarm)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	fmt.Println("-----------查询农场---unmarshal--------")
	err = json.Unmarshal(farmData, farm)
	if err != nil {
		errStr := fmt.Sprintf("json unmarshal  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	fmt.Println("-----------查询农场---unmarshal--success------")
	qRCodeResponse.FarmingRecordDetail.SectionInfo = land.SectionInfo

	qRCodeResponse.Province = land.Province
	qRCodeResponse.City = land.City
	qRCodeResponse.District = land.District
	qRCodeResponse.Address = land.Address
	qRCodeResponse.DetailedAddress = land.DetailedAddress
	qRCodeResponse.InspectionReport.QualityReportImageUrl = processInfo.QualityReportImageUrl
	qRCodeResponse.InspectionReport.QualityInspectionDate = processInfo.QualityInspectionDate
	qRCodeResponse.InspectionReport.SafeReportImageUrl = processInfo.SafeReportImageUrl
	qRCodeResponse.InspectionReport.SafeInspectionDate = processInfo.SafeInspectionDate
	if len(farm.WaterTestReportList) > 0 {
		for _, w := range farm.WaterTestReportList {
			qRCodeResponse.InspectionReport.WaterReportImageUrl += "," + w.ImageUrl
		}
		qRCodeResponse.InspectionReport.WaterReportImageUrl = qRCodeResponse.InspectionReport.WaterReportImageUrl[1:]
		qRCodeResponse.InspectionReport.WaterInspectionDate = farm.WaterTestReportList[0].InspectionDate
	}
	if len(farm.SoilTestReportList) > 0 {
		for _, s := range farm.SoilTestReportList {
			qRCodeResponse.InspectionReport.SoilReportImageUrl += "," + s.ImageUrl
		}
		qRCodeResponse.InspectionReport.SoilReportImageUrl = qRCodeResponse.InspectionReport.SoilReportImageUrl[1:]
		qRCodeResponse.InspectionReport.SoilInspectionDate = farm.SoilTestReportList[0].InspectionDate
	}
	qRCodeResponse.InspectionReport.BlockInfo.TxId = processInfo.TxId
	qRCodeResponse.InspectionReport.BlockInfo.BlockNumber, qRCodeResponse.InspectionReport.BlockInfo.TxCreateTime, err = sdk.GetBlockInfoByTxID(qRCodeResponse.InspectionReport.BlockInfo.TxId)

	res, err := json.Marshal(qRCodeResponse)
	if err != nil {
		errStr := fmt.Sprintf("json marshal  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryUploadRecord(c *gin.Context) {
	key := c.Param("uuid")
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryUploadRecord)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryUploadRecords(c *gin.Context) {
	key := c.Param("uuid")
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryUploadRecords)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryFarm(c *gin.Context) {
	key := c.Param("uuid")
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryFarm)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryLand(c *gin.Context) {
	key := c.Param("uuid")
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryLand)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryPlant(c *gin.Context) {
	key := c.Param("uuid")
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryPlant)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryFarmRecord(c *gin.Context) {
	key := c.Param("uuid")
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryFarmRecord)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryHarvest(c *gin.Context) {
	key := c.Param("uuid")
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryHarvest)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryDryInfo(c *gin.Context) {
	key := c.Param("uuid")
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryDryInfo)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryWarehouseInfo(c *gin.Context) {
	key := c.Param("uuid")
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryWarehouseInfo)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryProcessInfo(c *gin.Context) {
	key := c.Param("uuid")
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryProcessInfo)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryFeedSectionMap(c *gin.Context) {
	key := c.Param("uuid")
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryFeedSectionCrop)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryDict(c *gin.Context) {
	key := c.Param("uuid")
	if len(key) == 0 {
		key = "AgDict"
	}
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryDict)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func QueryPesticideFertilizer(c *gin.Context) {
	key := c.Param("uuid")
	if len(key) == 0 {
		key = "AgPesticideFertilizer"
	}
	utils.Logger.Debugf("query key: %v\n", key)
	res, err := sdk.Query(key, define.QueryPesticideFertilizer)
	if err != nil {
		errStr := fmt.Sprintf("FormQuery  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(res), c, http.StatusOK)
}

func GetDataByTxId(c *gin.Context) {
	key := c.Param("uuid")
	txData, err := sdk.GetTransactionByID(key, "")
	if err != nil {
		errStr := fmt.Sprintf("GetdataByTxid query  error : %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
		return
	}
	utils.Response(string(txData), c, http.StatusOK)
}

// 判断字符串是否在字符串列表中
func in(target string, strArray []string) bool {
	sort.Strings(strArray)
	i := sort.SearchStrings(strArray, target)
	if i < len(strArray) && strArray[i] == target {
		return true
	}
	return false
}
