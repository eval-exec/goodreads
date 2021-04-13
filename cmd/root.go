/*
Copyright © 2021 Slarsar <slarsar@yandex.com.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
	"github.com/spf13/cobra"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	csvW    *csv.Writer

	proxyServer string
	specificTag string
	debugf      bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:                    "goodreads",
	Aliases:                nil,
	SuggestFor:             nil,
	Short:                  "a simple goodread quote crawler",
	Example:                "./goodreads --proxy-server http://127.0.0.1:1081 --tag hope --debug",
	ValidArgs:              nil,
	ValidArgsFunction:      nil,
	Args:                   nil,
	ArgAliases:             nil,
	BashCompletionFunction: "",
	Deprecated:             "",
	Annotations:            nil,
	Version:                "",
	PersistentPreRun:       nil,
	PersistentPreRunE:      nil,
	PreRun: func(cmd *cobra.Command, args []string) {

		log.SetPrefix("goodreads")
		log.SetFlags(log.LstdFlags | log.Ltime)
		params := []func(collector *colly.Collector){
		}
		if debugf {
			params = append(params,
				colly.Debugger(&debug.LogDebugger{
					Output: nil,
					Prefix: "DEBUG ",
					Flag:   0,
				}))
		}

		c = colly.NewCollector(params...)
		if len(proxyServer) != 0 {

			if err := c.SetProxy(proxyServer); err != nil {
				log.Fatalf("set proxy %s\n", err)
			} else {
				log.Printf("set proxy %s success\n", proxyServer)
			}
		}
		c.AllowURLRevisit = false
		extensions.RandomUserAgent(c)

	},
	PreRunE: nil,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Create("goodreads.csv")
		if err != nil {
			log.Fatalf("fatal error: %s\n", err)
			return
		}
		defer file.Close()
		csvW = csv.NewWriter(file)
		defer csvW.Flush()

		allTags()

	},
	RunE:                       nil,
	PostRun:                    nil,
	PostRunE:                   nil,
	PersistentPostRun:          nil,
	PersistentPostRunE:         nil,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	TraverseChildren:           false,
	Hidden:                     false,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
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

	rootCmd.PersistentFlags().StringVar(&proxyServer, "proxy-server", "", "proxy server to use")
	rootCmd.PersistentFlags().StringVar(&specificTag, "tag", "", "optional: specify a tag to scrape?")
	rootCmd.PersistentFlags().BoolVar(&debugf, "debug", false, "debug mode?")
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

		// Search config in home directory with name ".goodreads" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".goodreads")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
