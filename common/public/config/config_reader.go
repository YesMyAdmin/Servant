package config;

import (
	"github.com/spf13/viper"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
    "crypto/sha256"
	"fmt"
	"os"
	"bytes"
)

var cfg *viper.Viper

func init() {

}

// DecryptConfig 解密配置文件,并读取配置到 viper 实例中
// masterPassword: 主密码，用于解密文件
// encryptedConfigPath: 加密的配置文件路径
func DecryptConfig(masterPassword string, encryptedConfigPath string) error {
	plaintext, err := doEncrtpt(masterPassword, encryptedConfigPath)
	if err != nil {
		return err
	}
	cfg = viper.New()
	cfg.ReadConfig(bytes.NewBuffer(plaintext))
	return nil
}

// EncryptConfig 加密配置文件,然后删除明文配置文件
// masterPassword: 主密码，用于加密文件
// plaintextConfigPath: 明文配置文件路径
// encryptedConfigPath: 加密后的配置文件路径
func EncryptConfig(masterPassword string, plaintextConfigPath string, encryptedConfigPath string) error {
	key := hashedMasterPassword(masterPassword)
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}
	plaintext, err := os.ReadFile(plaintextConfigPath)
	if err != nil {
		return err
	}
	nonce := make([]byte, gcm.NonceSize())
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	// 将加密后的数据写入文件
	err = os.WriteFile(encryptedConfigPath, ciphertext, 0644)
	if err != nil {
		return err
	}
	// 删除明文配置文件
	err = os.Remove(plaintextConfigPath)
	if err != nil {
		return err
	}
	return nil
}

// DumpConfigToFile 将配置文件解密后写入到指定路径
// masterPassword: 主密码，用于解密文件
// plaintextConfigPath: 明文配置文件路径
// encryptedConfigPath: 加密后的配置文件路径
func DumpConfigToFile(masterPassword string, plaintextConfigPath string, encryptedConfigPath string) error {
	plaintext, err := doEncrtpt(masterPassword, encryptedConfigPath)
	if err != nil {
		return err
	}
	return os.WriteFile(plaintextConfigPath, plaintext, 0644)
}

// hashedMasterPassword 生成主密码的哈希值，用于加密和解密配置文件
// masterPassword: 主密码
func hashedMasterPassword(masterPassword string) []byte {
	key := []byte(masterPassword)
	message := []byte("Yes,my admin.")

	h := hmac.New(sha256.New, key)
	h.Write(message)
	return h.Sum(nil)
}

func doEncrtpt(masterPassword string, encryptedConfigPath string) ([]byte, error) {
	key := hashedMasterPassword(masterPassword)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ciphertext, err := os.ReadFile(encryptedConfigPath)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	// 提取 nonce 和实际密文
	nonce, ciphertextActual := ciphertext[:nonceSize], ciphertext[nonceSize:]
	// 解密并验证
	plaintext, err := gcm.Open(nil, nonce, ciphertextActual, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// Get 获取配置项的值
// key: 配置项的键
func Get(key string) any {
	if cfg == nil {
		panic("config not initialized, please decrypt config file first")
	}
	return cfg.Get(key)
}