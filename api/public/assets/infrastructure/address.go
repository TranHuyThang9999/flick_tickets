package infrastructure

import (
	"context"
	"flick_tickets/api/public/assets/port"
	"flick_tickets/configs"
	"flick_tickets/core/adapter"

	"gorm.io/gorm"
)

type CollectionAddress struct {
	collection *gorm.DB
}

func NewCollectionAddress(cf *configs.Configs, sh *adapter.PostGresql) port.RepositoryExportAddress {
	return &CollectionAddress{
		collection: sh.CreateCollection(),
	}
}

func (c *CollectionAddress) GetAllCity(ctx context.Context) ([]*port.Cities, error) {
	var cities = make([]*port.Cities, 0)
	result := c.collection.Distinct("name").Find(&cities)
	return cities, result.Error

}
func (c *CollectionAddress) GetAllDistrictsByCityName(ctx context.Context, ciTyName string) ([]*port.Districts, error) {
	var listDistricts = make([]*port.Districts, 0)
	result := c.collection.Model(&port.Cities{}).Distinct("district").Where("name=?", ciTyName).Find(&listDistricts)
	return listDistricts, result.Error
}

func (c *CollectionAddress) GetAllCommunesByDistrictName(ctx context.Context, districtsName string) ([]*port.Communes, error) {
	var listCommunes = make([]*port.Communes, 0)
	result := c.collection.Model(&port.Cities{}).Distinct("commune").Where("district=?", districtsName).Find(&listCommunes)
	return listCommunes, result.Error
}
