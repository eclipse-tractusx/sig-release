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
	"os"
	"path"
	"testing"
)

const validHelmTestWorkflow = "helm_test_workflow_valid.yaml"
const invalidHelmTestWorkflow = "helm_test_workflow_invalid.yaml"

func TestShouldPassIfNoChartsDirPresent(t *testing.T) {
	dir := t.TempDir()
	result := NewHelmWorkflowCheck(dir).Test()

	if !result.Passed {
		t.Errorf("Should pass, no charts directory present.")
	}
}
func TestShouldPassIfWorkflowContainsHelmLintInstall(t *testing.T) {
	dir := t.TempDir()
	workflowDir := path.Join(dir, ".github", "workflows")
	_ = os.MkdirAll(workflowDir, 0770)
	copyTestFileTo(validHelmTestWorkflow, workflowDir, t)

	result := NewHelmWorkflowCheck(dir).Test()

	if !result.Passed {
		t.Errorf("Should pass, given workflow contains helm lint and install run commands.")
	}
}

func TestShouldFailIfWorkflowHasNoHelmLint(t *testing.T) {
	dir := t.TempDir()
	workflowDir := path.Join(dir, ".github", "workflows")
	_ = os.MkdirAll(workflowDir, 0770)
	copyTestFileTo(invalidHelmTestWorkflow, workflowDir, t)

	result := NewHelmWorkflowCheck(dir).Test()

	if result.Passed {
		t.Errorf("Should fail, given workflow doesn't have helm lint run command.")
	}
}

// Copy choosen test file to destination directory.
func copyTestFileTo(file string, dest string, t *testing.T) {
	data, err := os.ReadFile(path.Join("test", file))
	if err != nil {
		t.Errorf("Can't read provided file: %v", err)
	}
	err = os.WriteFile(path.Join(dest, file), data, 0770)
	if err != nil {
		t.Errorf("Can't write to provided directory: %v", err)
	}
}
