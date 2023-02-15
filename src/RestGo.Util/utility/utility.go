package utility

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
)

var machingIPAddress = ""

func GetContentFile(filename string) string {
	filename = fmt.Sprintf("./Content/%s", filename)
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return ""
	}
	defer file.Close()
	var rtn string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rtn += scanner.Text() + "\n"
	}
	return strings.Trim(rtn, "\n")
}

func UnmarshalJsonFile(filePath string, targetObjectType interface{}) error {
	// Open the file with utf-8 encoding
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new decoder
	decoder := json.NewDecoder(file)
	//Todo:要查一下，當初為什麼會加上限制
	//decoder.DisallowUnknownFields()

	// Unmarshal the JSON data into the target type
	if err := decoder.Decode(targetObjectType); err != nil {
		fmt.Printf("File name: %s\n", filePath)
		fmt.Printf("Object type: %t\n", targetObjectType)
		return err
	}
	return nil
}

func ConvertMapToObject(data map[string]interface{}, obj interface{}) error {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonStr, obj)
	if err != nil {
		return err
	}

	return nil
}

var MachingIPAddress = func() string {
	if machingIPAddress == "" {
		conn, err := net.Dial("udp", "8.8.8.8:80")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		machingIPAddress = conn.LocalAddr().(*net.UDPAddr).IP.String()
	}

	return machingIPAddress
}()

func GetPercentage(total int, value int) float32 {
	if (total == 0) || (value == 0) {
		return 0
	}
	result := float32(value) / float32(total) * 100
	return float32(math.Round(float64(result)*10) / 10)
}

func filterMap(m map[string]interface{}) {
	for k, v := range m {
		switch v := v.(type) {
		case nil:
			delete(m, k)
		case float64:
			if v == 0 {
				delete(m, k)
			}
		case []interface{}:
			newSlice := make([]interface{}, 0)
			for i := range v {
				if v[i] != nil {
					if innerMap, ok := v[i].(map[string]interface{}); ok {
						filterMap(innerMap)
					}
					newSlice = append(newSlice, v[i])
				}
			}
			if len(newSlice) == 0 {
				delete(m, k)
			} else {
				m[k] = newSlice
			}
		case string:
			if v == "" {
				delete(m, k)
			}
		case map[string]interface{}:
			filterMap(v)
			if len(v) == 0 {
				delete(m, k)
			}
		}
	}
}

func ConvertStructToMap(i interface{}) (map[string]interface{}, error) {
	byt, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(byt, &result); err != nil {
		return nil, err
	}

	filterMap(result)

	return result, nil
}

func ParserGetPutCommand(msg string) (objectNum int, objectName string, containNum int, containerName string) {
	parts := strings.Fields(msg)
	if len(parts) < 2 {
		return
	}

	objectNum, objectName = ParserObjectNumber(parts[0])
	if len(parts) < 3 || (parts[1] != "into" && parts[1] != "in") {
		containNum, containerName = ParserObjectNumber(parts[1])
		return
	}

	containNum, containerName = ParserObjectNumber(parts[2])
	return
}

func ParserObjectNumber(str string) (num int, name string) {
	obj := strings.Split(str, ".")
	if len(obj) == 2 {
		var err error
		num, err = strconv.Atoi(obj[0])
		if err != nil {
			return
		}
		name = obj[1]
	} else {
		num = 1
		name = str
	}
	return
}
