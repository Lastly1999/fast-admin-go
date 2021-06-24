package models

// CarSet 汽车车系类型
type CarSet struct {
	ID                uint   `json:"id" gorm:"column:id"`
	CarSetName        string `json:"carSetName" gorm:"column:car_set_name"`
	AutoHomeCarsetId  string `json:"autoHomeCarsetId" gorm:"column:AutoHomeCarsetId"`
	CarId             string `json:"carId" gorm:"column:car_id"`
	CarSetFirstletter string `json:"carSetFirstletter" gorm:"column:CarSetFirstletter"`
	CarSetSeriesstate string `json:"carSetSeriesstate" grom:"column:car_set_seriesstate"`
	CarSetSeriesorder string `json:"carSetSeriesorder" gorm:"column:car_set_seriesorder"`
	SetType           string `json:"setType" gorm:"column:set_type"`
	SetTypeName       string `json:"setTypeName" grom:"column:set_type_name"`
}

// TableName 重命名表名
func (s *CarSet) TableName() string {
	return "car_set"
}
