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
	"fmt"
	"os"
	"path"

	"github.com/eclipse-tractusx/tractusx-quality-checks/pkg/tractusx"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/engine"
	"k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

type ResourceMgmt struct {
	baseDir string
}

func NewResourceMgmt(baseDir string) tractusx.QualityGuideline {
	return &ResourceMgmt{baseDir}
}

func (r *ResourceMgmt) Name() string {
	return "TRG 5.04 - Resources Management"
}

func (r *ResourceMgmt) Description() string {
	return "The helm chart must contain app specific sane default CPU and Memory settings."
}

func (r *ResourceMgmt) ExternalDescription() string {
	return "https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-04"
}

func (r *ResourceMgmt) IsOptional() bool {
	return false
}

func (r *ResourceMgmt) Test() *tractusx.QualityResult {
	chartDir := path.Join(r.baseDir, "charts")
	if fi, err := os.Stat(chartDir); err != nil || !fi.IsDir() {
		return &tractusx.QualityResult{Passed: true}
	}

	helmCharts, err := os.ReadDir(chartDir)
	if err != nil || len(helmCharts) == 0 {
		return &tractusx.QualityResult{ErrorDescription: fmt.Sprintf("Can't read Helm Charts at %s.", chartDir)}
	}

	var errorDescription string
	for _, helmchart := range helmCharts {
		if !helmchart.IsDir() {
			continue
		}

		renderedChartManifests, errDesc := renderChart(path.Join(chartDir, helmchart.Name()))
		if renderedChartManifests == nil {
			errorDescription += errDesc
			continue
		}

		for manifestName, manifestContent := range renderedChartManifests {
			isValid, errMsg := validateResourceSetting(manifestContent)
			if !isValid {
				errorDescription += fmt.Sprintf("\n\t[%s]: %s", manifestName, errMsg)
				continue
			}
		}
	}

	if errorDescription != "" {
		return &tractusx.QualityResult{ErrorDescription: errorDescription}
	}
	return &tractusx.QualityResult{Passed: true}
}

func validateResourceSetting(k8sManifest string) (bool, string) {
	var containers []core.Container

	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, groupVersionKind, err := decode([]byte(k8sManifest), nil, nil)

	if err != nil {
		return true, ""
	}

	switch groupVersionKind.Kind {
	case "Deployment":
		containers = obj.(*v1.Deployment).Spec.Template.Spec.Containers
	case "StatefulSet":
		containers = obj.(*v1.StatefulSet).Spec.Template.Spec.Containers
	}

	for _, c := range containers {
		if c.Resources.Requests == nil {
			return false, "No resources requests found in the manifest."
		}
		if c.Resources.Requests.Cpu().IsZero() || c.Resources.Requests.Memory().IsZero() {
			return false, "CPU or Memory not defined in resources Requests."
		}
		if c.Resources.Limits == nil {
			return false, "No resources limits found in the manifest."
		}
		if c.Resources.Limits.Cpu().IsZero() || c.Resources.Limits.Memory().IsZero() {
			return false, "CPU or Memory not defined in resources Limits."
		}
		if c.Resources.Requests.Cpu().MilliValue() == c.Resources.Limits.Cpu().MilliValue() {
			return false, "Requested CPU can't be the same as Limit CPU. Limit should be 2-3 times higher."
		}
		if c.Resources.Requests.Memory().MilliValue() != c.Resources.Limits.Memory().MilliValue() {
			return false, "Requested Memory size must be equal to Limit Memory size."
		}
	}
	return true, ""
}

func renderChart(chartPath string) (map[string]string, string) {
	loadedChart, err := loader.Load(chartPath)
	if err != nil {
		return nil, fmt.Sprintf("\n\tCan't read %s helm chart.", chartPath)
	}

	finalValues := map[string]interface{}{
		"Values":  loadedChart.Values,
		"Release": map[string]string{"Namespace": "tractusx-check"},
	}

	renderedChart, err := engine.Render(loadedChart, finalValues)
	if err != nil {
		return nil, fmt.Sprintf("\n\tUnable to render helm chart %s.", chartPath)
	}
	return renderedChart, ""
}
