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

package container

import (
	"fmt"
	"os"
	"testing"
)

func TestShouldPassIfNoDockerfileFound(t *testing.T) {
	result := NewNonRootContainer("./").Test()

	if result == nil || result.Passed == false {
		t.Errorf("Non-Root Container Check should pass, if there is no Dockerfile found")
	}
}

// func TestValidateUser() validates the USER instruction in a Dockerfile to match lowercase or
func TestValidateUser(t *testing.T) {
	testCases := []struct {
		username         string
		groupname        string
		errorDescription string
		want             bool
	}{
		{"1000", "1010", "allowed USER instruction", true},
		{"username", "", "allowed USER instruction", true},
		{"65537", "123", "uid gt 65536 is not possible", false},
		{"65536", "65537", "gid gt 65536 is not possible", false},
		{"hiddenrootgrp", "0", "gid = 0 not allowed", false},
		{"user123", "123", "allowed USER instruction", true},
		{"USERNAME", "", "No upper case allowed", false},
		{"UserName", "", "No upper case allowed", false},
		{"userName", "", "No upper case allowed", false},
		//{NoUserFound, "", "empty user is invalid", false},
		{},
	}
	for _, tc := range testCases {
		result := validateUser(&user{tc.username, tc.groupname})
		if result != tc.want {
			t.Errorf("USER '%s':'%s' is not allowed: %s", tc.username, tc.groupname, tc.errorDescription)
		}
	}
}

func TestUserMethod(t *testing.T) {
	testCases := []struct {
		file *dockerfile
		want *user
	}{
		{correttoDockerfile(), &user{"corretto", ""}},                                              // USER corretto
		{temurinDockerfile(), &user{"temurin", ""}},                                                // USER temurin
		{newDockerfile().appendCommand("USER 0"), &user{"0", ""}},                                  // uid as int without gid defined
		{newDockerfile().appendCommand("USER root"), &user{"root", ""}},                            // uid as string without gid defined
		{newDockerfile().appendCommand("USER 0:1"), &user{"0", "1"}},                               // uid:gid as int
		{newDockerfile().appendCommand("USER 1000:0"), &user{"1000", "0"}},                         // uid:gid as int
		{newDockerfile().appendCommand("USER testuser:testgroup"), &user{"testuser", "testgroup"}}, // uid:gid as string
		//{newDockerfile().appendCommand("FROM distroless").appendEmptyLine(), &user{NoUserFound, ""}}, // No USER instruction in Dockerfile
	}

	for _, tc := range testCases {
		result := tc.file.user()

		if !tc.want.equals(result) {
			t.Errorf("expected user: '%s' group: '%s'", tc.want.user, tc.want.group)
		}
	}
}

func TestQualityCheckPass(t *testing.T) {
	testCases := []struct {
		file *dockerfile
		want bool
	}{
		{correttoDockerfile(), true},
		{temurinDockerfile(), true},
		{newDockerfile().appendCommand("FROM nginx").appendCommand("USER 102"), true},
		{newDockerfile().appendCommand("FROM nginx").appendCommand("USER 671234"), false},
		{newDockerfile().appendCommand("FROM nginx").appendCommand("USER 0"), false},
		//{newDockerfile().appendCommand("FROM nginx").appendEmptyLine(), false}, // No USER instruction in Dockerfile - not yet implemented
	}

	for _, tc := range testCases {
		if err := tc.file.writeTo("."); err != nil {
			fmt.Printf("could not write %s", tc.file.filename)
		}
		result := NewNonRootContainer("./").Test()

		if result.Passed != tc.want {
			t.Errorf("got '%t', expected '%t' as result", result.Passed, tc.want)
		}
		err := os.Remove(tc.file.filename)
		if err != nil {
			return
		}
	}
}

func correttoDockerfile() *dockerfile {
	return newDockerfile().
		appendCommand("FROM amazoncorreto:8").
		appendEmptyLine().
		appendCommand("USER corretto").
		appendCommand("COPY . .")
}

func temurinDockerfile() *dockerfile {
	return newDockerfile().
		appendCommand("FROM distroless").
		appendEmptyLine().
		appendCommand("COPY . .").
		appendCommand("FROM eclipse/temurin").
		appendCommand("USER temurin")
}
