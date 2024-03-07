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

 package exception

 import (
	"testing"
)

func TestShouldPassIfParsesExceptionDataFile(t *testing.T) {
	config, _ := GetExceptionsFromFile("test/exceptions_test.yaml")
	for _,e := range config.Exceptions {
		if e.Trg == "1.23" {
			for _, r := range e.Repositories {
				if r == "testing-repo" {
					return
				}
			}
		}
	}
	t.Errorf("Couldn't parse testing exceptions file.")
}

func TestShouldPassIfExceptionExistForSpecificRepository(t *testing.T) {
	config, _ := GetExceptionsFromFile("test/exceptions_test.yaml")
	if !config.IsExceptioned("1.23", "testing-repo") {
		t.Errorf("Test should pass, test data contains the exception.")
	}
}

func TestShouldFailIfExceptionExistForSpecificRepository(t *testing.T) {
	config, _ := GetExceptionsFromFile("test/exceptions_test.yaml")
	if config.IsExceptioned("1.23", "no-exception-repo") {
		t.Errorf("Test should fail, there is no exception for testing repo.")
	}
}
