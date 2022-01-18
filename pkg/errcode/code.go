package errcode

/**
 * @author jiangshangfang
 * @date 2021/7/27 6:28 PM
 * @note 响应体统一业务码
 **/
var (
	Success = NewError(0, "成功")

	ServerError        = NewError(10001, "内部服务器错误")
	TooManyRequest     = NewError(10002, "请求过多")
	ParamBindError     = NewError(10003, "参数信息错误")
	AuthorizationError = NewError(10004, "签名信息错误")
	UrlSignError       = NewError(10005, "参数签名错误")
	CacheSetError      = NewError(10006, "设置缓存失败")
	CacheGetError      = NewError(10007, "获取缓存失败")
	CacheDelError      = NewError(10008, "删除缓存失败")
	CacheNotExist      = NewError(10009, "缓存不存在")
	ResubmitError      = NewError(10010, "请勿重复提交")
	HashIdsEncodeError = NewError(10011, "HashID 加密失败")
	HashIdsDecodeError = NewError(10012, "HashID 解密失败")
	RBACError          = NewError(10013, "暂无访问权限")
	RedisConnectError  = NewError(10014, "Redis 连接失败")
	MySQLConnectError  = NewError(10015, "MySQL 连接失败")
	WriteConfigError   = NewError(10016, "写入配置文件失败")
	SendEmailError     = NewError(10017, "发送邮件失败")
	MySQLExecError     = NewError(10018, "SQL 执行失败")
	GoVersionError     = NewError(10019, "Go 版本不满足要求")
	InvalidTokenError  = NewError(10020, "Token 无效")
	TokenTimeoutError  = NewError(10021, "Token 过期")

	AuthorizedCreateError    = NewError(10101, "创建调用方失败")
	AuthorizedListError      = NewError(10102, "获取调用方列表失败")
	AuthorizedDeleteError    = NewError(10103, "删除调用方失败")
	AuthorizedUpdateError    = NewError(10104, "更新调用方失败")
	AuthorizedDetailError    = NewError(10105, "获取调用方详情失败")
	AuthorizedCreateAPIError = NewError(10106, "创建调用方 API 地址失败")
	AuthorizedListAPIError   = NewError(10107, "获取调用方 API 地址列表失败")
	AuthorizedDeleteAPIError = NewError(10108, "删除调用方 API 地址失败")

	AdminCreateError             = NewError(10201, "创建管理员失败")
	AdminListError               = NewError(10202, "获取管理员列表失败")
	AdminDeleteError             = NewError(10203, "删除管理员失败")
	AdminUpdateError             = NewError(10204, "更新管理员失败")
	AdminResetPasswordError      = NewError(10205, "重置密码失败")
	AdminLoginError              = NewError(10206, "登录失败")
	AdminLogOutError             = NewError(10207, "退出失败")
	AdminModifyPasswordError     = NewError(10208, "修改密码失败")
	AdminModifyPersonalInfoError = NewError(10209, "修改个人信息失败")
	AdminMenuListError           = NewError(10210, "获取管理员菜单授权列表失败")
	AdminMenuCreateError         = NewError(10211, "管理员菜单授权失败")
	AdminOfflineError            = NewError(10212, "下线管理员失败")
	AdminDetailError             = NewError(10213, "获取个人信息失败")

	MenuCreateError       = NewError(10301, "创建菜单失败")
	MenuUpdateError       = NewError(10302, "更新菜单失败")
	MenuDeleteError       = NewError(10303, "删除菜单失败")
	MenuListError         = NewError(10304, "获取菜单列表失败")
	MenuDetailError       = NewError(10305, "获取菜单详情失败")
	MenuCreateActionError = NewError(10306, "创建菜单栏功能权限失败")
	MenuListActionError   = NewError(10307, "获取菜单栏功能权限列表失败")
	MenuDeleteActionError = NewError(10308, "删除菜单栏功能权限失败")
)
