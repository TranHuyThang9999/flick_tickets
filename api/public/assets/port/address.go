package port

import "context"

type Cities struct {
	Name string `json:"Tỉnh Thành Phố"`
	Code string `json:"Mã TP"`
	// District     string `json:"Quận Huyện"`
	// DistrictCode string `json:"Mã QH"`
	// Commune      string `json:"Phường Xã"`
	// CommuneCode  string `json:"Mã PX"`
	// Level        string `json:"Cấp"`
}

type Districts struct {
	District     string `json:"Quận Huyện"`
	DistrictCode string `json:"Mã QH"`
}
type Communes struct {
	Commune     string `json:"Phường Xã"`
	CommuneCode string `json:"Mã PX"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CitiesResp struct {
	Result Result    `json:"result"`
	Cities []*Cities `json:"cities"`
}

type DistrictsResp struct {
	Result    Result       `json:"result"`
	Districts []*Districts `json:"districts"`
}
type CommunesResp struct {
	Result   Result      `json:"result"`
	Communes []*Communes `json:"communes"`
}

type RepositoryExportAddress interface {
	GetAllCity(ctx context.Context) ([]*Cities, error)
	GetAllDistrictsByCityName(ctx context.Context, districtName string) ([]*Districts, error)
	GetAllCommunesByDistrictName(ctx context.Context, districtsName string) ([]*Communes, error)
}
