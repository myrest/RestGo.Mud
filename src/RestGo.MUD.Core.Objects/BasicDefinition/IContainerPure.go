package BasicDefinition

// 可以存放物品的容器
type IContainerPure interface {
	PutIn(IObjectBasic)                  //放物品進Container
	GetOut(string) (IObjectBasic, error) //從Container拿出來
	IsExist(string) bool                 //物品是否存在
}
