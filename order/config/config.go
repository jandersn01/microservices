package config 

import(
	"log"
	"os"
	"strconv"
)

func GetEnv() string {
	return getEnviromentValue("ENV")
}

func GetDataSourceURL() string {
	return getEnviromentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := getEnviromentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Failed to parse application port: %v", err)
	}
	return port
}

func GetPaymentServiceUrl() string {
	return getEnviromentValue("PAYMENT_SERVICE_URL")
}

func getEnviromentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return os.Getenv(key)
}