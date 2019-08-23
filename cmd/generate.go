/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/ZhuPeng/weekly-report/pkg/github"
	v3 "github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

type Config struct {
	owner    string
	repo     string
	token    string
	template string
}

func (c Config) generateMeta() github.Meta {
	gc := github.NewClientWithToken(config.token)
	meta, err := gc.GetMeta(c.owner, c.repo)
	if err != nil {
		log.Println("generateMeta:", err)
	}
	return meta
}

type weeklyReportTemplate struct {
	Owner            string
	Repo             string
	WatcherCount     int
	StarCount        int
	StarDelta        string
	ForksCount       int
	ForkDelta        string
	Contributors     int
	ContributorDelta string
	MergedPrCount    int
	PRs              []*v3.PullRequest
}

var config = Config{}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		t, err := template.ParseFiles(config.template)
		if err != nil {
			log.Fatal(err)
		}
		gc := github.NewClientWithToken(config.token)

		meta := config.generateMeta()

		prs, _ := gc.GetPR(config.owner, config.repo, "closed")
		lastweek := time.Now().Add(-1 * time.Hour * 24 * 7)
		filterPRs := []*v3.PullRequest{}
		for _, p := range prs {
			// log.Println("pr:", p.MergedAt, *p.State, p.Merged)
			if p.MergedAt != nil && (*p.MergedAt).After(lastweek) {
				filterPRs = append(filterPRs, p)
			}
		}
		data := weeklyReportTemplate{
			Owner:         config.owner,
			Repo:          config.repo,
			ForksCount:    meta.ForkCount,
			StarCount:     meta.Stargazers.TotalCount,
			WatcherCount:  meta.Watchers.TotalCount,
			MergedPrCount: len(prs),
			PRs:           filterPRs,
		}
		err = t.Execute(os.Stdout, data)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")
	generateCmd.PersistentFlags().StringVarP(&config.token, "token", "", "xxyy", "Token for GitHub auth")
	generateCmd.PersistentFlags().StringVarP(&config.owner, "owner", "o", "ZhuPeng", "Owner of an GitHub repo")
	generateCmd.PersistentFlags().StringVarP(&config.repo, "repo", "r", "weekly-report", "Repo name of an GitHub repo")
	generateCmd.PersistentFlags().StringVarP(&config.template, "template", "t", "template/en.md", "Template of weekly repot of an GitHub repo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
