package main

import (
	"fmt"
	"strings"
	"testing"
	"os"

)


func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestThismain(t *testing.T) {
	fmt.Println("Hello, playground")
	t.Run("get-Project", GetResult)
}

func GetResult(t *testing.T){
	appNS       := "service-binding-demo"
	res         := "Shobith is created"
	result, checkFlag := getProjectCreationResult(t,res,appNS)
	
	fmt.Printf("########this is the result %s \n",result)
	fmt.Printf("########this is the checkFlag %v",checkFlag)
}

func getProjectCreationResult(t *testing.T, projectRes string, project string) (string, bool) {
	
	lstArr := strings.Split(projectRes, "\n")
	for _, item := range lstArr {
		if strings.Contains(item, project) {
			t.Logf("item matched as %s \n", item)
			return item, true
		}
	}
	return item, false
}
