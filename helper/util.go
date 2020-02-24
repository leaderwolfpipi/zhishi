package helper

// 提供公共的函数定义

import (
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
)

// 全局唯一id发生器
func GenerateIdBySnowflake(nodeId int) (int64, error) {
	// node_id is from node.yml
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		panic("generate id by snowflake error!")
	}

	// 生成64位整型id
	return node.Generate().Int64(), nil
}

// 加载全局配置文件
func LoadConfig() {
	// 初始化错误
	var err error = nil

	// 设置配置目录
	viper.AddConfigPath(".")
	viper.AddConfigPath("/Users/jonah/zhishi_proj/conf/")

	// 加载文件名
	configNameArr := []string{"db", "jwt", "node", "mongodb", "es", "redis"}
	for configItem := range configNameArr {
		viper.SetConfigName("configItem")
	}

	// 设置配置后缀
	viper.SetConfigType("yml")

	// 加载配置到内存
	err := viper.ReadInConfig()

	return err
}

// 解析配置项到数据结构
func GetConfig(config interface{}) interface{} {
	viper.Unmarshal(config)
	return config
}

// 根据键获取配置值
func GetConfigByKey(key string) interface{} {
	return viper.Get(key)
}
