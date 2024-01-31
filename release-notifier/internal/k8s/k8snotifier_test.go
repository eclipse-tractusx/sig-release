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

package k8s

import (
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestShouldPassIfNewReleaseIsSemVer(t *testing.T) {
	// Semantic Versioning schema regex
	const regexPattern = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

	newRelease := getLatestRel()
	if match, _ := regexp.MatchString(regexPattern, newRelease); !match {
		t.Errorf("Test should pass, it returns semver string.")
	}

}

func TestShouldPassIfNewReleaseGetsPersisted(t *testing.T) {
	fakeRelease := "0.0.1"
	_ = os.WriteFile(artifactName, []byte("0.0.1"), 0644)
	defer os.Remove(artifactName)
	IsNewRelease()
	newRelease := getRelFromArtifact()
	if strings.EqualFold(newRelease, fakeRelease) {
		t.Errorf("Test should pass, the new release should not be equal to fake.")
	}
}

func TestShouldPassIfEmailContainsRefReleases(t *testing.T) {
	fakeRelease, fakeAligned := "0.0.2", "0.0.1"
	_ = os.WriteFile(artifactName, []byte(fakeRelease), 0644)
	defer os.Remove(artifactName)
	os.Setenv("CURRENT_ALIGNED_K8S_VER", fakeAligned)
	mailContent, err := buildContent("../../templates/mail-k8s.html.tmpl")
	if err != nil {
		println(err)
	}
	if !strings.Contains(string(mailContent), fakeRelease) || !strings.Contains(string(mailContent), fakeAligned) {
		t.Errorf("Test should pass, content was given new and aligned releases.")
	}
}

