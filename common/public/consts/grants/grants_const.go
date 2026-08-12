package grants

//授权对象类型 (user/group/maid/logged_in/anonymous/global)
type AuthorizedTargetType string

const (
	User        AuthorizedTargetType = "user"       // 用户
	Group       AuthorizedTargetType = "group"      // 组
	Maid        AuthorizedTargetType = "maid"       // 女仆节点
	LoggedIn    AuthorizedTargetType = "logged_in" // 已登录用户
	Anonymous   AuthorizedTargetType = "anonymous"  // 匿名用户
	Global      AuthorizedTargetType = "global"     // 全局
)
//被授权访问的类型(page/api/file/datatable)
type AuthorizedContentType string

const (
	Page      	AuthorizedContentType = "page"      // 页面
	Api 		AuthorizedContentType = "api"	  // 接口
	File      	AuthorizedContentType = "file"      // 文件
	Datatable 	AuthorizedContentType = "datatable" // 数据表
)