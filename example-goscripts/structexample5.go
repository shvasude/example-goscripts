package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type SbrResponse struct {
	APIVersion string
	Kind       string
	Metadata   struct {
		Annotations struct { }
		Name	 string
		Namespace string
		
	}
	Spec struct {
		ApplicationSelector struct {
			Group       string
			Resource    string
			ResourceRef string
			Version     string
		}
		BackingServiceSelector struct {
			Group       string
			Kind        string
			ResourceRef string
			Version     string
		}
	}
}

func main() {
		
	a := `'{"apiVersion":"apps.openshift.io/v1alpha1","kind":"ServiceBindingRequest","metadata":{"annotations":{},"name":"binding-request","namespace":"service-binding-demo"},"spec":{"applicationSelector":{"group":"apps","resource":"deployments","resourceRef":"nodejs-rest-http-crud","version":"v1"},"backingServiceSelector":{"group":"postgresql.baiju.dev","kind":"Database","resourceRef":"db-demo","version":"v1alpha1"}}}
	'`
	if strings.Contains(a,"'"){
		a=strings.Trim(a,"'")
	}
	fmt.Println(a)
	
	resp := unmarshalJsonData(a)
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------")
	fmt.Println(resp)
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------")

}


func unmarshalJsonData(a string) SbrResponse {
	res := SbrResponse{}
	json.Unmarshal([]byte(a), &res)
	return res
}
