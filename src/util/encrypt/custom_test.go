package encrypt

import (
	"testing"
	"util/logger"
)

func TestEncryptID(t *testing.T) {
	id := "3555"
	cookedID, err := EncryptID(id)
	if err != nil {
		t.Error(err)
	}
	logger.Info.Println(cookedID)
}

func TestDecryptID(t *testing.T) {
	id := "hvqcefefeymtnhg76qe"
	cookedID, err := DecryptID(id)
	if err != nil {
		t.Error(err)
	}
	logger.Info.Println(cookedID)
}
