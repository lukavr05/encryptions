package main;

import (
  "fmt"
  "time"
  "strconv"
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

// pad the keystream so that it matches the length of the input string
func padKeystream(length int, keystream string) string {
  if (len(keystream) >= length) {
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

func main() {
  s := generateTimeSeed()
  plaintext := "hello"
  padded := padKeystream((len(plaintext) * 8), s)

  fmt.Printf("Plaintext:        %s\n", plaintext)
  fmt.Printf("Keystream:        %s\n", s)
  fmt.Printf("Padded Keystream: %s\n", padded)

}
