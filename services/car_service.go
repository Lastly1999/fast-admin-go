package services

import (
	"go-service/global"
	"go-service/models"
)

//	FindCarListById id查询汽车型号 支持模糊搜索 id为空时默认全部
func FindCarListById(typeId string) (carModel models.CarModel, total int64) {
	result := global.Db.Where("id LIKE ?", "%"+typeId+"%").First(&carModel)
	total = result.RowsAffected
	return carModel, total
}
