package resources

type NoContent204 struct {
	Status  int
	Message string
	Data    map[string]interface{}
}

func GetNoContent204Resource(data interface{}) IResource {
	dataJson := map[string]interface{}{
		"data":    data,
	}
	resource := NoContent204{Status: 204, Message: "", Data: dataJson}
	return &resource
}

func (NoContent204 NoContent204) GetStatus() int {
	return NoContent204.Status
}

func (NoContent204 NoContent204) GetData() map[string]interface{} {
	return NoContent204.Data
}
