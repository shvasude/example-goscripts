package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	body := `{
  "apiVersion":"apps.openshift.io/v1alpha1",
  "kind":"ServiceBindingRequest",
  "metadata":{
    "annotations":{},
    "name":"binding-request",
    "namespace":"service-binding-demo"
    },
    "spec":{
      "applicationSelector":{
        "group":"apps",
        "resource":"deployments",
        "resourceRef":"nodejs-rest-http-crud",
        "version":"v1"
        },
      "backingServiceSelector":{
        "group":"postgresql.baiju.dev",
        "kind":"Database",
        "resourceRef":"db-demo",
        "version":"v1alpha1"
        }
    }
}`

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
}

