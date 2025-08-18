/*******************************************************************************
 * Copyright (c) 2025 Fraunhofer-Gesellschaft zur Foerderung der angewandten Forschung e.V. (represented by Fraunhofer ISST)
 * Copyright (c) 2025 Contributors to the Eclipse Foundation
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

package docs

import (
	"log"
	"path"
	"path/filepath"
	"tractusx-release-automation/internal/tractusx"
)

var allowedArchitectureDocsGlob = []string{
	"docs/architecture/*.md",
	"docs/architecture/*.adoc",
}

type ArchitectureDocumentationExists struct {
	baseDir string
}

func NewArchitectureDocumentationExists(baseDir string) tractusx.QualityGuideline {
	return &ArchitectureDocumentationExists{baseDir}
}

func (r *ArchitectureDocumentationExists) IsOptional() bool {
	return false
}

func (r *ArchitectureDocumentationExists) Name() string {
	return "TRG 1.05 - Arc42 Documentation"
}

func (r *ArchitectureDocumentationExists) Description() string {
	return "Proper Architecture Documentation is necessary to support architects and operation teams for their decisions making processes and operative tasks. It therefore provides information like context, integration interfaces and intended usage of the component. We follow the good practice to document the Architecture of components."
}

func (r *ArchitectureDocumentationExists) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-05"
}

func (r *ArchitectureDocumentationExists) BaseDir() string {
	return r.baseDir
}

func (r *ArchitectureDocumentationExists) Test() *tractusx.QualityResult {
	checkPassed := false
	filepath.Glob(path.Join(r.baseDir, "*.md"))

	allowedGlobs := getArchitectureDocEntryPaths(r.baseDir)

	for _, allowedGlob := range allowedGlobs {
		matches, err := filepath.Glob(path.Join(r.baseDir, allowedGlob))
		if err != nil || len(matches) > 0 {
			log.Printf("Did FIND a file matching glob %s", allowedGlob)
			checkPassed = true
		} else {
			log.Printf("Did NOT find a file matching glob %s", allowedGlob)
		}
	}

	if checkPassed == false {
		return &tractusx.QualityResult{ErrorDescription: "Did not find a markdown or adoc file in folder docs/architecture OR specified path (.tractusx file)!"}
	}
	return &tractusx.QualityResult{Passed: true}
}

// return allowedArchitectureDocsGlob + ArchitectureDocEntryPath from .tractusx
func getArchitectureDocEntryPaths(dir string) []string {
	file, err := tractusx.MetadataFromLocalFile(dir)

	if err != nil {
		return allowedArchitectureDocsGlob
	}

	return append(allowedArchitectureDocsGlob, file.ArchitectureDocEntryPath)
}
