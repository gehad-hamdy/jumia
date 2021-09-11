package resources

type Error400 struct {
	Status  int
	Message string
	Data    map[string]interface{}
}

func GetError400Resource(message string) IResource {
	data := make(map[string]interface{})
	errors := make(map[string]interface{})
	data["message"] = message
	data["errors"] = errors
	resource := Error400{Status: 400, Message: "Bad Request", Data: data}
	return &resource
}

func GetError400DetailsResource(message string, errors map[string]interface{}) IResource {
	data := make(map[string]interface{})
	data["message"] = message
	data["errors"] = errors
	resource := Error400{Status: 400, Message: "Bad Request", Data: data}
	return &resource
}

func (error400 Error400) GetStatus() int {
	return error400.Status
}

func (error400 Error400) GetData() map[string]interface{} {
	return error400.Data
}
