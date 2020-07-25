package client

import (
	"testing"

	"github.com/filewalkwithme/go-gmp/pkg/9/gmp"
)

type mockConn struct{}

func (m *mockConn) Execute(command interface{}, response interface{}) error {
	if cmd, ok := command.(*gmp.AuthenticateCommand); ok {
		if cmd.Credentials.Username == "openvas" && cmd.Credentials.Password == "123" {
			(*response.(*gmp.AuthenticateResponse)).Status = "200"
		} else {
			(*response.(*gmp.AuthenticateResponse)).Status = "400"
		}
	}

	if cmd, ok := command.(*gmp.GetConfigsCommand); ok {
		if cmd.ConfigID == "bde773f3-2b3d-4fe6-81cb-6321ae2cc629" {
			(*response.(*gmp.GetConfigsResponse)).Status = "200"
		} else {
			(*response.(*gmp.GetConfigsResponse)).Status = "400"
		}
	}

	if cmd, ok := command.(*gmp.GetScannersCommand); ok {
		if cmd.ScannerID == "ee0311e7-3247-4425-bb9c-866d59f1e0e9" {
			(*response.(*gmp.GetScannersResponse)).Status = "200"
		} else {
			(*response.(*gmp.GetScannersResponse)).Status = "400"
		}
	}

	if cmd, ok := command.(*gmp.GetPreferencesCommand); ok {
		if cmd.ConfigID == "4b49617e-d1d8-44b8-af81-f4675b56f837" {
			(*response.(*gmp.GetPreferencesResponse)).Status = "200"
		} else {
			(*response.(*gmp.GetPreferencesResponse)).Status = "400"
		}
	}

	if cmd, ok := command.(*gmp.CreateConfigCommand); ok {
		if cmd.Name == "New Config" {
			(*response.(*gmp.CreateConfigResponse)).Status = "200"
		} else {
			(*response.(*gmp.CreateConfigResponse)).Status = "400"
		}
	}

	if cmd, ok := command.(*gmp.ModifyConfigCommand); ok {
		if cmd.Name == "Modified Config" {
			(*response.(*gmp.ModifyConfigResponse)).Status = "200"
		} else {
			(*response.(*gmp.ModifyConfigResponse)).Status = "400"
		}
	}

	if cmd, ok := command.(*gmp.CreateTaskCommand); ok {
		if cmd.Name == "New Task" {
			(*response.(*gmp.CreateTaskResponse)).Status = "200"
		} else {
			(*response.(*gmp.CreateTaskResponse)).Status = "400"
		}
	}

	if cmd, ok := command.(*gmp.CreateTargetCommand); ok {
		if cmd.Name == "New Target" {
			(*response.(*gmp.CreateTargetResponse)).Status = "200"
		} else {
			(*response.(*gmp.CreateTargetResponse)).Status = "400"
		}
	}

	if cmd, ok := command.(*gmp.StartTaskCommand); ok {
		if cmd.TaskID == "e512e2ca-9d0e-4bf3-bc73-7fbe6e9bbf31" {
			(*response.(*gmp.StartTaskResponse)).Status = "200"
		} else {
			(*response.(*gmp.StartTaskResponse)).Status = "400"
		}
	}

	if cmd, ok := command.(*gmp.GetTasksCommand); ok {
		if cmd.TaskID == "e512e2ca-9d0e-4bf3-bc73-7fbe6e9bbf31" {
			(*response.(*gmp.GetTasksResponse)).Status = "200"
		} else {
			(*response.(*gmp.GetTasksResponse)).Status = "400"
		}
	}

	if cmd, ok := command.(*gmp.StopTaskCommand); ok {
		if cmd.TaskID == "e512e2ca-9d0e-4bf3-bc73-7fbe6e9bbf31" {
			(*response.(*gmp.StopTaskResponse)).Status = "200"
		} else {
			(*response.(*gmp.StopTaskResponse)).Status = "400"
		}
	}

	if cmd, ok := command.(*gmp.DeleteTaskCommand); ok {
		if cmd.TaskID == "e512e2ca-9d0e-4bf3-bc73-7fbe6e9bbf31" {
			(*response.(*gmp.DeleteTaskResponse)).Status = "200"
		} else {
			(*response.(*gmp.DeleteTaskResponse)).Status = "400"
		}
	}

	if cmd, ok := command.(*gmp.GetResultsCommand); ok {
		if cmd.TaskID == "e512e2ca-9d0e-4bf3-bc73-7fbe6e9bbf31" {
			(*response.(*gmp.GetResultsResponse)).Status = "200"
		} else {
			(*response.(*gmp.GetResultsResponse)).Status = "400"
		}
	}

	return nil
}

func (m *mockConn) Close() error {
	return nil
}

func mockedConnection() gmp.Connection {
	return &mockConn{}
}

func TestNew(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}
}

func TestAuthenticate(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.AuthenticateCommand{}
	cmd.Credentials.Username = "openvas"
	cmd.Credentials.Password = "123"
	resp, err := cli.Authenticate(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during Authenticate: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}

func TestGetConfigs(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.GetConfigsCommand{}
	cmd.ConfigID = "bde773f3-2b3d-4fe6-81cb-6321ae2cc629"
	resp, err := cli.GetConfigs(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during GetConfigs: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}

func TestGetScanners(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.GetScannersCommand{}
	cmd.ScannerID = "ee0311e7-3247-4425-bb9c-866d59f1e0e9"
	resp, err := cli.GetScanners(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during GetScanners: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}

func TestGetPreferences(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.GetPreferencesCommand{}
	cmd.ConfigID = "4b49617e-d1d8-44b8-af81-f4675b56f837"
	resp, err := cli.GetPreferences(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during GetPreferences: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}

func TestCreateConfig(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.CreateConfigCommand{}
	cmd.Name = "New Config"
	resp, err := cli.CreateConfig(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during CreateConfig: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}

func TestModifyConfig(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.ModifyConfigCommand{}
	cmd.Name = "Modified Config"
	resp, err := cli.ModifyConfig(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during ModifyConfig: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}

func TestCreateTask(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.CreateTaskCommand{}
	cmd.Name = "New Task"
	resp, err := cli.CreateTask(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during CreateTask: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}

func TestCreateTarget(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.CreateTargetCommand{}
	cmd.Name = "New Target"
	resp, err := cli.CreateTarget(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during CreateTarget: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}

func TestStartTask(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.StartTaskCommand{}
	cmd.TaskID = "e512e2ca-9d0e-4bf3-bc73-7fbe6e9bbf31"
	resp, err := cli.StartTask(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during StartTask: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}

func TestGetTasks(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.GetTasksCommand{}
	cmd.TaskID = "e512e2ca-9d0e-4bf3-bc73-7fbe6e9bbf31"
	resp, err := cli.GetTasks(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during GetTasks: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}

func TestStopTask(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.StopTaskCommand{}
	cmd.TaskID = "e512e2ca-9d0e-4bf3-bc73-7fbe6e9bbf31"
	resp, err := cli.StopTask(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during StopTask: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}

func TestDeleteTask(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.DeleteTaskCommand{}
	cmd.TaskID = "e512e2ca-9d0e-4bf3-bc73-7fbe6e9bbf31"
	resp, err := cli.DeleteTask(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during DeleteTask: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}

func TestGetResults(t *testing.T) {
	cli := New(mockedConnection())
	if cli == nil {
		t.Fatalf("Client is nil")
	}

	cmd := &gmp.GetResultsCommand{}
	cmd.TaskID = "e512e2ca-9d0e-4bf3-bc73-7fbe6e9bbf31"
	resp, err := cli.GetResults(cmd)
	if err != nil {
		t.Fatalf("Unexpected error during GetResults: %s", err)
	}

	if resp.Status != "200" {
		t.Fatalf("Unexpected status. \nExpected: 200 \nGot: %s", resp.Status)
	}
}
