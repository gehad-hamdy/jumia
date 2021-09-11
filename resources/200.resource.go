package resources

type Success200 struct {
	Status  int
	Message string
	Data    map[string]interface{}
}

func GetSuccess200Resource(data interface{}, message string) IResource {
	dataJson := map[string]interface{}{
		"data":    data,
		"message": message,
	}
	resource := Success200{Status: 200, Message: "", Data: dataJson}
	return &resource
}

func (success200 Success200) GetStatus() int {
	return success200.Status
}

func (success200 Success200) GetData() map[string]interface{} {
	return success200.Data
}
