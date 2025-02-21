package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"os"
	"time"
)

var Cfg = &Config{}

// Config 配置结构体
type Config struct {
	App          App          `mapstructure:"app"`
	Mysql        Mysql        `mapstructure:"mysql"`
	Redis        Redis        `mapstructure:"redis"`
	MongoDB      MongoDB      `mapstructure:"mongodb"`
	RabbitMQ     RabbitMQ     `mapstructure:"rabbitmq"`
	KafkaMQ      KafkaMQ      `mapstructure:"kafka"`
	JwtAuth      JwtAuth      `mapstructure:"jwtAuth"`
	CloudStorage CloudStorage `mapstructure:"cloudStorage"`
}

/**
这里推荐使用mapstructure作为序列化标签
yaml不支持 AppSignExpire int64  `yaml:"app_sign_expire"` 这种下划线的标签
使用mapstructure值得注意的地方是，只要标签中使用了下划线等连接符，":"后就
不能有空格。
比如： AppSignExpire int64  `yaml:"app_sign_expire"`是可以被解析的
          AppSignExpire int64  `yaml: "app_sign_expire"` 不能被解析
*/

type App struct {
	CurrentLocation string         `mapstructure:"currentLocation"`
	Loc             *time.Location `mapstructure:"loc"`
	AppSignExpire   int64          `mapstructure:"app_sign_expire"`
	RunMode         string         `mapstructure:"run_mode"`
	HttpPort        string         `mapstructure:"http_port"`
	AppLogPath      string         `mapstructure:"app_log_path"`
	EnableLogging   bool           `mapstructure:"enable_logging"`
	AppUploadDir    string         `mapstructure:"app_upload_dir"`
	IsMigrate       bool           `mapstructure:"is_migrate"`
}

type Mysql struct {
	DBName            string        `mapstructure:"dbname"`
	User              string        `mapstructure:"user"`
	Password          string        `mapstructure:"password"`
	Host              string        `mapstructure:"host"`
	Ports             string        `mapstructure:"ports"`
	MaxOpenConn       int           `mapstructure:"max_open_conn"`
	MaxIdleConn       int           `mapstructure:"max_idle_conn"`
	ConnMaxLifeSecond time.Duration `mapstructure:"conn_max_life_second"`
	TablePrefix       string        `mapstructure:"table_prefix"`
	LogLevel          string        `mapstructure:"log_level"`
}

type Redis struct {
	Host        string `mapstructure:"host"`
	DB          int    `mapstructure:"db"`
	User        string `mapstructure:"user"`
	Password    string `mapstructure:"password"`
	Ports       string `mapstructure:"ports"`
	MinIdleConn int    `mapstructure:"min_idle_conn"`
	PoolSize    int    `mapstructure:"pool_size"`
	MaxRetries  int    `mapstructure:"max_retries"`
}

type MongoDB struct {
	DBname   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Ports    string `mapstructure:"ports"`
}

type RabbitMQ struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Ports    string `mapstructure:"ports"`
}

type KafkaMQ struct {
	Brokers []string `mapstructure:"brokers"` // Kafka 服务的地址列表
}

type JwtAuth struct {
	AccessSecret string `mapstructure:"accessSecret"`
	AccessExpire int64  `mapstructure:"accessExpire"`
}

type CloudStorage struct {
	BucketName  string `mapstructure:"bucketName"`
	Credentials string `mapstructure:"credentials"`
}

func InitConfig() {
	LoadConfig()
	InitLocation()
	InitMysql()
	InitRedis()
	//InitMongoDB()
	InitLogger()
}

// LoadConfig 加载配置，失败直接panic
func LoadConfig() {
	vipers := viper.New()

	// 1. 使用命令行参数指定配置文件路径
	configFile := flag.String("config", "", "Configuration file path")
	flag.Parse()

	// 2. 如果命令行参数指定了配置文件，则使用该路径
	if *configFile != "" {
		vipers.SetConfigFile(*configFile)
	} else {
		// 2.1 如果没有指定，则使用默认的配置文件路径   D:\WWW\GolandProjects\src\mangoTV\app\config\config.yml  /Users/yaosen/go/src/mangoTV/app/config/config.yml  /home/go/src/mangoTV/app/config/config.yml
		vipers.SetConfigFile("D:\\GolandProjects\\mangoTV\\app\\config\\config.yml")
	}

	//3.配置读取
	if err := vipers.ReadInConfig(); err != nil {
		panic(err)
	}

	//4.将配置映射成结构体
	if err := vipers.Unmarshal(Cfg); err != nil {
		panic(err)
	}

	//5. 监听配置文件变动,重新解析配置
	vipers.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(e.Name)
		if err := vipers.Unmarshal(Cfg); err != nil {
			panic(err)
		}
	})
}

// InitLogger **初始化全局 Logger**
func InitLogger() {
	logPath := Cfg.App.AppLogPath
	err := os.MkdirAll(logPath, os.ModePerm)
	if err != nil {
		logrus.Fatalf("Failed to create log directory: %v", err)
	}

	logFile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, time.Now().Format("2006-01-02")), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}

	mw := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(mw)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logrus.SetLevel(logrus.DebugLevel)
}

func InitLocation() {
	// 设置时区为配置时区
	loc, err := time.LoadLocation(Cfg.App.CurrentLocation)
	if err != nil {
		panic(err)
	}

	time.Local = loc // 全局设置为 UTC 时区
	Cfg.App.Loc = loc
}
