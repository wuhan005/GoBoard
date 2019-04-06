package main

import (
"crypto/aes"
"crypto/cipher"
"crypto/md5"
"crypto/sha1"
"encoding/base64"
"fmt"
"regexp"
"time"

//    "encoding/base64"
//    "strings"
"bytes"

"github.com/gin-gonic/gin"
)

/*
  sha1_encode
  产生哈希散列
*/
func sha1Encode(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

/*
  md5_encode
  产生md5散列
*/
func md5Encode(s string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	ret := fmt.Sprintf("%x", cipherStr)
	return ret
}

func pregCheck(p1 string, ept bool) error {
	//入口正则
	//"^[A-Za-z0-9]+$"英文字母数字//ept参数设置是否允许空值
	m1, err00 := regexp.MatchString("^[A-Za-z0-9%._]+$", p1)
	if err00 != nil {
		return fmt.Errorf("pregErr:%s", err00)
	}
	//空串或无效则返回错误
	if !(m1) {
		return fmt.Errorf("invaild params")
	}
	if ept {
		return nil
	}
	if p1 != "" {
		return nil
	}
	return fmt.Errorf("invaild empty params")
}
func pregCheckURL(p1 string, ept bool) error {
	//入口正则
	//"^[A-Za-z0-9]+$"英文字母数字//ept参数设置是否允许空值
	m1, err00 := regexp.MatchString("(http|ftp|https):\\/\\/[\\w\\-_]+(\\.[\\w\\-_]+)+([\\w\\-\\.,@?^=%&amp;:/~\\+#]*[\\w\\-\\@?^=%&amp;/~\\+#])?", p1)
	if err00 != nil {
		return fmt.Errorf("pregErr:%s", err00)
	}
	//空串或无效则返回错误
	if !(m1) {
		return fmt.Errorf("invaild params")
	}
	if ept {
		return nil
	}
	if p1 != "" {
		return nil
	}
	return fmt.Errorf("invaild empty params")
}

//(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?
func pregCheck2(p1 string, p2 string, ept bool) error {
	//入口正则
	//"^[A-Za-z0-9]+$"英文字母数字//ept参数设置是否允许空值
	m1, err00 := regexp.MatchString("^[A-Za-z0-9%._\\s]+$", p1)
	m2, err01 := regexp.MatchString("^[A-Za-z0-9%._\\s]+$", p2)
	if err00 != nil {
		return fmt.Errorf("pregErr:%s", err00)
	}
	if err01 != nil {
		return fmt.Errorf("pregErr:%s", err01)
	}
	//空串或无效则返回错误
	if !(m1 && m2) {
		return fmt.Errorf("invaild params")
	}
	if ept {
		return nil
	}
	if p1 != "" && p2 != "" {
		return nil
	}
	return fmt.Errorf("invaild empty params")
}

func pregCheck3(p1 string, p2 string, p3 string, ept bool) error {
	//入口正则
	//"^[A-Za-z0-9]+$"英文字母数字
	m1, err00 := regexp.MatchString("^[A-Za-z0-9%._]+$", p1)
	m2, err01 := regexp.MatchString("^[A-Za-z0-9%._]+$", p2)
	m3, err02 := regexp.MatchString("^[A-Za-z0-9%._]+$", p3)
	if err00 != nil {
		return fmt.Errorf("pregErr:%s", err00)
	}
	if err01 != nil {
		return fmt.Errorf("pregErr:%s", err01)
	}
	if err02 != nil {
		return fmt.Errorf("pregErr:%s", err02)
	}
	//空串或无效则返回错误
	if !(m1 && m2 && m3) {
		return fmt.Errorf("invaild params")
	}
	if ept {
		return nil
	}
	if p1 != "" && p2 != "" && p3 != "" {
		return nil
	}
	return fmt.Errorf("invaild empty params")
}
func pregCheck4(p1 string, p2 string, p3 string, p4 string, ept bool) error {
	//入口正则
	//"^[A-Za-z0-9]+$"英文字母数字
	m1, err00 := regexp.MatchString("^[A-Za-z0-9%._]+$", p1)
	m2, err01 := regexp.MatchString("^[A-Za-z0-9%._]+$", p2)
	m3, err02 := regexp.MatchString("^[A-Za-z0-9%._]+$", p3)
	m4, err03 := regexp.MatchString("^[A-Za-z0-9%._]+$", p4)
	if err00 != nil {
		return fmt.Errorf("pregErr:%s", err00)
	}
	if err01 != nil {
		return fmt.Errorf("pregErr:%s", err01)
	}
	if err02 != nil {
		return fmt.Errorf("pregErr:%s", err02)
	}
	if err03 != nil {
		return fmt.Errorf("pregErr:%s", err03)
	}
	//空串或无效则返回错误
	if !(m1 && m2 && m3 && m4) {
		return fmt.Errorf("invaild params")
	}
	if ept {
		return nil
	}
	if p1 != "" && p2 != "" && p3 != "" && p4 != "" {
		return nil
	}
	return fmt.Errorf("invaild empty params")
}

//app_key   nonce   school_code timestamp
func wxSign(appKey string, nonce string, schoolCode string, timestamp string) string {
	key := ""
	str := "app_key=" + appKey + "&nonce=" + nonce + "&school_code=" + schoolCode + "&timestamp=" + timestamp + "&key=" + key
	return md5Encode(str)
}

func wxSignCheck(appKey string, nonce string, schoolCode string, timestamp string, sign string) bool {
	key := ""
	str := "app_key=" + appKey + "&nonce=" + nonce + "&school_code=" + schoolCode + "&timestamp=" + timestamp + "&key=" + key
	sign2 := md5Encode(str)
	if sign == sign2 {
		return true
	}
	return false
}

//AesEncrypt from: http://www.baike.com/wiki/AES%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//AesDecrypt a
func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5Padding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

//PKCS5Padding s
func PKCS5Padding(ciphertext []byte) []byte {
	blockSize := 128
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func makeErrJSON(httpstus int, errcode int, errdata interface{}) (int, map[string]interface{}) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["error"] = errcode
	m["msg"] = fmt.Sprint(errdata)
	return httpstus, m
}

func makeErrJSONInside(c *gin.Context, httpstus int, errcode int, errdata interface{}) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["error"] = errcode
	m["msg"] = fmt.Sprint(errdata)
	c.JSON(httpstus, m)
}

func makeErrJSON3(httpstus int, errcode int, errdata interface{}) (int, map[string]interface{}, string) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["error"] = errcode
	m["msg"] = fmt.Sprint(errdata)
	//return httpstus, m, ""
	return 302, nil, "https://wx.hduhelp.com/error.php?code=" + fmt.Sprint(errcode) + "&msg=" + fmt.Sprint(errdata)
}

func makeSuccessJSON(httpstus int, data string) (int, map[string]interface{}) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["error"] = 0
	m["msg"] = data
	return httpstus, m
}

func between(a, small, big int) bool {
	if a >= small && a <= big {
		return true
	}
	return false
}

func getDateCode() (str string) {
	str = time.Now().Format("2006") + time.Now().Format("01") + time.Now().Format("02")
	return
}

func getWeek() (int, int) {
	return time.Now().ISOWeek()
}

func base64Decode(raw []byte) []byte {
	var buf bytes.Buffer
	decoded := make([]byte, 215)
	buf.Write(raw)
	decoder := base64.NewDecoder(base64.StdEncoding, &buf)
	decoder.Read(decoded)
	return decoded
}

func base64Encode(raw []byte) []byte {
	var encoded bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &encoded)
	encoder.Write(raw)
	encoder.Close()
	return encoded.Bytes()
}