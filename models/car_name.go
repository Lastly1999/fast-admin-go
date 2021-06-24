package models

// CarName 汽车型号模型
type CarName struct {
	ID           uint   `json:"id"`
	CarName      string `json:"carName" gorm:"column:car_name"`
	AutoHomeId   string `json:"autoHomeId" gorm:"column:auto_home_id"`
	Bfirstletter string `json:"bfirstletter" gorm:"column:bfirstletter"`
}

// TableName 重命名表名
func (s *CarName) TableName() string {
	return "car_name"
}
