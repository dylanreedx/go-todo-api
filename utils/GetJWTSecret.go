package utils

import "os"

func Secret() []byte {
	LoadEnv()
	return []byte(os.Getenv("SECRET"))
}
