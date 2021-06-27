package services

import (
	"go-service/global"
	"go-service/models"
	"strconv"
)

type PageRequest struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"pageSize"`
	Count     int64 `json:"count"`
	TotalPage int64 `json:"totalPage"`
}

//	FindCarListById id查询汽车型号 支持模糊搜索 id为空时默认全部
func FindCarListById(typeId string, page int, pageSize int) (err error, carModel []models.CarModel, pageReqest PageRequest) {
	//	查询汽车型号列表 id为空时默认查询全部
	err = global.Db.Where("id LIKE ?", "%"+typeId+"%").Limit(pageSize).Offset((page - 1) * pageSize).Order("id").Find(&carModel).Error
	//	错误时返回
	if err != nil {
		return err, nil, PageRequest{}
	}
	//	查询总条数
	global.Db.Model(&models.CarModel{}).Count(&pageReqest.Count)
	//	形参页数量转换为int64
	int64PageSize, _ := strconv.ParseInt(strconv.Itoa(pageSize), 10, 64)
	//	计算总页数
	pageReqest.TotalPage = pageReqest.Count / int64PageSize
	//	返回页数量
	pageReqest.PageSize = pageSize
	//	计算页数量
	if pageReqest.Count%int64PageSize != 0 {
		pageReqest.Page++
	}

	return err, carModel, pageReqest
}
