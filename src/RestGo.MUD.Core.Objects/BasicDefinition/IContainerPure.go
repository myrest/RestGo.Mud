package BasicDefinition

// 可以存放物品的容器
type IContainerPure interface {
	IContainerCanPut
	IContainerCanGet
}

type IContainerCanPut interface {
	PutIn(IObjectBasic) //放物品進Container
}

type IContainerCanGet interface {
	GetOut(string) (IObjectBasic, error) //從Container拿出來
	GetObjPoniter(num int, name string) (IObjectBasic, error)
}
