package helpers

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func Sha1(s string) string {
	h := sha1.New()
	h.Write(Stb(s))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func Sha256(s string) string {
	h := sha256.New()
	h.Write(Stb(s))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func Md5(s string) string {
	h := md5.New()
	_, err := io.WriteString(h, s)
	Panic(err)
	return fmt.Sprintf("%x", h.Sum(nil))
}
