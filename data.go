package slingr

func (c *App) GetRecord(entity, id string, queryParams map[string]string) ([]byte, error) {
	return c.Get("/data/"+entity+"/"+id, queryParams)
}

func (c *App) GetRecords(entity string, queryParams map[string]string) ([]byte, error) {
	return c.Get("/data/"+entity, queryParams)
}

func (c *App) CreateRecord(entity string, payload interface{}) ([]byte, error) {
	return c.Post("/data/"+entity, payload, nil)
}

func (c *App) UpdateRecord(entity, id string, payload interface{}) ([]byte, error) {
	return c.Put("/data/"+entity+"/"+id, payload, nil)
}

func (c *App) DeleteRecord(entity, id string) ([]byte, error) {
	return c.Delete("/data/"+entity+"/"+id, nil)
}

func (c *App) DeleteRecordAsync(entity, id string) ([]byte, error) {
	return c.Delete("/data/"+entity+"/"+id, map[string]string{"_async": "true"})
}

func (c *App) DeleteRecords(entity string, queryParams map[string]string) ([]byte, error) {
	return c.Delete("/data/"+entity, queryParams)
}

func (c *App) ExecuteActionOnRecord(entity, id, action string, payload interface{}) ([]byte, error) {
	return c.Put("/data/"+entity+"/"+id+"/"+action, payload, nil)
}

func (c *App) ExecuteGlobalAction(entity, action string, payload interface{}) ([]byte, error) {
	return c.Put("/data/"+entity+"/"+action, payload, nil)
}

func (c *App) ExecuteActionOnQuery(entity, action string, payload interface{}) ([]byte, error) {
	return c.Put("/data/"+entity+"/"+action, payload, nil)
}

func (c *App) Aggregate(entity string, payload interface{}) ([]byte, error) {
	return c.Put("/data/"+entity+"/aggregate", payload, nil)
}
