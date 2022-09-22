// SSL証明書と秘密鍵の生成(教育用)

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	// 証明書の構成を設定するための構造体Certificate
	// SSL証明書とは実質的にはX.509証明書の拡張鍵用途フィールドをサーバ認証に設定したもの
	template := x509.Certificate{
		SerialNumber: serialNumber, // ランダムに発生させたとても大きな整数
		Subject: subject,	    // 識別名
		NotBefore: time.Now(),
		NotAfter: time.Now().Add(365 * 24 * time.Hour), // 証明書が作成された日から1年間
		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // このX.509証明書がサーバ認証に使用されることを示す
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")}, // 証明書が127.0.0.1でだけ効力を持つ
	}

	// 証明書の作成には秘密鍵が必須。構造体には公開鍵が入っていてアクセスできる
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// DER形式のバイトデータのスライスを生成する
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create("cert.pem")
	// encoding/pemを使って、その証明書データを符号化してcert.pemというファイルにする
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	// 前に生成した鍵をPEM符号化してkey.pemというファイルにする
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
