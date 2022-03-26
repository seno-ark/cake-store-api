package config

type M map[string]interface{}

type Response struct {
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data"`
}

const (
	CONFIG_MAX_PAGINATION_COUNT = 10
	MSG_SUCCESS                 = "Success"
	MSG_ERROR_INVALID_DATA      = "Invalid Data"
	MSG_ERROR_CAKE_NOT_FOUND    = "Cake not found"
	MSG_ERROR_DATABASE          = "Database Error"
)

var LogFormat = "[${time_rfc3339} ${status} ${method} ${path} ${latency_human}]"
