package config

import (
	"acl-casbin/constant"
	"errors"
	"github.com/spf13/viper"
	"golang.org/x/time/rate"
	"log"
	"time"
)

type Config struct {
	Redis       RedisConfig  // Redis cache configuration
	Logger      LoggerConfig // Logger configuration
	App         AppConfig    // Application general settings
	Mongo       MongoConfig  // MongoDB configuration
	RateLimiter RateLimitConfig
	Token       TokenConfig
}
type TokenConfig struct {
	ExpiryAccessToken  time.Duration `json:"expiry_access_token"`
	ExpiryRefreshToken time.Duration `json:"expiry_refresh_token"`
	SecretKey          string        `json:"secret_key"`
}
type RateLimitConfig struct {
	Rate   rate.Limit
	Bursts int
}

// RedisConfig holds the configuration for connecting to a Redis server,
// including connection details, timeouts, and connection pooling.
type RedisConfig struct {
	Host               string        // Redis host
	Port               string        // Redis port
	Password           string        // Redis password
	Db                 string        // Redis database
	DialTimeout        time.Duration // Redis dial timeout
	ReadTimeout        time.Duration // Redis read timeout
	WriteTimeout       time.Duration // Redis write timeout
	IdleCheckFrequency time.Duration // Redis idle check frequency
	PoolSize           int           // Redis connection pool size
	PoolTimeout        time.Duration // Redis pool timeout
}

// MongoConfig holds configuration settings for MongoDB connection,
// such as the list of hosts, authentication details, and the database name.
type MongoConfig struct {
	Hosts      []string // List of MongoDB hosts
	Username   string   // MongoDB username for authentication
	Password   string   // MongoDB password for authentication
	Port       string   // MongoDB port
	Protocol   string   // MongoDB protocol (e.g., "mongodb")
	DbName     string   // MongoDB database name
	Collection string   // MongoDB collection name
	AuthSource string   // MongoDB authentication source
}

// AppConfig holds general application configuration details like
// application name, version, port, and host.
type AppConfig struct {
	Name    string // Application name
	Version string // Application version
	Port    int    // Port to run the application
	Host    string // Host of the application
}

// LoggerConfig contains the configuration for the application’s logger,
// including the file path, encoding type, log level, and the logger name.
type LoggerConfig struct {
	FilePath string // Path to the log file
	Encoding string // Log file encoding (e.g., "json")
	Level    string // Log level (e.g., "debug", "info")
	Logger   string // Logger name
}

// Get holds the global instance of the configuration.
// It is initialized only once and can be accessed throughout the application.
var Get *Config

// ExposeConfig loads and parses the configuration file based on the app mode (production, development, or test),
// and returns the global Config object. If the configuration is already loaded, it returns the existing instance.
func ExposeConfig(AppMode constant.AppMode) *Config {
	if Get == nil {
		cfgPath := getConfigPath(AppMode)
		v, err := LoadConfig(cfgPath, "yml")
		if err != nil {
			log.Fatalf("Error in load config %v", err)
		}
		cfg, err := ParseConfig(v)
		if err != nil {
			log.Fatalf("Error in parse config %v", err)
		}
		Get = cfg
	}
	return Get
}

// ParseConfig unmarshals the Viper instance into the global Config struct.
// It returns an error if unmarshalling fails.
func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to parse config: %v", err)
		return nil, err
	}
	return &cfg, nil
}

// LoadConfig reads the configuration file with the given filename and file type (e.g., yml),
// and returns a Viper instance. If the configuration file cannot be read, it returns an error.
func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config: %v", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

// getConfigPath determines the path of the configuration file based on the application's mode (production, development, test).
// It returns the appropriate path for the config file.
func getConfigPath(mode constant.AppMode) string {
	switch mode {
	case constant.App.Production:
		return constant.Path.ProductionConfig
	case constant.App.Development:
		return constant.Path.DevelopmentConfig
	case constant.App.Test:
		return constant.Path.TestConfig
	default:
		return constant.Path.DevelopmentConfig
	}
}
