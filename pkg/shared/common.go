package shared

import (
	"errors"
	"fmt"
	"os"
)

func getEnvironmentVariable(key string) (string, error) {
	envVarValue, envVarExists := os.LookupEnv(key)

	if !envVarExists {
		return "", errors.New(fmt.Sprintf("envVar %s does not exist", key))
	}

	return envVarValue, nil
}
