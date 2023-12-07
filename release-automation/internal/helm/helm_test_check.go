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

	return &tractusx.QualityResult{Passed: true}
}