package cmd

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/template"

	"github.com/cli/go-gh/v2"
	"github.com/spf13/cobra"

	"github.com/cli/cli/v2/pkg/iostreams"
)

type MultirepoOptions struct {
	IO      *iostreams.IOStreams
	Header  string
	Repos   []string
	Command []string
}

func NewCmdMultirepo() *cobra.Command {
	opts := &MultirepoOptions{
		IO: iostreams.System(),
	}

	cmd := &cobra.Command{
		Use:   "multirepo [flags] <repos> -- <command>",
		Short: "Run a Github CLI command on multiple repositories",
		Long: `Run a Github CLI command on multiple repositories.

Arguments:
  repos    One or more [HOST/]OWNER/REPO specifications, separated by commas
  command  A Github CLI command to run for each IFS repository
`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				cmd.SilenceUsage = false
				return fmt.Errorf("invalid arguments: provide a comma-separated list of repositories and a command to execute")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			noHeader, _ := cmd.Flags().GetBool("no-header")

			if !noHeader {
				headerTemplate, _ := cmd.Flags().GetString("header-template")
				opts.Header = headerTemplate
			}

			reposArg, commandArgs := args[0], args[1:]
			repos := strings.Split(reposArg, ",")
			opts.Repos = repos
			opts.Command = commandArgs

			return runMultirepo(opts)
		},
	}

	cmd.Flags().Bool("no-header", false, "Disable printing the repository header")
	cmd.Flags().String("header-template", "{{.Repo}}", "Go template used to render the repository header")
	cmd.MarkFlagsMutuallyExclusive("no-header", "header-template")

	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	return cmd
}

func Execute() error {
	return NewCmdMultirepo().Execute()
}

func runMultirepo(opts *MultirepoOptions) error {
	var tmpl *template.Template
	if opts.Header != "" {
		t, err := template.New("--header-template").Parse(opts.Header)
		if err != nil {
			return err
		}
		tmpl = t
	}

	ctx := context.Background()

	for _, repo := range opts.Repos {
		printHeader(opts.IO, tmpl, repo)

		commandArgs := append(opts.Command, "-R", repo)

		// We choose ExecInteractive so that we are just wrapping the chosen
		// command, with this program's stdout and stderr connected to the
		// command being run. This allows use of commands that need stdin,
		// and passes through any output formatting.
		err := gh.ExecInteractive(ctx, commandArgs...)
		if err != nil {
			return err
		}

		if tmpl != nil {
			fmt.Fprintln(opts.IO.Out, "")
		}
	}

	return nil
}

func printHeader(ios *iostreams.IOStreams, tmpl *template.Template, repo string) error {
	if tmpl != nil {
		header, err := renderTemplate(tmpl, repo)
		if err != nil {
			return err
		}
		fmt.Fprintln(ios.Out, RepoHeader(ios.ColorScheme(), header))
	}
	return nil
}

func renderTemplate(tmpl *template.Template, repo string) (string, error) {
	var buf bytes.Buffer

	err := tmpl.Execute(&buf, struct{ Repo string }{Repo: repo})
	if err != nil {
		return "", nil
	}

	return buf.String(), nil
}
