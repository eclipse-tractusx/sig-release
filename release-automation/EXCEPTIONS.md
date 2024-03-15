# Process Description for Raising Tractus-X Release Guideline Exceptions

## Objective

The process aims to establish a structured method for raising exceptions to the Tractus-X Release Guidelines (TRG) for the release automation tool. These exceptions are managed through pull requests (PRs) modifying the `exceptions.yaml` file. The system team will review these exceptions and make decisions accordingly.

1. **Identification of Exception:**
   - Before raising an exception, the requester must identify a specific guideline that needs to be exempted.

2. **Understanding Structure:**
   - Review the structure of the [exceptions.yaml](https://github.com/eclipse-tractusx/sig-release/blob/main/release-automation/internal/exception/exceptions.yaml) file to understand how exceptions are documented. Ensure familiarity with the format, including key-value pairs and any required fields. Proper exception needs to have repository url under respective TRG.

    ```yaml
    exceptions:
        - trg: "TRG 4.02 - Base images"
          repository:
          - "https://github.com/eclipse-tractusx/policy-hub"
    ```

3. **Creating a Branch:**
   - Create a new branch from the main repository to work on the exception request.

4. **Modifying `exceptions.yaml`:**
   - Navigate to the [exceptions.yaml](https://github.com/eclipse-tractusx/sig-release/blob/main/release-automation/internal/exception/exceptions.yaml) file within the repository.
   - Modify the file to include the necessary exception(s) following the predefined structure.

5. **Drafting a Pull Request (PR):**
   - Create a PR to propose the changes made to the [exceptions.yaml](https://github.com/eclipse-tractusx/sig-release/blob/main/release-automation/internal/exception/exceptions.yaml) file.
   - Provide a clear and descriptive title for the PR, summarizing the purpose of the exception(s).
   - In the PR description, elaborate on the reasons for requesting the exception(s) and any potential impact on the release process.

6. **Assigning Reviewers:**
   - Assign the appropriate members (committers role) to review the PR.
   - Ensure that reviewers are familiar with the Tractus-X Release Guidelines and have the authority to approve or decline exceptions.

7. **Review and Discussion:**
   - Assignees review the proposed exception(s) and provide feedback within the PR.
   - Discussion may occur regarding the justification, potential risks, and alternative solutions.

8. **Decision Making:**
   - Based on the feedback and discussion, members will make a decision to either accept or decline the exception(s).
   - If accepted, the exception(s) will be merged into the main branch.
   - If declined, reasons for the decision will be communicated, and the requester may revise and resubmit the exception request if necessary.

9. **Closing the PR:**
   - Once a decision has been reached, close the PR accordingly.
   - Ensure that all relevant stakeholders are informed of the outcome.

10. **Continuous Improvement:**
    - Periodically review the process for raising exceptions and incorporate any lessons learned or feedback from stakeholders to improve efficiency and effectiveness.
