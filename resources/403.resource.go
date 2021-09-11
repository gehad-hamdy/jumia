package resources

type Error403 struct {
	Status  int
	Message string
	Data    map[string]interface{}
}

func GetError403Resource(message string) IResource {
	data := make(map[string]interface{})
	errors := make(map[string]interface{})
	data["message"] = message
	data["errors"] = errors
	resource := Error403{Status: 403, Message: "Forbidden", Data: data}
	return &resource
}

func (error403 Error403) GetStatus() int {
	return error403.Status
}

func (error403 Error403) GetData() map[string]interface{} {
	return error403.Data
}
