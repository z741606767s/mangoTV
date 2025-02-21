package constants

type ErrCode struct {
	Code     int
	HTTPCode int
	Desc     string
}

type errCodes struct {
	ErrNo                   ErrCode
	ErrInternalServer       ErrCode
	ErrParams               ErrCode
	ErrAuthentication       ErrCode
	ErrNotFound             ErrCode
	ErrAuthenticationHeader ErrCode
	ErrAppKey               ErrCode
	ErrSign                 ErrCode
	ErrPermission           ErrCode
	ErrInvalidJson          ErrCode
	ErrTimeout              ErrCode
	ErrAuthExpired          ErrCode
	ErrElasticsearchServer  ErrCode
	ErrElasticsearchDSL     ErrCode
	ErrMysqlServer          ErrCode
	ErrMysqlSQL             ErrCode
	ErrMongoServer          ErrCode
	ErrMongoDSL             ErrCode
	ErrRedisServer          ErrCode
	ErrKafkaServer          ErrCode

	ErrWrongPassword                           ErrCode
	ErrUserDoesNotExist                        ErrCode
	ErrWrongUserNameOrPassword                 ErrCode
	ErrTheRequestCannotBeAccessedWithoutAToken ErrCode
	ErrTokenGenerationFailed                   ErrCode
	ErrAccountIsNotAvailable                   ErrCode
	ErrUploadingFile                           ErrCode
	ErrCreatingUploadDirectory                 ErrCode
	ErrSavingFile                              ErrCode
	ErrRequestsFrequent                        ErrCode
}

var ErrCodes = errCodes{
	ErrNo: ErrCode{
		Code:     200,
		HTTPCode: 200,
		Desc:     "Success",
	},
	ErrInternalServer: ErrCode{
		Code:     10010001,
		HTTPCode: 200,
		Desc:     "Internal server error",
	},
	ErrParams: ErrCode{
		Code:     10010002,
		HTTPCode: 200,
		Desc:     "Illegal params",
	},
	ErrNotFound: ErrCode{
		Code:     10010004,
		HTTPCode: 200,
		Desc:     "Page not found",
	},
	ErrAuthenticationHeader: ErrCode{
		Code:     10010005,
		HTTPCode: 200,
		Desc:     "Authentication header Illegal",
	},
	ErrAuthentication: ErrCode{
		Code:     10010003,
		HTTPCode: 200,
		Desc:     "Authentication failed",
	},
	ErrAuthExpired: ErrCode{
		Code:     10010011,
		HTTPCode: 200,
		Desc:     "Authentication expired",
	},
	ErrAppKey: ErrCode{
		Code:     10010006,
		HTTPCode: 200,
		Desc:     "Invalid app key",
	},
	ErrSign: ErrCode{
		Code:     10010007,
		HTTPCode: 200,
		Desc:     "Invalid secret key",
	},
	ErrPermission: ErrCode{
		Code:     10010008,
		HTTPCode: 200,
		Desc:     "Permission denied",
	},
	ErrInvalidJson: ErrCode{
		Code:     10010009,
		HTTPCode: 200,
		Desc:     "Invalid Json",
	},
	ErrTimeout: ErrCode{
		Code:     10010010,
		HTTPCode: 200,
		Desc:     "Server response timeout",
	},
	ErrElasticsearchServer: ErrCode{
		Code:     10010101,
		HTTPCode: 200,
		Desc:     "Elasticsearch server error",
	},
	ErrElasticsearchDSL: ErrCode{
		Code:     10010102,
		HTTPCode: 200,
		Desc:     "Elasticsearch  DSL error",
	},
	ErrMysqlServer: ErrCode{
		Code:     10010201,
		HTTPCode: 200,
		Desc:     "Mysql server error",
	},
	ErrMysqlSQL: ErrCode{
		Code:     10010202,
		HTTPCode: 200,
		Desc:     "Illegal SQL",
	},
	ErrMongoServer: ErrCode{
		Code:     10010301,
		HTTPCode: 200,
		Desc:     "MongoDB server error",
	},
	ErrMongoDSL: ErrCode{
		Code:     10010302,
		HTTPCode: 200,
		Desc:     "MongoDB DSL error",
	},
	ErrRedisServer: ErrCode{
		Code:     10010401,
		HTTPCode: 200,
		Desc:     "Redis server error",
	},
	ErrKafkaServer: ErrCode{
		Code:     10010501,
		HTTPCode: 200,
		Desc:     "Kafka server error",
	},

	ErrWrongPassword: ErrCode{
		Code:     10010502,
		HTTPCode: 200,
		Desc:     "Wrong password",
	},
	ErrUserDoesNotExist: ErrCode{
		Code:     10010503,
		HTTPCode: 200,
		Desc:     "User does not exist",
	},
	ErrWrongUserNameOrPassword: ErrCode{
		Code:     10010504,
		HTTPCode: 200,
		Desc:     "Wrong user name or password",
	},
	ErrTheRequestCannotBeAccessedWithoutAToken: ErrCode{
		Code:     10010505,
		HTTPCode: 200,
		Desc:     "The request cannot be accessed without a token",
	},
	ErrTokenGenerationFailed: ErrCode{
		Code:     10010506,
		HTTPCode: 200,
		Desc:     "Token generation failed",
	},
	ErrAccountIsNotAvailable: ErrCode{
		Code:     10010507,
		HTTPCode: 200,
		Desc:     "The current account is not available",
	},
	ErrUploadingFile: ErrCode{
		Code:     10010508,
		HTTPCode: 200,
		Desc:     "Error uploading file",
	},
	ErrCreatingUploadDirectory: ErrCode{
		Code:     10010509,
		HTTPCode: 200,
		Desc:     "Error creating upload directory",
	},
	ErrSavingFile: ErrCode{
		Code:     10010510,
		HTTPCode: 200,
		Desc:     "Error saving file",
	},
	ErrRequestsFrequent: ErrCode{
		Code:     10010511,
		HTTPCode: 200,
		Desc:     "Requests frequently failed",
	},
}
