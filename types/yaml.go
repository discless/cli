package types

type Config struct {
	Functions 	map[string]Function `yaml:"functions"`
}

type Function struct {
	File		string `yaml:"file"`
	Function	string `yaml:"function"`
	Category	string `yamle:"category"`
}