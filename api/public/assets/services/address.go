package services

import (
	"context"
	"encoding/json"
	"flick_tickets/api/public/assets/port"
	"flick_tickets/common/enums"
	"flick_tickets/core/events/caching/cache"
)

type ServiceAddress struct {
	addressRepo port.RepositoryExportAddress
	menory      cache.RepositoryCache
}

func NewServiceAddress(addressRepo port.RepositoryExportAddress, menory cache.RepositoryCache) *ServiceAddress {
	return &ServiceAddress{
		addressRepo: addressRepo,
		menory:      menory,
	}
}
func (us *ServiceAddress) GetAllCity(ctx context.Context) (*port.CitiesResp, error) {

	var citiesAddVaueCase []*port.Cities

	key := "getAllCity"
	exists, err := us.menory.KeyExists(ctx, key)
	if err != nil {
		return &port.CitiesResp{
			Result: port.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
			},
		}, nil
	}
	if !exists {
		resp, err := us.addressRepo.GetAllCity(ctx)
		if err != nil {
			return &port.CitiesResp{
				Result: port.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, err
		}

		err = us.menory.SetObjectById(ctx, key, resp)
		if err != nil {
			return &port.CitiesResp{
				Result: port.Result{
					Code:    enums.CACHE_ERR_CODE,
					Message: enums.CACHE_ERR_MESS,
				},
			}, nil
		}
		return &port.CitiesResp{
			Result: port.Result{
				Code:    enums.SUCCESS_CODE,
				Message: enums.SUCCESS_MESS,
			},
			Cities: resp,
		}, nil
	}

	dataString, err := us.menory.GetObjectById(ctx, key)
	if err != nil {
		return &port.CitiesResp{
			Result: port.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
			},
		}, nil
	}
	err = json.Unmarshal([]byte(dataString), &citiesAddVaueCase)
	if err != nil {
		return &port.CitiesResp{
			Result: port.Result{
				Code:    enums.ERROR_CONVERT_JSON_CODE,
				Message: enums.ERROR_CONVERT_JSON_MESS,
			},
		}, nil
	}
	return &port.CitiesResp{
		Result: port.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Cities: citiesAddVaueCase,
	}, nil
}
func (us *ServiceAddress) GetAllDistrictsByCityName(ctx context.Context, cityName string) (*port.DistrictsResp, error) {

	var district []*port.Districts

	exists, err := us.menory.KeyExists(ctx, cityName)
	if err != nil {
		return &port.DistrictsResp{
			Result: port.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
			},
		}, nil
	}
	if !exists {

		resp, err := us.addressRepo.GetAllDistrictsByCityName(ctx, cityName)
		if err != nil {
			return &port.DistrictsResp{
				Result: port.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, err
		}

		err = us.menory.SetObjectById(ctx, cityName, resp)
		if err != nil {
			return &port.DistrictsResp{
				Result: port.Result{
					Code:    enums.CACHE_ERR_CODE,
					Message: enums.CACHE_ERR_MESS,
				},
			}, nil
		}
		return &port.DistrictsResp{
			Result: port.Result{
				Code:    enums.SUCCESS_CODE,
				Message: enums.SUCCESS_MESS,
			},
			Districts: resp,
		}, nil
	}
	dataString, err := us.menory.GetObjectById(ctx, cityName)
	if err != nil {
		return &port.DistrictsResp{
			Result: port.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
			},
		}, nil
	}
	err = json.Unmarshal([]byte(dataString), &district)
	if err != nil {
		return &port.DistrictsResp{
			Result: port.Result{
				Code:    enums.ERROR_CONVERT_JSON_CODE,
				Message: enums.ERROR_CONVERT_JSON_MESS,
			},
		}, nil
	}
	return &port.DistrictsResp{
		Result: port.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Districts: district,
	}, nil
}
func (us *ServiceAddress) GetAllCommunesByDistrictName(ctx context.Context, districname string) (*port.CommunesResp, error) {

	var communes []*port.Communes

	exists, err := us.menory.KeyExists(ctx, districname)
	if err != nil {
		return &port.CommunesResp{
			Result: port.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
			},
		}, nil
	}
	if !exists {
		resp, err := us.addressRepo.GetAllCommunesByDistrictName(ctx, districname)
		if err != nil {
			return &port.CommunesResp{
				Result: port.Result{
					Code:    enums.DB_ERR_CODE,
					Message: enums.DB_ERR_MESS,
				},
			}, err
		}
		err = us.menory.SetObjectById(ctx, districname, resp)
		if err != nil {
			return &port.CommunesResp{
				Result: port.Result{
					Code:    enums.CACHE_ERR_CODE,
					Message: enums.CACHE_ERR_MESS,
				},
			}, nil
		}
		return &port.CommunesResp{
			Result: port.Result{
				Code:    enums.SUCCESS_CODE,
				Message: enums.SUCCESS_MESS,
			},
			Communes: resp,
		}, nil
	}
	dataString, err := us.menory.GetObjectById(ctx, districname)
	if err != nil {
		return &port.CommunesResp{
			Result: port.Result{
				Code:    enums.CACHE_ERR_CODE,
				Message: enums.CACHE_ERR_MESS,
			},
		}, nil
	}
	err = json.Unmarshal([]byte(dataString), &communes)
	if err != nil {
		return &port.CommunesResp{
			Result: port.Result{
				Code:    enums.ERROR_CONVERT_JSON_CODE,
				Message: enums.ERROR_CONVERT_JSON_MESS,
			},
		}, nil
	}
	return &port.CommunesResp{
		Result: port.Result{
			Code:    enums.SUCCESS_CODE,
			Message: enums.SUCCESS_MESS,
		},
		Communes: communes,
	}, nil
}
