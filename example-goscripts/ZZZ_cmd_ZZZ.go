package main 

import (
	"testing"
	//"os"
	//"fmt"
	"gotest.tools/v3/icmd"
	)

//var t = &testing.T{}

func TestExampleRunCmd(t *testing.T) {
	result := icmd.RunCmd(icmd.Command("cat", "/does/not/exist"))
	result.Assert(t, icmd.Expected{
		ExitCode: 1,
		Err:      "cat: /does/not/exist: No such file or directory",
	})

func TestExampleRunCommand(t *testing.T) {
	result := icmd.RunCommand("bash", "-c", "echo all good")
	result.Assert(t, icmd.Success)

func TestDisplayKUBECONFIG(t *testing.T) {
	var defaultKubeconfig string
	if os.Getenv("KUBECONFIG") != "" {
		defaultKubeconfig = os.Getenv("KUBECONFIG")
		fmt.Println("env path for KUBECONFIG")
	} else if usr, err := user.Current(); err == nil {
		defaultKubeconfig = path.Join(usr.HomeDir, ".kube/config")
	}	
}