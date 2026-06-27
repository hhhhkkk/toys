package expired_strategy

type IExpiredStrategy interface {
	Push(v string) (string, bool)
	Pop() (string, bool)
	Len() int
}
