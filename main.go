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

// simple XOR function, idc if it already exists
func xOR(a string, b string) string {
  i := Must(strconv.Atoi(a))
  j := Must(strconv.Atoi(b))

  r := (i + j) % 2

  result := strconv.Itoa(r)

  return result
}

// pad the keystream so that it matches the length of the input string
func padKeystream(plaintext string, keystream string) string {
  if (len(keystream) > len(plaintext)) {
    return keystream[0:len(plaintext)]
  }

  if (len(keystream) < len(plaintext)) {
    new_stream 
  }
}


func streamCipher(plaintext string, keystream string) string {
  
  for 
}

func main() {
  s := generateTimeSeed()

}
