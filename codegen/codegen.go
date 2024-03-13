package codegen

import (
    "time"
    "math/rand"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandSeq(n int) string {
    b := make([]rune, n)
        
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

