package main

import (
	"os"

	"github.com/k0kubun/pp"
	fvuls "github.com/vuls-saas/futurevuls-sdk-go"
)

// FVULS_API_KEY=xxx-xxx-xxxx-xxxx-xxxxxx FVULS_RAW_URL=https://rest.vuls.biz go run cve.go
func main() {
	client, err := fvuls.NewClientWithOptions(
		os.Getenv("FVULS_API_KEY"),
		os.Getenv("FVULS_RAW_URL"),
		true,
	)
	if err != nil {
		panic(err)
	}
	cve, err := client.GetCveDetail(fvuls.GetCveDetailParam{
		CveID: "CVE-2014-9474",
	})
	if err != nil {
		panic(err)
	}
	pp.Println(cve.CveID)

	// serverID := 9315
	// roleID := 1052
	// pkgID := 100
	// cpeID := 245

	cves, err := client.GetAllCveList(fvuls.GetCveListParam{
		// FilterServerID: &serverID,
		// FilterRoleID:   &roleID,
		// FilterPkgID:    &pkgID,
		// FilterCpeID:    &cpeID,
	})
	if err != nil {
		panic(err)
	}
	pp.Println(cves)
	pp.Println(len(cves))
}
