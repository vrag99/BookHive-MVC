package types

type YamlConfig struct {
	DbUser            string `yaml:"dbUser"`
	Password          string `yaml:"password"`
	Port              int    `yaml:"port"`
	Host              string `yaml:"host"`
	DbName            string `yaml:"dbName"`
	AccessTokenSecret string `yaml:"accessTokenSecret"`
}

type Claims struct {
	Id interface{}
	Username string
	IsAdmin bool
}

type UserData struct {
	Id           int
	Username     string
	Admin        int
	Hash         string
	RequestAdmin int
}

type Err struct {
	ErrMsg string
}

type Book struct {
	Id           int
	Name         string
	Quantity          int
	AvailableQuantity int
}

type UserViewData struct {
	Username string
	State    string
	Books    []Book
}

type AdminViewData struct {
	Username string
	State    string
	Books    []Book
	Error    string
}

type UserRequest struct {
	Id       int
	Username string
	BookName string
}

type UserRequestData struct {
	Username string
	State    string
	Requests []UserRequest
}

type MakeAdminRequest struct {
	Id       string
	Username string
}

type MakeAdminRequestData struct {
	Username string
	State    string
	Requests []MakeAdminRequest
}
