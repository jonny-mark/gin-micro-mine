/**
 * @author jiangshangfang
 * @date 2021/7/27 8:36 PM
 **/
package response

import "gin/common/constant"

var zhCNText = map[int]string{
	constant.ServerError:        "内部服务器错误",
	constant.TooManyRequests:    "请求过多",
	constant.ParamBindError:     "参数信息错误",
	constant.AuthorizationError: "签名信息错误",
	constant.UrlSignError:       "参数签名错误",
	constant.CacheSetError:      "设置缓存失败",
	constant.CacheGetError:      "获取缓存失败",
	constant.CacheDelError:      "删除缓存失败",
	constant.CacheNotExist:      "缓存不存在",
	constant.ResubmitError:      "请勿重复提交",
	constant.HashIdsEncodeError: "HashID 加密失败",
	constant.HashIdsDecodeError: "HashID 解密失败",
	constant.RBACError:          "暂无访问权限",
	constant.RedisConnectError:  "Redis 连接失败",
	constant.MySQLConnectError:  "MySQL 连接失败",
	constant.WriteConfigError:   "写入配置文件失败",
	constant.SendEmailError:     "发送邮件失败",
	constant.MySQLExecError:     "SQL 执行失败",
	constant.GoVersionError:     "Go 版本不满足要求",

	constant.AuthorizedCreateError:    "创建调用方失败",
	constant.AuthorizedListError:      "获取调用方列表失败",
	constant.AuthorizedDeleteError:    "删除调用方失败",
	constant.AuthorizedUpdateError:    "更新调用方失败",
	constant.AuthorizedDetailError:    "获取调用方详情失败",
	constant.AuthorizedCreateAPIError: "创建调用方 API 地址失败",
	constant.AuthorizedListAPIError:   "获取调用方 API 地址列表失败",
	constant.AuthorizedDeleteAPIError: "删除调用方 API 地址失败",

	constant.AdminCreateError:             "创建管理员失败",
	constant.AdminListError:               "获取管理员列表失败",
	constant.AdminDeleteError:             "删除管理员失败",
	constant.AdminUpdateError:             "更新管理员失败",
	constant.AdminResetPasswordError:      "重置密码失败",
	constant.AdminLoginError:              "登录失败",
	constant.AdminLogOutError:             "退出失败",
	constant.AdminModifyPasswordError:     "修改密码失败",
	constant.AdminModifyPersonalInfoError: "修改个人信息失败",
	constant.AdminMenuListError:           "获取管理员菜单授权列表失败",
	constant.AdminMenuCreateError:         "管理员菜单授权失败",
	constant.AdminOfflineError:            "下线管理员失败",
	constant.AdminDetailError:             "获取个人信息失败",

	constant.MenuCreateError:       "创建菜单失败",
	constant.MenuUpdateError:       "更新菜单失败",
	constant.MenuDeleteError:       "删除菜单失败",
	constant.MenuListError:         "获取菜单列表失败",
	constant.MenuDetailError:       "获取菜单详情失败",
	constant.MenuCreateActionError: "创建菜单栏功能权限失败",
	constant.MenuListActionError:   "获取菜单栏功能权限列表失败",
	constant.MenuDeleteActionError: "删除菜单栏功能权限失败",
}
