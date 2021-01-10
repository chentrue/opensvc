package e

var Msglags = map[int]string{
	200:	"ok",
	500:		"fail",
}

func GMsg(code int) string{
	msg, ok := Msglags[code]
	if ok {
		return msg
	}
	return Msglags[500]
}