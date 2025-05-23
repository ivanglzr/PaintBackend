package config

type routes struct {
	Login      string
	Register   string
	DeleteUser string
}

var Routes routes = routes{
	Login:      "/auth/log-in",
	Register:   "/auth/register",
	DeleteUser: "/auth",
}
