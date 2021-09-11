package resources

type Error404 struct {
	Status  int
	Message string
	Data    map[string]interface{}
}

func GetError404Resource(message string) IResource {
	data := make(map[string]interface{})
	errors := make(map[string]interface{})
	data["message"] = message
	data["errors"] = errors
	resource := Error404{Status: 404, Message: "NOT FOUND", Data: data}
	return &resource
}

func (error404 Error404) GetStatus() int {
	return error404.Status
}

func (error404 Error404) GetData() map[string]interface{} {
	return error404.Data
}
