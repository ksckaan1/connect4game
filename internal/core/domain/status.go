package domain

type Status uint

const (
	Playing Status = iota
	Draw
	XWin
	OWin
)

func (s Status) String() string {
	switch s {
	case Draw:
		return "Draw"
	case XWin:
		return "X Win"
	case OWin:
		return "O Win"
	default:
		return "Playing"
	}
}
