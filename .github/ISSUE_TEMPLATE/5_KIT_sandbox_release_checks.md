---
name: KIT Sandbox Release Checklist
about: Checklist for a KIT in Sandbox stage
title: '[KIT-SANDBOX] <KIT Name> - Sandbox Release Checklist'
labels: 'kit, sandbox'
assignees: ''
---


# KIT Sandbox Release Checklist

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

If changes to your KIT are not yet live on [our website](https://eclipse-tractusx.github.io/developer), link the open PR introducing the changes for this release.

If changes are already live, also point out the PRs witch included the changes for this version if possible.

- [PR number](pr-url)

---

## Pre-Requirements

- [ ] **CHANGELOG** available and maintained - A file documenting all notable changes, versions, and release dates ([Keep a Changelog](https://keepachangelog.com/))

### KIT TRG 10
- [ ] [TRG 10.01 - KIT Architecture](https://eclipse-tractusx.github.io/docs/release/trg-10/trg-10-01) - Defines the overall architectural structure and patterns for KITs
- [ ] [TRG 10. 02 - KIT Content](https://eclipse-tractusx.github.io/docs/release/trg-10/trg-10-02) - Specifies required content structure and organization
   - [ ] [TRG 7. 07 - Legal notice for non-code](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-07) - Image and media licensing requirements (CC BY 4.0)
   - [ ] [TRG 7. 08 - Legal notice for KIT documentation (CC-BY-4. 0)](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-08) - Documentation licensing requirements
- [ ] [TRG 10.03 - KIT Lifecycle](https://eclipse-tractusx.github.io/docs/release/trg-10/trg-10-03) - Describes lifecycle stages and transition criteria

### Master Data
- [ ] **Master Data Updated** (latest version & updated date) - Current version number and last update timestamp maintained

---

## General Requirements

- [ ] **Changelog** - Version history file (CHANGELOG.md) in repository root

## All Files

- [ ] **Notice** - Legal and compliance notice in every file with CC BY 4.0 license statement ([TRG 7. 07](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-07), [TRG 7.08](https://eclipse-tractusx.github.io/docs/release/trg-7/trg-7-08))
- [ ] **KIT Icon** - SVG icon stored at `/static/img/kits/<kit-id>/<kit-id>-kit.svg` following design guidelines

## Adoption View

- [ ] **Vision / Mission** - Clear description of the KIT's purpose, long-term goals, and what problem it solves
- [ ] **Business Value(s)** - Concrete benefits, ROI, and quantifiable value propositions for adopters
- [ ] **Use Case / Domain explanation** - Real-world scenarios, application contexts, and target industries

## Development View *(Optional at this stage)*

- [ ] **Architecture Overview** - High-level system design, component relationships, and integration points (recommended)

## Helpful links

- [KIT TRGs 10.XX](https://eclipse-tractusx.github.io/docs/category/trg-10---keep-it-together-kits)
- [KIT Office Hour](https://eclipse-tractusx.github.io/community/open-meetings#Eclipse%20Tractus-X%20KITs%20Community%20Office%20Hour)
- [KIT Template](https://github.com/eclipse-tractusx/eclipse-tractusx.github.io/tree/main/docs-kits/kit-template)