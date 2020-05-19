package main

import (

	"fmt"
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
	c := getSbrResponse()
	
	c.Metadata.Name = "binding-request"
	c.Metadata.Namespace = "service-binding-demo"
	c.Spec.ApplicationSelector.ResourceRef = "nodejs-rest-http-crud"
	c.Spec.ApplicationSelector.Resource = "deployments"
	c.Spec.BackingServiceSelector.Group = "postgresql.baiju.dev"
	c.Spec.BackingServiceSelector.Kind = "Database"	
	c.Spec.BackingServiceSelector.ResourceRef = "db-demo"

	fmt.Println("----------------------------------------------------------------------------------------------------------------------------")
	fmt.Println(c)
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------")

}


func getSbrResponse() SbrResponse {
	res := SbrResponse{
		APIVersion: "apps.openshift.io/v1alpha1",
		Kind:       "ServiceBindingRequest",
	}
	res.Spec.ApplicationSelector.Group = "apps"
	res.Spec.ApplicationSelector.Version = "v1"
	res.Spec.BackingServiceSelector.Version = "v1alpha1"	
	return res
}
