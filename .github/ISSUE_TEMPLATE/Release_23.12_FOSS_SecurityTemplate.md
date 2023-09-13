<!---
 * Copyright (c) 2021,2023 Contributors to the Eclipse Foundation
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
--->


# Release Compliance 23.12
Source in Catena-X Confluence + Expert Contacts [here](https://confluence.catena-x.net/x/DOZkBQ).

**Threat Modelling Analysis results**
Analysis completed (operations excluded):
> List of risks generated or updated, rated & actions defined
> Risks accepted or mitigation actions implemented and tested
> no high threats acceptable


**Artifact Repository:**
risk register
(decentral on Catena-X confluence)

**Prime Contacts:**
Security Team: SEC0

**Static Application Security Testing (SAST)**
>code must be scanned weekly with Veracode tool
>medium risks require mitigation statement
>high and above not accepted

**Best Practise:**
Confirm relevant repository as early as possible to SEC team to enable regular, automated scans. Evidence required for Gate approval.

**Artifact Repository:**
Veracode UI
(+ GitHub Action)

**Prime Contacts:**
Security Team: SEC1


**Dynamic Application Security Testing (DAST)**
incl API testing (if applicable)
=> all findings assessed
=> high & very high findings mitigated
=> evidence by re-scan

**Best Practise:**
Confirm relevant repository as early as possible to SEC team to enable regular, automated scans. Evidence required for Gate approval.

**Artifact Repository:**
INVICTI tool

**Prime Contacts:**
Security Team: SEC3 SEC4

**Secret scanning**

Scan executed centrally by SEC team
>ZERO valid findings

**Artifact Repository:**
Veracode or alternative tool
GitHub Secret Scanning
GitGuardian

Confirm relevant repository as early as possible to SEC team to enable regular, automated scans. Evidence required for Gate approval.

**Prime Contact:** SEC1

**Software Composition Analysis (SCA)**
Dependencies must be scanned with Veracode tool with regards to vulnerability
>high and above not accepted
>FOSS whitelist policy has to be passed

**Best Practise:**
Confirm relevant repository as early as possible to SEC team to enable regular, automated scans. Evidence required for Gate approval.

**Artifact Repository:**
Veracode UI
(& GitHub Action)

**Prime Contacts:**
Security Team: SEC1

**Container Scan conducted**
All containers in GitHub Packages must be scanned
>High / Critical findings not accepted

**Best Practise:**
Confirm relevant repository as early as possible to SEC team to enable regular, automated scans. Evidence required for Gate approval.

**Artifact Repository:**
Trivy
via nightly GitHub Action

**Prime Contacts:**
Security Team: SEC2

**Infrastructure as Code**
	
IaC code must be scanned
>Error findings not accepted

****Artifact Repository:****
KICS or alternative tool
via nightly GitHub Action

Confirm relevant repository as early as possible to SEC team to enable regular, automated scans. Evidence required for Gate approval.

**Prime Contacts:**
Security Team: SEC2