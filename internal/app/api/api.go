package api

// Base Api server instance description
type API struct {
	//UNEXPORTED FIELD
	config *Config
}

// Api constructor: build base API instance
func New(config *Config) *API {
	return &API{
		config: config,
	}
}

// Start http searver/configure/loggers/router/database connections etc...
func (api *API) Start() error {
	return nil
}
