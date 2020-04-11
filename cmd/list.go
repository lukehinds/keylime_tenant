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
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type JsonResults struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Results struct {
		OperationalState        int      `json:"operational_state"`
		V                       string   `json:"v"`
		IP                      string   `json:"ip"`
		Port                    int      `json:"port"`
		TpmPolicy               string   `json:"tpm_policy"`
		VtpmPolicy              string   `json:"vtpm_policy"`
		MetaData                string   `json:"meta_data"`
		ImaWhitelistLen         int      `json:"ima_whitelist_len"`
		TpmVersion              int      `json:"tpm_version"`
		AcceptTpmHashAlgs       []string `json:"accept_tpm_hash_algs"`
		AcceptTpmEncryptionAlgs []string `json:"accept_tpm_encryption_algs"`
		AcceptTpmSigningAlgs    []string `json:"accept_tpm_signing_algs"`
		HashAlg                 string   `json:"hash_alg"`
		EncAlg                  string   `json:"enc_alg"`
		SignAlg                 string   `json:"sign_alg"`
	} `json:"results"`
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the operational state of an Agent",
	Long:  `Lists the operational state of the an Agent`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set up agent operational state mapping
		states := map[int]string{
			0:  "REGISTERED",
			1:  "START",
			2:  "SAVED",
			3:  "GET_QUOTE",
			4:  "GET_QUOTE_RETRY",
			5:  "PROVIDE_V",
			6:  "PROVIDE_V_RETRY",
			7:  "FAILED",
			8:  "TERMINATED",
			9:  "INVALID_QUOTE",
			10: "TENANT_FAILED",
		}

		var urlreq string
		var verifier_ip = viper.GetString("verifier_ip")
		var verifier_port = viper.GetString("verifier_port")
		// fmt.Println("uuid exists?", viper.IsSet("uuid"))
		// fmt.Println("Get a flag:", viper.GetString("uuid"))
		// fmt.Println("Get a port:", viper.GetInt("agent_port"))

		//priv, err := ioutil.ReadFile("/var/lib/keylime/cv_ca/client-private.pem")
		//passbyte := []byte("default")
		//unenckey := decrypt(priv, passbyte)
		// stringkey := fmt.Sprint(unenckey)
		cert, err := tls.LoadX509KeyPair("/var/lib/keylime/cv_ca/client-cert.crt", "/var/lib/keylime/cv_ca/client-private-out.pem")
		if err != nil {
			log.Fatalf("Error reading cert: %s", err)
		}
		//
		caCert, err := ioutil.ReadFile("/var/lib/keylime/cv_ca/client-cert.crt")
		if err != nil {
			log.Fatalf("Error reading caCert: %s", err)
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		// Create a HTTPS client and supply the created CA pool and certificate
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
					RootCAs:            caCertPool,
					Certificates:       []tls.Certificate{cert},
				},
			},
		}

		var baseurl = fmt.Sprintf("https://%s:%s/agents/", verifier_ip, verifier_port)

		if viper.IsSet("uuid") {
			urlreq = fmt.Sprintf("%s%s", baseurl, viper.GetString("uuid"))

		} else {
			// var listing bool = true
			urlreq = fmt.Sprintf("%s", baseurl)
		}

		response, err := client.Get(urlreq)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println("HTTP Response Status:", response.StatusCode, http.StatusText(response.StatusCode))
		//fmt.Println("The Header:", response.Header)
		// Read the response body
		defer response.Body.Close()
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseBody))

		if response.StatusCode == 503 {
			log.Printf("Cannot connect to Verifier at %v with Port %v timed out. Connection refused.", verifier_ip, verifier_port)
		} else if response.StatusCode == 504 {
			log.Printf("Verifier at %v with Port %v timed out.", verifier_ip, verifier_port)
		} else if response.StatusCode == 404 {
			log.Printf("Agent %v does not exist on the verifier. Please try to add or update agent.", viper.GetString("uuid"))
		} else if response.StatusCode == 200 {
			var jsonresults JsonResults
			json.Unmarshal(responseBody, &jsonresults)
			log.Printf("Agent Status: %v", states[jsonresults.Results.OperationalState])
		} else {
			log.Printf("Unexpected response from Cloud Verifier: %v.", http.StatusText(response.StatusCode))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().String("uuid", "", "A help for uuid")
	viper.BindPFlag("uuid", listCmd.PersistentFlags().Lookup("uuid"))
}
