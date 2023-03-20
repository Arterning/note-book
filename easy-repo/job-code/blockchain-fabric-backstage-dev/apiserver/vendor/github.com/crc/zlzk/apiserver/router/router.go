package router

import (
	"github.com/crc/zlzk/apiserver/handler"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"sync"
)

// Router 全局路由
var router *gin.Engine
var onceCreateRouter sync.Once

func GetRouter() *gin.Engine {
	onceCreateRouter.Do(func() {
		router = createRouter()
	})
	return router
}

func createRouter() *gin.Engine {
	router := gin.Default()
	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			secureMiddleware := secure.New(secure.Options{
				FrameDeny: true,
			})
			err := secureMiddleware.Process(c.Writer, c.Request)

			// If there was an error, do not continue.
			if err != nil {
				c.Abort()
				return
			}

			// Avoid header rewrite if response is a redirection.
			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()
	router.Use(secureFunc)
	route := router.Group("/")
	{
		route.POST("kafkaData", handler.KafkaData) // kafka消息上链
		route.POST("putFarm", handler.PutFarm) // 农场数据上链
		route.POST("putLand", handler.PutLand) // 地块数据上链
		route.POST("putPlant", handler.PutPlant) // 种植数据上链
		route.POST("putFarmRecord", handler.PutFarmRecord) // 农事操作数据上链
		route.POST("putHarvest", handler.PutHarvest) // 收割数据上链
		route.POST("putDryInfo", handler.PutDryInfo) // 烘干数据上链
		route.POST("putWarehouseInfo", handler.PutWarehouseInfo) // 仓储数据上链
		route.POST("putProcessInfo", handler.PutProcessInfo) // 加工数据上链
		route.POST("putDict", handler.PutDict) // 农事字典数据上链
		route.POST("putPesticideFertilizer", handler.PutPesticideFertilizer) // 农药、化肥字典数据上链
		route.POST("putUploadRecord", handler.PutUploadRecord)

		route.GET("/key/:uuid", handler.Query) // h5溯源查询
		route.GET("/queryFarm/:uuid", handler.QueryFarm) // 地块信息查询
		route.GET("/queryLand/:uuid", handler.QueryLand) // 地块信息查询
		route.GET("/queryPlant/:uuid", handler.QueryPlant) // 种植信息查询
		route.GET("/queryFarmRecord/:uuid", handler.QueryFarmRecord) // 农事记录信息查询
		route.GET("/queryHarvest/:uuid", handler.QueryHarvest) // 收割信息查询
		route.GET("/queryDryInfo/:uuid", handler.QueryDryInfo) // 烘干信息查询
		route.GET("/queryWarehouseInfo/:uuid", handler.QueryWarehouseInfo) // 仓储信息查询
		route.GET("/queryProcessInfo/:uuid", handler.QueryProcessInfo) // 加工信息查询
		route.GET("/queryDict/:uuid", handler.QueryDict) // 农事字典信息查询
		route.GET("/queryPesticideFertilizer/:uuid", handler.QueryPesticideFertilizer) // 农药、化肥字典信息查询
		route.GET("/queryUploadRecord/:uuid", handler.QueryUploadRecord)
		route.GET("/queryUploadRecords/:uuid", handler.QueryUploadRecords)
		route.GET("/Transaction/:uuid", handler.GetDataByTxId) // 按交易Id查询交易信息
	}
	return router
}
