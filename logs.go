package slingr

import (
	"encoding/json"
)

type Log struct {
	Id             string      `json:"id"`
	Type           string      `json:"type"`
	SubType        string      `json:"subType"`
	Level          string      `json:"level"`
	Message        string      `json:"message"`
	UserEmail      string      `json:"userEmail"`
	AdminUserEmail string      `json:"adminUserEmail"`
	Timestamp      int64       `json:"timestamp"`
	IP             string      `json:"ip"`
	AdditionalInfo interface{} `json:"additionalInfo"`
}

type ManyLog struct {
	Total int   `json:"total"`
	Items []Log `json:"items"`
}

type LogPayload struct {
	Level          string      `json:"level"`
	Message        string      `json:"message"`
	AdditionalInfo interface{} `json:"additionalInfo"`
}

func (c *App) GetLogs(queryParams map[string]string) (*ManyLog, error) {
	r, err := c.Get("/status/logs", queryParams)
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

func (c *App) log(level, message string, data interface{}) (*Log, error) {
	payload := &LogPayload{
		Level:          level,
		Message:        message,
		AdditionalInfo: data,
	}
	r, err := c.Post("/status/logs", payload, nil)
	var log Log
	err = json.Unmarshal(r, &log)
	if err != nil {
		return nil, err
	}
	return &log, nil
}

func (c *App) LogInfo(message string, data interface{}) (*Log, error) {
	return c.log("INFO", message, data)
}

func (c *App) LogError(message string, data interface{}) (*Log, error) {
	return c.log("ERROR", message, data)
}

func (c *App) LogWarn(message string, data interface{}) (*Log, error) {
	return c.log("WARN", message, data)
}
