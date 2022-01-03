package structs

type Info struct {
	Server_Id       string `json:"server_id"`
	Server_Name     string `json:"server_name"`
	Expires         string `json:"expires"`
	Expires_Cytrus  string `json:"expires_cytrus"`
	Expires_Storage string `json:"expires_storage"`
	Param_Ram       int    `json:"param_ram,string"`
	Param_Disk      int    `json:"param_disk,string"`
	Lastlog_Panel   string `json:"lastlog_panel"`
}
