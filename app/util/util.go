package util

import "bytes"

func CombineString(strings ...interface{}) string {
	var buffer bytes.Buffer
	for i := 0; i < len(strings); i++ {
		buffer.WriteString(strings[i].(string))
	}
	return buffer.String()
}
