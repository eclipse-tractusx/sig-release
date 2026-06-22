---
name: KIT Graduated Release Checklist
about: Checklist for a Graduated KIT
title: '[KIT-GRADUATED] <KIT Name> - Graduated Release Checklist'
labels: 'kit, graduated'
assignees: ''
---

# KIT Graduated Release Checklist

This issue tracks all relevant release checks for your KIT.

> [!IMPORTANT]  
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

**Graduation Date:** *date*

If changes to your KIT are not yet live on [our website](https://eclipse-tractusx. github.io/developer), link the open PR introducing the changes for this release.

If changes are already live, also point out the PRs which included the changes for this version if possible. 

- [PR number](pr-url)

For graduated KITs all the requirements are required to be fullfilled and maintained:

---

## Pre-Requirements

- [ ] **CHANGELOG** available and maintained - Updated with every release documenting all changes ([Keep a Changelog](https://keepachangelog.com/))

### KIT TRG 10
- [ ] [TRG 10.01 - KIT Architecture](https://eclipse-tractusx.github.io/docs/release/trg-10/trg-10-01) - Defines the overall architectural structure and patterns for KITs
- [ ] [TRG 10.02 - KIT Content](https://eclipse-tractusx.github.io/docs/release/trg-10/trg-10-02) - Specifies required content structure and organization
   - [ ] [TRG 7.07 - Legal notice for non-code](https://eclipse-tractusx. github.io/docs/release/trg-7/trg-7-07) - Image and media licensing requirements (CC BY 4. 0)
   - [ ] [TRG 7.08 - Legal notice for KIT documentation (CC-BY-4.0)](https://eclipse-tractusx. github.io/docs/release/trg-7/trg-7-08) - Documentation licensing requirements
- [ ] [TRG 10.03 - KIT Lifecycle](https://eclipse-tractusx.github.io/docs/release/trg-10/trg-10-03) - Describes lifecycle stages and transition criteria

### Master Data
- [ ] **Master Data Updated** (latest version & updated date) - Current version number and last update timestamp maintained, reviewed quarterly

---

## General Requirements

- [ ] **Changelog** - Maintained with every version release following [semantic versioning](https://semver.org/) with all breaking changes documented
- [ ] **Code Owner** - Active code owners responding to issues, regular maintenance and updates, community engagement active ([GitHub CODEOWNERS docs](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners))

## All Files

- [ ] **Notice** - Regular audits conducted, new files checked in CI/CD ([TRG 7.07](https://eclipse-tractusx. github.io/docs/release/trg-7/trg-7-07), [TRG 7.08](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-08))
- [ ] **KIT Icon** - Current and follows branding guidelines, optimized for web usage

## Adoption View

- [ ] **Vision / Mission** - Reviewed annually, updated to reflect current goals, aligned with product roadmap
- [ ] **Business Value(s)** - Validated with real metrics, updated with new benefits, case studies linked
- [ ] **Use Case / Domain explanation** - Current and comprehensive with new use cases added as discovered, industry examples updated
- [ ] **Business Processes** - Reflects current workflows, updated with process improvements, BPMN diagrams maintained
- [ ] **Tutorials / Videos** - All links functional, content updated for current version, new tutorials added as needed, community contributions welcome
- [ ] **Whitepaper** - Updated with latest developments, new research incorporated, versioned appropriately

## Development View

- [ ] **Architecture Overview** - Reflects current system state, updated with architecture changes, migration guides for breaking changes, ADRs maintained
- [ ] **Component/Sequence Diagrams** - Accurate and up-to-date with new components documented, generated from code where possible
- [ ] **Standards** - All standards current, deprecated standards removed, new standards adopted and documented, compliance maintained
- [ ] **API-Specification / Protocols** - Open API Specifications updated with releases, backward compatibility maintained, deprecation policy followed
- [ ] **Access & Usage Policies** - Policies updated to reflect current requirements, documented in all relevant locations
- [ ] **Logic / Schema** - Schema migrations documented, business logic changes tracked, data dictionaries current
- [ ] **Semantic Models / Data Structures** - SAMM models versioned and current, backward compatibility considered, migration paths documented, linked to semantic hub
- [ ] **Test Cases** - Test suite maintained (>80% coverage), all tests passing, performance benchmarks tracked, regression tests for bugs, CI/CD integration active

## Operation View

- [ ] **Deployment Guide** - Tested with each release, updated for new dependencies, Helm charts versioned, Docker images published, migration guides included
- [ ] **Operation/Monitoring Guidelines** - Monitoring dashboards maintained, new metrics added as needed, Grafana dashboards published, alert runbooks updated, SLI/SLO defined and tracked
- [ ] **Security Guidelines** - Regular security scans, vulnerabilities patched promptly, security advisories published, penetration test results addressed, SBOM maintained

## Success Stories

- [ ] **Reference Implementations** - List of active implementations maintained, new implementations added, success metrics tracked, community showcase updated, case studies published

## Documentation

- [ ] **Extra Documentation / Links** - All links checked quarterly, broken links fixed, FAQ updated with community questions, external resources curated, community wiki maintained
- [ ] **Sample Data** - Data updated for current schema, additional datasets added, privacy compliance verified, data generation scripts maintained, synthetic data generators available

## Industry Extensions

- [ ] **At least one Industry Extension** - Catena-X standards current, new versions adopted, industry-specific updates incorporated, compliance certifications renewed, extension marketplace listing updated

---

## Helpful links

- [KIT TRGs 10.XX](https://eclipse-tractusx.github.io/docs/category/trg-10---keep-it-together-kits)
- [KIT Office Hour](https://eclipse-tractusx.github.io/community/open-meetings#Eclipse%20Tractus-X%20KITs%20Community%20Office%20Hour)
- [KIT Template](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/tree/main/docs-kits/kit-template)