package database

// Option -.
type Option func(*Database)

// MaxPoolSize -.
func MaxPoolSize(size int) Option {
	return func(c *Database) {
		c.maxPoolSize = size
	}
}
