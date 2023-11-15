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
	"tractusx-release-automation/internal/dashboard"
)

const buildOutputDir = "build"
const outputFileName = "index.html"

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "buildDashboard",
	Short: "Create a statically compiled dashboard with release check status",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		ensureOutputDirExists()
		copyAssets()

		var outputBuffer bytes.Buffer
		products, unhandledRepos, archivedRepos := dashboard.CheckProducts()

		dashboard.RenderHtmlTo(&outputBuffer, &dashboard.TemplateData{Config: getConfig(), CheckedProducts: products, UnhandledRepos: unhandledRepos, ArchivedRepos: archivedRepos})

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
	_, err = w.WriteString(buffer.String())
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

func getConfig() dashboard.Config {
	if os.Getenv("DASHBOARD_ASSETS_PATH") != "" {
		return dashboard.Config{AssetsPath: os.Getenv("DASHBOARD_ASSETS_PATH")}
	}
	return dashboard.Config{AssetsPath: "/assets"}
}
