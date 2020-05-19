package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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
	res := SbrResponse{}
	a := `{"apiVersion":"apps.openshift.io/v1alpha1","kind":"ServiceBindingRequest","metadata":{"annotations":{},"name":"binding-request","namespace":"service-binding-demo"},"spec":{"applicationSelector":{"group":"apps","resource":"deployments","resourceRef":"nodejs-rest-http-crud","version":"v1"},"backingServiceSelector":{"group":"postgresql.baiju.dev","kind":"Database","resourceRef":"db-demo","version":"v1alpha1"}}}`
	json.Unmarshal([]byte(a), &res)

	c := SbrResponse{
		APIVersion: "apps.openshift.io/v1alpha1",
		Kind:       "ServiceBindingRequest",
	}

	c.Metadata.Name = "binding-request"
	c.Metadata.Namespace =  "service-binding-demo"
	c.Spec.ApplicationSelector.Group = "apps"
	c.Spec.ApplicationSelector.Version = "v1"
	c.Spec.ApplicationSelector.ResourceRef = "nodejs-rest-http-crud"
	c.Spec.ApplicationSelector.Resource = "deployments"
	c.Spec.BackingServiceSelector.Group = "postgresql.baiju.dev"
	c.Spec.BackingServiceSelector.Kind = "Database"
	c.Spec.BackingServiceSelector.Version = "v1alpha1"	
	c.Spec.BackingServiceSelector.ResourceRef = "db-demo"


	fmt.Println(res)
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------")
	fmt.Println(c)
	fmt.Println(reflect.DeepEqual(res, c))
}

