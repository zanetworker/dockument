package main

import (
	"io"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/zanetworker/go-kubesanity/pkg/log"
	"github.com/zanetworker/go-kubesanity/pkg/network"
)

var networkCmdDesc = `
This is the command used to validate kubernetes network configuration (pods, services, etc)`

type networkCmdParams struct {
	duplicatePodIP, duplicateServiceIP bool
}

func newNetworkCmd(out io.Writer) *cobra.Command {
	networkCmdParams := &networkCmdParams{}

	networkCmd := &cobra.Command{
		Use:   "network",
		Short: "validate network configuration parameters",
		Long:  globalUsage,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Info("number of flags " + strconv.Itoa(len(args)))

			return networkCmdParams.run()
		},
	}

	f := networkCmd.Flags()
	networkCmd.GenBashCompletionFile("out.sh")
	f.BoolVar(&networkCmdParams.duplicatePodIP, "checkDuplicatePodIP", false, "if set to true, kubesanity will check for duplicate Pod IPs in all namespaces")
	f.BoolVar(&networkCmdParams.duplicateServiceIP, "checkDuplicateServiceIP", false, "if set to true, kubesanity will check for duplicate Service IPs in all namespaces")

	return networkCmd
}

func (n *networkCmdParams) run() error {

	kubernetesClient := network.NewKubeClient()
	if n.duplicatePodIP {
		hasDuplicatePodIPs, err := kubernetesClient.CheckDuplicatePodIP()
		if hasDuplicatePodIPs {
			log.Error("Discovered Duplicate Pod IPs in you Kubernetes Deployment")
		}
		return err
	}

	if n.duplicateServiceIP {
		hasDuplicateServiceIPs, err := kubernetesClient.CheckDuplicateServiceIP()
		if hasDuplicateServiceIPs {
			log.Error("Discovered Duplicate Service IPs in you Kubernetes Deployment")
		}
		return err
	}
	return nil
}
