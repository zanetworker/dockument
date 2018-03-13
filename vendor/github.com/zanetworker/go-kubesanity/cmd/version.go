package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/zanetworker/go-kubesanity/pkg/log"
	"github.com/zanetworker/go-kubesanity/pkg/version"
)

const versionDesc = `This command prints out the version of Kubesanity`

type versionCmdOpts struct {
	out   io.Writer
	short bool
}

func newVersionCmd(out io.Writer) *cobra.Command {

	versionCmdOpts := &versionCmdOpts{out: out}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "get version",
		Long:  versionDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			return versionCmdOpts.run(cmd)
		},
	}
	f := versionCmd.Flags()

	f.BoolVarP(&versionCmdOpts.short, "short", "s", false, "print only the version number and prefix of latest commit")

	return versionCmd
}

func (v *versionCmdOpts) run(cmd *cobra.Command) error {
	vOpts := version.Options{}
	version, err := vOpts.BuildVersion(v.short)
	if err != nil {
		return err
	}
	if _, err := fmt.Fprintf(v.out, "Kubesanity Version: %s", version); err != nil {
		log.Error("Failed to create version string")
	}
	return nil
}
