package secret_share

import (
	"2026CryptoStudy/is_prime"
	"math/big"
)

type Equation struct {
	Remainder *big.Int
	Modulus   *big.Int
}

func SolveSimpleCRT(equations []Equation) *big.Int {
	// finish SolveSimpleCRT. You can always assume that the moduli are pairwise coprime.
	//return nil
	M := big.NewInt(1)
	for _, eq := range equations {
		M.Mul(M, eq.Modulus)
	}
	result := big.NewInt(0)
	for _, eq := range equations {
		a := eq.Remainder
		m := eq.Modulus
		Mi := new(big.Int).Div(M, m)
		yi := new(big.Int).ModInverse(Mi, m)
		if yi == nil {
			return nil
		}
		term := new(big.Int).Mul(a, Mi)
		term.Mul(term, yi)
		result.Add(result, term)
	}
	return result.Mod(result, M)
}

func SolveSimple(equations []Equation) *big.Int {
	// TODO: finish SolveSimpleCRT. You should check whether modulus is prime.
	for _, eq := range equations {
		if !is_prime.IsPrime(eq.Modulus) {
			return nil
		}
	}
	return SolveSimpleCRT(equations)
}
