# Simple Slingr Go Client

Non-Official Slingr go module with helpers to communicate with any
Slingr App.

## Installation

```
go get github.com/espinosajuanma/slingr-go@latest
```

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/espinosajuanma/slingr-go"
)

func main() {
	appName := os.Getenv("SLINGR_APP_NAME")    // Your slingr app name
	environment := os.Getenv("SLINGR_APP_ENV") // Environment to login (dev, staging, prod)
	email := os.Getenv("SLINGR_APP_EMAIL")     // Your slingr email

	app := slingr.NewApp(appName, environment)

	// Login
	_, err := app.Login(email, os.Getenv("SLINGR_APP_PASSWORD"))
	if err != nil {
		panic(err)
	}
	// Use app.Token to cache the token
	fmt.Printf("connected to [%s-%s] with token [%s]\n", app.Name, app.Env, app.Token)

	// Get a record
	r, err := app.GetRecord("testEntity", "63d3bbef6f6e145fd3d363af", nil)
	if err != nil {
		panic(err)
	}
	var testRecord TestEntity
	err = json.Unmarshal(r, &testRecord)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Record ID: %s\n", testRecord.Id)
	fmt.Printf("Entity ID: %s\n", testRecord.Entity.Id)

	// You can logout to expire the token right away
	_, err = app.Logout()
	if err != nil {
		panic(err)
	}
	fmt.Println("Logged out")
}

type TestEntity struct {
	Id      string
	Version int
	Entity  s.EntityReference
}
```

