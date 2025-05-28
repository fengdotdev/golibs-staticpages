package helperfuncs_test

import (
	"testing"

	"github.com/fengdotdev/golibs-staticpages/helperfuncs"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestHash(t *testing.T) {
	data := []byte("example.com")

	SHA1 := "0caaf24ab1a0c33440c06afe99df986365b0781f"
	SHA256 := "a379a6f6eeafb9a55e378c118034e2751e682fab9f2d30ab13d2125586ce1947"
	CRC32 := "b6fa4eb9"
	MD5 := "5ababd603b22780302dd8d83498e5172"

	assert.Equal(t, helperfuncs.HashSHA1(data), SHA1)
	assert.Equal(t, helperfuncs.HashSHA256(data), SHA256)
	assert.Equal(t, helperfuncs.HashCrc32(data), CRC32)
	assert.Equal(t, helperfuncs.HashMD5(data), MD5)

}
