package util

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	data := []byte(str)
	return fmt.Sprintf("%x\n", md5.Sum(data))
}
