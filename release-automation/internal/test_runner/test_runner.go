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
	"errors"
	"fmt"
	"log"
	"tractusx-release-automation/internal/exception"
	"tractusx-release-automation/internal/repo"
	"tractusx-release-automation/internal/tractusx"

	"github.com/fatih/color"
)

type GuidelineTestRunner struct {
	guidelines []tractusx.QualityGuideline
	printer    tractusx.Printer
	metadata   tractusx.Metadata
}

func NewTestRunner(tests []tractusx.QualityGuideline, metadata tractusx.Metadata) *GuidelineTestRunner {
	return &GuidelineTestRunner{guidelines: tests, printer: &StdoutPrinter{}, metadata: metadata}
}

func (runner *GuidelineTestRunner) Run() error {
	var result *tractusx.QualityResult
	allPassed := true
	for i, guideline := range runner.guidelines {
		runner.printer.Print(fmt.Sprintf("\n%v. Testing Quality Guideline: %s", i+1, guideline.Name()))

		m, err := exception.GetData()
		exceptionPresent := false
		if err != nil {
			log.Println("Can't process exceptions.")
		} else {
			repoInfo := repo.GetRepoBaseInfo(guideline.BaseDir())
			// Exclude if
			// - exception configured for specific repository for specific TRG
			// - guideline does not apply to present repository
			if m.IsExceptioned(guideline.Name(), "https://github.com/eclipse-tractusx/"+repoInfo.Reponame) {
				runner.printer.Print(fmt.Sprintf("Test skippped. Exception exists for test '%s' and repo '%s'.", guideline.Name(), repoInfo.Reponame))
				result = &tractusx.QualityResult{Passed: true}
				exceptionPresent = true
			} else if !guideline.IsApplicableToCategory(runner.metadata.RepoCategory) {
				runner.printer.Print(fmt.Sprintf("Test skippped. Test '%s' is not applicable to repo category '%s'.", guideline.Name(), repoInfo.Reponame))
				result = &tractusx.QualityResult{Passed: true}
				exceptionPresent = true
			}
		}

		if !exceptionPresent {
			result = guideline.Test()
		}
		if guideline.IsOptional() && !result.Passed {
			runner.printer.LogWarning(
				fmt.Sprintf("Warning! Test failed, but test '%s' is marked as optional.\nMore infos:\n\t%s\n\t%s",
					guideline.Name(), result.ErrorDescription, guideline.ExternalDescription()),
			)
		} else if !result.Passed {
			runner.printer.LogError(
				fmt.Sprintf("Failed! Guideline description: %s\n\t%s\n\tMore infos: %s",
					guideline.Description(), result.ErrorDescription, guideline.ExternalDescription()),
			)
		} else {
			runner.printer.Print("Skipped or passed!")
		}

		allPassed = allPassed && (result.Passed || guideline.IsOptional())
	}

	if !allPassed {
		return errors.New("not all tests have passed")
	}

	return nil
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
