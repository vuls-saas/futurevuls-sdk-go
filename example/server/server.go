package main

import (
	"os"

	"github.com/k0kubun/pp"
	fvuls "github.com/vuls-saas/futurevuls-sdk-go"
)

// FVULS_API_KEY=xxx-xxx-xxxx-xxxx-xxxxxx FVULS_RAW_URL=https://rest.vuls.biz go run server.go
func main() {
	client, err := fvuls.NewClientWithOptions(
		os.Getenv("FVULS_API_KEY"),
		os.Getenv("FVULS_RAW_URL"),
		true,
	)
	if err != nil {
		panic(err)
	}
	// server, err := client.GetServerDetail(fvuls.GetServerDetailParam{
	// ServerID: "9398",
	// })
	// if err != nil {
	// panic(err)
	// }
	// pp.Println(server)

	// serverID := 9315
	// roleID := 1052
	// pkgID := 100
	// cpeID := 245

	servers, err := client.GetAllServerList(fvuls.GetServerListParam{})
	// Page:   1,
	// Limit: 1,
	// Offset: 3,
	// FilterServerID: &serverID,
	// FilterRoleID: &roleID,
	// FilterPkgID: &pkgID,
	// FilterCpeID: &cpeID,
	// })
	if err != nil {
		panic(err)
	}
	pp.Println(servers)

	// s := "c7-2-1"
	// server, err := client.UpdateServer(fvuls.UpdateServerParam{
	// ServerID:   14443,
	// ServerName: &s,
	// })
	// Page:   1,
	// Limit: 1,
	// Offset: 3,
	// FilterServerID: &serverID,
	// FilterRoleID: &roleID,
	// FilterPkgID: &pkgID,
	// FilterCpeID: &cpeID,
	// })
	// if err != nil {
	// panic(err)
	// }
	// pp.Println(server.ServerName)

	// err = client.DeleteServer(fvuls.DeleteServerParam{
	// ServerID: 13323,
	// })
	// if err != nil {
	// panic(err)
	// }
}
