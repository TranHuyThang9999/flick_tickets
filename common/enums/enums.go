package enums

const (
	SUCCESS_CODE   = 0
	SUCCESS_MESS   = "Success"
	LOGIN_ERR_CODE = 2
	LOGIN_ERR_MESS = "Login failed"

	DB_ERR_CODE = 4
	DB_ERR_MESS = "Database error"

	ADMIN_NOT_EXIST_CODE = 6
	ADMIN_NOT_EXIST_MESS = "ACCOUNT admin not exist"

	ERROR_CONVERT_JSON_CODE = 8
	ERROR_CONVERT_JSON_MESS = "error convert string to json"

	ERROR_LOAD_CONFIG_CODE = 10
	ERROR_LOAD_CONFIG_MESS = "error load coonfig"

	INVALID_REQUEST_CODE = 12
	INVALID_REQUEST_MESS = "invalid request"

	CONVERT_TO_NUMBER_CODE = 14
	CONVERT_TO_NUMBER_MESS = "error convert string to number"

	TRANSACTION_INVALID_CODE = 16
	TRANSACTION_INVALID_MESS = "invalid transaction"
)
