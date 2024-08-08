---
name: Product Release Checks
about: Use this template to track all release checks for your component
title: "[Product] Release Checks"
---

> [!IMPORTANT]
> Follow the [guidance on how to use the templates](https://github.com/eclipse-tractusx/sig-release/blob/main/README.md#release-management-acceptance-criteria).

## Release Info

Please provide information on what you want to be included in the Eclipse Tractus-X release.
If you are not owner of this issue, please provide the information as comment to the issue.
Make sure to assign this issue to expert(s) for their approval, as soon as you have finished preparation. Multiple assignees allowed; they will un-assign themselves once review completed.

**Version to be included in Eclipse Tractus-X release**: *version placeholder*

**Leading product repository:** *repository link*

# General Checks

- [ ] [**Tractus-X Release Guidelines**](https://eclipse-tractusx.github.io/docs/release)(TRGs) fulfilled

Make sure to open and fill in a separate **documentation issue** in your product repository using the [Quality Gate Checklist issue template](https://github.com/eclipse-tractusx/.github/blob/main/.github/ISSUE_TEMPLATE/qg-checklist.md)

> [!NOTE]
> Note: most criteria for documentation and security are now covered in TRGs

# Features
Please provide a list of all features that have been developed in the current phase. This list is essential for several reasons:

- **Test Preparation:** Identifying new features will help in the creation of new end-to-end (E2E) tests and ensure that the relevant aspects are thoroughly tested.
- **Awareness:** It will raise awareness of the features that need to be prioritized during the testing process.
- **Board Management:** This will assist in keeping the project board organized and up-to-date.

By documenting these features, you will contribute to a smoother and more efficient release process.

List your features here:

- [ ]

# Test Results

- [ ] **E2E Integration Test** passed
- [ ] **User Journey** approved

# Helpful Links

- [Eclipse Tractus-X openAPI specs on SwaggerHub](https://app.swaggerhub.com/search?owner=eclipse-tractusx-bot)
