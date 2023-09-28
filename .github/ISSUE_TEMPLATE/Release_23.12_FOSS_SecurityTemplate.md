---
name: Release Security 23.12
about: Use this template to track all Security-relevant topics for your component with regards to the upcoming Milestone.
title: "[FOSS NAME] Release Security 23.12"
labels: "security analysis"
assignees: kelaja
---

<!-- 
Thanks for your contribution! Please fill out this template as good as possible. 
Important: Contributing Guidelines can be found here: https://eclipse-tractusx.github.io/docs/oss/how-to-contribute
Checkout the repository README for process description. 
-->

# Release Security 23.12

Source in Catena-X Confluence and Expert Contacts [here](https://confluence.catena-x.net/x/DOZkBQ) (Source only accessible for Catena-X Consortia members in current transition phase).

- [ ] **Threat Modelling Analysis results**
  
    - Analysis completed (operations excluded):
    - List of risks generated or updated, rated & actions defined
    - Risks accepted or mitigation actions implemented and tested
    - No high threats acceptable

  _Artifact Repository:_
  
    - Risk Register (decentral on Catena-X confluence)

  _Contact:_
  
    - Security Team: SEC0

- [ ] **Static Application Security Testing (SAST)**

  - Code must be scanned weekly with Veracode tool
  - Medium risks require mitigation statement
  - High and above not accepted

  _Best Practise:_
  
    - Confirm relevant repository as early as possible to SEC team to enable regular, automated scans. Evidence required for Gate approval.

  _Artifact Repository:_
  
    - Veracode UI
    - (+ GitHub Action)

  _Contact:_
  
    - Security Team: SEC1

- [ ] **Dynamic Application Security Testing (DAST)**
  incl API testing (if applicable)
  - all findings assessed
  - high & very high findings mitigated
  - evidence by re-scan

  _Best Practise:_
  
    - Confirm relevant repository as early as possible to SEC team to enable regular, automated scans. Evidence required for Gate approval.

  _Artifact Repository:_
  
    - INVICTI tool

  _Prime Contacts:_
  
    - Security Team: SEC3 SEC4

- [ ] **Secret scanning**
  Scan executed centrally by SEC team and ZERO valid findings
  
  _Artifact Repository:_
  
    - Veracode or alternative tool
    - GitHub Secret Scanning
    - GitGuardian

  _Best Practise:_
  
    - Confirm relevant repository as early as possible to SEC team to enable regular, automated scans. Evidence required for Gate approval.

   _Contact:_
  
     - Security Team: SEC1

- [ ] **Software Composition Analysis (SCA)**
  Dependencies must be scanned with Veracode tool with regards to vulnerability
    - high and above not accepted
    - FOSS whitelist policy has to be passed

  _Best Practise:_
  
    - Confirm relevant repository as early as possible to SEC team to enable regular, automated scans. Evidence required for Gate approval.

  _Artifact Repository:_
  
    - Veracode UI
    - (& GitHub Action)

  _Contact:_
  
    - Security Team: SEC1

- [ ] **Container Scan conducted**
  All containers in GitHub Packages must be scanned
  
    - High / Critical findings not accepted

  _Best Practise:_
  
    - Confirm relevant repository as early as possible to SEC team to enable regular, automated scans. Evidence required for Gate approval.

  _Artifact Repository:_
  
    - Trivy
    - via nightly GitHub Action

  _Contact:_
  
    - Security Team: SEC2

- [ ] **Infrastructure as Code**
  IaC code must be scanned. 
    - Error findings not accepted

   _Best Practise:_
  
    - Confirm relevant repository as early as possible to SEC team to enable regular, automated scans. Evidence required for Gate approval.

  _Artifact Repository:_
  
    - KICS or alternative tool
    - via nightly GitHub Action

  _Prime Contacts:_
  
    - Security Team: SEC2
