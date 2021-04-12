package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

//SHA256生成哈希值
func GetFileSHA256HashCode(path string) string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("open", err)
		return ""
	}
	defer f.Close()
	hash := sha256.New()
	buf := make([]byte, 4096)
	for {
		nr, err := f.Read(buf[:])
		if (err != nil && err != io.EOF) || nr < 0 {
			fmt.Println("read", err)
			return ""
		}
		hash.Write(buf)
		if nr < 4096 {
			break
		}
	}
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}

//SHA256生成哈希值
func GetSHA256HashCode(file []byte) string {
	hash := sha256.New()
	hash.Write(file)
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}
