package main

import (
	"fmt"
	"testing"
	"os"
	"github.com/stretchr/testify/require"
	"regexp"
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
	cmdString := `/bin/bash -c '. ../../hack/examples-commons.sh && $1' EC uninstall_service_binding_operator_subscription_master
subscription.operators.coreos.com "service-binding-operator" deleted
clusterserviceversion.operators.coreos.com "service-binding-operator.v0.1.1-294" deleted
/bin/bash -c '. ../../hack/examples-commons.sh && $1' EC uninstall_service_binding_operator_source_master
operatorsource.operators.coreos.com "redhat-developer-operators" deleted`

	re := regexp.MustCompile(`.*subscription.operators.coreos.com\s.service-binding-operator.\sdeleted`)
	re1 := regexp.MustCompile(`.*operatorsource.operators.coreos.com\s.redhat-developer-operators.\sdeleted`)
	if re.MatchString(cmdString) && re1.MatchString(cmdString) {
		fmt.Printf("re is %s \n and re1 is %s \n",re,re1)
	}
	
	//require.Containsf(t, cmdString, []*regexp.Regexp{re, re1},  "")
	
}
