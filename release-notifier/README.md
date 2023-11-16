# Release notifier

Welcome to Release Notifier tool.

Release notifier tool was designed to keep community informed about new releases and updates.
Currently it supports PostgreSQL only with potential to grow and support other applications.

## Major Release Upgrade Process Description

**Objective:** To streamline the process of initiating and implementing a major release upgrade, ensuring community consensus, documentation alignment, and effective communication.

1. **Initiation of Discussion:**
   - A community member recognizes the need for a major release upgrade of a helm chart dependency, after receiving notification of a new release and initiates a discussion on the Tractus-X mailing list.
   - In the discussion, the member provides a brief summary of the major changes, improvements, or new features that the upgrade will introduce.

2. **Collaborate and Consensus:**
   - Community members actively participate in the mailing list discussion.
   - Collaboratively, the community outlines a consensus on the need for the upgrade and its potential impact on the project.
   - Discussions may include technical, resource, and compatibility considerations.

3. **Agreement to Proceed:**
   - If the community reaches a consensus and agrees to proceed with the major release upgrade, formal decision is made.

4. **Update TRG 5.07 Document:**
   - Community member updates PostgreSQL aligned version in the Tractus-X Release Guide (TRG) 5.07 following the [process](https://eclipse-tractusx.github.io/docs/release/).

5. **Update GitHub Workflow:**
   - Within the project's GitHub repository, update environment variable CURRENT_ALIGNED_PSQL_VER in the [psql-release-notifier workflow](https://github.com/eclipse-tractusx/sig-release/blob/main/.github/workflows/psql-release-notifier.yaml#L31) to reflect the agreed-upon version for the major release.

6. **Send Communication:**
   - Once the TRG and GitHub workflow are aligned with the new version, member prepares a formal communication.
   - The communication details the major release upgrade, its purpose, features, and any critical instructions for the community.

7. **Implementation:**
   - The major release upgrade is carried out as per the defined [upgrade plan](TBP), with close attention to the documentation, codebase, and any testing or validation procedures as required.

9. **Ongoing Feedback and Improvement:**
   - Continuous feedback from the community is encouraged and used to further improve the upgrade process for future releases.

By following this well-defined process, you ensure that major release upgrades are initiated, executed, and communicated effectively while maintaining community consensus and alignment with documentation. This promotes transparency, minimizes disruptions, and fosters a collaborative atmosphere within Tractus-X community.
