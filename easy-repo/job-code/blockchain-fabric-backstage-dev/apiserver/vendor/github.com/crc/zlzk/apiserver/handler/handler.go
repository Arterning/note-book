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
	uploadRecord := &define.ThirdPartUploadRecord{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, uploadRecord); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	// 查询历史数据并校验，暂时不做
	data, err := json.Marshal(uploadRecord)
	if err != nil {
		errStr := fmt.Sprintf("json marshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	// 更新链上数据
	info := []string{"PutUploadRecord", string(data)}
	utils.Logger.Debugf("put key: %v\n", info[0])
	resp, err := gohfc.GetHandler().SyncInvoke(info, "", "")
	if err != nil {
		errStr := fmt.Sprintf("数据上链失败 %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	utils.Response(resp.TxID, c, http.StatusOK)
	return
}

func PutFarm(c *gin.Context) {
	farm := &define.Farm{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, farm); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
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
	utils.Response(resp.TxID, c, http.StatusOK)
	return
}

func PutLand(c *gin.Context) {
	land := &define.Land{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, land); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
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
	utils.Response(resp.TxID, c, http.StatusOK)
	return
}

func PutPlant(c *gin.Context) {
	plant := &define.Plant{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, plant); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
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
	utils.Response(resp.TxID, c, http.StatusOK)
	return
}

func PutFarmRecord(c *gin.Context) {
	farmRecord := &define.FarmRecord{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, farmRecord); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
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
	utils.Response(resp.TxID, c, http.StatusOK)
	return
}

func PutHarvest(c *gin.Context) {
	harvest := &define.Harvest{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, harvest); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
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
	utils.Response(resp.TxID, c, http.StatusOK)
	return
}

func PutDryInfo(c *gin.Context) {
	dryInfo := &define.DryInfo{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, dryInfo); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
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
	utils.Response(resp.TxID, c, http.StatusOK)
	return
}

func PutWarehouseInfo(c *gin.Context) {
	warehouseInfo := &define.WarehouseInfo{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, warehouseInfo); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
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
	utils.Response(resp.TxID, c, http.StatusOK)
	return
}

func PutProcessInfo(c *gin.Context) {
	processInfo := &define.ProcessInfo{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errStr := fmt.Sprintf("get body error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, processInfo); err != nil {
		errStr := fmt.Sprintf("json unmarshal error %v", err)
		utils.Response(errStr, c, http.StatusBadRequest)
	}
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
	utils.Response(resp.TxID, c, http.StatusOK)
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

	fmt.Printf("kafkaMessage: %v\n",kafkaMessage)
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

		for _, land := range farm.Data {
			wg.Add(1)
			f := &define.Farm{}
			f.TenantId = reflect.ValueOf(land["tenant_id"]).String()
			f.FarmId = reflect.ValueOf(land["id"]).String()
			isActive := reflect.ValueOf(land["is_active"]).String()

			switch farm.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(f.FarmId, define.QueryFarm)
					if err != nil || len(d) == 0{
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
						case <- ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				go func(){
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
						case <- ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				for _, farmOld := range farm.Old {
					_, okTenantId := farmOld["tenant_id"]
					_, okFarmId := farmOld["id"]
					_, okIsActive := farmOld["is_active"]
					go func() {
						farmId := ""
						if okFarmId {
							farmId = reflect.ValueOf(farmOld["id"]).String()
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
							if okIsActive && isActive =="N"{
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
							case <- ctx.Done():
								return
							}
						}
					}()
				}
			}
		}
		wg.Wait()
		cancel()
	case "ag_section":
		for _, land := range farm.Data {
			f := &define.Farm{}
			f.TenantId = reflect.ValueOf(land["tenant_id"]).String()
			f.FarmId = reflect.ValueOf(land["farm_id"]).String()
			sectionId := reflect.ValueOf(land["id"]).String()
			isActive := reflect.ValueOf(land["is_active"]).String()

			switch farm.Type {
			case "DELETE":
				d, err := sdk.Query(f.FarmId, define.QueryFarm)
				if err != nil || len(d) == 0{
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
					if old.SectionIdList[i] == sectionId  {
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
					if old.Status == 0 && !in(sectionId, old.SectionIdList){
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
				for _, oldSection := range farm.Old {
					_, okSectionId := oldSection["id"]
					_, okTenantId := oldSection["tenant_id"]
					_, okFarmId := oldSection["farm_id"]
					_, okIsActive := oldSection["is_active"]
					farmId := ""
					if okFarmId {
						farmId = reflect.ValueOf(oldSection["farm_id"]).String()
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
								if old.SectionIdList[i] == sectionId  {
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
							oldSectionId = reflect.ValueOf(oldSection["id"]).String()
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
		}
	case "ag_test_report":
		for _, land := range farm.Data {
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
				if err != nil || len(f) == 0{
					return "", err
				}
				farm := &define.Farm{}
				err = json.Unmarshal(f, &farm)
				if err != nil {
					return "", err
				}
				if farm.Status == 0 {
					if testReportType == "1" {
						for i := 0; i < len(farm.SoilTestReportList); i++ {
							if farm.SoilTestReportList[i].ImageUrl == imageUrl  {
								farm.SoilTestReportList = append(farm.SoilTestReportList[:i], farm.SoilTestReportList[i+1:]...)
								i--
							}
						}
					} else if testReportType == "2" {
						for i := 0; i < len(farm.WaterTestReportList); i++ {
							if farm.WaterTestReportList[i].ImageUrl == imageUrl  {
								farm.WaterTestReportList = append(farm.WaterTestReportList[:i], farm.WaterTestReportList[i+1:]...)
								i--
							}
						}
					}
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
			case "INSERT":
				f, err := sdk.Query(farmId, define.QueryFarm)
				if err != nil || len(f) == 0{
					return "", err
				}
				farm := &define.Farm{}
				err = json.Unmarshal(f, &farm)
				if err != nil {
					return "", err
				}
				if farm.Status == 0 {
					testReportList := &define.TestReport{}
					testReportList.InspectionDate = beginDate
					testReportList.InspectionEndDate = endDate
					testReportList.ImageUrl = imageUrl
					if testReportType == "1" {
						existFlag := false
						for _, v := range farm.SoilTestReportList {
							if v.ImageUrl == imageUrl {
								existFlag = true
							}
						}
						if !existFlag {
							farm.SoilTestReportList = append(farm.SoilTestReportList, *testReportList)
						}
					} else if testReportType == "2" {
						existFlag := false
						for _, v := range farm.WaterTestReportList {
							if v.ImageUrl == imageUrl {
								existFlag = true
							}
						}
						if !existFlag {
							farm.WaterTestReportList = append(farm.WaterTestReportList, *testReportList)
						}
					}
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
			case "UPDATE":
				for _, oldLandSection := range farm.Old {
					_, okFarmId := oldLandSection["farm_id"]
					_, okTestReportType := oldLandSection["type"]
					_, okBeginDate := oldLandSection["begin_date"]
					_, okEndDate := oldLandSection["end_date"]
					_, okImageUrl := oldLandSection["attachment_id"]
					_, okIsActive := oldLandSection["is_active"]
					oldFarmId := ""
					if okFarmId {
						oldFarmId = reflect.ValueOf(oldLandSection["farm_id"]).String()
					} else {
						oldFarmId = farmId
					}
					f, err := sdk.Query(oldFarmId, define.QueryFarm)
					if err != nil{
						return "", err
					}
					farm := &define.Farm{}
					if len(f) > 0 {
						err = json.Unmarshal(f, &farm)
						if err != nil {
							return "", err
						}
						if okIsActive && isActive == "N" {
							if farm.Status == 0 {
								if testReportType == "1" {
									for i := 0; i < len(farm.SoilTestReportList); i++ {
										if farm.SoilTestReportList[i].ImageUrl == imageUrl {
											farm.SoilTestReportList = append(farm.SoilTestReportList[:i], farm.SoilTestReportList[i+1:]...)
											i--
										}
									}
								} else if testReportType == "2" {
									for i := 0; i < len(farm.WaterTestReportList); i++ {
										if farm.WaterTestReportList[i].ImageUrl == imageUrl {
											farm.WaterTestReportList = append(farm.WaterTestReportList[:i], farm.WaterTestReportList[i+1:]...)
											i--
										}
									}
								}
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
								continue
							}
						}
						if farm.Status == 0 {
							oldImageUrl := ""
							if okImageUrl {
								oldImageUrl = reflect.ValueOf(oldLandSection["attachment_id"]).String()
							}
							if testReportType == "1" {
								if len(farm.SoilTestReportList) > 0 {
									existFlag := false
									for i, v := range farm.SoilTestReportList {
										if v.ImageUrl == oldImageUrl || v.ImageUrl == imageUrl {
											farm.SoilTestReportList[i].ImageUrl = imageUrl
											if okBeginDate {
												farm.SoilTestReportList[i].InspectionDate = beginDate
											}
											if okEndDate {
												farm.SoilTestReportList[i].InspectionEndDate = endDate
											}
											existFlag = true
										}
									}
									if !existFlag {
										testReport := define.TestReport{
											ImageUrl:imageUrl,
											InspectionDate: beginDate,
											InspectionEndDate: endDate,
										}
										farm.SoilTestReportList = append(farm.SoilTestReportList, testReport)
									}
								} else {
									var testReport define.TestReport
									testReport.ImageUrl = imageUrl
									testReport.InspectionDate = beginDate
									testReport.InspectionEndDate = endDate
									farm.SoilTestReportList = append(farm.SoilTestReportList, testReport)
								}
							} else if testReportType == "2" {
								if len(farm.WaterTestReportList) > 0 {
									existFlag := false
									for i, v := range farm.WaterTestReportList {
										if v.ImageUrl == oldImageUrl || v.ImageUrl == imageUrl{
											farm.WaterTestReportList[i].ImageUrl = imageUrl
											if okBeginDate {
												farm.WaterTestReportList[i].InspectionDate = beginDate
											}
											if okEndDate {
												farm.WaterTestReportList[i].InspectionEndDate = endDate
											}
											existFlag = true
										}
									}
									if !existFlag {
										testReport := define.TestReport{
											ImageUrl:imageUrl,
											InspectionDate: beginDate,
											InspectionEndDate: endDate,
										}
										farm.WaterTestReportList = append(farm.WaterTestReportList, testReport)
									}
								} else {
									var testReport define.TestReport
									testReport.ImageUrl = imageUrl
									testReport.InspectionDate = beginDate
									testReport.InspectionEndDate = endDate
									farm.WaterTestReportList = append(farm.WaterTestReportList, testReport)
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
						farm.FarmId = farmId
						if testReportType == "1" {
							var testReport define.TestReport
							testReport.ImageUrl = imageUrl
							testReport.InspectionDate = beginDate
							testReport.InspectionEndDate = endDate
							farm.SoilTestReportList = append(farm.SoilTestReportList, testReport)
						} else if testReportType == "2" {
							var testReport define.TestReport
							testReport.ImageUrl = imageUrl
							testReport.InspectionDate = beginDate
							testReport.InspectionEndDate = endDate
							farm.WaterTestReportList = append(farm.WaterTestReportList, testReport)
						}
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

		for _, land := range landSection.Data {
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
					if err != nil || len(d) == 0{
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
						case <- ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				go func(){
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
						case <- ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				for _, oldLandSection := range landSection.Old {
					go func() {
						_, ok := oldLandSection["id"]
						_, okFarmId := oldLandSection["farm_id"]
						_, okName := oldLandSection["name"]
						_, okProvince := oldLandSection["province"]
						_, okCity := oldLandSection["city"]
						_, okDistrict := oldLandSection["district"]
						_, okAddress := oldLandSection["address"]
						_, okDetailedAddress := oldLandSection["detailed_address"]
						_, okImgUrl := oldLandSection["img_url"]
						_, okArea := oldLandSection["area"]
						_, okIsActive := oldLandSection["is_active"]
						sectionId := ""
						if ok {
							sectionId = reflect.ValueOf(oldLandSection["id"]).String()
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
							case <- ctx.Done():
								return
							}
						}
					}()
				}
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
func HandlePlant(plantData apiDefine.KafkaMessage) (string,error) {
	response := ""
	switch plantData.Table {
	case "ag_farming_record":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup

		// 农事记录id、农事操作人、处理时间
		for _, plant := range plantData.Data {
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
					if err != nil || len(d) == 0{
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
					if err != nil || len(f) == 0{
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
						case <- ctx.Done():
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
						case <- ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				for _, oldPlantData := range plantData.Old {
					_, okFarmRecordId := oldPlantData["id"]
					_, okSectionId := oldPlantData["section_id"]
					_, okSectionCropsId := oldPlantData["section_crops_id"]
					_, okHandleDate := oldPlantData["handle_date"]
					_, okOperatorName := oldPlantData["operator_name"]
					_, okStageCode := oldPlantData["stage_code"]
					_, okIsActive := oldPlantData["is_active"]
					go func() {
						oldFarmRecordId := ""
						if okFarmRecordId {
							oldFarmRecordId = reflect.ValueOf(oldPlantData["id"]).String()
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
						}else {
							sectionCropsId = l.SectionCropsId
						}
						p, err := sdk.Query(sectionCropsId, define.QueryPlant)
						if err != nil || len(p) == 0{
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
							oldPlant.SectionId = reflect.ValueOf(oldPlantData["section_id"]).String()
						}
						for i, v := range oldPlant.FarmRecordIdList {
							if v == reflect.ValueOf(oldPlantData["id"]).String() {
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
							case <- ctx.Done():
								return
							}
						}
					}()
				}
			}
		}
		if plantData.Type == "INSERT" {
			type  SectionPlant struct {
				SectionId string
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
						sectionCropsMap[sectionCropsId] = SectionPlant{SectionId: SecId, SectionCropsList:v.SectionCropsList}
					}
				} else {
					sectionCropsMap[sectionCropsId] = SectionPlant{SectionId: SecId, SectionCropsList:[]string{farmRecordId}}
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
						case <- ctx.Done():
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
		for _, plant := range plantData.Data {
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
					_, err = gohfc.GetHandler().SyncInvoke(info, "", "")
					if err != nil {
						wg.Done()
						return
					}
					p, err := sdk.Query(old.SectionCropsId, define.QueryPlant)
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
							if len(stageCode) < 5 {
								// stage_rice_16_34_prepare
								wg.Done()
								return
							}
							switch stageCode[4] {
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
				go func() {
					// 更新farmRecord上链数据
					//f.OperationCode = l.OperationCode

					// 使用物品需要从字典中获取映射值
					switch r.TypeCode {
					case "pesticideType", "fertilizerName": // 农药或化肥
						pesticideFertilizer, err := sdk.Query("AgPesticideFertilizer", define.QueryPesticideFertilizer)
						if err != nil || len(pesticideFertilizer) == 0 {
							//err := fmt.Errorf("数据不在链上，或查询链上数据失败：%v", err)
							wg.Done()
							return
						}
						agPesticideFertilizerInfo := &define.AgPesticideFertilizerInfo{}
						err = json.Unmarshal(pesticideFertilizer, agPesticideFertilizerInfo)
						if err != nil {
							wg.Done()
							return
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
							wg.Done()
							return
						}
						agDictInfo := &define.AgDictInfo{}
						err = json.Unmarshal(dict, agDictInfo)
						if err != nil {
							wg.Done()
							return
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
						wg.Done()
						return
					}
					// 更新plant上链数据
					fmt.Printf("")
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
				dd, err := sdk.Query(f.SectionCropsId, define.QueryPlant)
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
				if len(stageCode) < 5 {
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
				switch stageCode[4] {
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
				response = res.TxID

			case "UPDATE":
				for _, oldPlantData := range plantData.Old {
					_, okOperateId := oldPlantData["record_id"]
					//_, okOperationCode := oldPlantData["operation_code"]
					_, okRecordDetailId := oldPlantData["id"]
					_, okUseType := oldPlantData["use_type"]
					_, okUseGood := oldPlantData["use_good"]
					_, okTypeCode := oldPlantData["type_code"]
					_, okStartValue := oldPlantData["start_value"]
					_, okEndValue := oldPlantData["end_value"]
					_, okIsActive := oldPlantData["is_active"]
					go func() {
						operateId := ""
						if okOperateId {
							operateId = reflect.ValueOf(oldPlantData["record_id"]).String()
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
								if old.RecordDetailList[i].Id == r.Id  {
									old.RecordDetailList = append(old.RecordDetailList[:i],old.RecordDetailList[i+1:]...)
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
							for i, recordDetailId := range old.RecordDetailList{
								if recordDetailId.Id == r.Id {
									if okRecordDetailId {
										old.RecordDetailList[i].Id = r.Id
									}
									if okUseType {
										old.RecordDetailList[i].UseType = r.UseType
									}
									if okUseGood {
										old.RecordDetailList[i].UseGood = r.UseGood
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
							if okRecordDetailId || okOperateId || okUseType || okUseGood || okTypeCode || okStartValue ||
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

						p, err := sdk.Query(old.SectionCropsId, define.QueryPlant)
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
									if len(stageCode) < 5 {
										// stage_rice_16_34_prepare
										wg.Done()
										return
									}
									switch stageCode[4] {
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
									if v.TypeCode == q.TypeCode {
										old.RecordDetailList[j].UseGood = v.Name
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
						//fmt.Printf("old.OperationCode=%v\n", old.OperationCode)
						if len(stageCode) < 5 {
							// stage_rice_16_34_prepare
							wg.Done()
							return
						}
						//fmt.Printf("len(stateCode)=%d, stateCode=%v\n", len(stateCode), stateCode)
						switch stageCode[4] {
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
						response = res.TxID
						wg.Done()
						for {
							select {
							case <- ctx.Done():
								return
							}
						}
					}()
				}
			}
		}
		wg.Wait()
		cancel()
	case "sys_attachment":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		// 更新图片，subpath为压缩图、path为原始图
		for _, plant := range plantData.Data {
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
				go func(){
					d, err := sdk.Query(l.OperateId, define.QueryFarmRecord)
					if err != nil || len(d) == 0{
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
							if f.SubImageUrl[i] == subPath  {
								f.SubImageUrl = append(f.SubImageUrl[:i], f.SubImageUrl[i+1:]...)
								i--
							}
						}
						for i := 0; i < len(f.ImageUrl); i++ {
							if f.ImageUrl[i] == path  {
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
					dd, err := sdk.Query(f.SectionCropsId, define.QueryPlant)
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
					if len(stageCode) < 5 {
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
					switch stageCode[4] {
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
					}
					response = res.TxID
					wg.Done()
					for {
						select {
						case <- ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				d, err := sdk.Query(l.OperateId, define.QueryFarmRecord)
				if err != nil || len(d) == 0{
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
						case <- ctx.Done():
							return
						}
					}
				}()
				// 更新种植数据
				// 查询种植信息
				dd, err := sdk.Query(f.SectionCropsId, define.QueryPlant)
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
				if len(stageCode) < 5 {
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
				switch stageCode[4] {
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
				response = res.TxID

			// 更新path
			case "UPDATE":
				for _, oldPlantData := range plantData.Old {
					_, okOperateId := oldPlantData["pid"]
					//_, okFileClass := oldPlantData["file_class"]
					_, okSubPath := oldPlantData["sub_path"]
					_, okPath := oldPlantData["id"]
					_, okIsActive := oldPlantData["is_active"]
					go func() {
						operateId := ""
						if okOperateId {
							operateId = reflect.ValueOf(oldPlantData["pid"]).String()
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
							if len(r) == 0 && isActive == "Y"{
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
									if f.SubImageUrl[i] == subPath  {
										f.SubImageUrl = append(f.SubImageUrl[:i], f.SubImageUrl[i+1:]...)
										i--
									}
								}
								for i := 0; i < len(f.ImageUrl); i++ {
									if f.ImageUrl[i] == path  {
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
									oldSubPath = reflect.ValueOf(oldPlantData["sub_path"]).String()
								}
								oldPath := ""
								if okPath {
									oldPath = reflect.ValueOf(oldPlantData["id"]).String()
								}
								if in(oldSubPath, f.SubImageUrl) {
									for i, v := range f.SubImageUrl {
										if v == oldSubPath {
											f.SubImageUrl[i] = subPath
										}
									}
								} else {
									if !in(subPath, f.SubImageUrl){
										f.SubImageUrl = append(f.SubImageUrl,subPath)
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
						dd, err := sdk.Query(f.SectionCropsId, define.QueryPlant)
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
							if len(stageCode) < 5 {
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
							switch stageCode[4] {
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
						if len(stageCode) < 5 {
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
						switch stageCode[4] {
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
						for {
							select {
							case <- ctx.Done():
								return
							}
						}
					}()
				}
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
func HandleHarvest(harvestData apiDefine.KafkaMessage) (string,error) {
	response := ""
	switch harvestData.Table {
	case "ag_reap_batch":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		// 收割批次id（ag_reap_batch.id）、收割时间、Txid
		for _, harvest := range harvestData.Data {
			wg.Add(1)
			h := &define.Harvest{}
			h.HarvestId = reflect.ValueOf(harvest["id"]).String()
			h.ReapCreateDate, _ = time.Parse("2006-01-02 15:04:05", reflect.ValueOf(harvest["handle_date"]).String())
			isActive := reflect.ValueOf(harvest["is_active"]).String()
			switch harvestData.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(h.HarvestId, define.QueryHarvest)
					if err != nil || len(d) == 0{
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
						case <- ctx.Done():
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
						case <- ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				for _, oldHarvestData := range harvestData.Old {
					_, okHarvestId := oldHarvestData["id"]
					_, okReapCreateDate := oldHarvestData["create_date"]
					_, okIsActive := oldHarvestData["is_active"]
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
							case <- ctx.Done():
								return
							}
						}
					}()
				}
			}
		}
		wg.Wait()
		cancel()
	case "ag_section_crops_reap":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		// 种植id
		for _, harvest := range harvestData.Data {
			wg.Add(1)
			h := &define.Harvest{}
			h.HarvestId = reflect.ValueOf(harvest["reap_batch_id"]).String()
			sectionCropsId := reflect.ValueOf(harvest["section_crops_id"]).String()
			isActive := reflect.ValueOf(harvest["is_active"]).String()
			switch harvestData.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(h.HarvestId, define.QueryHarvest)
					if err != nil || len(d) == 0{
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
						if old.PlantIdList[i] == sectionCropsId  {
							old.PlantIdList = append(old.PlantIdList[:i],old.PlantIdList[i+1:]...)
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
						case <- ctx.Done():
							return
						}
					}
				}()

			case "INSERT":
				d, err := sdk.Query(h.HarvestId, define.QueryHarvest)
				if err != nil || len(d) == 0{
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
				for _, oldHarvestData := range harvestData.Old {
					_, okHarvestId := oldHarvestData["reap_batch_id"]
					_, okSectionCropsId := oldHarvestData["section_crops_id"]
					_, okIsActive := oldHarvestData["isActive"]
					go func() {
						harvestId := ""
						if okHarvestId {
							harvestId = reflect.ValueOf(oldHarvestData["reap_batch_id"]).String()
						} else {
							harvestId = h.HarvestId
						}
						d, err := sdk.Query(harvestId, define.QueryHarvest)
						if err != nil || len(d) == 0{
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
								if old.PlantIdList[i] == sectionCropsId  {
									old.PlantIdList = append(old.PlantIdList[:i],old.PlantIdList[i+1:]...)
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
							oldSectionId := reflect.ValueOf(oldHarvestData["section_crops_id"]).String()
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
							case <- ctx.Done():
								return
							}
						}
					}()
				}
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
func HandleDry(dryData apiDefine.KafkaMessage) (string,error) {
	response := ""
	switch dryData.Table {
	case "srv_dry_info":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		for _, dry := range dryData.Data {
			wg.Add(1)
			l := &define.DryInfo{}
			l.DryId = reflect.ValueOf(dry["id"]).String()
			l.HarvestId = reflect.ValueOf(dry["sg_id"]).String() // 收割id
			l.HgDate, _= time.Parse("2006-01-02 15:04:05", reflect.ValueOf(dry["hg_date"]).String())
			l.HgFactoryName = reflect.ValueOf(dry["hg_factory_name"]).String()
			l.ImageUrl = reflect.ValueOf(dry["attachment"]).String() // json 字符串
			isActive := reflect.ValueOf(dry["is_active"]).String()
			switch dryData.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(l.DryId, define.QueryDryInfo)
					if err != nil || len(d) == 0{
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
						case <- ctx.Done():
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
						case <- ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				for _, oldDryData := range dryData.Old {
					_, okDryId := oldDryData["id"]
					_, okHarvestId := oldDryData["sg_id"]
					_, okHgDate := oldDryData["hg_date"]
					_, okHgFactoryName := oldDryData["hg_factory_name"]
					_, okImageUrl := oldDryData["attachment"]
					_, okIsActive := oldDryData["is_active"]
					go func() {
						dryId := ""
						if okDryId {
							dryId = reflect.ValueOf(oldDryData["id"]).String()
						} else {
							dryId = l.DryId
						}
						d, err := sdk.Query(dryId, define.QueryDryInfo)
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
						if okDryId || okHarvestId || okHgDate || okHgFactoryName || okImageUrl {
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
						wg.Done()
						for {
							select {
							case <- ctx.Done():
								return
							}
						}
					}()
				}
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
func HandleWarehouse(warehouseData apiDefine.KafkaMessage) (string,error) {
	response := ""
	switch warehouseData.Table {
	case "srv_storage_info":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		for _, warehouse := range warehouseData.Data {
			wg.Add(1)
			l := &define.WarehouseInfo{}
			l.WarehouseId = reflect.ValueOf(warehouse["id"]).String()
			l.DryId = reflect.ValueOf(warehouse["hg_id"]).String()
			l.StartDate, _= time.Parse("2006-01-02 15:04:05", reflect.ValueOf(warehouse["start_date"]).String())
			l.DurationDate, _ = time.Parse("2006-01-02 15:04:05", reflect.ValueOf(warehouse["end_date"]).String())
			l.RefTemperature = reflect.ValueOf(warehouse["ref_temperature"]).String()
			l.CCFactoryName = reflect.ValueOf(warehouse["cc_factory_name"]).String()
			l.ImageUrl = reflect.ValueOf(warehouse["start_attachment"]).String() // json 字符串
			isActive := reflect.ValueOf(warehouse["is_active"]).String()
			switch warehouseData.Type {
			case "DELETE":
				go func() {
					d, err := sdk.Query(l.WarehouseId, define.QueryWarehouseInfo)
					if err != nil || len(d) == 0{
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
						case <- ctx.Done():
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
						case <- ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				for _, oldWarehouseData := range warehouseData.Old {
					_, okWarehouseId := oldWarehouseData["id"]
					_, okDryId := oldWarehouseData["hg_id"]
					_, okStartDate := oldWarehouseData["start_date"]
					_, okDurationDate := oldWarehouseData["end_date"]
					_, okRefTemperature := oldWarehouseData["ref_temperature"]
					_, okCCFactoryName := oldWarehouseData["cc_factory_name"]
					_, okImageUrl := oldWarehouseData["start_attachment"]
					_, okIsActive := oldWarehouseData["is_active"]
					go func() {
						warehouseId := ""
						if okWarehouseId {
							warehouseId = reflect.ValueOf(oldWarehouseData["id"]).String()
						} else {
							warehouseId = l.WarehouseId
						}
						d, err := sdk.Query(warehouseId, define.QueryWarehouseInfo)
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
						wg.Done()
						for {
							select {
							case <- ctx.Done():
								return
							}
						}
					}()
				}
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
func HandleProcess(processData apiDefine.KafkaMessage) (string,error) {
	response := ""
	switch processData.Table {
	case "srv_mach_packing":
		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		for _, process := range processData.Data {
			wg.Add(1)
			l := &define.ProcessInfo{}
			l.ProcessId = reflect.ValueOf(process["id"]).String()
			l.PreType = reflect.ValueOf(process["up_flow_type"]).String() // 上一环节类型仓储CC，烘干HG
			l.PreId = reflect.ValueOf(process["before_business_id"]).String() // 上一环节id
			l.ProcessDate, _ = time.Parse("2006-01-02 15:04:05", reflect.ValueOf(process["jg_date"]).String())
			l.ProcessEnterpriseName = reflect.ValueOf(process["jgc_enterprise_name"]).String() // 加工厂名称
			l.ProcessImageUrl = reflect.ValueOf(process["jg_attachment"]).String() // json字符串
			l.QualityReportImageUrl = reflect.ValueOf(process["pzjc_attachment"]).String() // json字符串
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
					if err != nil || len(d) == 0{
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
						case <- ctx.Done():
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
						case <- ctx.Done():
							return
						}
					}
				}()

			case "UPDATE":
				for _, oldProcessData := range processData.Old {
					_, okProcessId := oldProcessData["id"]
					_, okPreType := oldProcessData["up_flow_type"]
					_, okPreId := oldProcessData["before_business_id"]
					_, okProcessDate := oldProcessData["jg_date"]
					_, okProcessEnterpriseName := oldProcessData["jgc_enterprise_name"]
					_, okProcessImageUrl := oldProcessData["jg_attachment"]
					_, okQualityReportImageUrl := oldProcessData["pzjc_attachment"]
					_, okQualityInspectionDate := oldProcessData["pzjc_date"]
					_, okSafeReportImageUrl := oldProcessData["aqjc_attachment"]
					_, okSafeInspectionDate := oldProcessData["aqjc_date"]
					_, okPackDate := oldProcessData["pack_date"]
					_, okPackageImageUrl := oldProcessData["pack_attachment"]
					_, okIsActive := oldProcessData["is_active"]
					go func() {
						processId := ""
						if okProcessId {
							processId = reflect.ValueOf(oldProcessData["id"]).String()
						} else {
							processId = l.ProcessId
						}
						d, err := sdk.Query(processId, define.QueryProcessInfo)
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
						if okPreType {
							old.PreType = l.PreType
						}
						if okPreId {
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
						if okProcessId || okPreType || okPreId || okProcessDate || okProcessEnterpriseName ||
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
						wg.Done()
						for {
							select {
							case <- ctx.Done():
								return
							}
						}
					}()
				}
			}
		}
		wg.Wait()
		cancel()
	case "srv_mach_packing_norms":
	}
	return response, nil
}

// 数据字典
// key：AgDict
// value：id name value type_code
// 农事操作具体名字，比如灌溉，是喷灌、滴灌、还是满灌。
// 通过ag_farming_record_detail.type_code==ag_dict.type_code &&
// ag_farming_record_detail.use_good==ag_dict.value => ag_farming_record_detail.type_use=ag_dict.name
func HandleDict(dictData apiDefine.KafkaMessage) (string,error) {
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
				} else{
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

		for _, dict := range dictData.Data {
			dictId, _ := strconv.Atoi(reflect.ValueOf(dict["id"]).String())
			name := reflect.ValueOf(dict["name"]).String()
			value, _ := strconv.Atoi(reflect.ValueOf(dict["value"]).String())
			typeCode := reflect.ValueOf(dict["type_code"]).String()
			isActive := reflect.ValueOf(dict["is_active"]).String()
			switch dictData.Type {
			case "DELETE":
				d, err := sdk.Query("AgDict", define.QueryDict)
				if err != nil || len(d) == 0{
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
					if old.AgDictList[i].Id == int64(dictId)  {
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
				for _, oldDictData := range dictData.Old {
					_, okIsActive := oldDictData["is_active"]
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
								if old.AgDictList[i].Id == int64(dictId)  {
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
	}
	return response, nil
}

// 数据字典
// key：AgPesticideFertilizer
// value：id name type type_code value unit
// 化肥、农药的具体名字，计量单位。
// 通过ag_farming_record_detail.type_code==ag_pesticide_fertilizer.type_code &&
// ag_farming_record_detail.use_good==ag_pesticide_fertilizer.value => ag_farming_record_detail.type_use=ag_pesticide_fertilizer.name
func HandlePesticideFertilizer(pesticideFertilizerData apiDefine.KafkaMessage) (string,error) {
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
					if !existFlag{
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
		for _, pfData := range pesticideFertilizerData.Data {
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
				if err != nil || len(d) == 0{
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
					if old.AgPesticideFertilizerList[i].Id == pfId  {
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
				for _, oldPesticideFertilizerData := range pesticideFertilizerData.Old {
					_, okIsActive := oldPesticideFertilizerData["is_active"]
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
								if old.AgPesticideFertilizerList[i].Id == pfId  {
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
						if !existFlag{
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
	qRCodeResponse.ProcessRecordDetail.Process=processInfo.Process
	qRCodeResponse.ProcessRecordDetail.Package=processInfo.Package
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
	plant := &define.Plant{}
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
	fmt.Println("-----------查询种植---unmarshal--success------")
	qRCodeResponse.FarmingRecordDetail.PlantFarm = plant.PlantFarm
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
	qRCodeResponse.ReapCreateDate = harvest.ReapCreateDate
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
		qRCodeResponse.InspectionReport.WaterReportImageUrl = farm.WaterTestReportList[0].ImageUrl
		qRCodeResponse.InspectionReport.WaterInspectionDate = farm.WaterTestReportList[0].InspectionDate
	}
	if len(farm.SoilTestReportList) > 0 {
		qRCodeResponse.InspectionReport.SoilReportImageUrl = farm.SoilTestReportList[0].ImageUrl
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
func in (target string, strArray []string ) bool {
	sort.Strings(strArray)
	i := sort.SearchStrings(strArray, target)
	if i < len(strArray) && strArray[i] == target {
		return true
	}
	return false
}
