package Config

import "rest.com.tw/tinymud/src/RestGo.Util/utility"

func ConvertFromFile(filePath string, targetObjectType interface{}) error {
	return utility.UnmarshalJsonFile(filePath, targetObjectType)
}
