package structs

type Db struct {
	Mysql      string `json:"mysql"`
	Postgresql string `json:"psql"`
	Mongo      string `json:"mongo"`
}
