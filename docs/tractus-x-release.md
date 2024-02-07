# Tractus-X Release

## Artifacts

- Umbrella Helm Chart(s)
- Changelog (overarching feat)

## Versioning

- Semantic Versioning

## Process

<!-- How are products (considered to be) integrated in the Tractus-X release -->
- Prior to release, a Tractus-X umbrella Chart release candidate branch is created
- Producs issue PRs with their release candidate

## Quality Assurance

- User Journey tests (e2e) defined in e2e-testing repo
- User Journey tests automatically executed via `helm test` (nightly + on every product rc PR)
