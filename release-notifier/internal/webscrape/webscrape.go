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

package webscrape

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

// Semantic Versioning schema regex
const regexPattern = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`

func GetLatestPSQLRelease() string {
	var release string
	website := "https://bitnami.com/stack/postgresql"

	fmt.Println("Quering", website)
	c := colly.NewCollector()

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Can't load the page: ", err)
	})

	c.OnHTML("div.stack__header__properties__property", func(e *colly.HTMLElement) {
		if match, _ := regexp.MatchString(regexPattern, e.ChildText("p")); match {
			release = e.ChildText("p")
		}
	})

	c.Visit(website)
	return release
}
