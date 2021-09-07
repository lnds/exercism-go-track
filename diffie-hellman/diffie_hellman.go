package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

var Two = big.NewInt(2)
var r = rand.New(rand.NewSource(time.Now().Unix()))

func PrivateKey(p *big.Int) *big.Int {
	key := big.NewInt(0).Sub(p, Two)
	return key.Add(Two, key.Rand(r, key))
}

func PublicKey(a, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), a, p)
}

func SecretKey(a, B, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(B, a, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	b := PublicKey(a, p, g)
	return a, b
}
