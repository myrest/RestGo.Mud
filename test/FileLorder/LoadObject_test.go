package TestLoadObject

import (
	"testing"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/ObjectsImplementation/ObjectHelper"
)

func TestObjectLoader(t *testing.T) {
	const DocumentObjectRoot = "Documents/Objects/Objects"
	err := ObjectHelper.LoadObjectFromFolder(DocumentObjectRoot)
	if err != nil {
		t.Errorf("Load Object failed.")
	}
	if len(ObjectHelper.ObjectDefination) < 1 {
		t.Errorf("Load Object Got nothing.")
	}
}
