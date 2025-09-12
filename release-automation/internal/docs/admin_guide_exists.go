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
	"path"
	"path/filepath"
	"tractusx-release-automation/internal/tractusx"
)

type AdminGuideExists struct {
	baseDir string
}

func NewAdminGuideExists(baseDir string) tractusx.QualityGuideline {
	return &AdminGuideExists{baseDir: baseDir}
}

func (a *AdminGuideExists) Name() string { return "TRG 1.06 - docs/admin" }

func (a *AdminGuideExists) Description() string {
	return "Tractus-X components and applications may be operated in various productive environments. An Administrator`s guide shall help operators to install and operate the solution. It can also be a vehicle to share best-practices on maintenance etc."
}

func (a *AdminGuideExists) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-06"
}

func (a *AdminGuideExists) BaseDir() string {
	return a.baseDir
}

func (a *AdminGuideExists) Test() *tractusx.QualityResult {
	matches, err := filepath.Glob(path.Join(a.baseDir, "docs", "admin", "*.md"))

	if err != nil || len(matches) == 0 {
		return &tractusx.QualityResult{ErrorDescription: "A markdown file providing the admin guide has to be present."}
	}
	return &tractusx.QualityResult{Passed: true}
}

func (a *AdminGuideExists) IsOptional() bool {
	return true
}
