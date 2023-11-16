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

package dashboard

import "tractusx-release-automation/internal/tractusx"

type TemplateData struct {
	Config
	CheckedProducts []CheckedProduct
	UnhandledRepos  []Repository
	ArchivedRepos   []Repository
}

type CheckedProduct struct {
	Name                string
	LeadingRepo         string
	RepoCategory        string
	OverallPassed       bool
	CheckedRepositories []CheckedRepository
	ChartsDetails       []ChartDetails
}

type CheckedRepository struct {
	RepoName            string
	RepoUrl             string
	PassedAllGuidelines bool
	GuidelineChecks     []GuidelineCheck
}

type GuidelineCheck struct {
	GuidelineName    string
	GuidelineUrl     string
	Passed           bool
	Optional         bool
	ErrorDescription string
}

type Config struct {
	AssetsPath string
}

type Repository struct {
	Name, URL string
	Archived  bool
}

type Product struct {
	Name         string
	LeadingRepo  string
	RepoCategory string
	Repositories []Repository
}

type repoInfo struct {
	metadata tractusx.Metadata
	repoName string
	repoUrl  string
}

type ChartDetails struct {
	ChartName    string
	ChartVersion string
	AppVersion   string
}
