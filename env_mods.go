package env_mods

import (
	"os"

	"github.com/gin-gonic/gin"
)

const (
	Production  = "production"
	Development = "development"
	Stage       = "stage"
	Test        = "test"
)

var envMap = map[string]string{
	Production:  gin.ReleaseMode,
	Development: gin.DebugMode,
	Stage:       gin.DebugMode,
	Test:        gin.TestMode,
}

func GetMode(environment string) string {
	return envMap[environment]
}

// IsDebugMode checks if environment mode is debug
func IsDebugMode() bool {
	env := os.Getenv("ENV")
	if GetMode(env) == gin.DebugMode {
		return true
	}
	return false
}
