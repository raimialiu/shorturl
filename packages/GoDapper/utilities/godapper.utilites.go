package utilities

import "bytes"

func ConcatStrings(args ...string) string {
	var byteBuffer bytes.Buffer
	for v := range args {
		byteBuffer.WriteString(args[v])
	}

	return byteBuffer.String()
}
