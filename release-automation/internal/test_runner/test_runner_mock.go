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

package testrunner

import (
	"fmt"

	"tractusx-release-automation/internal/tractusx"
)

type Guideline struct {
	name                string
	description         string
	externalDescription string
	isOptional          bool
}

type FailingQualityGuideline struct {
	Guideline
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

func (f FailingQualityGuideline) Test() *tractusx.QualityResult {
	return &tractusx.QualityResult{Passed: false}
}

type PassingQualityGuideline struct {
	Guideline
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
