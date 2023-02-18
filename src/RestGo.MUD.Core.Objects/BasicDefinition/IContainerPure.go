package BasicDefinition

// 可以存放物品的容器
type IContainerPure interface {
	PutIn(IObjectBasic)                  //放物品進Container
	GetOut(string) (IObjectBasic, error) //從Container拿出來
	ItemListForDisplay() string          //列出內容物
	GetObjPoniter(num int, name string) (IObjectBasic, error)
}
