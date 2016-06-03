package utils
import (
    "encoding/base64"
)
//base64 encoding utility function
func Base64Encode(src []byte) []byte {
    return []byte(base64.StdEncoding.EncodeToString(src))
}

//base64 decoding utility function
func Base64Decode(src []byte) ([]byte, error) {
    return base64.StdEncoding.DecodeString(string(src))
}
