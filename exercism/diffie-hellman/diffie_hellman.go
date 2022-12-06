package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

var b2 = big.NewInt(2)

func PrivateKey(p *big.Int) *big.Int {
	m := new(big.Int).Sub(p, b2)
	n, err := rand.Int(rand.Reader, m)
	if err != nil {
		panic(err)
	}
	return n.Add(b2, n)
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	return private, PublicKey(private, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
