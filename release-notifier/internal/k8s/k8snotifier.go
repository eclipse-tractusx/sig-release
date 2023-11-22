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
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

 func GetLatestRel() string {
	var release []string
	website := "https://kubernetes.io/releases/"

	fmt.Println("Quering", website)
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Can't load the page: ", err)
	})

	c.OnHTML("span.release-inline-value", func(e *colly.HTMLElement) {
		release = append(release, e.Text)
	})

	c.Visit(website)
	return strings.Split(release[0], " ")[0]
}