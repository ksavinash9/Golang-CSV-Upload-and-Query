package utils
import (
    "encoding/base64"
    "io"
    "crypto/rand"
    "strconv"
)

// Find Substring util function
func SubString(s string, pos, length int) string {
    runes := []rune(s)
    l := pos + length
    if l > len(runes) {
        l = len(runes)
    }
    return string(runes[pos:l])
}
// GetGuid util for getting uuids
func GetGuid() string {
    b := make([]byte, 48)

    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        return ""
    }
    return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
// Basic String Util for interfacing strings
func ToString(args ...interface{}) string {
    result := ""
    for _, arg := range args {
        switch val := arg.(type) {
        case int:
            result += strconv.Itoa(val)
        case string:
            result += val
        }
    }
    return result
}

// Find Substring util function
func Substr(str string, start, length int) string {
    rs := []rune(str)
    rl := len(rs)
    end := 0

    if start < 0 {
        start = rl - 1 + start
    }
    end = start + length

    if start > end {
        start, end = end, start
    }

    if start < 0 {
        start = 0
    }
    if start > rl {
        start = rl
    }
    if end < 0 {
        end = 0
    }
    if end > rl {
        end = rl
    }

    return string(rs[start:end])
}
