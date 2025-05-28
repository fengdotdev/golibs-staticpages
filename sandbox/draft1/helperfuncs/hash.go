package helperfuncs

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash/crc32"
)

// HashSHA1 returns the SHA1 checksum of the data as a string
func HashSHA1(data []byte) string {
	hash := sha1.Sum(data)
	return fmt.Sprintf("%x", hash)
}

// HashSHA256 returns the SHA256 checksum of the data as a string
func HashSHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

// HashCrc32 returns the CRC32 checksum of the data as a string
func HashCrc32(data []byte) string {
	hash := crc32.ChecksumIEEE(data)
	return fmt.Sprintf("%08x", hash)
}

// HashMD5 returns the MD5 checksum of the data as a string
func HashMD5(data []byte) string {
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}