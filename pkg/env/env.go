package env

import (
	"log"
	"os"
	"strconv"
)

func GetString(key string) string {
	val, ok := os.LookupEnv(key)

	if !ok || val == "" {
		log.Fatalf("❌ Environment variable %s is required but was not found", key)
	}

	return val
}

func GetInt(key string) int {
	val, ok := os.LookupEnv(key)

	if !ok || val == "" {
		log.Fatalf("❌ Environment variable %s is required but was not found", key)
	}

	valAsInt, err := strconv.Atoi(val)

	if err != nil {
		log.Fatalf("❌ Error parsing Environment variable %s", key)
	}

	return valAsInt
}

func GetBool(key string, fallback bool) bool {
	val, ok := os.LookupEnv(key)

	if !ok || val == "" {
		log.Fatalf("❌ Environment variable %s is required but was not found", key)
	}

	boolVal, err := strconv.ParseBool(val)

	if err != nil {
		log.Fatalf("❌ Error parsing Environment variable %s", key)
	}

	return boolVal
}
