package entity

type Channels int

const (
	Local = Channels(1)
	Email = Channels(2)
)

func (c Channels) Value() int {
	return int(c)
}
