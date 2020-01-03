package main

type ErrorCode struct {
	errorno  int64
	errormsg string
}

var ErrorCodeMap = map[int64]string{}

var (
	CodeSucc = ErrorCode{errorno: 10000, errormsg: "success"}
	// base 40001 ~ 40020

	ReqParamsError      = ErrorCode{errorno: 40001, errormsg: "请求参数错误"}
	InvalidToken        = ErrorCode{errorno: 40002, errormsg: "登录失效，请重新登录"}
	ReqInternalError    = ErrorCode{errorno: 40003, errormsg: "请求超时，请重试"}
	AuthBasicEmptyErr   = ErrorCode{errorno: 40004, errormsg: "无效客户端请求"}
	AuthBasicConvertErr = ErrorCode{errorno: 40005, errormsg: "无效客户端请求"}
	AuthClientErr       = ErrorCode{errorno: 40006, errormsg: "请求客户端错误"}
	AuthClientPwdErr    = ErrorCode{errorno: 40007, errormsg: "请求客户端错误"}
	TokenEmptyErr       = ErrorCode{errorno: 40012, errormsg: "请先登录"}
	PermissionErr       = ErrorCode{errorno: 40013, errormsg: "您没有权限"}
	DataPermissionErr   = ErrorCode{errorno: 40014, errormsg: "您没有数据权限"}

	// for oauth 认证授权
	RegisterErr     = ErrorCode{errorno: 40051, errormsg: "Register error"}
	AuthorizeErr    = ErrorCode{errorno: 40052, errormsg: "Authorize error"}
	TokenErr        = ErrorCode{errorno: 40053, errormsg: "Token error"}
	TokenExpiresErr = ErrorCode{errorno: 40054, errormsg: "Token expires error"}
	// 40200 - 40299 Sku
	SkuDetailErr    = ErrorCode{errorno: 40201, errormsg: "获取商品详情失败"}
	SkuByBarcodeErr = ErrorCode{errorno: 40203, errormsg: "根据条形码获取商品列表失败"}
)

func main() {
	fmt

}
