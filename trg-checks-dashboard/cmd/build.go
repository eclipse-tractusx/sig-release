/*******************************************************************************
 * Copyright (c) 2023 Contributors to the Eclipse Foundation
 *
 * See the NOTICE file(s) distributed with this work for additional
 * information regarding copyright ownership.
 *
 * This program and the accompanying materials are made available under the
 * terms of the Apache License, Version 2.0 which is available at
 * https://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 ******************************************************************************/

package cmd

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"path"

	cp "github.com/otiai10/copy"
	"github.com/spf13/cobra"
	"trg-checks-dashboard/internal/templating"
	"trg-checks-dashboard/internal/tractusx"
)

const buildOutputDir = "build"
const outputFileName = "index.html"

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ensureOutputDirExists()
		copyAssets()

		var outputBuffer bytes.Buffer
		templating.RenderHtmlTo(&outputBuffer, &templating.TemplateData{Config: getConfig(), Checks: getChecks()})

		writeToFile(outputBuffer)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func writeToFile(buffer bytes.Buffer) {
	f, err := os.Create(path.Join(buildOutputDir, outputFileName))
	if err != nil {
		log.Fatalf("Could not create output file: %v", err)
	}

	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(buffer.Bytes()))
	if err != nil {
		log.Fatalf("Could not write to output file: %v", err)
	}

	if err = w.Flush(); err != nil {
		log.Fatalf("Could not flush output: %v", err)
	}
}

func ensureOutputDirExists() {
	if _, err := os.Stat(buildOutputDir); err != nil {
		if err := os.Mkdir(buildOutputDir, 0777); err != nil {
			log.Fatalf("Could not create build output dir: %v", err)
		}
	}
}

func copyAssets() {
	if err := cp.Copy("web/assets", path.Join(buildOutputDir, "assets")); err != nil {
		log.Fatalf("Could not copy CSS to build output dir: %v", err)
	}
}

func getConfig() templating.Config {
	return templating.Config{AssetsPath: "/assets"}
}

func getChecks() []templating.ProductCheck {
	for _, product := range tractusx.QueryProducts() {
		log.Printf("Would have found product: %s", product.Name)
	}
	return []templating.ProductCheck{
		{
			Product: tractusx.Product{Name: "Portal", LeadingRepo: "https://github.com/eclipse-tractusx/portal-cd", Repositories: []string{}},
			Checks:  []tractusx.ReleaseGuidelineCheck{},
		},
		{
			Product: tractusx.Product{Name: "EDC", LeadingRepo: "", Repositories: []string{}},
			Checks:  []tractusx.ReleaseGuidelineCheck{},
		},
		{
			Product: tractusx.Product{Name: "Trace-X", LeadingRepo: "", Repositories: []string{}},
			Checks:  []tractusx.ReleaseGuidelineCheck{},
		},
	}
}
