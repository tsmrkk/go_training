package main

import (
	"errors"
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	var err error
	var res float64
	if x >= 0 {
		res = math.Sqrt(x)
		return res, err
	}
	err = errors.New(ErrNegativeSqrt(x).Error())
	return x, err
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative nuber: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
