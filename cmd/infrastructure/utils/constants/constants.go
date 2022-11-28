package constants

const (
	// Config Constants
	GRPC_PORT           = "GRPC_PORT"
	CONFIG_PATH         = "CONFIG_PATH"
	DEFAULT_CONFIG_TYPE = "yaml"
	HTTP_PORT           = "HTTP_PORT"

	// Mongo DB Configs
	MONGO_DB_URI = "MONGO_URI"

	// Base Config Path
	BASE_CONFIG_PATH = "cmd/infrastructure/config/config.yaml"

	// GRPC Config Server
	MAX_CONNECTION_IDLE = 5
	GRPC_TIMEOUT        = 15
	MAX_CONNECTION_AGE  = 5
	GRPC_TIME           = 10
)
