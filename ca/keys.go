package ca

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	cloud_aws "github.com/ayush18023/Load_balancer_Fyp/internal/aws"
	"golang.org/x/crypto/ssh"
)

func RSAKeys() ([]byte, []byte) {
	bitSize := 4096

	privateKey, err := generatePrivateKey(bitSize)
	if err != nil {
		log.Fatal(err.Error())
	}

	publicKeyBytes, err := generatePublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatal(err.Error())
	}

	privateKeyBytes := encodePrivateKeyToPEM(privateKey)
	return publicKeyBytes, privateKeyBytes
}

func GenerateSSHKeyPairs(algo string, keyname string) (string, error) {
	// savePrivateFileTo := "./id_rsa_test"
	// savePublicFileTo := "./id_rsa_test.pub"
	var publicKey, privateKey []byte
	if strings.ToLower(algo) == "rsa" {
		publicKey, privateKey = RSAKeys()
	}
	pubreader := bytes.NewReader(publicKey)
	prireader := bytes.NewReader(privateKey)
	bucket := cloud_aws.NewS3()
	_, err := bucket.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(fmt.Sprintf("%s.pem", keyname)),
		Body:   prireader,
	})
	if err != nil {
		return string(privateKey), err
	}
	_, err = bucket.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(fmt.Sprintf("%s.pem.pub", keyname)),
		Body:   pubreader,
	})
	if err != nil {
		return string(privateKey), err
	}
	// err = writeKeyToFile(privateKeyBytes, savePrivateFileTo)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// err = writeKeyToFile([]byte(publicKeyBytes), savePublicFileTo)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	return string(privateKey), nil
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	log.Println("Private Key generated")
	return privateKey, nil
}

// encodePrivateKeyToPEM encodes Private Key from RSA to PEM format
func encodePrivateKeyToPEM(privateKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := pem.EncodeToMemory(&privBlock)

	return privatePEM
}

// generatePublicKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func generatePublicKey(privatekey *rsa.PublicKey) ([]byte, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return nil, err
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRsaKey)

	log.Println("Public key generated")
	return pubKeyBytes, nil
}

// writePemToFile writes keys to a file
func writeKeyToFile(keyBytes []byte, saveFileTo string) error {
	err := ioutil.WriteFile(saveFileTo, keyBytes, 0600)
	if err != nil {
		return err
	}

	log.Printf("Key saved to: %s", saveFileTo)
	return nil
}
