package profile

type Profile string

func (p Profile) AsString() string {
	return string(p)
}

const (
	DEFAULT Profile = "default"
	DEV     Profile = "dev"
	PROD    Profile = "prod"
)
