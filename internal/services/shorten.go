package services

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/stierprogrammer/goshort/internal/storage"
)

func Shorten(url string) string {
	hash := md5.Sum([]byte(url));
	shortCode := hex.EncodeToString(hash[:])[:6];

	storage.Save(shortCode, url);

	return shortCode;
}