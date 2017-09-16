package main

import (
	"fmt"
	"io"
	"math"
	"math/big"
	"os"
)

func main() {
	pi(math.MaxInt64, os.Stdout)
}

func pi(digits int64, r io.Writer) {
	rem := big.NewRat(0, 1)
	for k := int64(0); k < digits; k++ {
		rem.Mul(rem, big.NewRat(16, 1))

		rem.Add(rem, big.NewRat(4, 8*k+1))
		rem.Sub(rem, big.NewRat(2, 8*k+4))
		rem.Sub(rem, big.NewRat(1, 8*k+5))
		rem.Sub(rem, big.NewRat(1, 8*k+6))

		flt, _ := rem.Float64()
		digit := int64(flt)
		rem.Sub(rem, big.NewRat(digit, 1))

		fmt.Fprintf(r, "%x", digit)
	}
}
