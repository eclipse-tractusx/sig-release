---
name: Product Release Checks
about: Use this template to track all release checks for your component
title: "[Product] Release Checks"
labels: compliance, documentation, "security analysis", "test results"
---

# Compliance Verifications

This issue tracks all compliance related checks, that need to be performed for a product release in Eclipse Tractus-X.

- [ ] **Gaia-X** compliance confirmed
- [ ] **GDPR** compliance confirmed (personal data, data protection + privacy DPP)
- [ ] **Interoperability** checks performed
- [ ] **Data Sovereignty** checks performed
- [ ] Compliant with relevant _published_ **CX Standards** (see the [Catena-X standard library](https://catena-x.net/de/standard-library))

# Documentation

- [ ] **Arc24** documentation up-to-date
- [ ] **Administrators Guide** up-to-date
- [ ] **End-User manual** up-to-date
- [ ] **Interface documentation** up-to-date

# Security Checks

- [ ] **Thread Modelling Analysis** passed
- [ ] **Static Application Security Testing** (SAST) scans passed
- [ ] **Dynamic Application Security Testing** (DAST) tests passed
- [ ] **Secret Scans** passed
- [ ] **Software Composition Analysis** (SCA) passed
- [ ] **Container Scans** passed
- [ ] **Infrastructure as Code** (IaC) scans passed

# General Checks

- [ ] [**Tractus-X Release Guidelines**](https://eclipse-tractusx.github.io/docs/release) fulfilled
- [ ] Compliant with the **Style Guide**

# Test Results

- [ ] **E2E Integration Test** passed
- [ ] **User Journey** approved

# Helpful Links

- [Eclipse Tractus-X openAPI specs on SwaggerHub](https://app.swaggerhub.com/search?owner=eclipse-tractusx-bot)
- [Catena-X standard library](https://catena-x.net/de/standard-library)