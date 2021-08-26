package generator

import (
	"com.nguyenonline/formipro/internal"
	"log"
	"math/rand"
	"os"
	"time"
)

var digits = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandDir() (string, error) {
	dirName := time.Now().Format("20060102150405") + "_" + randString(10)
	err := os.Mkdir(internal.TmpDir+dirName, 0777)
	if err != nil {
		log.Printf("Could not create directory '%s', error is '%s'\n", dirName, err)
		return "", err
	}
	return dirName, nil
}

func randString(n int) string {
	result := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = digits[rand.Intn(len(digits))]
	}
	return string(result)
}
