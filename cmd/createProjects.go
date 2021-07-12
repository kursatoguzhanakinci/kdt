/*
Copyright © 2021 Kondukto

*/

package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/kondukto-io/kdt/client"
	"github.com/spf13/cobra"
)

// createProjectCmd represents the create project command
var createProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "creates a new project on Kondukto",
	Run:   createProjectsRootCommand,
}

func init() {
	createCmd.AddCommand(createProjectCmd)

	createProjectCmd.Flags().Bool("force-create", false, "ignore if the URL is used by another project")
	createProjectCmd.Flags().String("repo-id", "", "ALM project repository url")
	createProjectCmd.Flags().StringP("team", "t", "", "project team name")
	createProjectCmd.Flags().StringP("labels", "l", "", "comma separated labels")
	createProjectCmd.Flags().StringP("alm", "a", "", "ALM tool name")

}

func createProjectsRootCommand(cmd *cobra.Command, _ []string) {
	parseLabels := func(labels []client.ProjectLabel) string {
		var l string
		for i, label := range labels {
			if i == 0 {
				l = label.Name
				continue
			}
			l += fmt.Sprintf(",%s", label.Name)
		}
		return l
	}

	projectRows := []Row{
		{Columns: []string{"NAME", "ID", "TEAM", "LABELS", "UI Link"}},
		{Columns: []string{"----", "--", "----", "------", "-------"}},
	}

	c, err := client.New()
	if err != nil {
		qwe(1, err, "could not initialize Kondukto client")
	}

	repositoryID, err := cmd.Flags().GetString("repo-id")
	if err != nil {
		qwe(1, err, "failed to parse the repo url flag")
	}

	if repositoryID == "" {
		qwm(1, "missing required flag repo-id")
	}

	force, err := cmd.Flags().GetBool("force-create")
	if err != nil {
		qwm(1, fmt.Sprintf("failed to parse the force-create flag: %v", err))
	}

	if !force {
		var alm = repositoryID
		projects, err := c.ListProjects("", alm)
		if err != nil {
			qwe(1, err, "failed to check project with alm info")
		}

		if len(projects) > 0 {
			for _, project := range projects {
				projectRows = append(projectRows, Row{Columns: []string{project.Name, project.ID, project.Team.Name, parseLabels(project.Labels), project.Links.HTML}})
			}
			TableWriter(projectRows...)
			qwm(1, "project with the same repo-id already exists. for force creation pass --force-create flag")
		}
	}

	team, err := cmd.Flags().GetString("team")
	if err != nil {
		qwe(1, err, "failed to parse the team flag: %v")
	}

	labels, err := cmd.Flags().GetString("labels")
	if err != nil {
		qwe(1, err, "failed to parse the labels flag")
	}

	parsedLabels := make([]client.ProjectLabel, 0)
	for _, s := range strings.Split(labels, ",") {
		parsedLabels = append(parsedLabels, client.ProjectLabel{Name: s})
	}

	tool, err := cmd.Flags().GetString("alm")
	if err != nil {
		qwe(1, err, "failed to parse the alm flag")
	}

	pd := client.ProjectDetail{
		Source: func() client.ProjectSource {
			s := client.ProjectSource{Tool: tool}
			_, err := url.Parse(repositoryID)
			if err != nil {
				s.ID = repositoryID
			} else {
				s.URL = repositoryID
			}
			return s
		}(),
		Team: client.ProjectTeam{
			Name: team,
		},
		Labels:   parsedLabels,
		Override: force,
	}

	project, err := c.CreateProject(pd)
	if err != nil {
		qwe(1, err, "failed to create project")
	}

	projectRows = append(projectRows, Row{Columns: []string{project.Name, project.ID, project.Team.Name, parseLabels(project.Labels), project.Links.HTML}})

	TableWriter(projectRows...)
	qwm(0, "project created successfully")
}
