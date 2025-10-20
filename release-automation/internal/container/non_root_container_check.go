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

package container

import (
	"fmt"
	"regexp"
	"strings"
	"tractusx-release-automation/internal/tractusx"
)

// validateUserRegex is used to match valid username/uid, group-name/gid
const validateUserRegex = `(^(\d|[1-9]\d{1,3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-6])$)|(^[a-z_][a-z0-9_-]*[$]?$)`

type NonRootContainer struct {
	baseDir string
}

// NewNonRootContainer returns a new check based on tractusx.QualityGuideline interface.
func NewNonRootContainer(baseDir string) tractusx.QualityGuideline {
	return &NonRootContainer{baseDir}
}

func (n NonRootContainer) Name() string {
	return "TRG 4.03 - Non-root container"
}

func (n NonRootContainer) Description() string {
	return "Container images shall not run as root for security reasons."
}

func (n NonRootContainer) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-4/trg-4-03"
}

func (n NonRootContainer) BaseDir() string {
	return n.baseDir
}

func (n NonRootContainer) Test() *tractusx.QualityResult {
	checkPassed := true
	errorDescription := ""
	dockerfiles := findDockerfilesAt(n.baseDir)

	for _, dockerfilePath := range dockerfiles {
		file, err := dockerfileFromPath(dockerfilePath)
		if err != nil {
			fmt.Printf("Could not read Dockerfile from path %s\n", dockerfilePath)
		}

		if !validateUser(file.user()) {
			checkPassed = false
			if n.baseDir != "./" {
				dockerfilePath = strings.Split(dockerfilePath, n.baseDir)[1]
			}
			errorDescription += "\nInvalid user specified in Dockerfile: " + dockerfilePath[1:]
			if tractusx.CliErrOutputFormat == tractusx.WebErrOutputFormat {
				errorDescription += "<br>"
			}
		}
	}

	return &tractusx.QualityResult{Passed: checkPassed, ErrorDescription: errorDescription}
}

func (n NonRootContainer) IsOptional() bool {
	return false
}

// validateUser validates USER instruction in Dockerfiles and return a bool if a valid USER has been found.
// To return true username/group-name must not be root or contain upper case letters, or must not be 0 or greater 65536
// for uid/gid.
func validateUser(u *user) bool {
	if u == nil {
		return false
	}

	validUser, _ := regexp.Match(validateUserRegex, []byte(u.user))

	var validGroup bool

	if u.group == "" {
		validGroup = true
	} else {
		validGroup, _ = regexp.Match(validateUserRegex, []byte(u.group))
	}

	// return false if user is malformed (uppercase, or uid > 65536) or a root or missing USER instruction is detected.
	return validUser && validGroup && !isRootUser(u.user) && !isRootGroup(u.group)
}

func isRootUser(user string) bool {
	return user == "root" || user == "0"
}

func isRootGroup(group string) bool {
	return group == "root" || group == "0"
}

func (a *NonRootContainer) IsApplicableToCategory(category tractusx.RepoCategory) bool {
	return category == tractusx.RepoCategoryProduct
}
