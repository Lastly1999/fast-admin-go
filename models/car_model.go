package models

// CarModel 汽车类型表、包含一些汽车的详细参数
type CarModel struct {
	ID               uint   `json:"id" gorm:"column:id"`
	CarModelName     string `json:"carModelName" gorm:"column:car_model_name"`
	CarSetId         string `json:"carSetId" gorm:"column:car_set_id"`
	CarModelId       string `json:"carModelId" gorm:"column:car_model_id"`
	Type             string `json:"type" gorm:"column:type"`
	TypeName         string `json:"typeName" gorm:"column:type_name"`
	CarModelState    string `json:"carModelState" gorm:"column:car_model_state"`
	CarModelMinprice string `json:"carModelMinprice" gorm:"car_model_minprice"`
	CarModelMaxprice string `json:"carModelMaxprice" gorm:"column:car_model_maxprice"`
}

// TableName 重命名表名
func (s *CarModel) TableName() string {
	return "hema_car_model"
}
