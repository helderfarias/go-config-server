package domain

type SpringCloudConfig struct {
	Server struct {
		Port int `yml:"port"`
	} `yml:"server"`

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
				} `yml:"server"`
			} `yml:"cloud"`
		} `yml:"cloud"`
	} `yml:"spring"`
}
