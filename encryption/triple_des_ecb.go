package encryption

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/errors"
	"strings"
)

//Des encryption
func encrypt(origData, key []byte) ([]byte, error) {
	if len(origData) < 1 || len(key) < 1 {
		return nil, errors.InternalServerError("wrong data or key")
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(origData)%bs != 0 {
		return nil, errors.InternalServerError("wrong padding")
	}
	out := make([]byte, len(origData))
	dst := out
	for len(origData) > 0 {
		block.Encrypt(dst, origData[:bs])
		origData = origData[bs:]
		dst = dst[bs:]
	}
	return out, nil
}

//Des decryption
func decrypt(crypted, key []byte) ([]byte, error) {
	if len(crypted) < 1 || len(key) < 1 {
		return nil, errors.InternalServerError("wrong data or key")
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	out := make([]byte, len(crypted))
	dst := out
	bs := block.BlockSize()
	if len(crypted)%bs != 0 {
		return nil, errors.InternalServerError("wrong crypted size")
	}

	for len(crypted) > 0 {
		block.Decrypt(dst, crypted[:bs])
		crypted = crypted[bs:]
		dst = dst[bs:]
	}

	return out, nil
}

//[golang ECB 3DES Encrypt]
func TripleDesECBEncrypt(origData, key []byte) ([]byte, error) {
	tkey := make([]byte, 24, 24)
	copy(tkey, key)
	k1 := tkey[:8]
	k2 := tkey[8:16]
	k3 := tkey[16:]

	block, err := des.NewCipher(k1)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	origData = PKCS5Padding(origData, bs)

	buf1, err := encrypt(origData, k1)
	if err != nil {
		return nil, err
	}
	buf2, err := decrypt(buf1, k2)
	if err != nil {
		return nil, err
	}
	out, err := encrypt(buf2, k3)
	if err != nil {
		return nil, err
	}
	return out, nil
}

//[golang ECB 3DES Decrypt]
func TripleDesECBDecrypt(crypted, key []byte) ([]byte, error) {
	tkey := make([]byte, 24, 24)
	copy(tkey, key)
	k1 := tkey[:8]
	k2 := tkey[8:16]
	k3 := tkey[16:]
	buf1, err := decrypt(crypted, k3)
	if err != nil {
		return nil, err
	}
	buf2, err := encrypt(buf1, k2)
	if err != nil {
		return nil, err
	}
	out, err := decrypt(buf2, k1)
	if err != nil {
		return nil, err
	}
	out = PKCS5UnPadding(out)
	return out, nil
}

func StringToTripleDesECBEncrypt(plainText, key string) (*string, error) {
	if len(key) < 24 {
		return nil, errors.InternalServerError("Wrong key size, size key must be 24 character")
	}

	crypted, err := TripleDesECBEncrypt([]byte(plainText), []byte(key))

	if err != nil {
		return nil, errors.InternalServerError("Something wrong when encripted triple des")
	}

	encriptedString := strings.ToUpper(fmt.Sprintf("%x \n", string(crypted)))

	return &encriptedString, err
}

func StructToTripleDesECBEncrypt(plainData interface{}, key string) (*string, error) {
	if len(key) < 24 {
		return nil, errors.InternalServerError("Wrong key size, size key must be 24 character")
	}

	marshaledData, err := json.Marshal(plainData)

	if err != nil {
		return nil, err
	}

	crypted, err := TripleDesECBEncrypt(marshaledData, []byte(key))

	if err != nil {
		return nil, errors.InternalServerError("Something wrong when encripted triple des")
	}

	encriptedString := strings.ToUpper(fmt.Sprintf("%x \n", string(crypted)))

	return &encriptedString, err
}

func StringToTripleDesECBDecrypt(encrypted, key string) (*string, error) {
	if len(key) < 24 {
		return nil, errors.InternalServerError("Wrong key size, size key must be 24 character")
	}

	toHex, _ := hex.DecodeString(encrypted)
	decrypted, err := TripleDesECBDecrypt(toHex, []byte(key))

	if err != nil {
		return nil, errors.InternalServerError("Something wrong when decrypted triple des")
	}

	decryptedString := fmt.Sprintf("%v \n", string(bytes.Trim(decrypted, "\x00")))

	return &decryptedString, err
}
