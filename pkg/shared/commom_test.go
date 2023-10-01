package shared

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvironmentVariable_Exists(t *testing.T) {
	key := "EXISTING_VARIABLE"
	value := "value123"
	os.Setenv(key, value)

	defer os.Unsetenv(key)

	result, err := getEnvironmentVariable(key)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	assert.Equal(t, value, result, "getEnvironmentVariable returns correct value when key exists")

	if result != value {
		t.Errorf("Expected %s, but got %s", value, result)
	}
}

func TestGetEnvironmentVariable_NotExists(t *testing.T) {
	key := "NON_EXISTING_VARIABLE"
	_, err := getEnvironmentVariable(key)

	expectedError := fmt.Sprintf("envVar %s does not exist", key)
	if err == nil || err.Error() != expectedError {
		t.Errorf("Expected error: %s, but got: %v", expectedError, err)
	}

	assert.Equal(t, expectedError, err.Error(), "getEnvironmentVariable returns correct error")
}
