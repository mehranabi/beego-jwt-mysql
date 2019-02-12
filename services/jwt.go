package services

import (
	"crypto/rsa"
	"github.com/SermoDigital/jose/crypto"
	"github.com/gbrlsnchs/jwt"
	"io/ioutil"
	"time"
)

// JWT default/fixed claims
var (
	iss = "api"
	sub = "uid-"
	aud = "client"
	exp = 24 * 30 * time.Hour // 30 days
	nbf = 30 * time.Second    // 30 seconds
)

// Get public/private keys
func GetKeyPair() (private *rsa.PrivateKey, public *rsa.PublicKey, err error) {
	// Read private key file
	privKeyBytes, e := ioutil.ReadFile("./keys/private.txt")
	if e != nil {
		return nil, nil, e
	}
	// Create private key object
	privKey, e := crypto.ParseRSAPrivateKeyFromPEM(privKeyBytes)
	if e != nil {
		return nil, nil, e
	}
	// Read public key file
	pubKeyBytes, e := ioutil.ReadFile("./keys/public.txt")
	if e != nil {
		return nil, nil, e
	}
	// Create public key object
	pubKey, e := crypto.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if e != nil {
		return nil, nil, e
	}

	return privKey, pubKey, nil
}

// Validate JWT token
func ValidateToken(token string, uid int64) (valid bool, err error) {
	// Get keys
	privKey, pubKey, e := GetKeyPair()
	if e != nil {
		return false, e
	}

	// Validating
	now := time.Now()
	signer := jwt.NewRS512(privKey, pubKey)

	// First, extract the payload and signature.
	// This enables un-marshaling the JWT first and verifying it later or vice versa.
	payload, sig, e := jwt.Parse(token)
	if e != nil {
		return false, e
	}

	// Check signature
	if e = signer.Verify(payload, sig); e != nil {
		return false, e
	}

	// Un-Marshal token
	var jot jwt.JWT
	if e = jwt.Unmarshal(payload, &jot); e != nil {
		return false, e
	}

	// Validators
	audV := jwt.AudienceValidator(aud)
	expV := jwt.ExpirationTimeValidator(now)
	issV := jwt.IssuerValidator(iss)
	nbfV := jwt.NotBeforeValidator(now)
	subV := jwt.SubjectValidator(sub + string(int(uid)))

	// Validate
	if e := jot.Validate(audV, expV, issV, nbfV, subV); e != nil {
		return false, e
	}

	return true, nil
}

// Generate JWT token for user
func MakeToken(uid int64) (token string, err error) {
	// Get keys
	privKey, pubKey, e := GetKeyPair()
	if e != nil {
		return "", e
	}

	// Generate JWT Token for user
	now := time.Now()
	signer := jwt.NewRS512(privKey, pubKey)
	jot := &jwt.JWT{
		Issuer:         iss,
		Subject:        sub + string(int(uid)),
		Audience:       aud,
		ExpirationTime: now.Add(exp).Unix(),
		NotBefore:      now.Add(nbf).Unix(),
		IssuedAt:       now.Unix(),
	}
	jot.SetAlgorithm(signer)
	payload, e := jwt.Marshal(jot)
	if e != nil {
		return "", e
	}
	t, e := signer.Sign(payload)
	if e != nil {
		return "", e
	}

	// Return token
	return string(t), nil
}
