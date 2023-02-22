package BasicDefinition

import "google.golang.org/grpc/codes"

type IObjectBasic interface {
	GetObjectBasic() *ObjectBasic
	HaveCapability(cab BasicCapability) bool
}

const (
	ObjectCannotMove codes.Code = iota + 1000
)
