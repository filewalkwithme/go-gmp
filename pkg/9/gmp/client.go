package gmp

type Client interface {
	Authenticate(cmd *AuthenticateCommand) (resp *AuthenticateResponse, err error)
	GetConfigs(cmd *GetConfigsCommand) (resp *GetConfigsResponse, err error)
	GetScanners(cmd *GetScannersCommand) (resp *GetScannersResponse, err error)
	GetPreferences(cmd *GetPreferencesCommand) (resp *GetPreferencesResponse, err error)
	CreateConfig(cmd *CreateConfigCommand) (resp *CreateConfigResponse, err error)
	ModifyConfig(cmd *ModifyConfigCommand) (resp *ModifyConfigResponse, err error)
	CreateTask(cmd *CreateTaskCommand) (resp *CreateTaskResponse, err error)
	CreateTarget(cmd *CreateTargetCommand) (resp *CreateTargetResponse, err error)
	StartTask(cmd *StartTaskCommand) (resp *StartTaskResponse, err error)
	GetTasks(cmd *GetTasksCommand) (resp *GetTasksResponse, err error)
	StopTask(cmd *StopTaskCommand) (resp *StopTaskResponse, err error)
	DeleteTask(cmd *DeleteTaskCommand) (resp *DeleteTaskResponse, err error)
	GetResults(cmd *GetResultsCommand) (resp *GetResultsResponse, err error)
}
