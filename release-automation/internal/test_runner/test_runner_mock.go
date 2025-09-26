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

package testrunner

import (
	"fmt"
	"tractusx-release-automation/internal/tractusx"
)

const DEFAULT_REPO_CATEOGRY = tractusx.RepoCategorySpecial

func DefaultMetadata() tractusx.Metadata {
	return tractusx.Metadata{
		ProductName:       "sig-infra",
		LeadingRepository: "https://github.com/eclipse-tractusx/sig-infra",
		RepoCategory:      DEFAULT_REPO_CATEOGRY,
		Repositories: []tractusx.Repository{
			{
				Name:             "sig-infra",
				UsageDescription: "Contains issue templates and workflows",
				Url:              "https://github.com/eclipse-tractusx/sig-infra",
			},
			{
				Name:             "charts",
				UsageDescription: "Central Helm repository containing all released charts of Tractus-X",
				Url:              "https://github.com/eclipse-tractusx/charts",
			},
		},
		SkipReleaseChecks: tractusx.SkipReleaseChecks{
			AlignedBaseImages: []string{
				"/path/to/non-published/Dockerfile",
			},
		},
		ConfigureReleaseChecks: tractusx.ConfigureReleaseChecks{},
	}
}

type Guideline struct {
	name                 string
	description          string
	externalDescription  string
	baseDir              string
	isOptional           bool
	metadataRepoCategory tractusx.RepoCategory
}

type FailingQualityGuideline struct {
	Guideline
}

func (f FailingQualityGuideline) IsApplicableToCategory(category tractusx.RepoCategory) bool {
	// fallback: return true if no category has been set
	if f.metadataRepoCategory == tractusx.RepoCategoryUnknown {
		return true
	}
	return f.metadataRepoCategory == category
}

func (f FailingQualityGuideline) IsOptional() bool {
	return f.isOptional
}

func (f FailingQualityGuideline) Name() string {
	return f.name
}

func (f FailingQualityGuideline) Description() string {
	return f.description
}

func (f FailingQualityGuideline) ExternalDescription() string {
	return f.externalDescription
}

func (p FailingQualityGuideline) BaseDir() string {
	return p.baseDir
}

func (f FailingQualityGuideline) Test() *tractusx.QualityResult {
	return &tractusx.QualityResult{Passed: false}
}

type PassingQualityGuideline struct {
	Guideline
}

func (p PassingQualityGuideline) IsApplicableToCategory(category tractusx.RepoCategory) bool {
	// fallback: return true if no category has been set
	if p.metadataRepoCategory == tractusx.RepoCategoryUnknown {
		return true
	}
	return p.metadataRepoCategory == category
}

func (p PassingQualityGuideline) IsOptional() bool {
	return p.isOptional
}

func (p PassingQualityGuideline) Name() string {
	return p.name
}

func (p PassingQualityGuideline) Description() string {
	return p.description
}

func (p PassingQualityGuideline) ExternalDescription() string {
	return p.externalDescription
}

func (p PassingQualityGuideline) BaseDir() string {
	return p.baseDir
}

func (p PassingQualityGuideline) Test() *tractusx.QualityResult {
	return &tractusx.QualityResult{Passed: true}
}

type PrinterMock struct {
	messages []string
}

func (p *PrinterMock) Print(message string) {
	fmt.Println(message)
	p.messages = append(p.messages, message)
}

func (p *PrinterMock) PrintTitle(title string) {
	p.Print(title)
}

func (p *PrinterMock) LogWarning(warning string) {
	p.Print(warning)
}

func (p *PrinterMock) LogError(err string) {
	p.Print(err)
}

func (p *PrinterMock) LogInfo(info string) {
	p.Print(info)
}

func (p *PrinterMock) LogSuccess(msg string) {
	p.Print(msg)
}
