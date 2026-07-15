package dto

type DatabaseHealthOutput struct {
	Status            string `json:"status" example:"up" doc:"Database Health Status"`
	Message           string `json:"message" example:"It's healthy" doc:"Database Status Message"`
	Error             string `json:"error" example:"db down: ..." doc:"Database Error Message"`
	OpenConnections   string `json:"open_connections" example:"1" doc:"The number of established connections both in use and idle."`
	InUse             string `json:"in_use" example:"2" doc:"The number of connections currently in use."`
	Idle              string `json:"idle" example:"3" doc:"The number of idle connections."`
	WaitCount         string `json:"wait_count" example:"4" doc:"The total number of connections waited for."`
	WaitDuration      string `json:"wait_duration" example:"0h1m0.3s" doc:"The total time blocked waiting for a new connection."`
	MaxIdleClosed     string `json:"max_idle_closed" example:"10" doc:"The total number of connections closed due to SetMaxIdleConns."`
	MaxLifetimeClosed string `json:"max_lifetime_closed" example:"10" doc:"The total number of connections closed due to SetConnMaxLifetime."`
}
