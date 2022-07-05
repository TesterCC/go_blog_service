package global

import "github.com/testercc/blog-service/pkg/setting"

// 进行全局变量的声明，便于在接下来的步骤将其关联起来，并且提供给应用程序内部调用。
// 全局变量的初始化，是会随着应用程序的不断演进不断改变的。
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)

