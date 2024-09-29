package protocol

type PackageType byte

type Packet struct {
	Type PackageType
	Len  uint32
	Body any
}
