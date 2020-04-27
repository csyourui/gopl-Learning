package ch3

import (
	"bytes"
	"fmt"
)

//IntsToString is like fmt.Sprint(values) but adds commas.
func IntsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
