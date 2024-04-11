package controllers

import (
	"flick_tickets/api/public/assets/services"

	"github.com/gin-gonic/gin"
)

type ControllerAddress struct {
	serviceAdddress *services.ServiceAddress
	*baseController
}

func NewControllerAddress(serviceAdddress *services.ServiceAddress, base *baseController) *ControllerAddress {
	return &ControllerAddress{
		serviceAdddress: serviceAdddress,
		baseController:  base,
	}
}
func (ac *ControllerAddress) GetAllCity(ctx *gin.Context) {
	resp, err := ac.serviceAdddress.GetAllCity(ctx)
	ac.Response(ctx, resp, err)
}
func (ac *ControllerAddress) GetAllDistrictsByCityName(ctx *gin.Context) {
	name := ctx.Query("name")
	resp, err := ac.serviceAdddress.GetAllDistrictsByCityName(ctx, name)
	ac.Response(ctx, resp, err)
}
func (ac *ControllerAddress) GetAllCommunesByDistrictName(ctx *gin.Context) {
	name := ctx.Query("name")
	resp, err := ac.serviceAdddress.GetAllCommunesByDistrictName(ctx, name)
	ac.Response(ctx, resp, err)

}
