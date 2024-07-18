package types

type Action struct {
	Actions []struct {
		Title       string `yaml:"title"`
		Description string `yaml:"description"`

		Shell     string `yaml:"shell"`
		Arguments []struct {
			Name    string `yaml:"name"`
			Default string `yaml:"default"`
		} `yaml:"arguments,omitempty"`
	} `yaml:"actions"`
}
