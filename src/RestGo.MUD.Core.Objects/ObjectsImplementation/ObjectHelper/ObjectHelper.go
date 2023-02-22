package ObjectHelper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Objects/BasicDefinition"
)

var ObjectDefination = make(map[string]*BasicDefinition.IObjectBasic)

func LoadObjectFromFolder(DocumentRoot string) error {
	err := filepath.Walk(DocumentRoot, loadObjectFromFile)
	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", DocumentRoot, err)
	}
	return err
}

func loadObjectFromFile(path string, f os.FileInfo, err error) error {
	if !f.IsDir() && strings.HasSuffix(path, ".json") {
		// 路徑第一層為ObjectType
		parts := strings.Split(path, string(os.PathSeparator))
		objectType := BasicDefinition.ParseObjectType(parts[1])
		var object []map[string]interface{}

		// Open the file with utf-8 encoding
		file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create a new decoder
		decoder := json.NewDecoder(file)

		// Unmarshal the JSON data into the target type
		if err := decoder.Decode(&object); err != nil {
			fmt.Printf("File name: %s\n", path)
			fmt.Printf("Object type: %s\n", objectType.String())
			return err
		}
		for _, value := range object {
			item := GetIndividualItem(value)
			objectCode := fmt.Sprintf("%v", value["Object_Code"])
			if _, ok := ObjectDefination[objectCode]; ok {
				return fmt.Errorf("物件重覆定義[%s], %s", objectCode, path)
			} else {
				ObjectDefination[objectCode] = &item
				return nil
			}

		}
	}
	return nil
}
