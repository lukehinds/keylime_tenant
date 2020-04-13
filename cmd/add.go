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

/*
The following functions are required:

mytenant.init_add(vars(args))
mytenant.preloop()
mytenant.do_cv()
mytenant.do_quote()
if args.verify:
mytenant.do_verify()
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an agent to the Verifier",
	Long: `The add flag is used to set up an Agent for measurement
	by the Keylime Verifier. Further sub commands can be used to 
	load objects such as whitelists of payloads.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("Get a flag:", viper.GetString("tpm_policy"))

		if viper.IsSet("whitelist") {
			whitelist := viper.GetString("whitelist")
			processWhitelist(whitelist)

		}
	},
}

func initAdd() {

}

func doCV() {

}

func preLoop() {

}

func doQuote() {

}

func processWhitelist(whitelist string) {
	fileBytes, err := ioutil.ReadFile(whitelist)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")
	fmt.Printf("%T\n", sliceData)

	fmt.Println(sliceData)
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().String("whitelist", "", "The IMA whitelist")
	viper.BindPFlag("whitelist", addCmd.PersistentFlags().Lookup("whitelist"))
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
