package utils

import (
	"crypto/md5"
	"encoding/hex"
	"crypto/rand"
	"io"
	"encoding/base64"
	"time"
)
/*
 * md5方法
 *@auth heyibo
 *@date 2018-5-30
 */
const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)


func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

/*
 * Guid方法
 *@auth heyibo
 *@date 2018-5-30
 */
func GetGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

/*
 * GetDate方法
 *@auth heyibo
 *@date 2018-5-30
 */
func GetDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04")
}
/*
 * GetDateMHGuid方法
 *@auth heyibo
 *@date 2018-5-30
 */
func GetDateMH(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("01-02 03:04")
}

/**
 * base64加密
 */
func Base64Encode(src []byte) []byte {
	var coder = base64.NewEncoding(base64Table)
	return []byte(coder.EncodeToString(src))
}


/**
 * base64解密
 */
func Base64Decode(src []byte) ([]byte, error) {
	var coder = base64.NewEncoding(base64Table)
	return coder.DecodeString(string(src))
}



