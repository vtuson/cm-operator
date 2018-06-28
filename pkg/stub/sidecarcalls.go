package stub

import (
	"fmt"
	"os"
)

func getServiceURL() string {
	domain := os.Getenv("CMSERVICE")
	port := os.Getenv("CMPORT")
	return fmt.Sprintf("http://%s:%s/", domain, port)
}

func addRepo(name string, url string) bool {
	fmt.Println("adding repo:", name, url)
	p := make(map[string]string)
	p["name"] = name
	p["git"] = url
	call := RequestParms{
		Endpoint: getServiceURL() + "repo/new",
		Method:   HTTP_POST,
		Params:   p,
	}
	resp, err := Curl(call)
	if err != nil {
		fmt.Println(err)
		return false
	}
	code, _ := HttpStatus(resp)
	if code > 299 {
		fmt.Println(err)
		return false
	}
	return true
}

func deleteRepo(name string) bool {
	p := make(map[string]string)
	p["name"] = name
	call := RequestParms{
		Endpoint: getServiceURL() + "repo/" + name,
		Method:   HTTP_DELETE,
		Params:   p,
	}
	resp, err := Curl(call)
	if err != nil {
		fmt.Println(err)
		return false
	}
	code, _ := HttpStatus(resp)
	if code > 299 {
		fmt.Println(err)
		return false
	}
	return true
}

func updateRepo(name string) bool {
	p := make(map[string]string)
	p["name"] = name
	call := RequestParms{
		Endpoint: getServiceURL() + "repo/" + name + "/update",
		Method:   HTTP_GET,
		Params:   p,
	}
	resp, err := Curl(call)
	if err != nil {
		fmt.Println(err)
		return false
	}
	code, _ := HttpStatus(resp)
	if code > 299 {
		fmt.Println(err)
		return false
	}
	return true
}

func addDependency(name string, url string) bool {
	fmt.Println("Adding dependency:", name, url)
	p := make(map[string]string)
	p["name"] = name
	p["addr"] = url
	call := RequestParms{
		Endpoint: getServiceURL() + "repo/dependency",
		Method:   HTTP_POST,
		Params:   p,
	}
	resp, err := Curl(call)
	if err != nil {
		fmt.Println(err)
		return false
	}
	code, _ := HttpStatus(resp)
	if code > 299 {
		fmt.Println(err)
		return false
	}
	return true
}
