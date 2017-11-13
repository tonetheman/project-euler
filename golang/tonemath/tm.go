package tonemath

import "math"

func IsPrime(n int) bool {

  // the are prime
  if n == 1 || n == 2 || n ==3 {
    return true
  }

  // even is not prime ever (except 2)
  if n%2==0 {
    return false
  }

  // something is broke here
  for i:=3;i<int(math.Sqrt(float64(n)))+1;i+=2 {
    if n%i == 0 {
      return false
    }
  }

  return true
}
