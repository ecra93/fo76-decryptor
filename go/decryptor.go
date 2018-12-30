package main

import (
  "strings"
  "sort"
  "io/ioutil"
  "fmt"
)

const Alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Keycard struct {
  letter  byte
  number  byte
}

func NewKeycard(letter byte, number byte) *Keycard {
  k := new(Keycard)
  k.letter = letter
  k.number = number
  return k
}

type Encoder struct {
  encoding map[byte]byte
}

func NewEncoder(keycards []Keycard) *Encoder {
  e := new(Encoder)
  e.encoding = make(map[byte]byte)
  for _, k := range keycards {
    e.encoding[k.letter] = byte(k.number)
  }
  return e
}

func (e Encoder) Encode(b byte) byte {
  return e.encoding[b]
}

type Decryptor struct {
  alphabet    string
  alternative string
}

func NewDecryptor(keyword string) *Decryptor {
  d := new(Decryptor)
  d.alphabet = Alphabet
  filter := func(r rune) rune {
    if strings.IndexRune(keyword, r)  < 0 {
      return r
    }
    return -1
  }
  d.alternative = keyword + strings.Map(filter, d.alphabet)
  return d
}

func (d Decryptor) Decrypt(word string) string {
  runeDecrypt := func(r rune) rune {
    return rune(d.alphabet[strings.IndexRune(d.alternative, r)])
  }
  decrypt := func(encrypted string) string {
    return strings.Map(runeDecrypt, encrypted)
  }
  return decrypt(word)
}

func (d Decryptor) Encrypt(word string) string {
  runeEncrypt := func(r rune) rune {
    return rune(d.alternative[strings.IndexRune(d.alphabet, r)])
  }
  encrypt := func(decrypted string) string {
    return strings.Map(runeEncrypt, decrypted)
  }
  return encrypt(word)
}

func StringSort(word string) string {
  tmp := strings.Split(word, "")
  sort.Strings(tmp)
  return strings.Join(tmp, "")
}

func Decrypt(keyword string, keycards []Keycard) (string, string) {

  decryptor := NewDecryptor(keyword)

  letters := ""
  for _, keycard := range keycards {
    letters = letters + string(keycard.letter)
  }
  fmt.Printf(letters + "\n")

  decrypted := StringSort(decryptor.Decrypt(letters))
  fmt.Printf(decrypted + "\n")

  anagram := ""
  dictionary, _ := ioutil.ReadFile("./dictionary.txt")
  words := strings.Fields(string(dictionary))
  for _, word := range words {
    if decrypted == StringSort(strings.ToUpper(word)) {
      _ = decrypted
      anagram = strings.ToUpper(word)
      fmt.Printf(anagram + "\n")
      break
    }
  }

  arranged := decryptor.Encrypt(anagram)
  fmt.Printf(arranged + "\n")

  encoder := NewEncoder(keycards)

  code := ""
  for _, letter := range arranged {
    code = code + string(encoder.Encode(byte(letter)))
  }
  fmt.Printf(code + "\n")

  return anagram, code
}
/*
func main() {
  slice := []Keycard{
    Keycard {
      letter: 'R',
      number: '6',
    },
    Keycard {
      letter: 'A',
      number: '6',
    },
    Keycard {
      letter: 'C',
      number: '3',
    },
    Keycard {
      letter: 'F',
      number: '5',
    },
    Keycard {
      letter: 'J',
      number: '6',
    },
    Keycard {
      letter: 'N',
      number: '0',
    },
    Keycard {
      letter: 'T',
      number: '2',
    },
    Keycard {
      letter: 'W',
      number: '3',
    },
  }
  Decrypt("WHITEBOARDS", slice)
}
*/
