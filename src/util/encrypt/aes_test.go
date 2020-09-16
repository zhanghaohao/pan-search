package encrypt

import (
	"testing"
	"util/logger"
)

func TestAesCBCEncrypt(t *testing.T) {

}

func TestEncrypt(t *testing.T) {
	id := []byte("1")
	key := []byte("fjincb56u95bza10")
	encryptedID, err := Encrypt(id, key)
	if err != nil {
		t.Error(err)
	}
	logger.Info.Println(encryptedID)
}

func TestAesCBCDecrypt(t *testing.T) {

}

func TestDecrypt(t *testing.T) {
	encryptedID := "atLKOrf0uiASGh/lyUaQetRwYANtV9O49614tF62QUE="
	key := []byte("fjincb56u95bza10")
	id, err := Decrypt(encryptedID, key)
	if err != nil {
		t.Error(err)
	}
	logger.Info.Println(id)
}

