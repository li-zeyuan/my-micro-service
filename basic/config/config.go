package config

import (
	"database/sql"
	"fmt"
	"github.com/go-log/log"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/file"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	inited    bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

// InitConfig 初始化配置
func InitConfig() {

	m.Lock()
	defer m.Unlock()

	if inited {
		log.Log(fmt.Errorf("[InitConfig] 配置已经初始化过"))
		return
	}

	// 加载yml配置
	// 先加载基础配置
	appPath, _ := filepath.Abs(filepath.Dir(filepath.Join("./", string(filepath.Separator))))

	pt := filepath.Join(appPath, "conf")
	os.Chdir(appPath)

	// 找到application.yml文件
	if err := config.Load(file.NewSource(file.WithPath(pt + "/application.yml"))); err != nil {
		panic(err)
	}

	// 找到需要引入的新配置文件
	if err := config.Get(defaultRootPath, "profiles").Scan(&profiles); err != nil {
		panic(err)
	}

	log.Infof("[InitConfig] 加载配置文件：path: %s, %+v\n", pt+"/application.yml", profiles)

	// 开始导入新文件
	if len(profiles.GetInclude()) > 0 {
		include := strings.Split(profiles.GetInclude(), ",")

		sources := make([]source.Source, len(include))
		for i := 0; i < len(include); i++ {
			filePath := pt + string(filepath.Separator) + defaultConfigFilePrefix + strings.TrimSpace(include[i]) + ".yml"
			fmt.Printf(filePath + "\n")
			sources[i] = file.NewSource(file.WithPath(filePath))
		}

		// 加载include的文件
		if err := config.Load(sources...); err != nil {
			panic(err)
		}
	}

	// 赋值
	config.Get(defaultRootPath, "etcd").Scan(&etcdConfig)
	config.Get(defaultRootPath, "mysql").Scan(&mysqlConfig)

	// 标记已经初始化
	inited = true
}
