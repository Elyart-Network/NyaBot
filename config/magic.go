package config

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"github.com/Elyart-Network/NyaBot/logger"
)

var MagicKey = hex.EncodeToString([]byte("WulYAa8M4FaSITGNUJyk6gOifgzV80uh"))

func EncodeMagic(ctx string) string {
	b64e := base64.StdEncoding.EncodeToString([]byte(ctx))
	key, _ := hex.DecodeString(MagicKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Errorf("Magic", err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		logger.Errorf("Magic", err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(b64e), nil)
	return hex.EncodeToString(ciphertext)
}

func DecodeMagic(ctx string) string {
	key, _ := hex.DecodeString(MagicKey)
	enc, _ := hex.DecodeString(ctx)
	block, err := aes.NewCipher(key)
	if err != nil {
		logger.Errorf("Magic", err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		logger.Errorf("Magic", err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		logger.Errorf("Magic", err.Error())
	}
	b64d, err := base64.StdEncoding.DecodeString(string(plaintext))
	if err != nil {
		logger.Errorf("Magic", err.Error())
	}
	return string(b64d)
}
