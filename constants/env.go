package constants

type Env string

const (
	Production  Env = "production"
	Staging     Env = "staging"
	Development Env = "development"
)

func (e Env) String() string {
	return string(e)
}

func (e Env) IsProduction() bool {
	return e == Production
}

func (e Env) IsStaging() bool {
	return e == Staging
}

func (e Env) IsDevelopment() bool {
	return e == Development
}
