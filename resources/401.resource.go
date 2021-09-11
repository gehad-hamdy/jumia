package resources

type Error401 struct {
	Status  int
	Message string
	Data    map[string]interface{}
}

func GetError401Resource(message string) IResource {
	data := make(map[string]interface{})
	errors := make(map[string]interface{})
	data["message"] = message
	data["errors"] = errors
	resource := Error401{Status: 401, Message: "Unauthorized", Data: data}
	return &resource
}

func (error401 Error401) GetStatus() int {
	return error401.Status
}

func (error401 Error401) GetData() map[string]interface{} {
	return error401.Data
}
