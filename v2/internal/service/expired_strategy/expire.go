package expired_strategy

type Config struct {
	Name   string
	MaxLen int
}

type IExpiredStrategy interface {
	Push(v string) (string, bool)
	Pop() (string, bool)
	Len() int
}
