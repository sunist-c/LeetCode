package add_binary

import "math/big"

func addBinary(a string, b string) string {
	numberA, _ := new(big.Int).SetString(a, 2)
	numberB, _ := new(big.Int).SetString(b, 2)
	numberA.Add(numberA, numberB)
	return numberA.Text(2)
}
