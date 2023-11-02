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

package container

import (
	"fmt"
	"strings"

	"tractusx-release-automation/internal/tractusx"
)

var baseImageAllowList = []string{
	"eclipse-temurin",
	"nginxinc/nginx-unprivileged",
	"mcr.microsoft.com/dotnet/runtime",
	"mcr.microsoft.com/dotnet/aspnet",
	"alpine",
}

type AllowedBaseImage struct {
	baseDir string
}

func NewAllowedBaseImage(baseDir string) tractusx.QualityGuideline {
	return &AllowedBaseImage{baseDir: baseDir}
}

func (a *AllowedBaseImage) Name() string {
	return "TRG 4.02 - Base images"
}

func (a *AllowedBaseImage) Description() string {
	return "We are aligning all product Docker images to a set of approved ones. This also makes it easier to properly annotate the used images as dependency"
}

func (a *AllowedBaseImage) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-4/trg-4-02"
}

func (a *AllowedBaseImage) Test() *tractusx.QualityResult {
	foundDockerFiles := findDockerfilesAt(a.baseDir)
	dockerfilesToSkip := getDockerfilePathsToIgnore(a.baseDir)

	checkPassed := true
	var deniedBaseImages []string
	for _, dockerfilePath := range foundDockerFiles {
		if containsString(dockerfilesToSkip, dockerfilePath) {
			fmt.Printf("Dockerfile at path %s configured to skip", dockerfilePath)
			continue
		}

		file, err := dockerfileFromPath(dockerfilePath)
		if err != nil {
			fmt.Printf("Could not read dockerfile from Path %s\n", dockerfilePath)
			continue
		}

		if !isAllowedBaseImage(file.baseImage()) {
			checkPassed = false
			deniedBaseImages = append(deniedBaseImages, file.baseImage())
		}
	}

	return &tractusx.QualityResult{Passed: checkPassed, ErrorDescription: buildErrorDescription(deniedBaseImages)}
}

func (a *AllowedBaseImage) IsOptional() bool {
	return false
}

// Function to return error message of failed test.
// There are two types of formatting: cli (default) and web
func buildErrorDescription(deniedImages []string) string {
	if len(deniedImages) == 0 {
		return ""
	}
	if tractusx.ErrorOutputFormat == tractusx.WebErrOutputFormat {
		return "Dockerfile(s) use not approved images:<br>" + strings.Join(deniedImages, "<br>")
	}
	return "We want to align on docker base images. We detected a Dockerfile specifying " +
		strings.Join(deniedImages, ", ") + "\n\tAllowed images are: \n\t - " +
		strings.Join(baseImageAllowList, "\n\t - ")
}

func isAllowedBaseImage(image string) bool {
	for _, imageFromAllowList := range baseImageAllowList {
		if strings.Contains(image, imageFromAllowList) {
			return true
		}
	}
	return false
}

func getDockerfilePathsToIgnore(dir string) []string {
	file, err := tractusx.MetadataFromLocalFile(dir)
	if err != nil {
		return []string{}
	}

	return file.AlignedBaseImages
}

func containsString(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
