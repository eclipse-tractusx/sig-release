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

import (
	"html/template"
	"io"
	"log"
	"path"
)

const templateDir = "web/templates"
const indexTemplate = "index.html.tmpl"

var allTemplates = []string{
	"index.html.tmpl",
	"header.html.tmpl",
	"body.html.tmpl",
	"footer.html.tmpl",
}

// RenderHtmlTo does par take an w io.Writer and
func RenderHtmlTo(w io.Writer, data *TemplateData) {
	templates := template.Must(template.New(indexTemplate).ParseFiles(allTemplatePaths()...))
	if err := templates.ExecuteTemplate(w, indexTemplate, data); err != nil {
		log.Fatalf("Could not execute template: %v", err)
	}
}

func allTemplatePaths() []string {
	var result []string
	for _, t := range allTemplates {
		result = append(result, path.Join(templateDir, t))
	}
	return result
}
