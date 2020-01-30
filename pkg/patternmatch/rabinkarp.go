// References:
// https://www.youtube.com/watch?v=H4VrKHVG5qI
// https://www.topcoder.com/community/data-science/data-science-tutorials/introduction-to-string-searching-algorithms/


// Rabin carp algorithm to find pattern p in text t.
// Approach :- Follow naive solution with one addition - Use hash codes of pattern and same size char set from set iteratively
// and compare hash codes, if match then only compare char by char of both.
// TODO- Use modulo for keeping hash codes smalled which comes with a cost called - collision, there fore we should use
// double hashing using two primes
package patternmatch

import (
  "fmt"
)

const PRIME = 101

func RabinKarp(t, p string) int {
  tn := len(t)
  pn := len(p)

  // Only one time we need to calculate hash function from rollingHash
  hp := rollingHash(p)
  h0 := rollingHash(t[0:pn])
  if h0 == hp {
    // Check for each character
    match := true
    fmt.Println ("Hash code match", hp)
    for j:=0; j<pn; j++ {
      if p[j] != t[j] {
        match = false
        break
      }
    }

    if match {
      return 0
    }

  }

  // Next time we can evaluate from previosuly calculated Hashcode (Rolling Hashing) 
  // Refer Tushar Youtube link given
  // h1 = h0 - ascii(prev window's last Char)
  // h1 = h1 / PRIME
  // h1 = h1 + ascii(next/current window's last character) * pow(prime, sizeOfPattern -1)
  var i int
  match := false
  for i = 1; i <= tn-pn; i++ {
    ht := (h0 - int(t[i-1])) / PRIME
    ht = ht + int(t[i+pn-1])*pow(PRIME, pn-1)
    if ht == hp {
      // Check for each character
      match = true
      for j:=0; j<pn; j++ {
          if p[j] != t[j+i] {
            match = false
            break
          }
      }
    }
    if match {
      break
    }
    h0 = ht
  }


  if match {
    return i
  } else {
    return -1
  }


}

func rollingHash(s string) (int) {
  var h int
  for i:=0; i<len(s); i++ {
    h = h + int(s[i])*pow(PRIME, i)
  }
  return h
}

func pow(m, n int) int {
  p := 1
  for n > 0 {
    p = m * p
    n--
  }

  return p
}

func main() {
  t := "abdebctfgebctfebctfebctfebctfebctfebctfebctf ebctfebctfebctfebctfebctfebctfebctfebctfebctfebct febctfebctfebctfebctfebctfebctfebctfebctfebctfebct febctfebctfebctfebctfebctfebctfebctfebctfebctfebctf ebctfebctfebctfebctfebctfebctfebctfebctfebctfebctf ebctfebctfebctfebctfebctfebctfebctfebctfebctfebctfebctfebctfebctfepctf"
  p := "ctfepct"
  im := RabinKarp(t, p)
  if im < 0 {
    fmt.Println("Error: No pattern matched")
    return
  }
  fmt.Printf("matching string at index %d is %s", im, t[im:im+len(p)])
}
