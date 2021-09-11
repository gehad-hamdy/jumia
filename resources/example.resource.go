package resources

type ExampleResource struct {
	Status  int
	Message string
	Data map[string] interface{}
}
func GetExampleResource(data string) IResource {
	resource := ExampleResource{Status: 200, Message: data, Data: make(map[string] interface{})}
	resource.Data["message"] = data
	return &resource
}

func (exampleResource ExampleResource)  GetStatus() int{
	return exampleResource.Status
}

func (exampleResource ExampleResource)  GetData() map[string] interface{} {
	return exampleResource.Data
}
