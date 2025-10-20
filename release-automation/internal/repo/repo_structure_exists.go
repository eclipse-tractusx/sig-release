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

package repo

import (
	"fmt"
	"path"
	"strings"
	"tractusx-release-automation/internal/filesystem"
	"tractusx-release-automation/internal/tractusx"
)

type RepoStructureExists struct {
	baseDir string
}

func NewRepoStructureExists(baseDir string) tractusx.QualityGuideline {
	return &RepoStructureExists{baseDir}
}

func (c RepoStructureExists) IsOptional() bool {
	return false
}

func (c RepoStructureExists) Name() string {
	return "TRG 2.03 - Repo structure"
}

func (c RepoStructureExists) Description() string {
	return "All repositories must follow specified files and folders structure."
}

func (c RepoStructureExists) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-3"
}

func (c RepoStructureExists) BaseDir() string {
	return c.baseDir
}

func (c RepoStructureExists) Test() *tractusx.QualityResult {
	// Slices containing required files and folders in the repo structure.
	// Before modification make sure you align to TRG 2.03 guideline.
	listOfOptionalFilesToBeChecked := []string{
		path.Join(c.baseDir, "AUTHORS.md"),
		path.Join(c.baseDir, "INSTALL.md"),
	}

	listOfMandatoryFilesToBeChecked := []string{
		path.Join(c.baseDir, "CODE_OF_CONDUCT.md"),
		path.Join(c.baseDir, "CONTRIBUTING.md"),
		path.Join(c.baseDir, "LICENSE"),
		path.Join(c.baseDir, "NOTICE.md"),
		path.Join(c.baseDir, "README.md"),
		path.Join(c.baseDir, "SECURITY.md"),
	}

	mandatoryForLeadingRepo := []string{
		path.Join(c.baseDir, "docs"),
	}

	mandatoryForLeadingProductRepo := []string{
		path.Join(c.baseDir, "charts"),
	}

	printer := &tractusx.StdoutPrinter{}

	if isLeadingRepo(c.baseDir) {
		listOfMandatoryFilesToBeChecked = append(listOfMandatoryFilesToBeChecked, mandatoryForLeadingRepo...)

		if isProductRepo(c.baseDir) {
			listOfMandatoryFilesToBeChecked = append(listOfMandatoryFilesToBeChecked, mandatoryForLeadingProductRepo...)
		}
	}

	missingMandatoryFiles := filesystem.CheckMissingFiles(listOfMandatoryFilesToBeChecked)
	missingOptionalFiles := filesystem.CheckMissingFiles(listOfOptionalFilesToBeChecked)
	if dependencyFiles := filesystem.FindPrefixedFiles(c.baseDir, "DEPENDENCIES"); dependencyFiles == nil {
		missingMandatoryFiles = append(missingMandatoryFiles, path.Join(c.baseDir, "DEPENDENCIES"))
	}

	if len(missingOptionalFiles) > 0 {
		printer.LogWarning(
			fmt.Sprintf("Warning! Guideline description: %s\n\t%s\n\tMore infos: %s\n",
				c.Description(), "The check detected following optional files missing: "+strings.Join(missingOptionalFiles, ", "),
				c.ExternalDescription()))
	}

	if len(missingMandatoryFiles) > 0 {
		cleanMissingFiles := []string{}
		for _, missingFile := range missingMandatoryFiles {
			// if a mandatory folder is missing, the path does not contain the base dir and thus mustn't be removed
			if strings.HasPrefix(missingFile, c.baseDir) {
				cleanMissingFiles = append(cleanMissingFiles, strings.Split(missingFile, c.baseDir)[1][1:])
			} else {
				cleanMissingFiles = append(cleanMissingFiles, missingFile)
			}
		}
		return &tractusx.QualityResult{ErrorDescription: "The check detected following mandatory files missing: " + strings.Join(cleanMissingFiles, ", ")}
	}

	return &tractusx.QualityResult{Passed: true}
}

func (a *RepoStructureExists) IsApplicableToCategory(category tractusx.RepoCategory) bool {
	return category == tractusx.RepoCategoryProduct || category == tractusx.RepoCategorySupport || category == tractusx.RepoCategorySpecial
}
