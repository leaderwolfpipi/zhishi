package helper

// 提供公共的函数定义

import (
	"github.com/bwmarrin/snowflake"
	"github.com/spf13/viper"
)

// 全局唯一id发生器
func GenerateIdBySnowflake(nodeId int64) (uint64, error) {
	// node_id is from node.yml
	node, err := snowflake.NewNode(nodeId)
	if err != nil {
		panic("generate id by snowflake error!")
	}

	// 生成64位整型id
	return (uint64)(node.Generate().Int64()), nil
}

// 加载全局配置文件
func LoadConfig() error {
	// 初始化错误
	var err error = nil

	// 实例化viper
	config := viper.New()

	// 设置配置目录
	config.AddConfigPath("./conf")
	config.AddConfigPath("/Users/jonah/github_code/zhishi/conf/")

	// 加载文件名
	configNameArr := []string{"mongodb", "es", "redis", "db", "jwt", "node"}
	for _, configItem := range configNameArr {
		config.SetConfigName(configItem)
		config.SetConfigType("yml")     // 设置配置类型
		config.MergeInConfig()          // 合并配置
		err = config.ReadInConfig()     // 加载配置到内存
		configs := config.AllSettings() // 获取当前文件配置
		for k, v := range configs {     // 全部写入默认值
			viper.SetDefault(k, v)
		}
	}

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
