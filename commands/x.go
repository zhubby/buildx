package commands

import (
	"github.com/docker/buildx/util/cobrautil/completion"
	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"
)

type xOptions struct {

}


func runX() {

}

func xCmd(dockerCli command.Cli,opt *xOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "x",
		Short: "打包镜像一条龙",
		Args:  cli.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVersion(dockerCli)
		},
		ValidArgsFunction: completion.Disable,
	}
	
	return cmd
}
