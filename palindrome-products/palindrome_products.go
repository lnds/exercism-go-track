package palindrome

import "fmt"

type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {

	if fmin > fmax {
		err = fmt.Errorf("fmin > fmax")
		return
	}

	pmin.Product = fmax * fmax
	pmax.Product = fmin * fmin
	palindromes := 0
	for x := fmin; x <= fmax; x++ {
		for y := fmin; y <= fmax; y++ {
			value := x * y
			if !isPalindrome(value) {
				continue
			}
			palindromes++
			if value == pmin.Product {
				pmin.Factorizations = insert(pmin.Factorizations, [2]int{x, y})
			} else if value < pmin.Product {
				pmin = New(x, y)
			}
			if value == pmax.Product {
				pmax.Factorizations = insert(pmax.Factorizations, [2]int{x, y})
			} else if value > pmax.Product {
				pmax = New(x, y)
			}
		}
	}
	if palindromes == 0 {
		err = fmt.Errorf("no palindromes")
	}
	return
}

func New(a, b int) Product {
	return Product{
		Product:        a * b,
		Factorizations: [][2]int{{a, b}},
	}
}

func insert(factors [][2]int, factor [2]int) [][2]int {
	for _, f := range factors {
		if (f[0] == factor[1] && f[1] == factor[0]) || (f[0] == factor[0] && f[1] == factor[1]) {
			return factors
		}
	}
	return append(factors, factor)
}

func isPalindrome(value int) bool {
	reversed := 0
	n := value
	for n > 0 {
		reversed = 10*reversed + n%10
		n /= 10
	}
	return reversed == value
}
