# Tractus-X Release

Details about a Tractus-X release can be found in the dedicated repository [eclipse-tractusx/tractus-x-release](https://github.com/eclipse-tractusx/tractus-x-release).
There, you can also find details about the [release process](https://github.com/eclipse-tractusx/tractus-x-release/blob/main/RELEASE.md).

## Artifacts

- Changelog (overarching features)
- Umbrella Helm Chart(s)
- Overarching architecture documentation (if available)

## Versioning

- Calendar versioning

## Process

<!-- How are products (considered to be) integrated in the Tractus-X release -->
- Prior to release, a Tractus-X umbrella Chart release candidate branch is created
- Producs issue PRs with their release candidate

### Quality Assurance

- User Journey tests (e2e) defined in e2e-testing repo
- User Journey tests automatically executed via `helm test` (nightly + on every product rc PR)
