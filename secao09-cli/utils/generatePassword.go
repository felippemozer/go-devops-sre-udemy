package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func GeneratePassword(size int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	if size < 8 {
		fmt.Println("Tamanho mínimo de 8 caracteres")
		return ""
	}
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		fmt.Println("Erro ao gerar senha")
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}
