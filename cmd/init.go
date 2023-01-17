package cmd

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type options struct {
	Name string
}

// initCmd creates a new context and sets the current context to it.
// When no name is provided a random name is generated for it.
// This function is idempotent, so creating a context with existing names has no impact
// and no error is returned.
func initCmd(out io.Writer) *cobra.Command {
	opts := &options{}

	cmd := &cobra.Command{
		Use:   "init",
		Short: "create a new context",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Name == "" {
				opts.Name = fmt.Sprintf("init-%d", time.Now().Unix())
			}

			userhome, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			contextDIR := fmt.Sprintf("%s/.asd/%s", userhome, opts.Name)
			err = os.MkdirAll(contextDIR, os.ModePerm)
			if err != nil {
				return err
			}
			//TODO: save current context
			// add it to existing context list
			_, _ = fmt.Fprintf(out, "context created in: %s\n", contextDIR)
			return nil
		},
	}

	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "context name")

	return cmd
}
