package prototype

type Config struct {
	workDir string
	user    string
}

func NewConfig(workdir, user string) *Config {
	return &Config{
		workDir: workdir,
		user:    user,
	}
}
func (c *Config) WithWorkDir(dir string) *Config {
	c.workDir = dir
	return c
}
func (c *Config) WithUser(user string) *Config {
	c.user = user
	return c
}
