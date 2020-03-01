package helper

// 提供公共的结构体定义
// 统一登陆模型
type LoginParams struct {
	UserName string `json:"username" form:"username"` // 用户名或者账号
	Password string `json:"password" form:"password"` // 密码
	Code     string `json:"code" form:"code"`         // 验证码
}

// 单页结构
type PageResult struct {
	PageNum  int         // 当前页数
	PageSize int         // 单页数目
	Total    int         // 总记录数
	Rows     interface{} // 返回记录
}

// 结果集定义
type JsonResult struct {
	Code    int         `json:"code"`    // 结果码
	Message string      `json:"message"` // 提示信息
	Result  interface{} `json:"result"`  // 结果集
}

// 数据库配置结构
type MysqlConfig struct {
	Driver       string `json:"driver"`       // 数据库类型；eg: mysql
	Url          string `json:"url"`          // 连接url
	Username     string `json:"username"`     // 用户名
	Password     string `json:"password"`     // 密码
	MaxOpenConns int    `json:"maxOpenConns"` // 最大打开连接数
	MaxIdleConns int    `json:"maxIdleConns"` // 最大空闲连接数
	Singular     bool   `json:"singular"`     // 表名不可为复数形式
}

// mysql配置
type DB struct {
	Mysql MysqlConfig
}

// jwt配置结构
type JwtConfig struct {
	SignKey            string `json:"signKey"`            // token生成秘钥
	ActiveTime         string `json:"activeTime"`         // 签名生效时间单位秒
	ExpiredTime        string `json:"expiredTime"`        // token有效期单位天
	RefreshExpiredTime string `json:"refreshExpiredTime"` // 刷新token有效时间
	Issuer             string `json:"issuer"`             // 发行人
}

// 节点配置结构
type NodeConfig struct {
	NodeId     string `json:"nodeId"`     // 节点id
	NodeName   string `json:"nodeName"`   // 节点名
	NodeIp     string `json:"nodeIp"`     // 节点内网ip地址
	NodeDomain string `json:"nodeDomain"` // 节点域名地址
}

// token结构
type JwtToken struct {
	Token JwtConfig
}

// 节点集群结构
type Nodes struct {
	Server []NodeConfig
}

// 定义实体处理函数
type EntityFunc func() interface{}

// 获取mysql配置信息
func GetMysqlConfig() *DB {
	mysql := DB{}
	return GetConfig(&mysql).(*DB)
}

// 获取jwt-token配置信息
func GetTokenConfig() *JwtToken {
	token := JwtToken{}
	return GetConfig(&token).(*JwtToken)
}

// 获取节点集群配置信息
func GetNodesConfig() *Nodes {
	nodes := Nodes{}
	return GetConfig(&nodes).(*Nodes)
}
