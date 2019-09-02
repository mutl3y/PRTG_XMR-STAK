/*
Copyright © 2019 Mutl3y

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
	"../stats"
	"fmt"

	"github.com/spf13/cobra"
)

// generalCmd represents the general command
var generalCmd = &cobra.Command{
	Use:   "Stats",
	Short: "General stats",
	Long: `Stats 

Basic XMR-Stak stats outputted as XML for ingestion by PRTG Network Monitor

Examples
PRTG_XMR-STAK.exe Stats -T 500ms -H 192.168.1.201 -P 420

./PRTG_XMR-STAK Stats -T 500ms -H 192.168.1.201 -P 420
`,
	Run: func(cmd *cobra.Command, args []string) {
		flags := cmd.Flags()
		h, err := flags.GetString("host")
		if err != nil {
			fmt.Println(err)
		}

		p, err := flags.GetInt("port")
		if err != nil {
			fmt.Println(err)

		}
		t, err := flags.GetDuration("timeout")
		if err != nil {
			fmt.Println(err)
		}
		url := fmt.Sprintf("http://%v:%v/Api.json", h, p)
		stats.PrtgStats(t, url)
	},
}

func init() {
	rootCmd.AddCommand(generalCmd)
}
