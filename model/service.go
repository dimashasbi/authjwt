package model

//Properties is used to map config file from json
type Properties struct {
	IP          string `json:"ip_database"`
	Port        string `json:"port_database"`
	User        string `json:"user_database"`
	Password    string `json:"password_database"`
	DBName      string `json:"name_database"`
	Config      string `json:"config_name"`
	LogPath     string `json:"log_path"`
	TimeOut     string `json:"timeout"`
	ExpireToken string `json:"expire_token"`
	ServiceName string `json:"service_name"`
}

//Message type is used to marshal user password request
type Message struct {
	Header       Header      `json:"header"`
	Body         interface{} `json:"Body"`
	PathVariable []string    `json:"PathVariable"`
	QueryParams  []string    `json:"QueryParams"`
}

//Header is used to mapping header request
type Header struct {
	URL        string `json:"url"`
	Auth       string `json:"auth"`
	Query      string `json:"query"`
	StatusCode string `json:"status_code"`
}

// Configuration type is used to define common configuration to get db or to get server configuration
type Configuration struct {
	IP       string      `json:"ip"`
	Port     string      `json:"port"`
	User     string      `json:"user"`
	Password string      `json:"password"`
	Name     string      `json:"name"`
	Engine   interface{} //varible to create instance of object , this variable could be anything. Example : Engine.(*sql.Conn) => onject sql connetion
}

//InterConn type is used to define a struct for interconnection
type InterConn struct {
	SrcService  string      `json:"source_service"`
	DestService string      `json:"destination_service"`
	Description interface{} `json:"desc"`
}
