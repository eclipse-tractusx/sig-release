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

package docs

import (
	"os"
	"path"

	"tractusx-release-automation/internal/tractusx"
)

type InstallExists struct {
	baseDir string
}

// NewInstallExists returns a new check based on QualityGuideline interface.
func NewInstallExists(baseDir string) tractusx.QualityGuideline {
	return &InstallExists{baseDir}
}

func (i *InstallExists) Name() string {
	return "TRG 1.02 - INSTALL.md"
}

func (i *InstallExists) Description() string {
	return "File INSTALL.md contains comprehensive instructions for installation."
}

func (i *InstallExists) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-2"
}

func (i *InstallExists) Test() *tractusx.QualityResult {
	_, err := os.Stat(path.Join(i.baseDir, "INSTALL.md"))

	if err != nil {
		return &tractusx.QualityResult{
			ErrorDescription: "Optional file INSTALL.md not found in current directory.",
		}
	}
	return &tractusx.QualityResult{Passed: true}
}

func (i *InstallExists) IsOptional() bool {
	return true
}
