// Copyright © 2018 Dhananjay Balan
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package cmd

import (
	"fmt"
	"os"

	"bufio"
	"github.com/spf13/cobra"
	"net/url"
)

var (
	strict bool
	part   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gpurl [flags] [urls]",
	Short: "gpurl is a url parser-extractor",
	Long: `gpurl extracts parts from a url. It uses parser from go standard library.
Urls can be passed as arguments or piped via stdin`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {
			// we use args from cli
			for _, u := range args {
				_gpurl(u, part, strict)
			}
		} else {
			// read from stdin
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				_gpurl(scanner.Text(), part, strict)
			}
			handleError(scanner.Err())
		}
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
	rootCmd.PersistentFlags().StringVarP(&part, "part", "p", "host", "part of the url to extract, could be host, scheme, path")
	rootCmd.PersistentFlags().BoolVarP(&strict, "empty-error", "e", false, "error if extracted part is empty")
}

// this is the main for core functionality
func _gpurl(link, partname string, strict bool) {
	res, err := parseURL(link, partname)
	handleError(err)
	if res != "" {
		fmt.Println(res)
	} else if strict {
		handleError(fmt.Errorf("%s is empty and strict checking is enabled", partname))
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(-1)
	}
}

func findPart(partname string, purl *url.URL) (res string, err error) {
	switch partname {
	case "host":
		res = purl.Hostname()
	case "scheme":
		res = purl.Scheme
	case "path":
		res = purl.Path
	default:
		err = fmt.Errorf("wrong partname: %s", partname)
	}
	return
}

func parseURL(link, partname string) (string, error) {
	purl, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	return findPart(partname, purl)
}
