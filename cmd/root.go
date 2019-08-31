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
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "PRTG_XMR-STAK",
	Short: "PRTG Sensor for XMR-STAK",
	Long:  `PRTG Sensor for XMR-STAK`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
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

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.PRTG_XMR-STAK.yaml)")
	rootCmd.PersistentFlags().StringP("host", "H", "127.0.0.1", "hostname / IP")
	rootCmd.PersistentFlags().IntP("port", "P", 420, "port")
	rootCmd.PersistentFlags().DurationP("timeout", "T", 500*time.Millisecond, "timeout string eg 500ms")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".PRTG_XMR-STAK" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".PRTG_XMR-STAK")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
