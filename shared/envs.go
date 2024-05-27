package shared

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type ApplicationEnvs struct {
	ApplicationName    string `envconfig:"APPLICATION_NAME"`
	ApplicationVersion string `envconfig:"APPLICATION_VERSION"`
	ApplicationPort    string `envconfig:"APPLICATION_PORT"`
}

type LoggingEnvs struct {
	LogLevel     string `envconfig:"LOG_LEVEL"`
	LogFilePath  string `envconfig:"LOG_FILE_PATH"`
	LogFormatter string `envconfig:"LOG_FORMATTER"`
	LogMaxSize   int    `envconfig:"LOG_MAX_SIZE"`
	LogMaxBackup int    `envconfig:"LOG_MAX_BACKUP"`
	LogMaxAge    int    `envconfig:"LOG_MAX_AGE"`
	LogCompress  bool   `envconfig:"LOG_COMPRESS"`
}

type DatabaseNebulaEnvs struct {
	DatabaseNebulaName           string `envconfig:"DATABASE_NEBULA_NAME"`
	DatabaseNebulaUser           string `envconfig:"DATABASE_NEBULA_USER"`
	DatabaseNebulaPass           string `envconfig:"DATABASE_NEBULA_PASS"`
	DatabaseNebulaHost           string `envconfig:"DATABASE_NEBULA_HOST"`
	DatabaseNebulaPort           string `envconfig:"DATABASE_NEBULA_PORT"`
	DatabaseNebulaMaxIdleConn    int    `envconfig:"DATABASE_NEBULA_MAX_IDLE_CONNECTION"`
	DatabaseNebulaMaxOpenConn    int    `envconfig:"DATABASE_NEBULA_MAX_OPEN_CONNECTION"`
	DatabaseNebulaConMaxLifetime string `envconfig:"DATABASE_NEBULA_CON_MAX_LIFETIME"`
	DatabaseNebulaLowSqlQuery    bool   `envconfig:"DATABASE_NEBULA_LOW_SQL_QUERY"`
	DatabaseNebulaDriver         string `envconfig:"DATABASE_NEBULA_DRIVER"`
	DatabaseNebulaPathMigration  string `envconfig:"DATABASE_NEBULA_PATH_MIGRATION"`
}

type Envs struct {
	SetMode string `envconfig:"SET_MODE"`
	ApplicationEnvs
	LoggingEnvs
	DatabaseNebulaEnvs
}

func NewEnvs() (*Envs, error) {
	var environments = Envs{}
	filename := "/home/xoxo/go/src/github.com/xoxoist/ut-tutor/.env"

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		err := envconfig.Process("", &environments)
		if err != nil {
			return nil, err
		}
	}

	err = godotenv.Load(filename)
	if err != nil {
		return nil, err
	}

	err = envconfig.Process("", &environments)
	if err != nil {
		return nil, err
	}

	return &environments, nil
}
