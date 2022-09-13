package config

type (
	Config struct {
		Database Database
	}

	Database struct {
		Host       string
		Port       string
		User       string
		Password   string
		DB         string
		Collection string
	}
)
