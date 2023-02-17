package BasicDefinition

type IObjectBasic interface {
	GetObjectBasic() ObjectBasic
	HaveCapability(cab BasicCapability) bool
}
