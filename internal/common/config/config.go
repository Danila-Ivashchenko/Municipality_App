package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	configFiles = []string{}
)

func addConfigFile(file string) int {
	configFiles = append(configFiles, file)
	return len(configFiles)
}

type Config struct {
	postgresUser       string
	postgresPass       string
	postgresHost       string
	postgresPort       string
	postgresDB         string
	postgresSSLMode    string
	storageFileBaseURL string

	storagePath    string
	migrationsPath string

	env string

	useHttp  bool
	httpPort string
	httpHost string

	useHttps  bool
	httpsPort string
	httpsHost string

	httpsCertPath string
	httpsKeyPath  string
}

func (c *Config) GetPsqlURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", c.postgresUser, c.postgresPass, c.postgresHost, c.postgresPort, c.postgresDB, c.postgresSSLMode)
}

func (c *Config) GetHTTPPort() string {
	return c.httpPort
}

func (c *Config) GetFileStorageBaseURL() string {
	return c.storageFileBaseURL
}

func (c *Config) GetStoragePath() string {
	return c.storagePath
}

func (c *Config) GetMigrationsPath() string {
	return c.migrationsPath
}

func (c *Config) GetEnv() string {
	return c.env
}

func LoadEnv(filenames ...string) error {
	const op = "common.config.LoadEnv"
	err := godotenv.Load(filenames...)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

func (c *Config) UseHttp() bool {
	return c.useHttp
}

func (c *Config) HttpPort() string {
	return c.httpPort
}

func (c *Config) HttpHost() string {
	return c.httpHost
}

func (c *Config) UseHttps() bool {
	return c.useHttps
}

func (c *Config) HttpsPort() string {
	return c.httpsPort
}

func (c *Config) HttpsHost() string {
	return c.httpsHost
}

func (c *Config) CertPath() string {
	return c.httpsCertPath
}

func (c *Config) CertKeyPath() string {
	return c.httpsKeyPath
}

func GetConfig() *Config {
	if len(configFiles) > 0 {
		err := LoadEnv(configFiles...)
		if err != nil {
			panic(err)
		}
	}
	cfg := &Config{
		postgresUser:       "",
		postgresPass:       "",
		postgresHost:       "localhost",
		postgresPort:       "5432",
		postgresDB:         "",
		env:                "local",
		postgresSSLMode:    "disable",
		storageFileBaseURL: "http://localhost:6060",
		storagePath:        "storage",
		migrationsPath:     "migrations",

		useHttp:  true,
		httpPort: "8080",
		httpHost: "0.0.0.0",

		useHttps:  false,
		httpsPort: "8443",
		httpsHost: "0.0.0.0",

		httpsCertPath: "./cert/cert.crt",
		httpsKeyPath:  "./cert/secret.key",
	}

	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	db := os.Getenv("POSTGRES_DB")
	ssl := os.Getenv("POSTGRES_SSL_MODE")
	env := os.Getenv("ENV")

	storageFileBaseURL := os.Getenv("FILE_STORAGE_BASE_URL")
	storagePath := os.Getenv("FILE_STORAGE_PATH")
	migrationsPath := os.Getenv("MIGRATIONS_PATH")

	useHttp := os.Getenv("USE_HTTP")
	httpPort := os.Getenv("HTTP_PORT")
	httpHost := os.Getenv("HTTP_HOST")

	useHttps := os.Getenv("USE_HTTPS")
	httpsPort := os.Getenv("HTTPS_PORT")
	httpsHost := os.Getenv("HTTPS_HOST")
	httpsCertPath := os.Getenv("HTTPS_CERT_PATH")
	httpsKeyPath := os.Getenv("HTTPS_KEY_PATH")

	if env != "" {
		cfg.env = env
	}
	if user != "" {
		cfg.postgresUser = user
	}
	if pass != "" {
		cfg.postgresPass = pass
	}
	if host != "" {
		cfg.postgresHost = host
	}
	if port != "" {
		cfg.postgresPort = port
	}
	if db != "" {
		cfg.postgresDB = db
	}
	if ssl != "" {
		cfg.postgresSSLMode = ssl
	}
	if storageFileBaseURL != "" {
		cfg.storageFileBaseURL = storageFileBaseURL
	}
	if storagePath != "" {
		cfg.storagePath = storagePath
	}
	if migrationsPath != "" {
		cfg.migrationsPath = migrationsPath
	}

	if useHttp != "" {
		cfg.useHttp = useHttp == "true"
	}
	if httpPort != "" {
		cfg.httpPort = httpPort
	}
	if httpHost != "" {
		cfg.httpHost = httpHost
	}

	if useHttps != "" {
		cfg.useHttps = useHttps == "true"
	}
	if httpsPort != "" {
		cfg.httpsPort = httpsPort
	}
	if httpsHost != "" {
		cfg.httpsHost = httpsHost
	}
	if httpsCertPath != "" {
		cfg.httpsCertPath = httpsCertPath
	}
	if httpsKeyPath != "" {
		cfg.httpsKeyPath = httpsKeyPath
	}

	return cfg
}
