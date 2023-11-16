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
	"errors"
	"os"
	"path"
	"strings"

	"tractusx-release-automation/internal/filesystem"
	"tractusx-release-automation/internal/tractusx"
)

var helmStructureFiles = []string{
	".helmignore",
	"LICENSE",
	"README.md",
	"values.yaml",
}

type HelmStructureExists struct {
	baseDir string
}

func NewHelmStructureExists(baseDir string) tractusx.QualityGuideline {
	return &HelmStructureExists{baseDir}
}

func (r *HelmStructureExists) Name() string {
	return "TRG 5.02 - Chart structure"
}

func (r *HelmStructureExists) Description() string {
	return "Helm Chart should follow a specific structure."
}

func (r *HelmStructureExists) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-02"
}

func (r *HelmStructureExists) IsOptional() bool {
	return false
}

func (r *HelmStructureExists) Test() *tractusx.QualityResult {
	chartDir := path.Join(r.baseDir, "charts")
	helmCharts, err := os.ReadDir(chartDir)
	if err != nil || len(helmCharts) == 0 {
		return &tractusx.QualityResult{Passed: true}
	}

	var missingFiles []string
	var chartYamlFiles []string
	for _, hc := range helmCharts {
		if !IsChartDirectory(path.Join(chartDir, hc.Name())) {
			continue
		}
		chartYamlFiles = append(chartYamlFiles, path.Join(chartDir, hc.Name(), "Chart.yaml"))
		missingFiles = getMissingChartFiles(path.Join(chartDir, hc.Name()))
	}

	errorDescriptionCharts := ""
	chartsValid := true
	if len(chartYamlFiles) > 0 {
		for _, fpath := range chartYamlFiles {
			errorDescriptionCharts = "\n\t+ Analysis for " + strings.Split(fpath, "charts")[1][1:] + ": "
			err = validateChart(fpath)
			if err != nil {
				errorDescriptionCharts += err.Error()
				chartsValid = false
			}
		}
	}

	if len(missingFiles) > 0 || !chartsValid {
		errMsg := "+ Following Helm Chart structure files are missing: " + strings.Join(missingFiles, ", ") + errorDescriptionCharts
		if tractusx.ErrorOutputFormat == tractusx.WebErrOutputFormat {
			return &tractusx.QualityResult{ErrorDescription: strings.ReplaceAll(errMsg, "\n", "<br>")}
		}
		return &tractusx.QualityResult{ErrorDescription: errMsg}
	}
	return &tractusx.QualityResult{Passed: true}
}

// IsChartDirectory evaluates, if the given directory can be seen as helm chart directory
func IsChartDirectory(dir string) bool {
	chartYamlPath := path.Join(dir, "Chart.yaml")
	_, err := os.Stat(chartYamlPath)
	return err == nil
}

func getMissingChartFiles(chartPath string) []string {
	var result []string
	for _, fileToCheck := range helmStructureFiles {
		missingFile := filesystem.CheckMissingFiles([]string{path.Join(chartPath, fileToCheck)})
		if missingFile != nil {
			result = append(result, []string{strings.Split(missingFile[0], "charts")[1][1:]}...)
		}
	}
	return result
}

func validateChart(chartYamlFile string) error {
	errorMessage := ""

	cyf := ChartYamlFromFile(chartYamlFile)
	missingFields := cyf.GetMissingMandatoryFields()

	if len(missingFields) > 0 {
		errorMessage += "\n\t\t - Missing mandatory fields: " + strings.Join(missingFields, ", ")
	}
	if !cyf.IsVersionValid() {
		errorMessage += "\n\t\t - " + cyf.Version + " Version of the Helm Chart is incorrect. It needs to follow Semantic Version."
	}

	if errorMessage != "" {
		return errors.New(errorMessage)
	}
	return nil
}
