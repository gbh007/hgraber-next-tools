package config

type Config struct {
	MasterAPI   MasterAPI   `envconfig:"MASTER_API" yaml:"master_api"`
	Application Application `envconfig:"APPLICATION" yaml:"application"`
}

func DefaultConfig() Config {
	return Config{
		MasterAPI:   DefaultMasterAPI(),
		Application: DefaultApplication(),
	}
}

type MasterAPI struct {
	Addr  string `envconfig:"ADDR" yaml:"addr"`
	Token string `envconfig:"TOKEN" yaml:"token"`
}

func DefaultMasterAPI() MasterAPI {
	return MasterAPI{}
}

type Application struct {
	Debug bool `envconfig:"DEBUG" yaml:"debug"`
}

func DefaultApplication() Application {
	return Application{
		Debug: false,
	}
}
