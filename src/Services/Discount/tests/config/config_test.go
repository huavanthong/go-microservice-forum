package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"

	"github.com/huavanthong/microservice-golang/src/Services/Discount/internal/config"
	"github.com/stretchr/testify/assert"
)

var debugMode bool

func TestLoadConfig(t *testing.T) {

	if debugMode {
		// Do something in debug mode
	}

	// Define test cases
	testCases := []struct {
		name          string
		configContent string
		expected      *config.Config
		shouldError   bool
	}{
		// Case 1:
		{
			name: "valid config",
			configContent: `database:
            host: localhost
            port: 5432
            user: postgres
            password: mysecretpassword
            dbname: discount_service`,

			expected: &config.Config{
				Database: config.DatabaseConfig{
					Host:     "localhost",
					Port:     5432,
					User:     "postgres",
					Password: "mysecretpassword",
					DBName:   "discount_service",
				},
			},
			shouldError: false,
		},
		// Case 2:
		{
			name: "invalid config",
			configContent: `database:
            port: 5432
            user: postgres
            password: mysecretpassword
            dbname: discount_service`,
			expected:    nil,
			shouldError: true,
		},
	}

	// Iterate over test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a temporary config file
			tmpFile, err := createTempConfigFile(tc.configContent)
			assert.NoError(t, err)

			// Cleanup the temporary file after the test is done
			defer cleanupTempConfigFile(tmpFile)

			// Load the configuration from the file
			config, err := config.LoadConfig(tmpFile.Name())

			// Assert the test results
			if tc.shouldError {
				assert.Error(t, err)
				assert.Nil(t, config)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, config)
			}
		})
	}
}

// Create temp file yaml config for testing
func createTempConfigFile(content string) (*os.File, error) {

	// Convert string sang []byte
	yamlBytes := []byte(content)

	// Đọc YAML từ []byte
	var config config.Config
	err := yaml.Unmarshal(yamlBytes, &config)
	if err != nil {
		panic(err)
	}

	// Convert struct sang YAML và ghi vào file
	yamlBytes, err = yaml.Marshal(&config)
	if err != nil {
		panic(err)
	}

	// Create tempoture file
	tmpFile, err := ioutil.TempFile("./", "test_config*.yml")
	if err != nil {
		return nil, err
	}

	// Write content yanl to temp file
	err = ioutil.WriteFile(tmpFile.Name(), yamlBytes, 0644)
	if err != nil {
		os.Remove(tmpFile.Name())
		return nil, err
	}

	// DEBUG: Enable debugger info
	if true {
		// Read content of file
		data, err := ioutil.ReadFile(tmpFile.Name())
		if err != nil {
			// handle error
		}

		fmt.Println(string(data))
		// Print location of file
		fmt.Printf("Temporary file created: %s\n", tmpFile.Name())
	}

	return tmpFile, nil
}

// Clean up file after finished test
func cleanupTempConfigFile(tmpFile *os.File) {
	tmpFile.Close()
	os.Remove(tmpFile.Name())
}
