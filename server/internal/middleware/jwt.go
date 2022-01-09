package middleware

var _ Jwter = (*Jwt)(nil)

type Jwter interface {
	Init()
}

type Jwt struct{}

func (j *Jwt) Init() {

}
