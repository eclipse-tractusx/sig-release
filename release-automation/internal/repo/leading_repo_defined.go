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

package repo

import (
	"strings"

	"tractusx-release-automation/internal/tractusx"
)

type LeadingRepositoryDefined struct {
	baseDir string
}

func NewLeadingRepositoryDefined(baseDir string) tractusx.QualityGuideline {
	return &LeadingRepositoryDefined{baseDir}
}

func (l *LeadingRepositoryDefined) Name() string {
	return "TRG 2.04 - Leading product repository"
}

func (l *LeadingRepositoryDefined) Description() string {
	return "The definition of a leading product repository makes it easy for all stakeholders to identify the entrypoint to a product"
}

func (l *LeadingRepositoryDefined) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-2/trg-2-4"
}

func (l *LeadingRepositoryDefined) IsOptional() bool {
	return false
}

func (l *LeadingRepositoryDefined) Test() *tractusx.QualityResult {
	metadata, err := tractusx.MetadataFromLocalFile(l.baseDir)
	if err != nil {
		return &tractusx.QualityResult{ErrorDescription: "Failed! The leadingRepository property must be defined in .tractusx metadata file. Could not load metadata"}
	}

	if strings.TrimSpace(metadata.LeadingRepository) == "" {
		return &tractusx.QualityResult{ErrorDescription: "Failed! The leadingRepository property must be defined in .tractusx metadata file"}
	}

	return &tractusx.QualityResult{Passed: true}
}
