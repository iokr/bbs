package conf

type Option func(c *Config)

// WithFilePath from file path load config file.
func WithFilePath(filePath string) Option {
	return func(c *Config) {
		c.filePath = filePath
	}
}
