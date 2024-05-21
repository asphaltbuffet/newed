package config

// Key is a type for configuration keys.
type Key string

const (
	// General configuration keys.
	LogLevelKey     Key = "logging"       // Configuration key for logging level.
	TemplateDirsKey Key = "template.dirs" // Configuration key for template directories.
)

func (k Key) String() string {
	return string(k)
}
