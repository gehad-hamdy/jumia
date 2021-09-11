package resources

type Error500 struct {
	Status  int
	Message string
	Data    map[string]interface{}
}

func GetError500Resource(message string) IResource {
	data := make(map[string]interface{})
	errors := make(map[string]interface{})
	data["message"] = message
	data["errors"] = errors
	resource := Error500{Status: 500, Message: "Internal server error", Data: data}
	return &resource
}

func (error500 Error500) GetStatus() int {
	return error500.Status
}

func (error500 Error500) GetData() map[string]interface{} {
	return error500.Data
}
