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

	AES_ENCRYPT_AES_CODE = 18
	AES_ENCRYPT_AES_MESS = "AES encryption error"

	HASH_PASSWORD_ERR_CODE = 20
	HASH_PASSWORD_ERR_MESS = "password error"

	USER_EXITS_CODE      = 22
	USER_EXITS_CODE_MESS = "ACCOUNT  exist"

	USER_NOT_EXIST_CODE = 24
	USER_NOT_EXIST_MESS = "ACCOUNT user not exist"

	CREATE_TOKEN      = 26
	CREATE_TOKEN_MESS = "create token error"

	UPLOAD_FILE_ERR_CODE = 28
	UPLOAD_FILE_ERR_MESS = "upload file error"

	CONVERT_STRING_TO_ARRAY_CODE = 30
	CONVERT_STRING_TO_ARRAY_MESS = "convert string string to array error"

	SEND_EMAIL_ERR_CODE = 32
	SEND_EMAIL_ERR_MESS = "error send email"

	TICKETS_REGISTERED_ERR_CODE = 34
	TICKETS_REGISTERED_ERR_MESS = "error tickets regisedred"

	AES_DECRYPT_AES_CODE = 18
	AES_DECRYPT_AES_MESS = "error QRcode is invalid"

	DATA_EMPTY_ERR_CODE = 20
	DATA_EMPTY_ERR_MESS = "data empty"

	OTP_ERR_VERIFY_CODE = 22
	OTP_ERR_VERIFY_MESS = "error verify otp send email"

	CACHE_ERR_CODE = 24
	CACHE_ERR_MESS = "error set cache"

	SHOW_TIME_CODE = 26
	SHOW_TIME_MESS = "error show existing"

	CLIENT_ERROR_CODE = 28
	CLIENT_ERROR_MESS = "error client"

	ROOM_EXSTIS_CODE = 30
	ROOM_EXSTIS_MESS = "room exists"

	ACCOUNT_STAFF_LOCK_CODE = 32
	ACCOUNT_STAFF_LOCK_MESS = "account locked"
)
const (
	ROLE_ADMIN     = 1
	ROLE_STAFF     = 3
	ACCOUNT_ACTIVE = 5
)
