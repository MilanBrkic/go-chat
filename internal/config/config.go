package config

import (
	"go-chat/internal/util"
)

var SERVER_PORT = util.GetEnv("SERVER_PORT", "8080")
var MONGO_URL = util.GetEnv("MONGO_URL", "mongodb://localhost:27017")
