package main

import (
	"os"

	"github.com/k0kubun/pp"
	fvuls "github.com/vuls-saas/futurevuls-sdk-go"
)

// FVULS_API_KEY=7f2156ab-96f8-11e8-b515-xxxxxx FVULS_RAW_URL=https://rest.vuls.biz go run cve.go
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

	// serverID := 2
	// roleID := 105
	// pkgID := 100
	// cpeID := 245

	cves, err := client.GetCveList(fvuls.GetCveListParam{
		// Page:   1,
		// Limit:  1,
		// Offset: 3,
		// FilterServerID: &serverID,
		// FilterRoleID: &roleID,
		// FilterPkgID: &pkgID,
		// FilterCpeID: &cpeID,
	})
	if err != nil {
		panic(err)
	}
	pp.Println(len(cves.Cves))
}
