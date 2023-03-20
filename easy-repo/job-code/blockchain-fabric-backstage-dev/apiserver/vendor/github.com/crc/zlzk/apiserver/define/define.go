package define

type Response struct {
	Code    int    `json:"ResponseCode"`
	Message string `json:"ResponseMsg"`
	Data    string `json:"data"`
}

type KafkaMessage struct {
	Data []map[string]interface{} `json:"data"` // 业务具体字段
	Database string `json:"database"` // 数据库表名
	ES int64 `json:"es"`
	ID int `json:"id"`
	IsDdl bool `json:"isDdl"`
	MysqlType interface{} `json:"mysqlType"` // mysql表字段
	Old []map[string]interface{} `json:"old"`
	PkNames interface{} `json:"pkNames"`
	Sql string `json:"sql"`
	SqlType interface{} `json:"sqlType"`
	Table string `json:"table"` // mysql表名称，ag_section_crops，
	Ts int64 `json:"ts"`
	Type string `json:"type"` // 数据库动作，更新、删除、新增。INSERT、DELETE、UPDATE
}
