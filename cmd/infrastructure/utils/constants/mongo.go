package constants

import "time"

const (
	CONNECT_TIMEOUT    = 30 * time.Second
	MAX_CONN_IDLE_TIME = 3 * time.Minute
	MIN_POOL_SIZE      = 20
	MAX_POOL_SIZE      = 300
)
