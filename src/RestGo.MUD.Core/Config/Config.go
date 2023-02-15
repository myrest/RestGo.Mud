package Config

import "rest.com.tw/tinymud/src/RestGo.Util/utility"

func convertFromFile(filePath string, targetObjectType interface{}) error {
	return utility.UnmarshalJsonFile(filePath, targetObjectType)
}
