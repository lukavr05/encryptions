package main

import (
	"fmt"
	"strconv"
	"time"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func generateTimeSeed() string {
	t := time.Now()
	t_now := t.Format("20060102150405")

	bin_t := ""
	for _, c := range t_now {
		bin_t = fmt.Sprintf("%s%.8b", bin_t, c)
	}

	return bin_t
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// simple XOR function, idc if it already exists
func xOR(a string, b string) string {
	i := Must(strconv.Atoi(a))
	j := Must(strconv.Atoi(b))

	r := (i + j) % 2

	result := strconv.Itoa(r)

	return result
}

func stringToBin(s string) string {
	binString := ""

	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return binString
}

// pad the keystream so that it matches the length of the input string
func padKeystream(length int, keystream string) string {
	if len(keystream) >= length {
		return keystream[0:length]
	} else {
		reversed := reverseString(keystream)
		padded := ""

		useOriginal := true

		for len(padded) < length {
			if useOriginal {
				padded += keystream
			} else {
				padded += reversed
			}
			useOriginal = !useOriginal
		}
		return padded[:length]
	}
}

func streamCipherEncrypt(plaintext string, keystream string) string {
	ciphertext := ""
	p_runes := []rune(plaintext)
	k_runes := []rune(keystream)

	for i := 0; i < len(keystream); i++ {
		a := string(p_runes[i])
		b := string(k_runes[i])
		ciphertext += xOR(a, b)
	}

	return ciphertext
}

func main() {
	k := generateTimeSeed()
	plaintext := "this is a secret message"
  p := stringToBin(plaintext)
	padded := padKeystream(len(p), k)

	fmt.Printf("Plaintext:        %s\n", plaintext)
	fmt.Printf("Keystream:        %s\n", k)
	fmt.Printf("Plaintext(b):     %s\n", p)
	fmt.Printf("Padded Keystream: %s\n", padded)

  c := streamCipherEncrypt(p, padded)

  fmt.Printf("Encrypted:        %s\n", c)
}
