package types

type Config struct {
	FrontendUri string `yaml:"frontend_uri"`
	System      struct {
		Debug    bool   `yaml:"debug"`
		Postgres string `yaml:"postgres"`
		Redis    string `yaml:"redis"`
	} `yaml:"system"`
	Misskey struct {
		Instance string `yaml:"instance"`
		Token    struct {
			Admin  string `yaml:"admin"`
			Notify string `yaml:"notify"`
		} `yaml:"token"`
	} `yaml:"misskey"`
}
