package stringx

import (
	"bytes"
	"fmt"
	"strings"
)

// BuilderConcat 字符串拼接.
func BuilderConcat(vals []string) string {
	var builder strings.Builder
	for _, val := range vals {
		if val == "" {
			continue
		}
		builder.WriteString(val)
	}
	return builder.String()
}

func BufferConcat(vals []string) string {
	buffer := new(bytes.Buffer)
	for _, val := range vals {
		if val == "" {
			continue
		}
		buffer.WriteString(val)
	}
	return buffer.String()
}

func ByteConcat(vals []string) string {
	buffer := make([]byte, 0, len(vals))
	for _, val := range vals {
		if val == "" {
			continue
		}
		buffer = append(buffer, val...)
	}
	return string(buffer)
}

func SprintfConcat(vals []string) string {
	s := ""
	for _, val := range vals {
		if val == "" {
			continue
		}
		s = fmt.Sprintf("%s%s", s, val)
	}
	return s
}

func PlusConcat(vals []string) string {
	s := ""
	for _, val := range vals {
		s += val
	}
	return s
}
