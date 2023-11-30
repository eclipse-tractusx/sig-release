# Release notifier

Welcome to Release Notifier tool.

Release notifier tool was designed to keep community informed about new releases and updates.
Currently it supports PostgreSQL & Kubernetes with potential to grow and support other applications.

## Upgrade Alignment Process Description

**Objective:** To streamline the process of initiating and implementing an upgrade, ensuring community consensus, documentation alignment, and effective communication.

1. **Initiation of Discussion:**
   - A community member recognizes the need for an upgrade of a component, after receiving notification of a new release and initiates a discussion on the Tractus-X mailing list.
   - In the discussion, the member provides a brief summary of the major changes, improvements, or new features that the upgrade will introduce.

2. **Collaborate and Consensus:**
   - Community members actively participate in the mailing list discussion.
   - Collaboratively, the community outlines a consensus on the need for the upgrade and its potential impact on the project.
   - Discussions may include technical, resource, and compatibility considerations.

3. **Agreement to Proceed:**
   - If the community reaches a consensus and agrees to proceed with the major release upgrade, formal decision is made.

4. **Update Aligned Version in the TRG Document:**
   - Community member updates aligned version in the respective Tractus-X Release Guide following the [process](https://eclipse-tractusx.github.io/docs/release/).
      - PostgreSQL: [TRG 5.07](https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-07)
      - Kubernetes: [TRG 5.10](https://eclipse-tractusx.github.io/docs/release/trg-5/trg-5-10)

5. **Update GitHub Workflow:**
   - Within the project's GitHub repository, update environment variable to reflect the agreed-upon version for the major release.
      - PostgreSQL: CURRENT_ALIGNED_PSQL_VER at [release-notifier workflow](https://github.com/eclipse-tractusx/sig-release/blob/main/.github/workflows/release-notifier.yaml#L35) 
      - Kubernetes: CURRENT_ALIGNED_K8S_VER at [release-notifier workflow](https://github.com/eclipse-tractusx/sig-release/blob/main/.github/workflows/release-notifier.yaml#L36) 

6. **Send Communication:**
   - Once the TRG and GitHub workflow are aligned with the new version, member prepares a formal communication.
   - The communication details the component relevant release upgrade, its purpose, features, and any critical instructions for the community.

7. **Implementation:**
   - The upgrade is carried out as per the defined upgrade plan, with close attention to the documentation, codebase, and any testing or validation procedures as required.
      - PostgreSQL: [upgrade plan](TBP)
      - Kubernetes: [upgrade plan](TBP)

9. **Ongoing Feedback and Improvement:**
   - Continuous feedback from the community is encouraged and used to further improve the upgrade process for future releases.

By following this well-defined process, you ensure that major release upgrades are initiated, executed, and communicated effectively while maintaining community consensus and alignment with documentation. This promotes transparency, minimizes disruptions, and fosters a collaborative atmosphere within Tractus-X community.
