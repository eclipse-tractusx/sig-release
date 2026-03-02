---
name: KIT Incubating In Progress Release Checklist
about: Checklist for a KIT in Incubating In Progress stage
title: '[KIT-IN-PROGRESS] <KIT Name> - Incubating In Progress Release Checklist'
labels: 'kit, incubating, in-progress'
assignees: ''
---

# KIT Incubating In Progress Release Checklist

This issue tracks all relevant release checks for your KIT.

> [! IMPORTANT]  
> Follow the [guidance on how to use the templates](https://github.com/eclipse-tractusx/sig-release/blob/main/README.md#release-management-acceptance-criteria).

## Release Info

Please provide information on what you want to be included in the Eclipse Tractus-X release.
If you are not owner of this issue, please provide the information as comment to the issue.
Have your prepared content reviewed by a professional peer group before submitting it for committer approval.
Best practice: have a professional Peer review your prepared content before you submit for expert approvals. 
Please link the open or merged PR that includes the changes.

**KIT Name:** *name placeholder*

**Version to be included in Eclipse Tractus-X release:** *version placeholder*
- KIT semantic version of change log.
- Date of KIT publication must be aligned with overall release date.

If changes to your KIT are not yet live on [our website](https://eclipse-tractusx.github.io/developer), link the open PR introducing the changes for this release.

If changes are already live, also point out the PRs which included the changes for this version if possible.

- [PR number](pr-url)

---

## Pre-Requirements

- [ ] **CHANGELOG** available and maintained - A file documenting all notable changes, versions, and release dates ([Keep a Changelog](https://keepachangelog.com/))

### KIT TRG 10
- [ ] [TRG 10. 01 - KIT Architecture](https://eclipse-tractusx.github.io/docs/release/trg-10/trg-10-01) - Defines the overall architectural structure and patterns for KITs
- [ ] [TRG 10.02 - KIT Content](https://eclipse-tractusx.github.io/docs/release/trg-10/trg-10-02) - Specifies required content structure and organization
   - [ ] [TRG 7.07 - Legal notice for non-code](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-07) - Image and media licensing requirements (CC BY 4.0)
   - [ ] [TRG 7.08 - Legal notice for KIT documentation (CC-BY-4.0)](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-08) - Documentation licensing requirements
- [ ] [TRG 10. 03 - KIT Lifecycle](https://eclipse-tractusx. github.io/docs/release/trg-10/trg-10-03) - Describes lifecycle stages and transition criteria

### Master Data
- [ ] **Master Data Updated** (latest version & updated date) - Current version number and last update timestamp maintained

---

## General Requirements

- [ ] **Changelog** - Version history file (CHANGELOG.md) in repository root following semantic versioning
- [ ] **Code Owner** ⚠️ **REQUIRED** - CODEOWNERS file with designated maintainers for review and approval ([GitHub CODEOWNERS docs](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners))

## All Files

- [ ] **Notice** - Legal and compliance notice in all source, documentation, and configuration files ([TRG 7.07](https://eclipse-tractusx. github.io/docs/release/trg-7/trg-7-07), [TRG 7.08](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-08))
- [ ] **KIT Icon** - SVG icon stored at `/static/img/kits/<kit-id>/<kit-id>-kit.svg` following design guidelines and branding

## Adoption View

- [ ] **Vision / Mission** - Clear, peer-reviewed description of KIT's purpose and long-term goals
- [ ] **Business Value(s)** - Concrete benefits with quantifiable metrics, clear ROI, and defined target audience
- [ ] **Use Case / Domain explanation** - Complete real-world scenarios with examples, diagrams, and applicable industries
- [ ] **Business Processes** ⚠️ **REQUIRED** - Business workflows with BPMN 2.0 diagrams (may not be relevant for Network Services)
- [ ] **Tutorials / Videos** ⚠️ **REQUIRED** - At least 2-3 comprehensive tutorials with video demonstrations, step-by-step guides with screenshots, and hands-on exercises
- [ ] **Whitepaper** ⚠️ **REQUIRED** - Complete technical/business analysis with research findings, case studies, and peer review for quality

## Development View

- [ ] **Architecture Overview** ⚠️ **REQUIRED** - Complete system design with C4 model diagrams, integration points, dependencies, reviewed by architect
- [ ] **Component/Sequence Diagrams** ⚠️ **REQUIRED** - All major interactions documented in UML 2.0 or PlantUML format with sequence diagrams for key workflows
- [ ] **Standards** ⚠️ **REQUIRED** - All relevant standards documented, linked, and compliance verified with version numbers specified
- [ ] **API-Specification / Protocols** ⚠️ **REQUIRED** - Complete Open API Specifications with all endpoints, examples, error codes
- [ ] **Access & Usage Policies** ⚠️ **REQUIRED** - If DSP is used (via Connector), the policies which are required
- [ ] **Logic / Schema** ⚠️ **REQUIRED** - Complete business logic documentation, database schemas with ERD diagrams, data flows, and validation rules
- [ ] **Semantic Models / Data Structures** ⚠️ **REQUIRED** - SAMM definitions, JSON Schema, linked to standard ontologies, version-controlled and published
- [ ] **Test Cases** ⚠️ **REQUIRED** - Comprehensive coverage (>80%) including unit, integration, E2E, performance, and security tests with documented scenarios

## Operation View

- [ ] **Deployment Guide** ⚠️ **REQUIRED** - Complete step-by-step instructions with Docker Compose, Kubernetes/Helm charts, environment config, troubleshooting, tested on clean environment
- [ ] **Operation/Monitoring Guidelines** *(recommended)* - Prometheus metrics, Grafana dashboards, log aggregation (ELK/Loki), alerting rules, backup/disaster recovery
- [ ] **Security Guidelines** *(recommended)* - OWASP compliance, TLS/SSL config, OAuth2/OIDC setup, secrets management (Vault), security scanning, incident response

## Success Stories

- [ ] **Reference Implementations** ⚠️ **REQUIRED** - At least one production/pilot implementation with working code on GitHub, deployment docs, contact info, success metrics and testimonials

## Documentation

- [ ] **Extra Documentation / Links** *(recommended)* - Comprehensive FAQ, troubleshooting guides, community forum/Slack/Discord links, related KITs and integrations
- [ ] **Sample Data** ⚠️ **REQUIRED** - Representative datasets in multiple formats (JSON, CSV, XML), sufficient for testing, GDPR/privacy compliant, with data dictionary

## Industry Extensions

- [ ] **At least one Industry Extension** ⚠️ **REQUIRED** - Catena-X standard version specified, industry data models, compliance verified, standards submitted/published, implementation examples

---

## Helpful links

- [KIT TRGs 10.XX](https://eclipse-tractusx.github.io/docs/category/trg-10---keep-it-together-kits)
- [KIT Office Hour](https://eclipse-tractusx.github.io/community/open-meetings#Eclipse%20Tractus-X%20KITs%20Community%20Office%20Hour)
- [KIT Template](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/tree/main/docs-kits/kit-template)