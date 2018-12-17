package main

import (
	"os"

	"github.com/k0kubun/pp"
	fvuls "github.com/vuls-saas/futurevuls-sdk-go"
)

// FVULS_API_KEY=xxx-xxx-xxxx-xxxx-xxxxxx FVULS_RAW_URL=https://rest.vuls.biz go run task.go
func main() {
	client, err := fvuls.NewClientWithOptions(
		os.Getenv("FVULS_API_KEY"),
		os.Getenv("FVULS_RAW_URL"),
		true,
	)
	if err != nil {
		panic(err)
	}
	pp.Println(client.Health())
}
