# Weekly Report of {{.Owner}}/{{.Repo}}

This is a weekly report of {{.Owner}}/{{.Repo}}. It summarizes what have changed in the project during the passed week, including pr merged, new contributors, and more things in the future.


## Repo Overview

### Basic data

Baisc data shows how the watch, star, fork and contributors count changed in the passed week.

| Watch | Star | Fork | Contributors |
|:-----:|:----:|:----:|:------------:|
| {{.WatcherCount}} | {{.StarCount}} ({{.StarDelta}}) | {{.ForksCount}} ({{.ForkDelta}}) | {{.Contributors}} ({{.ContributorDelta}}) |

### Issues & PRs

Issues & PRs show the new/closed issues/pull requests count in the passed week.


## PR Overview
Thanks to contributions from community, {{.Owner}}/{{.Repo}} team merged **{{.MergedPrCount}}** pull requests in the repository last week. They are:

{{range .PRs}}
* [{{.Title}}]({{.HTMLURL}})  `{{.MergedAt.Format "Jan 02"}}`
{{end}}
