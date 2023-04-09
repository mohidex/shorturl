package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyString(t *testing.T) {
	emptyStr := EmptyString("")
	withStr := EmptyString("Some string")
	assert.Equal(t, emptyStr, true)
	assert.Equal(t, withStr, false)
}

func TestSha256Of(t *testing.T) {
	bytes, err := Sha256Of("https://randomurl.com")
	res := []byte{
		164, 139, 71, 195, 16, 222, 76, 141, 220, 5, 8,
		18, 93, 147, 163, 209, 117, 118, 34, 131, 101,
		236, 210, 33, 135, 193, 225, 152, 142, 88, 133, 173,
	}
	assert.Equal(t, bytes, res)
	assert.Equal(t, err, nil)
}

func TestBase58Encode(t *testing.T) {
	res := []byte{
		164, 139, 71, 195, 16, 222, 76, 141, 220, 5, 8,
		18, 93, 147, 163, 209, 117, 118, 34, 131, 101,
		236, 210, 33, 135, 193, 225, 152, 142, 88, 133, 173,
	}
	encodedStr := Base58Encode(res)
	assert.Equal(t, encodedStr, "C5K3UVL5mrppUeivi38CZR4qAtPg5SDezvAeR8LHf2cC")
}

func TestGenerateShortLinkNoError(t *testing.T) {
	str, err := GenerateShortLink("https://randomurl.com")
	assert.Equal(t, str, "C5K3UVL5")
	assert.Equal(t, err, nil)
}

func TestGenerateShortLinkError(t *testing.T) {
	str, err := GenerateShortLink("")
	assert.Equal(t, str, "")
	assert.NotEqual(t, err, nil)
}
