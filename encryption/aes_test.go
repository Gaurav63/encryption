package encryption_test

import (
	"testing"

	"github.com/Gaurav63/security/encryption"
	"github.com/stretchr/testify/assert"
)

const plainText = "plain text"
const cipherKey = "secret_key"

func TestAES(t *testing.T) {
	cipherKeyBytes := encryption.Byte32([]byte(cipherKey))
	actualCipherBytes, err := encryption.AESEncrypt([]byte(plainText), cipherKeyBytes)
	if err != nil {
		t.Errorf("Error in AES encryption: %v\n", err)
	} else {
		actualPlainTextBytes, err := encryption.AESDecrypt(actualCipherBytes, cipherKeyBytes)
		if err != nil {
			t.Errorf("Error in AES decryption: %v\n", err)
		}
		assert.Equal(t, plainText, string(actualPlainTextBytes))
	}
}
