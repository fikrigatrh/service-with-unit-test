package models

import "time"

type ServerConfig struct {
	Name          string `env:"NAME_SERVER"`
	Port          string `env:"PORT_SERVER,required"`
	Host          string `env:"HOST_SERVER,required"`
	ServiceHost   string `env:"SERVICE_HOST"`
	Protocol      string `env:"PROTOCOL_SERVER"`
	JWTSecret     string `env:"SECRET"`
	JSONPathFile  string `env:"JSON_PATHFILE,required"`
	DBConfig      DBConfig
	ElasticConfig ElasticConfig
}

type DBConfig struct {
	Name          string `env:"NAME_POSTGRES,required"`
	Host          string `env:"HOST_POSTGRES,required"`
	Port          string `env:"PORT_POSTGRES,required"`
	User          string `env:"USER_POSTGRES"`
	Password      string `env:"PASS_POSTGRES"`
	NameLogDb     string `env:"NAME_POSTGRES_LOG,required"`
	HostLogDb     string `env:"HOST_POSTGRES_LOG,required"`
	PortLogDb     string `env:"PORT_POSTGRES_LOG,required"`
	UserLogDb     string `env:"USER_POSTGRES_LOG"`
	PasswordLogDb string `env:"PASS_POSTGRES_LOG"`
}

type ElasticConfig struct {
	Host     string `env:"HOST_ELASTICSEARCH,required"`
	Port     string `env:"PORT_ELASTICSEARCH,required"`
	User     string `env:"USER_ELASTICSEARCH"`
	Password string `env:"PASS_ELASTICSEARCH"`
	Index    string `env:"INDEX_ELASTICSEARCH,required"`
}

type Logs struct {
	ID           uint      `json:"id" gorm:"column:id"`
	Level        string    `json:"level" gorm:"column:level"`
	Message      string    `json:"message" gorm:"column:message"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at"`
	RequestBE    string    `json:"request_be" gorm:"type:JSONB NULL DEFAULT '{}'::JSONB"`
	ResponseBE   string    `json:"response_be" gorm:"type:JSONB NULL DEFAULT '{}'::JSONB"`
	PathError    string    `json:"path_error"`
	ResponseTime string    `json:"response_time"`
	TraceHeader  string    `json:"trace_header" gorm:"type:JSONB NULL DEFAULT '{}'::JSONB"`
}

type ResponseCustomErr struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type ErrMeta struct {
	ServiceCode string
	FieldErr    string
}

const (
	ServiceCode = "10"
)

type ResponseSuccess struct {
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	Data            interface{} `json:"data"`
}

type DaerahApi struct {
	Provinsi []Detail `json:"provinsi"`
}

type Detail struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type Provinsi struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type KotaKabupaten struct {
	KotaKabupaten []KotaKabDetail `json:"kota_kabupaten"`
}

type KotaKabDetail struct {
	ID         int    `json:"id"`
	IdProvinsi string `json:"id_provinsi"`
	Nama       string `json:"nama"`
}

type KotaKab struct {
	ID         int    `json:"id"`
	IdProvinsi string `json:"id_provinsi"`
	Nama       string `json:"nama"`
}
