package config

type (
	Config struct {
		Database Database
	}

	Database struct {
		Url        string
		Password   string
		DB         string
		Collection string
	}
)
