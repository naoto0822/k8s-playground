package pkg

type Config struct {
}

func InitConfig() (*Config, error) {
	cfg := &Config{}
	return cfg, nil
}
