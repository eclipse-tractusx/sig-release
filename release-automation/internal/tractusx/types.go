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

package tractusx

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

// ErrorOutputFormat is used to control the output format of error message.
// Before running QualityGuideline.Test(), set the value to your preferred output format
var ErrorOutputFormat = CliErrOutputFormat

const (
	CliErrOutputFormat = iota
	WebErrOutputFormat
)

// QualityGuideline represents the QualityGuideline to check as an interface.
//
// The interface provide information about Name, Description, ExternalDescription,
// the Test and the IsOptional bool.
type QualityGuideline interface {
	// Name returns the Name to the QualityGuideline to test.
	Name() string
	// Description returns a brief description of the tested QualityGuideline.
	Description() string
	// ExternalDescription returns a URL to more information about the QualityGuideline.
	ExternalDescription() string
	// BaseDir returns repository local path.
	BaseDir() string
	// Test executes the test and returns QualityResult. If any error occurs it
	// returns QualityResult.Passed false.
	Test() *QualityResult
	// IsOptional returns a bool it the test or QualityGuideline is optional or not.
	IsOptional() bool
	// IsApplicableToCategory returns true if the QualityGuideline must be applied to the RepoCategory
	IsApplicableToCategory(category RepoCategory) bool
}

// QualityResult implements test result via Passed bool and in case of error a
// ErrorDescription.
type QualityResult struct {
	Passed           bool
	ErrorDescription string
}

type Printer interface {
	Print(message string)
	LogWarning(warning string)
	LogError(err string)
}

type StdoutPrinter struct {
}

func (p *StdoutPrinter) Print(message string) {
	fmt.Println(message)
}

func (p *StdoutPrinter) LogWarning(warning string) {
	color.Yellow(warning)
}

func (p *StdoutPrinter) LogError(err string) {
	color.Red(err)
}

type RepoCategory int

const (
	RepoCategoryUnknown RepoCategory = iota
	RepoCategoryProduct
	RepoCategorySpecial
	RepoCategorySupport
	// Add more categories as needed
)

// String method for better debugging/logging
func (rc RepoCategory) String() string {
	switch rc {
	case RepoCategoryProduct:
		return "product"
	case RepoCategorySpecial:
		return "special"
	case RepoCategorySupport:
		return "support"
	default:
		return "unknown"
	}
}

// Implement YAML marshaling
func (rc RepoCategory) MarshalYAML() (interface{}, error) {
	return rc.String(), nil
}

func (rc *RepoCategory) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	*rc = ParseRepoCategory(s)
	return nil
}

// ParseRepoCategory converts string to RepoCategory
func ParseRepoCategory(s string) RepoCategory {
	switch strings.ToLower(s) {
	case "product":
		return RepoCategoryProduct
	case "special":
		return RepoCategorySpecial
	case "support":
		return RepoCategorySupport
	default:
		return RepoCategoryUnknown
	}
}
