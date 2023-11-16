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

type ChangeLogExists struct {
	baseDir string
}

func NewChangelogExists(baseDir string) *ChangeLogExists {
	return &ChangeLogExists{baseDir}
}

func (c ChangeLogExists) IsOptional() bool {
	return false
}

func (c ChangeLogExists) Name() string {
	return "TRG 1.03 - CHANGELOG.md"
}

func (c ChangeLogExists) Description() string {
	return "Tracking changes in Open Source is critical to have a way of knowing what new features have been introduced, what bugs have been fixed, what security CVEs have been mitigated. In Eclipse Tractus-X we use a CHANGELOG.md to document this"
}

func (c ChangeLogExists) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-3"
}

func (c ChangeLogExists) Test() *tractusx.QualityResult {
	_, err := os.Stat(path.Join(c.baseDir, "CHANGELOG.md"))

	if err != nil {
		return &tractusx.QualityResult{ErrorDescription: "A CHANGELOG.md file has to be present, describing the changes on between your releases"}
	}
	return &tractusx.QualityResult{Passed: true}
}
