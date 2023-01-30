package slingr

import "encoding/json"

type Job struct {
	Id    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`
	RunBy struct {
		Id       string `json:"id"`
		FullName string `json:"fullName"`
	} `json:"runBy"`
	ParentJob             RecordReference `json:"parentJob"`
	RootJob               RecordReference `json:"rootJob"`
	CreateDate            int64           `json:"createDate"`
	StartDate             int64           `json:"startDate"`
	EndDate               int64           `json:"endDate"`
	Status                string          `json:"status"`
	HasErrors             bool            `json:"hasErrors"`
	Progress              int             `json:"progress"`
	RecordsCount          int             `json:"recordsCount"`
	RecordsProcessed      int             `json:"recordsProcessed"`
	ChildrenJobsCount     int             `json:"childrenJobsCount"`
	ChildrenJobsProcessed int             `json:"childrenJobsProcessed"`
	Stoppable             bool            `json:"stoppable"`
	LowPriority           bool            `json:"lowPriority"`
	Data                  struct {
		PushChanges  bool   `json:"pushChanges"`
		WakingUp     bool   `json:"wakingUp"`
		FileName     string `json:"fileName"`
		EntityName   string `json:"entityName"`
		ActionName   string `json:"actionName"`
		ListenerName string `json:"listenerName"`
		NotifyUsers  bool   `json:"notifyUsers"`
		EndpointName string `json:"endpointName"`
	} `json:"data"`
	Results JobGenericResults `json:"results"`
}

type ManyJob struct {
	Offset string `json:"offset"`
	Total  int    `json:"total"`
	Items  []Job  `json:"items"`
}

type JobGenericResults struct {
	RowsImported    int    `json:"rowsImported"`
	RowsWithErrors  int    `json:"rowsWithErrors"`
	FileLink        string `json:"fileLink"`
	FileId          string `json:"fileId"`
	RecordsExported int    `json:"recordsExported"`
}

type JobRecordsResults map[string]struct {
	Status       string `json:"status"`
	Response     string `json:"response"`
	ErrorMessage string `json:"errorMessage"`
}

func (c *App) GetJob(id string) (*Job, error) {
	r, err := c.Get("/status/job/"+id, nil)
	if err != nil {
		return nil, err
	}
	var job Job
	err = json.Unmarshal(r, &job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (c *App) GetJobLogs(id string) (*ManyLog, error) {
	r, err := c.Get("/status/job/"+id+"/logs", nil)
	if err != nil {
		return nil, err
	}
	var manyLog ManyLog
	err = json.Unmarshal(r, &manyLog)
	if err != nil {
		return nil, err
	}
	return &manyLog, nil
}

func (c *App) GetJobs(queryParams map[string]string) (*Job, error) {
	r, err := c.Get("/status/jobs", queryParams)
	if err != nil {
		return nil, err
	}
	var job Job
	err = json.Unmarshal(r, &job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (c *App) StopJob(id string) (*Job, error) {
	r, err := c.Put("/status/jobs/"+id+"/stop", nil, nil)
	if err != nil {
		return nil, err
	}
	var job Job
	err = json.Unmarshal(r, &job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (c *App) ForceStopJob(id string) (*Job, error) {
	r, err := c.Put("/status/jobs/"+id+"/forceStop", nil, nil)
	if err != nil {
		return nil, err
	}
	var job Job
	err = json.Unmarshal(r, &job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (c *App) CancelJob(id string) (*Job, error) {
	r, err := c.Put("/status/jobs/"+id+"/cancel", nil, nil)
	if err != nil {
		return nil, err
	}
	var job Job
	err = json.Unmarshal(r, &job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (c *App) ResumeJob(id string) (*Job, error) {
	r, err := c.Put("/status/jobs/"+id+"/resume", nil, nil)
	if err != nil {
		return nil, err
	}
	var job Job
	err = json.Unmarshal(r, &job)
	if err != nil {
		return nil, err
	}
	return &job, nil
}
