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

type ReadmeExists struct {
	baseDir string
}

func NewReadmeExists(baseDir string) tractusx.QualityGuideline {
	return &ReadmeExists{baseDir}
}

func (r *ReadmeExists) IsOptional() bool {
	return false
}

func (r *ReadmeExists) Name() string {
	return "TRG 1.01 - README.md"
}

func (r *ReadmeExists) Description() string {
	return "A good README.md file is the starting point for everyone opening a repository. It should help them find all critical information in an easy way."
}

func (r *ReadmeExists) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-1/trg-1-1"
}

func (r *ReadmeExists) Test() *tractusx.QualityResult {
	_, err := os.Stat(path.Join(r.baseDir, "README.md"))

	if err != nil {
		return &tractusx.QualityResult{ErrorDescription: "Did not find a README.md file in current directory!"}
	}
	return &tractusx.QualityResult{Passed: true}
}
