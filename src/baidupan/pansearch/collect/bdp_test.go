package collect

import (
	"testing"
	"fmt"
)

func Test_handle(t *testing.T) {
	url := handle("http://pan.baidu.com/s/dfjie7fkd")
	fmt.Println(url)
}
