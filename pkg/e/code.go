package e

/**
返回码
*/
const (
	// 成功
	SUCCESS = 200
	// 失败
	FAIL = 0
	// 无效参数
	INVALID_PARAMS = 400

	// 内部异常
	ERROR = 500
	// 存在TAG
	ERROR_EXIST_TAG = 10001
	// 不存在TAG
	ERROR_NOT_EXIST_TAG = 10002
	// 存在ARTICLE
	ERROR_NOT_EXIST_ARTICLE = 10003

	/**
	校验TOKEN出错
	*/
	ERROR_AUTH_CHECK_TOKEN_FAIL = 20001
	/**
	校验TOKEN超时
	*/
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	/**
	  错误TOKEN
	*/
	ERROR_AUTH_TOKEN = 20004
	/**
	  错误权限
	*/
	ERROR_AUTH = 20005
)
