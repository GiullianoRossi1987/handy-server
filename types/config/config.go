package types

import "os"

type PsConfig struct {
	Host     string
	Username string
	Password string
	Db       string
}

func (c *PsConfig) FromEnv() {
	c.Host = os.Getenv("HOSTNAME")
	c.Username = os.Getenv("USERNAME")
	c.Password = os.Getenv("PASSWORD")
	c.Db = os.Getenv("DATABASE")
}
