package BasicDefinition

import uuid "github.com/satori/go.uuid"

type ObjectBasic struct {
	ID                 string //GUID
	Name_EN            string //英文名，可以用來Get or Look，比如 Knife
	Name_CH            string //中文名，用來組合訊息用，比如小刀，可以用來顯示 xxx手持著小刀
	Level              int    //至少需要等級多少，才能使用該物品
	Description_Ground string //在地上的描述，比如：一把在地上的小刀
	Description_Look   string //這是一把不怎麼利的小刀，可能沒啥用
	Weight             int    //重量
	Pricing            int    //售價

	//能力類
	Capability []BasicCapability //物件可允許的動作類別
	ObjectType ObjectType        //物件的䫴型，反序列化時，要依此類別製造物件

	//以下為未用到
	SystemCode           string //系統物件碼
	CommonObjectType     string
	DestroyWhenZeroQuota bool
	AllowExecuteTimes    int
	Decoration           []string
}

func (o *ObjectBasic) GetObjectBasic() ObjectBasic {
	return *o
}

func (o *ObjectBasic) HaveCapability(cab BasicCapability) bool {
	for _, v := range o.Capability {
		if v == cab {
			return true
		}
	}
	return false
}

func (o *ObjectBasic) SetObjectType(objType ObjectType) {
	o.SystemCode = objType.String()
	o.ID = uuid.NewV4().String()
}
