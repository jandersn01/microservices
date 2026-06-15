package config 

import(
	"log"
	"os"
	"string"
)

func GetEnv() string {
	return getEnviromentValue("ENV")
}

func GetSourceUrl() string {
	return getEnviromentValue("DATA_SOURCE_URL")
}

func GetAplicationPort() int {
	portStr := getEnviromentValue("APLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Failed to parse application port: %v", err)
	}
	return port
}

func getEnviromentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return os.Getenv(key)
}

