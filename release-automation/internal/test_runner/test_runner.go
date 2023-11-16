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
	"errors"
	"fmt"

	"github.com/fatih/color"
	"tractusx-release-automation/internal/tractusx"
)

type GuidelineTestRunner struct {
	guidelines []tractusx.QualityGuideline
	printer    tractusx.Printer
}

func NewTestRunner(tests []tractusx.QualityGuideline) *GuidelineTestRunner {
	return &GuidelineTestRunner{guidelines: tests, printer: &StdoutPrinter{}}
}

func (runner *GuidelineTestRunner) Run() error {
	allPassed := true
	for i, guideline := range runner.guidelines {
		runner.printer.Print(fmt.Sprintf("\n%v. Testing Quality Guideline: %s", i+1, guideline.Name()))

		result := guideline.Test()
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
