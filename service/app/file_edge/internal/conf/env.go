package conf

type Env struct {
	TemporaryFileDirectory string
}

var env *Env = &Env{}

func GetEnv() *Env {
	return env
}
