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
		false,
	)
	if err != nil {
		panic(err)
	}
	task, err := client.GetTaskDetail(fvuls.GetTaskDetailParam{
		TaskID: 556574,
	})
	if err != nil {
		panic(err)
	}
	pp.Println(task)

	pri := "high"
	t, err := client.UpdateTask(fvuls.UpdateTaskParam{
		TaskID:   556574,
		Priority: &pri,
	})
	if err != nil {
		panic(err)
	}
	pp.Println(t)

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
}
