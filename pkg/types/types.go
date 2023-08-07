package types

type YamlConfig struct {
	DbUser            string `yaml:"dbUser"`
	Password          string `yaml:"password"`
	Port              int    `yaml:"port"`
	Host              string `yaml:"host"`
	DbName            string `yaml:"dbName"`
	AccessTokenSecret string `yaml:"accessTokenSecret"`
	AdminPassword     string `yaml:"adminPassword"`
}

type UserData struct {
	Id       int
	Username string
	Admin    int
	Hash     string
}

type Err struct {
	ErrMsg string
}

type Book struct {
	Id int
	Name string
	Qty int
	AvailableQty int
}

type UserViewData struct {
	Username string
	State string
	Books []Book
}

