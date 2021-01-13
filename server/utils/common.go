/**
 * @Author: li.zhang
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2020/12/11 下午4:57
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
