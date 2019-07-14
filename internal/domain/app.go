package domain

type SpringCloudConfig struct {
	Server struct {
		Host string `yml:"host"`
		Port int    `yml:"port"`
	} `yml:"server"`

	Security struct {
		Basic struct {
			Enabled  bool   `yml:"enabled"`
			User     string `yml:"user"`
			Password string `yml:"password"`
		} `yml:"basic"`
		APIKey struct {
			Enabled   bool   `yml:"enabled"`
			KeyLookup string `yml:"keylookup"`
			Token     string `yml:"token"`
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
					Vault  map[string]interface{} `yml:"vault"`
				} `yml:"server"`
			} `yml:"cloud"`
		} `yml:"cloud"`
		Nats struct {
			Servers string `yml:"servers"`
			Subject string `yml:"subject"`
			Auth    struct {
				Type     string `yml:"type"`
				Token    string `yml:"token"`
				User     string `yml:"user"`
				Password string `yml:"password"`
			} `yml:"auth"`
		} `yml:"profiles"`
	} `yml:"spring"`
}
