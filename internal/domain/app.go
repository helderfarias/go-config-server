package domain

type SpringCloudConfig struct {
	Server struct {
		Port int `yml:"port"`
	} `yml:"server"`

	Security struct {
		Basic struct {
			Enabled  bool   `yml:"enabled"`
			User     string `yml:"user"`
			Password string `yml:"password"`
		} `yml:"basic"`
		APIKey struct {
			Enabled bool   `yml:"enabled"`
			Token   string `yml:"token"`
		} `yml:"apikey"`
	} `yml:"security"`

	Encrypt struct {
		Key string `yml:"key"`
	} `yml:"encrypt"`

	Logging struct {
		Level struct {
			Root string `yml:"root"`
		} `yml:"level"`
	} `yml:"logging"`

	Spring struct {
		Profiles struct {
			Active string `yml:"active"`
		} `yml:"profiles"`
		Cloud struct {
			Config struct {
				Server struct {
					Native map[string]interface{} `yml:"native"`
					Git    map[string]interface{} `yml:"git"`
				} `yml:"server"`
			} `yml:"cloud"`
		} `yml:"cloud"`
	} `yml:"spring"`
}
