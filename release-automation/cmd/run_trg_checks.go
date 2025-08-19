/*******************************************************************************
 * Copyright (c) 2023 Contributors to the Eclipse Foundation
 * Copyright (c) 2025 Fraunhofer-Gesellschaft zur Foerderung der angewandten Forschung e.V. (represented by Fraunhofer ISST)
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
	"fmt"
	"os"

	"tractusx-release-automation/internal/container"
	"tractusx-release-automation/internal/docs"
	"tractusx-release-automation/internal/helm"
	"tractusx-release-automation/internal/open_source"
	"tractusx-release-automation/internal/repo"
	"tractusx-release-automation/internal/test_runner"
	"tractusx-release-automation/internal/tractusx"

	"github.com/spf13/cobra"
)

// checkLocalCmd represents the checkLocal command
var checkLocalCmd = &cobra.Command{
	Use:   "checkLocal",
	Short: "Does run a quality check on local files",
	Long:  `Execute the checkLocal command in any directory you want to check for quality compliance with eclipse-tractusx rules`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running local checks of eclipse-tractusx release guidelines")

		basedir := os.Getenv("CHECKLOCAL_BASEDIR")
		if basedir == "" {
			basedir = "./"
		}

		var releaseGuidelines = []tractusx.QualityGuideline{
			docs.NewReadmeExists(basedir),
			docs.NewInstallExists(basedir),
			docs.NewChangelogExists(basedir),
			docs.NewAdminGuideExists(basedir),
			docs.NewArchitectureDocumentationExists(basedir),
			repo.NewDefaultBranch(basedir),
			repo.NewRepoStructureExists(basedir),
			repo.NewLeadingRepositoryDefined(basedir),
			container.NewAllowedBaseImage(basedir),
			container.NewNonRootContainer(basedir),
			helm.NewHelmStructureExists(basedir),
			helm.NewResourceMgmt(basedir),
			helm.NewHelmWorkflowCheck(basedir),
			open_source.NewNoticeForNonCodeExists(basedir),
		}

		runner := testrunner.NewTestRunner(releaseGuidelines)
		err := runner.Run()

		if err != nil {
			fmt.Println("Error occurred! Check command output for details on failed checks")
			os.Exit(1)
		}

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(checkLocalCmd)
}
