package utils
import (
    "encoding/hex"
    "crypto/md5"
)
//md5 Util function - get md5sum of string
func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}
