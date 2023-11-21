package config

type (
	Credential struct {
		Product string
		Key     string
		Secret  string
		Region  string
	}

	Aws struct {
		Credentials []Credential
	}
)
