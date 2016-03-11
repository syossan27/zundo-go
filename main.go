package main

import (
  "fmt"
  "math/rand"
  "time"
  "reflect"
  "os"
)

func main() {
  zun := "ズン"
  doko := "ドコ"
  zundoko := []string{zun, doko}
  zzzzdoko := []string{zun, zun, zun, zun, doko}
  lyrics := []string{}
  rand.Seed(time.Now().UnixNano())

  for {
    time.Sleep(1 * time.Second)
    word := zundoko[rand.Intn(2)]
    fmt.Print(word + " ")
    lyrics = append(lyrics, word)
    if len(lyrics) >= 5 {
      length := len(lyrics)
      some_lyrics := lyrics[length-5:]
      if reflect.DeepEqual(some_lyrics, zzzzdoko) {
        fmt.Print("キ・ヨ・シ！")
        os.Exit(0)
      }
    }
  }
}
