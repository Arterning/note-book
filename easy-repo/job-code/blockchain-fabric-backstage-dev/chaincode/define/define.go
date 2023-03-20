package define

import "time"

type StatusType int

const (
	Query                    = "Query"
	QueryUploadRecord        = "QueryUploadRecord"  // 查询上传记录
	QueryUploadRecords       = "QueryUploadRecords" // 查询历史记录
	QueryFarm                = "QueryFarm"
	QueryLand                = "QueryLand"
	QueryPlant               = "QueryPlant"
	QueryFarmRecord          = "QueryFarmRecord"
	QueryHarvest             = "QueryHarvest"
	QueryDryInfo             = "QueryDryInfo"
	QueryWarehouseInfo       = "QueryWarehouseInfo"
	QueryProcessInfo         = "QueryProcessInfo"
	QueryFeedSectionCrop     = "QueryFeedSectionCrop"
	QueryDict                = "QueryDict"
	QueryPesticideFertilizer = "QueryPesticideFertilizer"

	PutUploadRecord        = "PutUploadRecord" // 上传记录
	PutFarm                = "PutFarm"
	PutLand                = "PutLand"
	PutPlant               = "PutPlant"
	PutFarmRecord          = "PutFarmRecord"
	PutHarvest             = "PutHarvest"
	PutDryInfo             = "PutDryInfo"
	PutWarehouseInfo       = "PutWarehouseInfo"
	PutProcessInfo         = "PutProcessInfo"
	PutFeedSectionCrop     = "PutFeedSectionCrop"
	PutDict                = "PutDict"
	PutPesticideFertilizer = "PutPesticideFertilizer"
)

// ThirdPartUploadRecord 第三方上传记录
type ThirdPartUploadRecord struct {
	UploadRecordId         string    `json:"upload_record_id"`          // 记录ID
	ThirdPartId            string    `json:"third_part_id"`             // 第三方企业ID
	ThirdPartName          string    `json:"third_part_name"`           // 第三方企业名称
	UploadBusinessId       string    `json:"upload_business_id"`        // 业务Id
	UploadBusinessModuleId string    `json:"upload_business_module_id"` // 业务模块Id
	UploadContext          string    `json:"upload_context"`            // 内容
	CreateDate             time.Time `json:"create_date"`               // 创建时间
	ModifyDate             time.Time `json:"modify_date"`               // 修改时间
	Creator                string    `json:"creator"`                   // 创建人
	Modifier               string    `json:"modifier"`                  // 修改人
	UploadKey              string    `json:"upload_key"`                // key
}

// 公司信息，农场编码、种植的农作物信息上链
type Tenant struct {
	TenantId   string    `json:"tenant_id"`    // 公司id
	FarmIdList []string  `json:"farm_id_list"` // 农场ID列表
	Crops      []AgCrops `json:"crops"`        // 种植的农作物信息
	Status     int       `json:"status"`       // 数据状态，0正常，1删除。
}

// AgCrops ag_crops种植的农作物信息
type AgCrops struct {
	Id         string    `json:"id"`          // 作物ID
	Name       string    `json:"name"`        // 作物名称
	IsActive   string    `json:"is_active"`   // 是否有效
	CreateDate time.Time `json:"create_date"` // 创建时间
	ModifyDate time.Time `json:"modify_date"` // 修改时间
	Creator    string    `json:"creator"`     // 创建人
	Modifier   string    `json:"modifier"`    // 修改人
	TenantId   string    `json:"tenant_id"`   // 公司id
}

// AgFarm ag_farm农场信息
type AgFarm struct {
	Id            string    `json:"id"`             // 农场ID
	IsActive      string    `json:"is_active"`      // 是否有效
	CreateDate    time.Time `json:"create_date"`    // 创建时间
	ModifyDate    time.Time `json:"modify_date"`    // 更新日期
	Creator       string    `json:"creator"`        // 创建人
	Modifier      string    `json:"modifier"`       // 更新人
	TenantId      string    `json:"tenant_id"`      // 公司ID
	FarmName      string    `json:"farm_name"`      // 农场名称
	FarmLocation  string    `json:"farm_location"`  // 农场地址
	FarmLongitude string    `json:"farm_longitude"` // 农场经度
	FarmLatitude  string    `json:"farm_latitude"`  // 农场维度
	Sn            int       `json:"sn"`             // 序号
	StandardId    int       `json:"standard_id"`    // 气象标准坐标
	ProvinceCode  string    `json:"province_code"`  // 省编码
	CountyCode    string    `json:"county_code"`
	CityCode      string    `json:"city_code"` // 市编码
	FarmType      string    `json:"farm_type"`
}

// AgFarmUser ag_farm_user农场用户信息
type AgFarmUser struct {
	Id          string    `json:"id"`          // id
	IsActive    string    `json:"is_active"`   // 是否有效
	CreateDate  time.Time `json:"create_date"` // 创建日期
	ModifyDate  time.Time `json:"modify_date"` // 更新日期
	Creator     string    `json:"creator"`     // 创建人
	Modifier    string    `json:"modifier"`    // 更新人
	TenantId    string    `json:"tenant_id"`   // 公司ID
	FarmId      string    `json:"farm_id"`     // 农场ID
	UserName    string    `json:"user_name"`   // 用户姓名
	UserType    string    `json:"user_type"`   // 1. 管理员；2.专家；0.普通用户
	Phone       string    `json:"phone"`       // 电话号码
	UserId      string    `json:"user_id"`     // 用户编号
	HeadImg     string    `json:"head_img"`    // 照片地址
	CreatorName string    `json:"create_name"` // 创建人姓名
	Sn          int       `json:"sn"`          // 排序号
}

// Varieties 农作物品类信息上链
type Varieties struct {
	AgVarietiesList []AgVarieties `json:"varieties_list"` // 农作物品类
}

// AgVarieties ag_varieties 农作物品类
type AgVarieties struct {
	Id          int       `json:"id"`           // 品种ID
	Name        string    `json:"name"`         // 品种名称
	CropId      int       `json:"crop_id"`      // 作物ID
	GenericSpec string    `json:"generic_spec"` // 基础属性
	IsActive    string    `json:"is_active"`    // 是否有效
	CreateDate  time.Time `json:"create_date"`  // 创建时间
	ModifyDate  time.Time `json:"modify_date"`  // 更新时间
	Creator     string    `json:"creator"`      // 创建人
	Modifier    string    `json:"modifier"`     // 更新人
	Sort        int       `json:"sort"`
	Status      int       `json:"status"` // 数据状态，0正常，1删除。
}

// 地块上链信息，土壤检测及检测结果、地块种植、农事记录数据上链。
type Section struct {
	Id             string          `json:"id"`           // 地块ID
	TenantId       string          `json:"tenant_id"`    // 公司id
	FarmId         string          `json:"farm_id"`      // 农场id
	SectionId      string          `json:"section_id"`   // 地块id
	AgSection      AgSection       `json:"section_info"` // 地块信息
	AgSoil         []AgSoil        `json:"soil"`         //
	AgSectionCrops AgSectionCrops  `json:"section_crops"`
	FarmingRecord  []FarmingRecord `json:"farming_record"`
	Status         int             `json:"status"` // 数据状态，0正常，1删除。
}

// AgSection ag_section地块信息
type AgSection struct {
	Id                 int64     `json:"id"`                  // 地块ID
	IsActive           string    `json:"is_active"`           // 是否有效
	CreateDate         time.Time `json:"create_date"`         // 创建日期
	ModifyDate         time.Time `json:"modify_date"`         // 更新日期
	Creator            string    `json:"creator"`             // 创建人
	Modifier           string    `json:"modifier"`            // 更新人
	TenantId           string    `json:"tenant_id"`           // 公司ID
	FarmId             string    `json:"farm_id"`             // 农场ID
	Code               string    `json:"code"`                // 地块编码
	Name               string    `json:"name"`                // 地块名称
	Status             string    `json:"status"`              // 状态
	Remark             string    `json:"remark"`              // 备注
	Mu                 float64   `json:"mu"`                  // 亩数
	Fen                float64   `json:"fen"`                 // 分数
	Area               float64   `json:"area"`                // 面积
	Circumference      float64   `json:"circumference"`       // 周长
	AdministrativeCode string    `json:"administrative_code"` // 行政编码
	ImgUrl             string    `json:"img_url"`             // 图片地址
	Province           string    `json:"province"`            // 省
	City               string    `json:"city"`                // 市
	District           string    `json:"district"`            // 区/县
	Address            string    `json:"address"`             // 地址
	DetailedAddress    string    `json:"detailed_address"`    // 详细地址
	StandardId         int       `json:"standard_id"`         // 气象标准坐标id
	BatchNumber        int       `json:"batch_number"`        // 批次号
}

// AgSoil ag_soil土壤检测
type AgSoil struct {
	Id           int       `json:"id"`             // id
	SectionId    string    `json:"section_id"`     // 地块ID
	ColPointCode int       `json:"col_point_code"` // 采集点编号
	CollectTime  string    `json:"collect_time"`   // 采样时间
	Area         string    `json:"area"`           // 种植面积
	Longitude    float64   `json:"longitude"`      // 经度
	Latitude     float64   `json:"latitude"`       // 纬度
	Depth        string    `json:"depth"`          // 取土深度
	Batch        string    `json:"batch"`          // 化验批次
	VarId        string    `json:"var_id"`         // 品种ID
	IsActive     string    `json:"is_active"`      // 是否有效
	CreateDate   time.Time `json:"create_date"`    // 创建时间
	ModifyDate   time.Time `json:"modify_date"`    // 更新时间
	Creator      string    `json:"creator"`        // 创建人
	Modifier     string    `json:"modifier"`       // 更新人
	FarmId       string    `json:"farm_id"`        // 农场ID
	AgSoilExt    AgSoilExt `json:"ag_soil_ext"`    // 检测结果
}

// AgSoilExt ag_soil_ext土壤检测结果
type AgSoilExt struct {
	Id            int     `json:"id"`              // id
	SoilId        int     `json:"soil_id"`         // 土壤id
	SoilElementId int     `json:"soil_element_id"` // 元素编码id
	ExtValue      float64 `json:"ext_value"`       // 指标值
	ExtLevel      string  `json:"ext_level"`       // 处于水平,1:低,2:中,3:高
	CreateData    string  `json:"create_data"`     // 入库时间
}

// AgSectionCrops ag_section_crops地块种植农作物信息
type AgSectionCrops struct {
	Id          string    `json:"id"`           // 种植id
	IsActive    string    `json:"is_active"`    // 是否有效
	CreateDate  time.Time `json:"create_date"`  // 创建日期
	ModifyDate  time.Time `json:"modify_date"`  // 更新日期
	Creator     string    `json:"creator"`      // 创建人
	Modifier    string    `json:"modifier"`     // 更新人
	TenantId    string    `json:"tenant_id"`    // 公司ID
	SectionId   string    `json:"section_id"`   // 地块ID
	CropsId     int       `json:"crops_id"`     // 作物id
	VarietiesId int       `json:"varieties_id"` // 品种id
	PlantDate   string    `json:"plant_date"`   // 种植时间
	Status      string    `json:"status"`       // 用来标记是否是最新种植作物 （Y是，N否）
	EndDate     string    `json:"end_date"`     // 收割时间
	IncomeTotal float64   `json:"income_total"` // 总收益
	YieldTotal  float64   `json:"yield_total"`  // 总产量
	TargetYield float64   `json:"target_yield"` // 目标产量（每亩产量）
	PlantYear   string    `json:"plant_year"`   // 种植年度
	SowTypeId   int       `json:"sow_type_id"`
}

// 农事记录 ag_farming_record
type FarmingRecord struct {
	Id                    string                `json:"id"`                       // 农事记录id,record_id
	IsActive              string                `json:"is_active"`                // 是否有效
	CreateDate            time.Time             `json:"create_date"`              // 创建日期
	ModifyDate            time.Time             `json:"modify_date"`              // 修改日期
	Creator               string                `json:"creator"`                  // 创建人
	Modifier              string                `json:"modifier"`                 // 修改人
	TenantId              string                `json:"tenant_id"`                // 公司ID
	FarmId                string                `json:"farm_id"`                  // 农场ID
	SectionId             string                `json:"section_id"`               // 地块ID
	Type                  string                `json:"type"`                     // 农事类型
	Content               string                `json:"content"`                  // 内容
	HandleDate            time.Time             `json:"handle_date"`              // 处理日期
	HandleEndDate         time.Time             `json:"handle_end_date"`          // 处理截至日期
	Remark                string                `json:"remark"`                   // 备注
	OperatorId            string                `json:"operator_id"`              // 处理人
	OperatorName          string                `json:"operator_name"`            // 处理人姓名
	Expense               float64               `json:"expense"`                  // 本次农事费用
	SectionCropsId        string                `json:"section_crops_id"`         // 种植id
	ReapBatchId           string                `json:"reap_batch_id"`            // 收割批次id
	IsGuide               string                `json:"is_guide"`                 // 是否来源农事指导
	FarmingOperationId    string                `json:"farming_operation_id"`     // 农事指导id
	HandleType            string                `json:"handle_type"`              // 是否按农事指导处理（0按指导处理，1不安指导处理）
	AgFarmingRecordDetail AgFarmingRecordDetail `json:"ag_farming_record_detail"` // 农事记录详情
}

// AgFarmingRecordDetail ag_farming_record_detail 农事记录详情
type AgFarmingRecordDetail struct {
	Id            string    `json:"id"`             // id
	IsActive      string    `json:"is_active"`      // 是否有效
	CreateDate    time.Time `json:"create_date"`    // 创建日期
	ModifyDate    time.Time `json:"modify_date"`    // 修改日期
	Creator       string    `json:"creator"`        // 创建人
	Modifier      string    `json:"modifier"`       // 修改人
	TenantId      string    `json:"tenant_id"`      // 公司id
	RecordId      string    `json:"record_id"`      // 农事记录id
	Type          string    `json:"type"`           // 农事类型
	OperationCode string    `json:"operation_code"` // 农事操作code
	OperationName string    `json:"operation_name"` // 农事操作名称
	UseType       string    `json:"use_type"`       // 农事方式或使用物品
	UseGood       string    `json:"use_good"`       // 使用物品
	StartValue    float64   `json:"start_value"`    // 指标值开始值
	EndValue      float64   `json:"end_value"`      // 指标值结束值
	Remark        string    `json:"remark"`         // 备注
	OrderNo       string    `json:"order_no"`       // 显示顺序
	TypeCode      string    `json:"type_code"`      // 字典类型
}

// 收割批次信息、收割批次关联、烘干批次信息、加工信息数据上链
type Reap struct {
	TenantId           string             `json:"tenant_id"`          // 公司id
	FarmId             string             `json:"farm_id"`            // 农场id
	ReapBatchId        string             `json:"reap_batch_id"`      // 收割批次id
	AgReapBatch        AgReapBatch        `json:"ag_reap_batch"`      // 收割批次信息
	AgSectionCropsReap AgSectionCropsReap `json:"section_crops_reap"` // 地块农作物收割关联
	DryingBatch        []AgDryingBatch    `json:"drying_batch"`       // 烘干批次
	Status             int                `json:"status"`             // 数据状态，0正常，1删除。
}

// AgReapBatch ag_reap_batch 收割批次信息、收割批次关联、烘干批次信息、加工信息数据上链
type AgReapBatch struct {
	Id                              string    `json:"id"`                                 // id
	IsActive                        string    `json:"is_active"`                          // 是否有效
	CreateDate                      time.Time `json:"create_date"`                        // 创建日期
	ModifyDate                      time.Time `json:"modify_date"`                        // 更新日期
	Creator                         string    `json:"creator"`                            // 创建人
	Modifier                        string    `json:"modifier"`                           // 更新人
	TenantId                        string    `json:"tenant_id"`                          // 公司id
	FarmId                          string    `json:"farm_id"`                            // 农场id
	Code                            string    `json:"code"`                               // 收割批次编码
	Name                            string    `json:"name"`                               // 收割批次名称
	DryingCompanyId                 string    `json:"drying_company_id"`                  // 烘干公司id
	DryingCompanyName               string    `json:"drying_company_name"`                // 烘干公司名称
	Remark                          string    `json:"remark"`                             // 备注
	Amount                          float64   `json:"amount"`                             // 数量
	HandleDate                      string    `json:"handle_date"`                        // 处理日期
	IsMultiple                      string    `json:"is_multiple"`                        // 是否来源多块地
	StartDryingRemainingAmount      float64   `json:"start_drying_remaining_amount"`      // 剩余待开始烘干数量
	EndDryingRemainingAmount        float64   `json:"end_drying_remaining_amount"`        // 剩余待结束烘干数量
	EndRefrigerationRemainingAmount float64   `json:"end_refrigeration_remaining_amount"` // 剩余待结束冷藏数量
	EndProcessRemainingAmount       float64   `json:"end_process_remaining_amount"`       // 剩余待结束加工数量
}

// AgSectionCropsReap ag_section_crops_reap 地块农作物收割关联
type AgSectionCropsReap struct {
	Id             string    `json:"id"`               // ID
	IsActive       string    `json:"is_active"`        // 是否有效
	CreateDate     time.Time `json:"create_date"`      // 创建日期
	ModifyDate     time.Time `json:"modify_date"`      // 更新日期
	Creator        string    `json:"creator"`          // 创建人
	Modifier       string    `json:"modifier"`         // 更新人
	TenantId       string    `json:"tenant_id"`        // 公司ID
	SectionCropsId string    `json:"section_crops_id"` // 种植id
	ReapBatchId    string    `json:"reap_batch_id"`    // 收割批次id
}

// AgDryingBatch ag_drying_batch烘干批次
type AgDryingBatch struct {
	Id                   string         `json:"id"`                     // id
	IsActive             string         `json:"is_active"`              // 是否有效
	CreateDate           time.Time      `json:"create_date"`            // 创建日期
	ModifyDate           time.Time      `json:"modify_date"`            // 修改日期
	Creator              string         `json:"creator"`                // 创建人
	Modifier             string         `json:"modifier"`               // 修改人
	TenantId             string         `json:"tenant_id"`              // 公司Id
	FarmId               string         `json:"farm_id"`                // 农场Id
	Code                 string         `json:"code"`                   // 烘干批次编码
	Name                 string         `json:"name"`                   // 烘干批次名称
	ReapBatchId          string         `json:"reap_batch_id"`          // 收割批次id
	Remark               string         `json:"remark"`                 // 备注
	HandleDate           string         `json:"handle_date"`            // 处理时间
	DryingCurve          string         `json:"drying_curve"`           // 烘干曲线
	Amount               float64        `json:"amount"`                 // 数量
	IsEndDrying          string         `json:"is_end_drying"`          // 是否已结束烘干
	IsStartRefrigeration string         `json:"is_start_refrigeration"` // 是否已开始冷藏
	IsEndRefrigeration   string         `json:"is_end_refrigeration"`   // 是否已结束冷藏
	IsStartProcess       string         `json:"is_start_process"`       // 是否已开始加工
	IsEndProcess         string         `json:"is_end_process"`         // 是否已结束加工
	RefrigerationId      string         `json:"refrigeration_id"`       // 冷藏地点id
	RefrigerationName    string         `json:"refrigeration_name"`     // 冷藏地点名称
	ProcessCompanyId     string         `json:"process_company_id"`     // 加工公司id
	ProcessCompanyName   string         `json:"process_company_name"`   // 加工公司名称
	HandleRecord         AgHandleRecord `json:"handle_record"`          // 加工记录
}

// AgHandleRecord ag_handle_record 加工记录
type AgHandleRecord struct {
	Id            string    `json:"id"`              // id
	IsActive      string    `json:"is_active"`       // 是否有效
	CreateDate    time.Time `json:"create_date"`     // 创建日期
	ModifyDate    time.Time `json:"modify_date"`     // 修改日期
	Creator       string    `json:"creator"`         // 创建人
	Modifier      string    `json:"modifier"`        // 修改人
	TenantId      string    `json:"tenant_id"`       // 公司ID
	FarmId        string    `json:"farm_id"`         // 农场ID
	Type          string    `json:"type"`            // 农事类型
	ReapBatchId   string    `json:"reap_batch_id"`   // 收割批次id
	DryingBatchId string    `json:"drying_batch_id"` // 烘干批次id
	HandleDate    string    `json:"handle_date"`     // 处理时间
	EndProductNum int       `json:"end_product_num"` // 成品数量
	WeightPerBag  int       `json:"weight_per_bag"`  // 每袋重量
	Price         float64   `json:"price"`           // 成交价格
	Remark        string    `json:"remark"`          // 备注
}

/*type KafkaInfo struct {
	TenantId string `json:"tenant_id"` // 公司id
	Farm []Farm `json:"farm"` // 农场
	Crops []AgCrops `json:"crops"` // 农作物基本信息
}*/

// AgDict ag_dict
type AgDict struct {
	Id          int64     `json:"id"`            // 主键，递增
	IsActive    string    `json:"is_active"`     // 是否有效
	CreateDate  time.Time `json:"create_date"`   // 创建时间
	ModifyDate  time.Time `json:"modify_date"`   // 更新时间
	Creator     string    `json:"creator"`       // 创建人
	Modifier    string    `json:"modifier"`      // 更新人
	Name        string    `json:"name"`          // 名称
	Value       int       `json:"value"`         // 值
	TypeCode    string    `json:"type_code"`     // 字典分类编码，里面编码，详细解释见AgDictType结构体记录。
	Description string    `json:"description"`   // 描述
	Code        string    `json:"code"`          // 编码（唯一）
	ParentCode  string    `json:"parent_code"`   // 上级字典编码
	IsFrontShow string    `json:"is_front_show"` // 是否可前端展示
	SortingNum  int       `json:"sorting_num"`   // 排序
}

// AgDictType ag_dict_type
type AgDictType struct {
	Id          string    `json:"id"`          // 主键
	IsActive    string    `json:"is_active"`   // 是否有效
	CreateDate  time.Time `json:"create_date"` // 创建时间
	ModifyDate  time.Time `json:"modify_date"` // 更新时间
	Creator     string    `json:"creator"`     // 创建人
	Modifier    string    `json:"modifier"`    // 更新人
	Name        string    `json:"name"`        // 名称
	Code        string    `json:"code"`        // 编码
	Description string    `json:"description"` // 描述
}

// ag_soil_element_dic
type AgSoilElementDic struct {
	Id          int    `json:"id"`
	ElementType int    `json:"element_type"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	CreateDate  string `json:"create_date"`
}

// ag_pesticide_fertilizer 化肥、农药字典表
type AgPesticideFertilizer struct {
	Id          string    `json:"id"`          // 主键
	IsActive    string    `json:"is_active"`   // 是否有效
	CreateDate  time.Time `json:"create_date"` // 创建时间
	ModifyDate  time.Time `json:"modify_date"` // 更新时间
	Creator     string    `json:"creator"`     // 创建人
	Modifier    string    `json:"modifier"`    // 更新人
	TenantId    string    `json:"tenant_id"`   // 公司ID
	UserId      string    `json:"user_id"`     // 用户id
	Code        string    `json:"code"`        // 编码
	Name        string    `json:"name"`        // 名称
	Type        string    `json:"type"`        // 分类（0农药，1肥料）
	TypeCode    string    `json:"type_code"`   // 类型编码
	Value       string    `json:"value"`       // 值
	Description string    `json:"description"` // 描述
	Apply       string    `json:"apply"`       // 适用作物
	Dan         string    `json:"dan"`         // 氮
	Lin         string    `json:"lin"`         // 磷
	Jia         string    `json:"jia"`         // 钾
	Unit        string    `json:"unit"`        // 计量单位
	IsSystem    string    `json:"is_system"`   // 是否系统
}

//------------------------------------------------------------------------------------------------
// 农场基本信息、用户信息、地块编号、农作物品类、收割批次数据上链，土壤检测报告、水质检测报告
type Farm struct {
	TenantId string `json:"tenant_id"` // 公司ID
	FarmId   string `json:"farm_id"`   // 农场id
	//AgFarm AgFarm `json:"farm_info"` // 农场信息
	//AgFarmUser AgFarmUser `json:"farm_user_info"` // 农场用户信息
	SectionIdList []string `json:"section_id_list"` // 地块列表
	//Section []Section `json:"section"`
	//ReapBatchIdList []string `json:"reap_batch_id_list"` // 收割批次列表
	//ReapBatch []Reap `json:"reap_batch"` // 收割批次
	//FarmTime int64 `json:"farm_time"` // 稳定性或压力测试时，用来测试落账时间
	SoilTestReportList  []TestReport `json:"soil_test_report_list"`  // 检测报告
	WaterTestReportList []TestReport `json:"water_test_report_list"` // 检测报告
	Status              int          `json:"status"`                 // 数据状态，0正常，1删除。
}

type FarmBatch struct {
	FarmList []Farm `json:"farm_list"` // 手动批量农场信息
}

// 收割Id--种植ID--地块ID
//       |-种植ID--地块ID
// 地块上链信息
// key: 地块ID
// value：地块ID、地块名称、地块位置、地块图片URL
type Land struct {
	SectionId string `json:"section_id"` // 地块ID
	FarmId    string `json:"farm_id"`    // 农场id
	SectionInfo
	Status int `json:"status"` // 0正常，1删除
}

type LandBatch struct {
	LandList []Land `json:"land_list"` // 手动批量地块信息
}

// 种植上链信息
// key: 种植ID
// value: 种植ID、地块ID、农事记录、农事记录详情
// 农事记录按农事八周期分类（大田准备、移栽、返青期、分蘖期、孕穗期、抽穗扬花、灌浆期、成熟期收获）
type Plant struct {
	SectionCropsId         string   `json:"section_crops_id"`           // 种植ID
	SectionId              string   `json:"section_id"`                 // 地块ID
	FarmRecordIdList       []string `json:"farm_record_id_list"`        // 农事记录ID列表
	FarmRecordIdDeleteList []string `json:"farm_record_id_delete_list"` // 农事记录ID删除列表
	PlantFarm
	Status int `json:"status"` // 0正常，1删除
}

type PlantBatch struct {
	PlantList []Plant `json:"plant_list"` // 手动种植信息
}

// 农事记录上链
// key：农事记录id
// value：种植ID、农事记录ID、操作者、处理时间、操作码、使用物品、物品用量、状态
type FarmRecord struct {
	SectionCropsId string `json:"section_crops_id"` // 种植ID
	OpDetail
	Status int `json:"status"` // 0正常，1删除
}

type FarmRecordBatch struct {
	FarmRecordList []FarmRecord `json:"farm_record_list"` // 手动批量农事记录
}

// 收割上链信息
// key: 收割批次ID
// value：收割批次ID、种植ID列表（按地块面积排序）、交易ID（作为农事相关区块链证书）
type Harvest struct {
	HarvestId      string    `json:"harvest_id"`       // 收割批次ID
	ReapCreateDate time.Time `json:"reap_create_date"` // 收割时间
	PlantIdList    []string  `json:"plant_id_list"`    // 种植ID列表（按地块面积排序）
	TxId           string    `json:"tx_id"`            // 交易ID，作为农事相关区块链证书
	Status         int       `json:"status"`           // 0正常，1删除
}

type HarvestBatch struct {
	HarvestList []Harvest `json:"harvest_list"` // 手动批量收割信息
}

type FeedSectionCrop struct {
	FeedListId     string `json:"feed_list_id"`     // 进粮清单ID
	SectionCropsId string `json:"section_crops_id"` // 种植ID
	IsActive       string `json:"is_active"`        // 是否有效
	TxId           string `json:"tx_id"`            // 交易ID，作为农事相关区块链证书
	Status         int    `json:"status"`           // 0正常，1删除
}

// 烘干上链信息
// key: 烘干批次ID
// value： 烘干批次ID、收割批次ID、烘干信息
type DryInfo struct {
	DryId     string `json:"dry_id"`     // 烘干批次ID
	HarvestId string `json:"harvest_id"` // 收割批次ID
	Dry
	FeedListId   int64     `json:"feed_list_id"`  // 进粮清单ID
	HarvestYear  string    `json:"harvest_year"`  // 收割年份
	ReachingDate time.Time `json:"reaching_date"` // 收割日期（取进粮日期）
	Status       int       `json:"status"`        // 0正常，1删除
}

type DryInfoBatch struct {
	DryInfoList []DryInfo `json:"dry_info_list"` // 手动批量烘干信息
}

// 仓储上链信息
// key: 仓储批次ID
// value：仓储批次ID、烘干批次ID、收割批次ID、仓储信息
type WarehouseInfo struct {
	WarehouseId string `json:"warehouse_id"` // 仓储批次ID
	DryId       string `json:"dry_id"`       // 烘干批次ID
	Warehouse
	Status int `json:"status"` // 0正常，1删除
}

type WarehouseInfoBatch struct {
	WarehouseInfoList []WarehouseInfo `json:"warehouse_info_list"` // 手动批量仓储信息
}

// 加工过程上链数据
// key: 加工批次ID
// value: 加工批次ID、上一环节类型（烘干或仓储）、上一环节批次ID、加工信息、交易ID（作为加工、检测报告相关区块链证书）
type ProcessInfo struct {
	ProcessId string `json:"process_id"` // 加工批次ID
	PreType   string `json:"pre_type"`   // 上一环节类型（烘干或仓储）
	PreId     string `json:"pre_id"`     // 上一环节批次ID
	Process
	QualityReportImageUrl string    `json:"quality_report_image_url"` // 品质检测报告图片存放地址
	QualityInspectionDate time.Time `json:"quality_inspection_date"`  // 品质检测报告里的检测日期
	SafeReportImageUrl    string    `json:"safe_report_image_url"`    // 安全检测报告图片存放地址
	SafeInspectionDate    time.Time `json:"safe_inspection_date"`     // 安全检测报告里的检测日期
	Package
	TxId   string `json:"tx_id"`  // 交易ID（作为加工、检测报告相关区块链证书）
	Status int    `json:"status"` // 0正常，1删除
}

type ProcessInfoBatch struct {
	ProcessInfoList []ProcessInfo `json:"process_info_list"` // 手动批量加工信息
}

// 字典上链数据
// key: AgDict
// value: []AgDict
type AgDictInfo struct {
	AgDictList []AgDict `json:"ag_dict_list"` // AgDict列表
}

// 农药、化肥字典上链数据
// key: AgPesticideFertilizer
// value: []AgPesticideFertilizer
type AgPesticideFertilizerInfo struct {
	AgPesticideFertilizerList []AgPesticideFertilizer `json:"ag_pesticide_fertilizer_list"` // AgPesticideFertilizer列表
}

// 前端h5/小程序二维码溯源查询，数据一起返回给前端
// 数据包括：地块农事操作收割、烘干仓储加工包装、检测报告信息、区块链证书
// 分别查询上述四种数据，将数据拼接后一并返回给前端
type QRCodeResponse struct {

	// 产品基本信息，从数据中台获取
	// 安全与品质，等级：优质一等，从数据中台获取，点击查看质检报告
	// 种植管理，收割时间、产地、点击查看种植过程
	ReapCreateDate  time.Time `json:"reap_create_date"` // 收割时间
	Province        string    `json:"province"`         // 省
	City            string    `json:"city"`             // 市
	District        string    `json:"district"`         // 区/县
	Address         string    `json:"address"`          // 地址
	DetailedAddress string    `json:"detailed_address"` // 详细地址
	/*
		// 加工管理，烘干时间、加工时间、冷藏温度、点击查看加工过程
		HgDate time.Time `json:"hg_date"` // 烘干时间
		JgDate time.Time `json:"jg_date"` // 加工时间
		RefTemperature string `json:"ref_temperature"` // 冷藏温度
	*/
	// 质检报告，点击查看区块链证书，品质、安全、水质、土壤图片及日期
	InspectionReport InspectionReport `json:"inspection_report"`
	// 区块链证书，区块高度，交易码，上链时间、上传者身份（统一显示品牌商）
	// 种植管理详情，地块农事操作，多个地块，只显示面积最大地块的信息
	FarmingRecordDetail FarmingRecordDetail `json:"farming_record_detail"`
	// 加工管理详情，烘干、仓储、加工、包装
	ProcessRecordDetail ProcessRecordDetail `json:"process_record_detail"`
}

// 质检报告
// 品质检测报告图片、检测日期，来自srv_mach_packing加工包装信息表
// 安全检测报告图片、检测日期，来自srv_mach_packing加工包装信息表
// 水质检测报告图片、检测日期，来自智农云
// 土壤检测报告图片、检测日期，来自智农云
type InspectionReport struct {
	QualityReportImageUrl string    `json:"quality_report_image_url"` // 品质检测报告图片存放地址
	QualityInspectionDate time.Time `json:"quality_inspection_date"`  // 品质检测报告里的检测日期
	SafeReportImageUrl    string    `json:"safe_report_image_url"`    // 安全检测报告图片存放地址
	SafeInspectionDate    time.Time `json:"safe_inspection_date"`     // 安全检测报告里的检测日期
	WaterReportImageUrl   string    `json:"water_report_image_url"`   // 水质检测报告图片存放地址
	WaterInspectionDate   time.Time `json:"water_inspection_date"`    // 水质检测报告里的检测日期
	SoilReportImageUrl    string    `json:"soil_report_image_url"`    // 土壤检测报告图片存放地址
	SoilInspectionDate    time.Time `json:"soil_inspection_date"`     // 土壤检测报告里的检测日期
	BlockInfo
}

// 区块链证书
// 区块高度，交易码，上链时间、上传者身份（统一显示品牌商）
type BlockInfo struct {
	BlockNumber  uint64 `json:"block_number"`   // 区块高度
	TxId         string `json:"tx_id"`          // 交易哈希
	TxCreateTime string `json:"tx_create_time"` // 交易创建时间
	Creator      string `json:"creator"`        // 上传者身份（统一显示品牌商）
}

type FarmingRecordDetail struct {
	// 种植地块，同收割批次只显示面积最大地块信息
	// 收割批次
	SectionInfo SectionInfo `json:"section_info"`
	// 种植信息
	// 大田准备、移栽、返青期、分蘖期、孕穗期、抽穗扬花、灌浆期、成熟期
	PlantFarm    PlantFarm  `json:"plant_farm"`
	OpDetailList []OpDetail `json:"op_detail_list"` // 农事操作汇总展示
	BlockInfo
}

// 种植地块，同收割批次只显示面积最大地块信息
// 地块名称、地块位置、地块图片
type SectionInfo struct {
	Name                string       `json:"name"`                   // 地块名称
	Province            string       `json:"province"`               // 省
	City                string       `json:"city"`                   // 市
	District            string       `json:"district"`               // 区/县
	Address             string       `json:"address"`                // 地址
	DetailedAddress     string       `json:"detailed_address"`       // 详细地址
	ImgUrl              string       `json:"img_url"`                // 图片地址
	Area                float64      `json:"area"`                   // 面积
	SoilTestReportList  []TestReport `json:"soil_test_report_list"`  // 检测报告
	WaterTestReportList []TestReport `json:"water_test_report_list"` // 检测报告
}

type TestReport struct {
	ImageUrl          string    `json:"image_url"`           // 检测报告图片存放地址
	InspectionDate    time.Time `json:"inspection_date"`     // 检测报告里的检测日期
	InspectionEndDate time.Time `json:"inspection_end_date"` // 检测报告里有效截至日期
}

// 种植信息
// 阶段名称(state_code)
// 大田准备（stage_rice_16_34_prepare）、移栽（stage_rice_16_34_sowing）、
// 返青期（stage_rice_16_34_rejuvenation）、分蘖期（stage_rice_16_34_tillering）、
// 孕穗期（stage_rice_16_34_booting）、抽穗扬花（stage_rice_16_34_heading）、
// 灌浆期（stage_rice_16_34_filling）、成熟期收获（stage_rice_16_34_harvest）
// 16代表品种“美香占2号”，34代表属地省份“安徽省”
type PlantFarm struct {
	FieldPreparation   []OpDetail `json:"field_preparation"`    // 田地准备
	Transplant         []OpDetail `json:"transplant"`           // 移栽
	ReGreening         []OpDetail `json:"re_greening"`          // 返青期
	Tillering          []OpDetail `json:"tillering"`            // 分蘖期
	Booting            []OpDetail `json:"booting"`              // 孕穗期
	HeadingAndBlooming []OpDetail `json:"heading_and_blooming"` // 抽穗扬花
	GrainFillingStage  []OpDetail `json:"grain_filling_stage"`  // 灌浆期
	Maturity           []OpDetail `json:"maturity"`             // 成熟期收获
}

// 农事操作详情
// 操作人员、操作时间
// 使用物品、指标用量
// 图片地址url
// 农事操作code(operation_code)
// oper_rice_16_34_prepare_001（施肥）、oper_rice_16_34_prepare_002（灭茬翻地）、oper_rice_16_34_prepare_003（晒田）
// oper_rice_16_34_prepare_004（泡田）、oper_rice_16_34_prepare_005（整地）、oper_rice_16_34_prepare_006（喷药）
// 如果start和end都有值且不等就是范围，如果相等就是固定值。如果只有start 那就是小于等于，如果只有end就是大于等于
type OpDetail struct {
	OperateId        string         `json:"operate_id"`         // 农事记录id
	HandleDate       time.Time      `json:"handle_date"`        // 操作时间
	OperatorName     string         `json:"operator_name"`      // 操作人员
	StageCode        string         `json:"stage_code"`         // 农事阶段
	RecordDetailList []RecordDetail `json:"record_detail_list"` // 农事记录详情
	// 农事记录在dbo.sys_attachment.pid = agronomy-back.ag_farming_record.id；dbo.sys_attachment.file_class = farmingRecord
	// subpath为压缩图地址，path为原图地址，图片会有多个
	SubImageUrl []string `json:"sub_image_url"` // 压缩图片地址
	ImageUrl    []string `json:"image_url"`     // 正常图片地址
}

type RecordDetail struct {
	Id            string  `json:"id"`             // id
	Type          string  `json:"type"`           // type
	OperationCode string  `json:"operation_code"` // 农事操作码
	OperationName string  `json:"operation_name"` // 农事操作名称
	UseType       string  `json:"use_type"`       // 农事方式或使用物品
	UseGood       string  `json:"use_good"`       // 使用物品
	TypeCode      string  `json:"type_code"`      // ag_dict字典码
	OrderNo       string  `json:"order_no"`       // 显示顺序
	StartValue    float64 `json:"start_value"`    // 指标值开始值
	EndValue      float64 `json:"end_value"`      // 指标值结束值
	Unit          string  `json:"unit"`           //计量单位kg、ml等
}

// 加工管理详情，烘干、仓储、加工、包装
type ProcessRecordDetail struct {
	Dry       Dry       `json:"dry"`       // 烘干
	Warehouse Warehouse `json:"warehouse"` // 仓储
	Process   Process   `json:"process"`   // 加工
	Package   Package   `json:"package"`   // 包装
	BlockInfo
}

// 烘干
// 烘干时间、烘干公司、图片地址
type Dry struct {
	HgDate        time.Time `json:"hg_date"`         // 烘干日期
	HgFactoryName string    `json:"hg_factory_name"` // 烘干企业
	ImageUrl      string    `json:"image_url"`       // 图片地址
}

// 仓储
// 仓储开始时间、用时、冷藏温度、冷藏地点、图片地址
type Warehouse struct {
	StartDate      time.Time `json:"start_date"`      // 仓储开始时间
	DurationDate   time.Time `json:"duration_date"`   // 仓储持续时间
	RefTemperature string    `json:"ref_temperature"` // 冷藏温度
	CCFactoryName  string    `json:"cc_factory_name"` // 仓储厂名称（冷藏地点）
	ImageUrl       string    `json:"image_url"`       // 图片地址
}

// 加工
// 加工时间、加工公司、图片地址
type Process struct {
	ProcessDate           time.Time `json:"process_date"`            // 加工日期
	ProcessEnterpriseName string    `json:"process_enterprise_name"` // 加工厂名称（加工公司）
	ProcessImageUrl       string    `json:"process_image_url"`       // 图片地址
}

// 包装
// 包装时间、包装规则、图片地址
type Package struct {
	PackDate        time.Time `json:"pack_date"`         // 包装日期
	EachBagWeight   string    `json:"each_bag_weight"`   // 包装规格（出厂重量）
	PackageImageUrl string    `json:"package_image_url"` // 图片地址
}
