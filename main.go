package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

// Wallet represents an Ethereum wallet
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    common.Address
}

// NewWallet generates a new Ethereum wallet
func NewWallet() *Wallet {
	privateKey, _ := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	publicKey := &privateKey.PublicKey
	address := crypto.PubkeyToAddress(*publicKey)
	return &Wallet{privateKey, publicKey, address}
}

// ExportPrivateKey exports the private key of an Ethereum wallet as a hexadecimal string
func (w *Wallet) ExportPrivateKey() string {
	return fmt.Sprintf("%x", crypto.FromECDSA(w.PrivateKey))
}

// ImportPrivateKey imports an Ethereum wallet from a private key hexadecimal string
func ImportPrivateKey(privateKeyHex string) *Wallet {
	privateKey, _ := crypto.HexToECDSA(privateKeyHex)
	publicKey := &privateKey.PublicKey
	address := crypto.PubkeyToAddress(*publicKey)
	return &Wallet{privateKey, publicKey, address}
}

func main() {
	// Generate a new wallet
	wallet := NewWallet()

	// Export the private key of the wallet
	privateKeyHex := wallet.ExportPrivateKey()
	fmt.Println("Private Key:", privateKeyHex)

	// Import the wallet from the private key
	wallet = ImportPrivateKey(privateKeyHex)
	fmt.Println("Address:", wallet.Address.Hex())
}
