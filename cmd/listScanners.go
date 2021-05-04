package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/kondukto-io/kdt/client"
	"github.com/spf13/cobra"
)

// listScannersCmd represents the listScanners command
var listScannersCmd = &cobra.Command{
	Use:   "scanners",
	Short: "lists scanners in Kondukto",
	Run:   scannersRootCommand,
}

func init() {
	listCmd.AddCommand(listProjectsCmd)
	listCmd.AddCommand(listScannersCmd)
}

func scannersRootCommand(cmd *cobra.Command, args []string) {
	c, err := client.New()
	if err != nil {
		qwe(1, err, "could not initialize Kondukto client")
	}

	var arg string
	if len(args) != 0 {
		arg = args[0]
	}

	scanners, err := c.ListScanners(arg)
	if err != nil {
		qwe(1, err, "could not retrieve scanners")
	}

	if len(scanners) < 1 {
		qwm(1, "no scanners found")
	}

	w := tabwriter.NewWriter(os.Stdout, 8, 8, 4, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "NAME\tID")
	fmt.Fprintln(w, "---\t---")
	for _, project := range scanners {
		fmt.Fprintf(w, "%s\t%s\n", project.Name, project.ID)
	}
}
