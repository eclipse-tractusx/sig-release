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
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"tractusx-release-automation/internal/tractusx"
)

type HelmWorkflowCheck struct {
	baseDir string
}

func NewHelmWorkflowCheck(baseDir string) tractusx.QualityGuideline {
	return &HelmWorkflowCheck{baseDir}
}

func (r *HelmWorkflowCheck) Name() string {
	return "TRG 5.09 - Helm test"
}

func (r *HelmWorkflowCheck) Description() string {
	return "Helm test is the technical solution helm provides to verify that a helm chart works as expected."
}

func (r *HelmWorkflowCheck) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-09"
}

func (r *HelmWorkflowCheck) IsOptional() bool {
	return false
}

func (r *HelmWorkflowCheck) Test() *tractusx.QualityResult {
	if fi, err := os.Stat(path.Join(r.baseDir, "charts")); err != nil || !fi.IsDir() {
		return &tractusx.QualityResult{Passed: true}
	}
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

		w, err := loadGitHubWorkflowFromFile(path.Join(r.baseDir, ".github/workflows", workflow.Name()))
		if err != nil {
			return &tractusx.QualityResult{ErrorDescription: err.Error()}
		}
		if isValidHelmTestWorkflow(w) {
			return &tractusx.QualityResult{Passed: true}
		}
	}

	return &tractusx.QualityResult{ErrorDescription: "Helm test workflow not found."}
}

func isValidHelmTestWorkflow(workflow GitHubWorkflow) bool {
	var (
		hasDispatchTrigger    bool
		hasChartTestingAction bool
		hasKindAction         bool
		hasHelmTestRun        bool
		hasHelmInstallRun     bool
	)

	for trigger := range workflow.On {
		if strings.EqualFold(trigger, "workflow_dispatch") {
			hasDispatchTrigger = true
		}
	}

	helmTestPattern := `\b(ct|helm)\s*lint\b.*`
	testRegexp := regexp.MustCompile(helmTestPattern)
	helmInstallPattern := `\b(ct|helm)\s*install\b`
	installRegexp := regexp.MustCompile(helmInstallPattern)

	for _, job := range workflow.Jobs {
		for _, step := range job.Steps {
			if strings.Contains(step.Uses, "chart-testing-action") {
				hasChartTestingAction = true
			}
			if strings.Contains(step.Uses, "kind-action") {
				hasKindAction = true
			}
			if testRegexp.MatchString(step.Run) {
				hasHelmTestRun = true
			}
			if installRegexp.MatchString(step.Run) {
				hasHelmInstallRun = true
			}
		}
	}
	allCheckPassed := hasDispatchTrigger && hasChartTestingAction && hasKindAction && hasHelmTestRun && hasHelmInstallRun
	return allCheckPassed
}

func loadGitHubWorkflowFromFile(filePath string) (GitHubWorkflow, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Unable to read %v.\n", filePath)
		return GitHubWorkflow{}, fmt.Errorf("unable to read workflow: %v", err)
	}

	var w GitHubWorkflow
	err = yaml.Unmarshal(data, &w)
	if err != nil {
		log.Printf("Unable to parse YAML file: %v.\n", filePath)
		return GitHubWorkflow{}, fmt.Errorf("unable to parse workflow: %v", err)
	}

	return w, nil
}
