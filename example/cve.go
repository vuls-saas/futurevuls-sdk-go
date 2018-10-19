package main

import (
	"os"

	"github.com/k0kubun/pp"
	fvuls "github.com/vuls-saas/futurevuls-sdk-go"
)

func main() {
	client := fvuls.NewClient(os.Getenv("FVULS_API_KEY"))
	cve, err := client.GetCveDetail(fvuls.GetCveDetailPayload{
		CveID: "CVE-2014-9474",
	})
	if err != nil {
		panic(err)
	}
	pp.Println(cve)
}
