package payDemo

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sort"
	"testing"
)

func TestSign(t *testing.T) {
	message := "Hello, world!" // 要签名的数据

	// 生成 RSA 密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 使用 SHA-256 哈希算法计算数据的摘要
	hash := sha256.Sum256([]byte(message))

	// 使用私钥对摘要进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		fmt.Println(err)
		return
	}

	// 将签名结果进行 Base64 编码，方便传输和存储
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	fmt.Println("Message:", message)
	fmt.Println("Signature:", signatureBase64)
}

func Verify(params map[string]string, pub *rsa.PublicKey, sign string) (err error) {
	// 和签名步骤相同，对收到的请求参数按照字母顺序进行排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var signature string = ""
	i := 0
	for _, k := range keys {
		v := params[k]
		if i != 0 {
			signature += "&"
		}
		signature += k
		signature += "="
		signature += v
		i++
	}
	// 和签名步骤相同，将排序后的signature进行hash操作
	h := sha256.New()
	h.Write([]byte(signature))
	Sha256Code := h.Sum(nil)
	// 对签名进行base64解码
	decodeSignature, err := base64.StdEncoding.DecodeString(sign)
	// 使用rsa验签函数
	// 第一个参数是公钥
	// 第二个参数是hash函数
	// 第三个参数是被hash函数处理过的原始输入
	// 第四个参数是被处理过的签名
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, Sha256Code, decodeSignature)
	if err != nil { // 验证失败
		return err
	}
	return nil // 验证成功
}
