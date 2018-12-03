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
	// cpeID := 448
	// pkgCpe, err := client.GetCpeDetail(fvuls.GetCpeDetailParam{
	// CpeID: cpeID,
	// })
	// if err != nil {
	// panic(err)
	// }
	// pp.Println(pkgCpe)

	// pkgID := 2610003
	// pkgCpe, err := client.GetPkgDetail(fvuls.GetPkgDetailParam{
	// PkgID: pkgID,
	// })

	// if err != nil {
	// panic(err)
	// }
	// pp.Println(pkgCpe)

	// serverID := 9315
	// roleID := 105
	// pkgID := 100
	// cpeID := 245
	// ignore := true
	// stats := []string{
	// "patch_applied",
	// "new",
	// }

	// tasks, err := client.GetTaskList(fvuls.GetTaskListPayload{
	// // Page:   1,
	// // Limit: 1,
	// // Offset: 3,
	// FilterStatus: stats,
	// // FilterServerID: &serverID,
	// // // FilterRoleID: &roleID,
	// // // FilterPkgID: &pkgID,
	// // // FilterCpeID: &cpeID,

	// // FilterIgnore: &ignore,
	// // FilterMainUserIDs: []int{112},
	// })
	// if err != nil {
	// panic(err)
	// }
	// pp.Println(tasks)
	// pp.Println(len(tasks.Tasks))

	// cveID := "CVE-2016-7117"
	// taskID := 556578
	// serverID := 9315
	roleID := 1311
	pkgCpes, err := client.GetAllPkgCpeList(fvuls.GetPkgCpeListParam{
		Limit: 1000,
		// FilterCveID: &cveID,
		// FilterTaskID: &taskID,
		// FilterServerID: &serverID,
		FilterRoleID: &roleID,
	})
	if err != nil {
		panic(err)
	}
	// pp.Println(pkgCpes)
	for _, p := range pkgCpes {
		// if len(p.Tasks) != 0 {
		pp.Println(p)
		// }
	}
}
