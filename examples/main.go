package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/filewalkwithme/go-gmp/pkg/9/gmp"
	"github.com/filewalkwithme/go-gmp/pkg/9/gmp/client"
	"github.com/filewalkwithme/go-gmp/pkg/9/gmp/connections"
)

// https://docs.greenbone.net/API/GMP/gmp-9.0.html
func main() {
	// Connect to GVMD
	fmt.Println("Connecting to GVMD...")
	conn, err := connections.NewUnixConnection("/tmp/openvas-socks/gvmd.sock")
	if err != nil {
		panic(err)
	}
	fmt.Println("-> ok")
	defer conn.Close()

	// Instantiate  a new GMP Client
	fmt.Println("Instantiate  a new GMP Client")
	gmpClient := client.New(conn)
	fmt.Println("-> ok")

	// Authenticate
	fmt.Println("Authenticating...")
	auth := &gmp.AuthenticateCommand{}
	auth.Credentials.Username = "openvas"
	auth.Credentials.Password = "openvas"
	_, err = gmpClient.Authenticate(auth)
	if err != nil {
		panic(err)
	}
	fmt.Println("-> ok")

	// Get the Default scanner
	fmt.Println("Getting the Default scanner...")
	s := &gmp.GetScannersCommand{}
	s.Filter = `name="OpenVAS Default"`
	getScannersResp, err := gmpClient.GetScanners(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("-> ok (scanner id: %s)\n", getScannersResp.Scanner[0].ID)

	// Get the configuration named "Full and fast"
	fmt.Println("Getting the configuration named \"Full and fast\"...")
	c := &gmp.GetConfigsCommand{}
	c.Filter = `name="Full and fast"`
	configResp, err := gmpClient.GetConfigs(c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("-> ok (config id: %s)\n", configResp.Config[0].ID)

	// Create a target
	fmt.Println("Creating a new target...")
	ct := &gmp.CreateTargetCommand{}
	ct.Name = "localhost"
	ct.Hosts = "127.0.0.1"
	createTargetResp, err := gmpClient.CreateTarget(ct)
	if err != nil {
		panic(err)
	}
	fmt.Printf("-> ok (target id: %s)\n", createTargetResp.ID)

	// Create a new task
	fmt.Println("Creating a new task...")
	newTask := &gmp.CreateTaskCommand{}
	newTask.Name = "New Task"
	newTask.Config = new(gmp.CreateTaskConfig)
	newTask.Target = new(gmp.CreateTaskTarget)
	newTask.Scanner = new(gmp.CreateTaskScanner)
	newTask.Config.ID = configResp.Config[0].ID
	newTask.Target.ID = createTargetResp.ID
	newTask.Scanner.ID = getScannersResp.Scanner[0].ID
	newTaskResp, err := gmpClient.CreateTask(newTask)
	if err != nil {
		panic(err)
	}
	fmt.Printf("-> ok (task id: %s)\n", newTaskResp.ID)

	// Start the task
	fmt.Println("Start the task")
	st := &gmp.StartTaskCommand{}
	st.TaskID = newTaskResp.ID
	_, err = gmpClient.StartTask(st)
	if err != nil {
		panic(err)
	}

	// Monitoring task progress
	fmt.Println("Monitoring task progress...")
	for {
		gt := &gmp.GetTasksCommand{}
		gt.TaskID = newTaskResp.ID
		getTasksResp, err := gmpClient.GetTasks(gt)
		if err != nil {
			panic(err)
		}
		time.Sleep(10 * time.Second)
		fmt.Printf("Monitoring task progress: %s%%\n", getTasksResp.Task[0].Progress.Value)

		if x, _ := strconv.Atoi(getTasksResp.Task[0].Progress.Value); x >= 100 {
			break
		}
	}

	// Get results
	getResults := &gmp.GetResultsCommand{}
	getResults.TaskID = newTaskResp.ID
	getResults.Filter = `min_qod=0 rows=1000`
	results, err := gmpClient.GetResults(getResults)
	if err != nil {
		panic(err)
	}

	// Show results
	for i := 0; i < len(results.Result); i++ {
		fmt.Printf("Result[%d]: %s (score: %s)\n", i, results.Result[i].Name, results.Result[i].Severity)
	}
}
