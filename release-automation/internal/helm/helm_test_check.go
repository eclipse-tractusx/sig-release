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

package helm

import (
	"bufio"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"tractusx-release-automation/internal/tractusx"
)

type HelmTestCheck struct {
	baseDir string
}

func NewHelmTestCheck(baseDir string) tractusx.QualityGuideline {
	return &HelmTestCheck{baseDir}
}

func (r *HelmTestCheck) Name() string {
	return "TRG 5.09 - Helm test"
}

func (r *HelmTestCheck) Description() string {
	return "Helm test is the technical solution helm provides to verify that a helm chart works as expected."
}

func (r *HelmTestCheck) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-09"
}

func (r *HelmTestCheck) IsOptional() bool {
	return false
}

func (r *HelmTestCheck) Test() *tractusx.QualityResult {
	workflowsDir := path.Join(r.baseDir, ".github/workflows")
	if fi, err := os.Stat(workflowsDir); err != nil || !fi.IsDir() {
		return &tractusx.QualityResult{ErrorDescription: "Expected GitHub workflows directory doesn't exist in repository."}
	}
	workflows, err := os.ReadDir(workflowsDir)
	if err != nil || len(workflows) == 0 {
		return &tractusx.QualityResult{ErrorDescription: "No workflows found in repository."}
	}
	for _, workflow := range workflows {
		if workflow.IsDir() {
			continue
		}
		file, err := os.Open(path.Join(r.baseDir, ".github/workflows", workflow.Name()))
		if err != nil {
			log.Println("Can't read workflow: ", err)
			continue
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var lines []string

		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		file.Close()
		helmTestCommandRegexp := `^.*\bct\s*lint\b.*`
		re := regexp.MustCompile(helmTestCommandRegexp)
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if  re.MatchString(trimmed) && trimmed[0] != '#' {
				return &tractusx.QualityResult{Passed: true}
			}
		}
	}

	return &tractusx.QualityResult{ErrorDescription: "Helm test workflow not found."}
}
