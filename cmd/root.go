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
	"fmt"
	"github.com/mutl3y/PRTG_XMR-STAK/stats"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "PRTG_XMR-STAK",
	Short: "PRTG Sensor for XMR-STAK",
	Long: `
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
		th, err := flags.GetBool("threads")
		url := fmt.Sprintf("http://%v:%v/Api.json", h, p)
		stats.PrtgStats(t, url, th)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("host", "H", "127.0.0.1", "hostname / IP")
	rootCmd.PersistentFlags().IntP("port", "P", 420, "port")
	rootCmd.PersistentFlags().DurationP("timeout", "T", 500*time.Millisecond, "timeout string eg 500ms")
	rootCmd.PersistentFlags().BoolP("threads", "t", false, "include thread info")
}
