package config

// 属性配置文件
type defaultProfiles struct {
	Include string `json:"include"`
}

// 默认etcd配置
type defaultEtcdConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

// mysql配置
type defaultMysqlConfig struct {
	URL               string `json:"url"`
	Enable            bool   `json:"enable"`
	MaxIdleConnection int    `json:"max_idle_connection"`
	MaxOpenConnection int    `json:"max_open_connection"`
}
