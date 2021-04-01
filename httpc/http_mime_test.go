package httpc

import (
	"fmt"
	"testing"
)

func TestHTTPTypeByExtension(t *testing.T) {
	mime := HTTPTypeByExtension("/test/test.html")
	fmt.Println(mime)
}
