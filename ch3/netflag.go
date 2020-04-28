package ch3

//Flags ...
type Flags uint

const (
	//FlagUp = 00001
	FlagUp Flags = 1 << iota
	//FlagBroadcast = 00010
	FlagBroadcast
	//FlagLoopback = 00100
	FlagLoopback
	//FlagPointToPoint = 01000
	FlagPointToPoint
	//FlagMulticast = 10000
	FlagMulticast
)

//IsUp :return v & FlagUp == FlagUP
func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

//TurnDown :v = v &^ FlagUp
func TurnDown(v *Flags) {
	*v &^= FlagUp
}

//SetBroadcast :v |= FlagBroadcast
func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

// IsCast :return v&(FlagBroadcast|FlagMulticast)
func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}
