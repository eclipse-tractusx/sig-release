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
	"strings"
	"testing"

	"tractusx-release-automation/internal/tractusx"
)

func TestShouldPassIfNotTestsAreRun(t *testing.T) {
	err := NewTestRunner([]tractusx.QualityGuideline{}, DefaultMetadata()).Run()

	if err != nil {
		t.Errorf("Running no tests should succeed")
	}
}

func TestShouldFailWhenRunningFailingQualityTage(t *testing.T) {
	err := NewTestRunner([]tractusx.QualityGuideline{&FailingQualityGuideline{}}, DefaultMetadata()).Run()

	if err == nil {
		t.Errorf("Expected error if only a failing quality guideline is run")
	}
}

func TestShouldPassIfOnlyPassingQualityGuidelinesAreRun(t *testing.T) {
	err := NewTestRunner([]tractusx.QualityGuideline{&PassingQualityGuideline{}}, DefaultMetadata()).Run()

	if err != nil {
		t.Errorf("Should not raise an error if only succeeding quality guidelines are run")
	}
}

func TestShouldFailIfAtLeastOneFailingTestIsRun(t *testing.T) {
	runner := NewTestRunner([]tractusx.QualityGuideline{
		&PassingQualityGuideline{},
		&PassingQualityGuideline{},
		&FailingQualityGuideline{},
		&PassingQualityGuideline{},
	}, DefaultMetadata())
	err := runner.Run()

	if err == nil {
		t.Errorf("Expected error, if at least one of the quality guideline is failing")
	}
}

func TestShouldLogTheGuidelineNameWhenRunningTheTest(t *testing.T) {
	runner := NewTestRunner(
		[]tractusx.QualityGuideline{&PassingQualityGuideline{Guideline{name: "TRG 1 - README test"}}},
		DefaultMetadata(),
	)
	printerMock := &PrinterMock{}
	runner.printer = printerMock

	_ = runner.Run()

	if len(printerMock.messages) < 1 {
		t.Errorf("Expected at least one printed message when guidelines are tested, but got none")
	}

	if !strings.Contains(printerMock.messages[0], "TRG 1 - README test") {
		t.Errorf("Expected the printed messages to contain the guideline name, but printed was %s", printerMock.messages[0])
	}
}

func TestShouldLogDescriptionOfGuidelineIfTheTestIsFailing(t *testing.T) {
	failingGuideline := &FailingQualityGuideline{
		Guideline{
			name:                "TRG 3000 - ChuckNorris",
			description:         "Every project should define counter measures against roundhouse kicks",
			externalDescription: "https://en.wikipedia.org/wiki/Chuck_Norris",
		},
	}
	runner := NewTestRunner([]tractusx.QualityGuideline{failingGuideline}, DefaultMetadata())
	printerMock := &PrinterMock{}
	runner.printer = printerMock

	_ = runner.Run()

	if len(printerMock.messages) < 2 {
		t.Errorf("On failing guidelines test, at least a message for the test run and a description of it should be printed")
	}

	if !strings.Contains(printerMock.messages[0], "TRG 3000 - ChuckNorris") {
		t.Errorf("Did not print guideline name when testing")
	}

	if !strings.Contains(printerMock.messages[1], "Failed! ") ||
		!strings.Contains(printerMock.messages[1], failingGuideline.description) ||
		!strings.Contains(printerMock.messages[1], failingGuideline.externalDescription) {
		t.Errorf("Runner should print description and external description when a quality guideline check fails. Did print: \n%s", allMessages(printerMock.messages))
	}
}

func TestShouldOnlyLogAdditionalDescriptionForFailingTests(t *testing.T) {
	failingGuideline := &FailingQualityGuideline{
		Guideline{
			name:                "TRG 3000 - ChuckNorris",
			description:         "Every project should define counter measures against roundhouse kicks",
			externalDescription: "https://en.wikipedia.org/wiki/Chuck_Norris",
		},
	}
	passingGuideline := &PassingQualityGuideline{
		Guideline{
			name:                "TRG 4711 - auto-pass",
			description:         "automatically passing. result ignored",
			externalDescription: "https://de.wikipedia.org/wiki//dev/null",
		},
	}
	runner := NewTestRunner([]tractusx.QualityGuideline{failingGuideline, passingGuideline}, DefaultMetadata())
	printerMock := &PrinterMock{}
	runner.printer = printerMock

	_ = runner.Run()

	if len(printerMock.messages) != 4 {
		t.Errorf("Expected exactly 4 logged messages. \n1. Name of guideline; 2. Additional info for failing; 3 Name of passing guideline. Got %d messages", len(printerMock.messages))
	}
}

func TestShouldOnlyRunForMatchingRepoCategory(t *testing.T) {
	executedGuideline := &PassingQualityGuideline{
		Guideline{
			name:                 "TRG 4711 - Applicable as Repo Type is 'special' matching to default metadata in test",
			metadataRepoCategory: tractusx.RepoCategorySpecial,
		},
	}
	excludedGuideline := &PassingQualityGuideline{
		Guideline{
			name:                 "TRG 3000 - Not applicable as Repo Type is 'support' NOT matching to default metadata in test",
			metadataRepoCategory: tractusx.RepoCategorySupport,
		},
	}
	runner := NewTestRunner([]tractusx.QualityGuideline{executedGuideline, excludedGuideline}, DefaultMetadata())
	printerMock := &PrinterMock{}
	runner.printer = printerMock

	_ = runner.Run()

	if len(printerMock.messages) != 5 {
		t.Errorf("Expected exactly 5 logged messages. \n1. Name of guideline; 2. Skipped or passed! 3. Name of guideline 4. Test skippped. Test 's' is not applicable to repo category 's'.; 5. Skipped or passed!  - BUT got %d messages", len(printerMock.messages))
	}
}

func allMessages(messages []string) string {
	var result string
	for _, m := range messages {
		result += fmt.Sprintln(m)
	}

	return result
}

func TestShouldNotFailIfOptionalTestFail(t *testing.T) {
	failingGuideline := &FailingQualityGuideline{
		Guideline{
			name:                "TRG 1.02 - INSTALL.md",
			description:         "Optional content should not fail a test",
			externalDescription: "https://github.com/",
			isOptional:          true,
		},
	}

	runner := NewTestRunner([]tractusx.QualityGuideline{failingGuideline}, DefaultMetadata())
	err := runner.Run()

	if err != nil {
		t.Errorf("Optional Guidlines should not make the runner fail!")
	}
}
