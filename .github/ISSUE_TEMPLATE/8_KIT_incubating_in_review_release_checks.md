---
name: KIT Incubating In Review Release Checklist
about: Checklist for a KIT in Incubating In Review stage
title: '[KIT-REVIEW] <KIT Name> - Incubating In Review Release Checklist'
labels: 'kit, incubating, in-review'
assignees: ''
---

# KIT Incubating In Review Release Checklist

This issue tracks all relevant release checks for your KIT.

> [!IMPORTANT]  
> Follow the [guidance on how to use the templates](https://github.com/eclipse-tractusx/sig-release/blob/main/README. md#release-management-acceptance-criteria).

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

**Review Committee:** *reviewer names*

If changes to your KIT are not yet live on [our website](https://eclipse-tractusx.github. io/developer), link the open PR introducing the changes for this release. 

If changes are already live, also point out the PRs which included the changes for this version if possible.

- [PR number](pr-url)

---

## Pre-Requirements

- [ ] **CHANGELOG** available and maintained - A file documenting all notable changes, versions, and release dates ([Keep a Changelog](https://keepachangelog.com/))

### KIT TRG 10
- [ ] [TRG 10.01 - KIT Architecture](https://eclipse-tractusx.github.io/docs/release/trg-10/trg-10-01) - Defines the overall architectural structure and patterns for KITs
- [ ] [TRG 10.02 - KIT Content](https://eclipse-tractusx.github. io/docs/release/trg-10/trg-10-02) - Specifies required content structure and organization
   - [ ] [TRG 7.07 - Legal notice for non-code](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-07) - Image and media licensing requirements (CC BY 4.0)
   - [ ] [TRG 7.08 - Legal notice for KIT documentation (CC-BY-4.0)](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-08) - Documentation licensing requirements
- [ ] [TRG 10.03 - KIT Lifecycle](https://eclipse-tractusx.github.io/docs/release/trg-10/trg-10-03) - Describes lifecycle stages and transition criteria

### Master Data
- [ ] **Master Data Updated** (latest version & updated date) - Current version number and last update timestamp maintained

---

## General Requirements

- [ ] **Changelog** - Complete version history with all releases documented using semantic versioning
- [ ] **Code Owner** - Active and responsive code owners with regular review and approval activity ([GitHub CODEOWNERS docs](https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/about-code-owners))

## All Files

- [ ] **Notice** - All files audited for compliance with copyright and licensing headers ([TRG 7.07](https://eclipse-tractusx. github.io/docs/release/trg-7/trg-7-07), [TRG 7.08](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-08))
- [ ] **KIT Icon** - SVG icon meets all design guidelines, properly versioned and optimized

## Adoption View

- [ ] **Vision / Mission** 🔍 *Under Review* - Clear, compelling, well-articulated, and aligned with organizational strategy
- [ ] **Business Value(s)** 🔍 *Under Review* - Quantifiable and demonstrable benefits with ROI clearly stated and target audience well-defined
- [ ] **Use Case / Domain explanation** 🔍 *Under Review* - Comprehensive with real-world examples applicable across target industries with technical accuracy verified
- [ ] **Business Processes** 🔍 *Under Review* - Complete BPMN diagrams with clear process flows and stakeholder roles defined (may not be relevant for Network Services)
- [ ] **Tutorials / Videos** 🔍 *Under Review* - Sufficient coverage of all features, easy to follow, technically accurate with working code examples
- [ ] **Whitepaper** 🔍 *Under Review* - Technically sound and well-researched with properly cited references and professional quality

## Development View

- [ ] **Architecture Overview** 🔍 *Under Review* - Sound architectural design with scalability/performance considerations, clear integration points, best practices followed
- [ ] **Component/Sequence Diagrams** 🔍 *Under Review* - Accurate and complete following UML/PlantUML standards with sufficient detail for implementation
- [ ] **Standards** 🔍 *Under Review* - All standards properly referenced with current versions, compliance verifiable, industry alignment confirmed
- [ ] **API-Specification / Protocols** 🔍 *Under Review* - Complete Open API Specifications, OpenAPI 3.x compliant, complete with examples, versioning strategy clear
- [ ] **Access & Usage Policies** 🔍 *Under Review* - If DSP is used (via Connector), the policies are complete and correct
- [ ] **Logic / Schema** 🔍 *Under Review* - Well-documented business logic, complete database schemas, clear data flows
- [ ] **Semantic Models / Data Structures** 🔍 *Under Review* - SAMM models validated, standards-compliant, extensible design, version-controlled
- [ ] **Test Cases** 🔍 *Under Review* - Adequate coverage (>80%), all tests passing, edge cases covered, performance benchmarks met

## Operation View

- [ ] **Deployment Guide** 🔍 *Under Review* - Clear and complete instructions with all prerequisites listed, successfully tested independently, comprehensive troubleshooting
- [ ] **Operation/Monitoring Guidelines** 🔍 *Under Review* - Monitoring setup clearly explained, metrics well-defined, alerting rules appropriate, dashboard templates included
- [ ] **Security Guidelines** 🔍 *Under Review* - Security measures adequate, threat model complete, OWASP best practices followed, vulnerabilities addressed, acceptable security scan results

## Success Stories

- [ ] **Reference Implementations** 🔍 *Under Review* - Implementations verified operational with confirmed contacts, testimonials obtained, adoption path clear, value demonstrated with metrics

## Documentation

- [ ] **Extra Documentation / Links** 🔍 *Under Review* - All links functional and current, documentation complete and accurate, FAQ comprehensive, community resources available
- [ ] **Sample Data** 🔍 *Under Review* - Data representative of real use cases, privacy and GDPR compliant, sufficient for comprehensive testing, well-documented structure

## Industry Extensions

- [ ] **At least one Industry Extension** 🔍 *Under Review* - Standards verified current, Catena-X process followed, published or submitted, documentation complete, compliance verified

---

## Review Status Tracker

- [ ] Documentation review completed
- [ ] Technical review completed
- [ ] Security review completed (if applicable)
- [ ] Business review completed
- [ ] All feedback addressed
- [ ] Ready for graduation decision

---

## Helpful links

- [KIT TRGs 10.XX](https://eclipse-tractusx.github.io/docs/category/trg-10---keep-it-together-kits)
- [KIT Office Hour](https://eclipse-tractusx.github.io/community/open-meetings#Eclipse%20Tractus-X%20KITs%20Community%20Office%20Hour)
- [KIT Template](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/tree/main/docs-kits/kit-template)