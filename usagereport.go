package main

import "github.com/cloudfoundry/cli/plugin"

//UsageReportCmd the plugin
type UsageReportCmd struct {
}

type organization struct {
	url       string
	name      string
	quotaURL  string
	spacesURL string
}

//GetMetadata returns metatada
func (cmd *UsageReportCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "usage-report",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		Commands: []plugin.Command{
			{
				Name:     "usage-report",
				HelpText: "Report AI and memory usage for orgs and spaces",
				UsageDetails: plugin.Usage{
					Usage: "cf usage-report",
				},
			},
		},
	}
}

//UsageReportCommand doer
func (cmd *UsageReportCmd) UsageReportCommand(cli plugin.CliConnection, args []string) {
	//Do the things
}

//Run runs the plugin
func (cmd *UsageReportCmd) Run(cli plugin.CliConnection, args []string) {
	if args[0] == "usage-report" {
		cmd.UsageReportCommand(cli, args)
	}
}

func main() {
	plugin.Start(new(UsageReportCmd))
}
