/*
MIT License

Copyright (c) 2020 Luke Hinds

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "keylime_tenant",
	Short: "The Keylime Tenant CLI tool",
	Long: `The Keylime Tenant CLI tool is used for the configuration
of KeyLime Agent nodes.`,
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.keylime_tenant.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().String("v", "", "The IP Address of the Verifier")
	viper.BindPFlag("v", rootCmd.PersistentFlags().Lookup("v"))

	rootCmd.PersistentFlags().String("t", "", "The IP Address of the Agent")
	viper.BindPFlag("t", rootCmd.PersistentFlags().Lookup("t"))

	rootCmd.PersistentFlags().Int("tp", 9002, "The Port of the Agent")
	viper.BindPFlag("tp", rootCmd.PersistentFlags().Lookup("tp"))

	rootCmd.PersistentFlags().String("uuid", "", "The UUID of the Agent to list")
	viper.BindPFlag("uuid", rootCmd.PersistentFlags().Lookup("uuid"))
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
			log.Fatalln(err)
		}

		// Search config in home directory with name ".pask" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".keylime_tenant")
		viper.SetConfigType("yaml")

		// In the succeeding line, the reader will note that
		// ReadInConfig returns an error, and I am intentionally
		// ignoring it.
		//
		// If the config file couldn't be read, *I don't care*. But
		// if it can be, it should be.
		err = viper.ReadInConfig()
		if err != nil {
			log.Println("Error reading config file: ", err)
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		log.Println("Cannot locate config file:", viper.ConfigFileUsed())
	}
}
