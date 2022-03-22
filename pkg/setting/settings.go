package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

/*
NewSetting 方法，用于初始化本项目的配置的基础属性，
设定配置文件的名称为 config，配置类型为 yaml，并且设置其配置路径为相对路径 configs/，
以此确保在项目目录下执行运行时能够成功启动。
*/
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")    // 可以不断地调用 AddConfigPath 方法, 设置多个配置路径的, 尽可能的尝试解决路径查找的问题
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
