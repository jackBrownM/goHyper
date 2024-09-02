package svcLib

type Config struct {
	Name string
	Host string
	Port int
	Mode string
}

func (c *Config) IsProd() bool {
	return c.Mode == "prod"
}

func (c *Config) IsTest() bool {
	return c.Mode == "test"
}

func (c *Config) IsDev() bool {
	return c.Mode == "dev"
}
