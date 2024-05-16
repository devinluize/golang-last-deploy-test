package config

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	Database setupDatabase
}

type setupDatabase struct {
	Port          int    `mapstructure:"SERVER_PORT"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPass        string `mapstructure:"DB_PASS"`
	DBName        string `mapstructure:"DB_NAME"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        int    `mapstructure:"DB_PORT"`
	SSHUser       string `mapstructure:"SSH_USER"`
	SSHPass       string `mapstructure:"SSH_PASS"`
	SSHHost       string `mapstructure:"SSH_HOST"`
	SSHPort       int    `mapstructure:"SSH_PORT"`
	MemoryBackend string `mapstructure:"MEMORY_BACKEND"`
	MemoryIndex   string `mapstructure:"MEMORY_INDEX"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     int    `mapstructure:"REDIS_PORT"`
	RedisUsername string `mapstructure:"REDIS_USERNAME"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	AccountName   string `mapstructure:"ACCOUNT_NAME"`
	Issuer        string `mapstructure:"ISSUER"`
	TempFilePath  string `mapstructure:"TEMP_FILE_PATH"`
	ClientOrigin  string `mapstructure:"CLIENT_ORIGIN"`
	SmtpEmailFrom string `mapstructure:"SMTP_EMAIL_FROM"`
	SmtpHost      string `mapstructure:"SMTP_HOST"`
	SmtpUser      string `mapstructure:"SMTP_USER"`
	SmtpPass      string `mapstructure:"SMTP_PASS"`
	SmtpPort      int    `mapstructure:"SMTP_PORT"`
	GeneralAPI    string `mapstructure:"GENERAL_API"`
}

var EnvConfigs *setupDatabase

func InitEnvConfigs(gen bool, env string) {
	EnvConfigs = SetupConfiguration(gen, env)
}

func SetupConfiguration(gen bool, env string) (config *setupDatabase) {
	// Tell viper the path/location of your env file. If it is root just add "."
	if env == "prod" {
		if gen {
			viper.AddConfigPath("../../../.production")
		} else {
			viper.AddConfigPath(".production")
		}
	} else {
		if gen {
			viper.AddConfigPath("../../../.development")
		} else {
			viper.AddConfigPath(".development")
		}
	}

	// Tell viper the name of your file
	viper.SetConfigName("app")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	return
}

var JWT_KEY = []byte("dasdasdasdas")

type JWTClaim struct {
	UserName string
	UserRole int
	jwt.RegisteredClaims
}
