package db

type SatisDb struct {
	Name         string            `json:"name"`
	Homepage     string            `json:"homepage"`
	Repositories []SatisRepository `json:"repositories"`
	RequireAll   bool              `json:"require-all"`
}

type SatisRepository struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}
