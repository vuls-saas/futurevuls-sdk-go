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

	// role, err := client.GetRoleDetail(fvuls.GetRoleDetailParam{
	// RoleID: 1311,
	// })
	// if err != nil {
	// panic(err)
	// }
	// pp.Println(role)

	//TODO test with cveID param
	cveID := "CVE-2017-14167"
	roles, err := client.GetAllRoleList(fvuls.GetRoleListParam{
		FilterCveID: &cveID,
	})
	if err != nil {
		panic(err)
	}
	pp.Println(roles)

	// s := "web"
	// role, err := client.UpdateRole(fvuls.UpdateRoleParam{
	// RoleID:   1311,
	// RoleName: &s,
	// })
	// if err != nil {
	// panic(err)
	// }
	// pp.Println(role)

	// err = client.DeleteRole(fvuls.DeleteRoleParam{
	// RoleID: 1526,
	// })
	// if err != nil {
	// panic(err)
	// }
}
