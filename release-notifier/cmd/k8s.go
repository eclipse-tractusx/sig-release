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
	"log"
	"os"
	"regexp"
	"release-notifier/internal/k8s"

	"github.com/spf13/cobra"
)

// k8sCmd represents the k8s command
var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Sends out notification of new Kubernetes release",
	Long: `Sends out notification of new Kubernetes release
	to a mailinglist:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Kubernetes release notifier called.")
		k8sNotifier()
	},
}

func init() {
	rootCmd.AddCommand(k8sCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// k8sCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// k8sCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func k8sNotifier() {
	// Semantic Versioning schema regex
	const regexPattern = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

	latestRelease := k8s.GetLatestRel()
	prevRelease := k8s.GetPrevRelFromArtifact()

	if match, _ := regexp.MatchString(regexPattern, latestRelease); match && latestRelease != prevRelease {
		log.Printf("New release is out: %v\n", latestRelease)
		alignedRelease := os.Getenv("CURRENT_ALIGNED_K8S_VER")
		log.Printf("Current aligned version: %v\n", alignedRelease)
		k8s.Notify(latestRelease, alignedRelease)
		k8s.SaveLatestRel(latestRelease)
	} else {
		log.Println("No new release :(")
	}
}