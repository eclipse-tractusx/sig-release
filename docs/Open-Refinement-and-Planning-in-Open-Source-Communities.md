Welcome to our Wiki page on the open refinement and planning processes for our open source community. Effective planning and transparent communication are critical for the success of our open source project. Here, we'll outline how our community can participate in and leverage these processes to ensure smooth and efficient collaboration.

## Overview

Open refinement and planning allow the community to:

1. **Discuss** - Exchange thoughts, feedback, and suggestions on potential new features.
2. **Refine** - Hone and polish ideas into actionable items.
3. **Prioritize** - Decide on the features that should be worked upon in the upcoming releases.

### Refinement

* Related UserStories/Tasks will be created within product repository (task breakdown)
* Link the related UserStories/Tasks/Issues to the feature
* Set appropriate state ("Input" or "Backlog")

### During Planning

* Release Planning Board will be used from GitHub [Release Planning / Release View](https://github.com/orgs/eclipse-tractusx/projects/26)
* Finally attach all related UserStories/Tasks/Issues that have been planned for that feature
* Set the target Release milestone

## How to Create Features in the `sig-release` Repository

1. **Issue Creation**: Anyone in the community can suggest a new feature by creating an issue in the `sig-release` repository.
2. **Issue Type**: When creating your issue, it's crucial to specify the `Issue Type`:
   * `feature`: For new open features.
   * `bug`: For bugs or issues with existing features.
3. **Labels**: Use relevant labels to further categorize and provide context to your feature (in any cases, please select also your specific product label). Other than that, please select the `Prep-PI11` label to indicate that the feature is being prepared for the next release.
4. **Milestone**: Use relevant milestone for the next release, if you already know, it will be implemented. Currently, its `24.03`. Officially, the milestone is set during the planning meeting.

**Good to know:** Use the [Release Planning / Inbox](https://github.com/orgs/eclipse-tractusx/projects/26/views/9) board to add new open features. There is the last row with the functionality to add new issues.

## Task lists on Feature Level in `sig-releases`

When creating a `feature`, it's recommended to use a task list to break down the feature into smaller, actionable tasks. This not only provides clarity but also allows collaborators to pick up specific tasks and also link the different tasks to the related product repositories.

**BUT: These tasks should be created in the related product repository and should be linked / mentioned to the feature in `sig-release` repository.** 

```markdown
## Tasks for Feature XYZ:

- [ ] Task 1 description.
- [ ] Task 2 description.
- [ ] Task 3 description.
```

## Linking Issues from Different Repositories

**Important:** Please don`t use the "create issue" functionality from GitHub (provided with a task list). This will create a new issue in the "sig-release" repository. Moving this ( to the specific product repository ) and all the links is a mess. Please create the issue in the related product repository and link it to the feature in the "sig-release" repository. see also [Linking Issues from Different Repositories](https://github.com/eclipse-tractusx/sig-release/discussions/227)

To reference issues from different repositories within your task list, follow these steps:

1. Ensure you have the full path: `owner/repo#issue_number`. For instance, `eclipse-tractusx/item-relationship-service#187`.
2. Embed this in your markdown as you would any other link or mention.

Example:

Feature [PCWM Transfer](https://github.com/eclipse-tractusx/sig-release/issues/212). Two Issues from different repositories are linked to this feature.

```markdown
### Related Issues:

- [ ] eclipse-tractusx/item-relationship-service#187
- [ ] eclipse-tractusx/item-relationship-service#188
```

## Boards

We utilize several boards to manage and track the progress of issues and features:

1. **[Release Planning / Inbox](https://github.com/orgs/eclipse-tractusx/projects/26/views/9)**: Used for adding new open features.
2. **[Release Planning / Bug Board](https://github.com/orgs/eclipse-tractusx/projects/26/views/18)**: Tracks ongoing work on bugs.
3. **[Release Planning / Release 23.12 View](https://github.com/orgs/eclipse-tractusx/projects/26/views/8)**: Tracks ongoing work on all features for the next release.

## General guidelines

```java
while (FOSSactivity) {
    beNice();
    communicate();
    value(code).over(hierarchy);
    expandMyPortfolio(events);
    act(responsibly);
    ManageUsers(Interests[]);
    have(greatTime);}
```

## Community Participation

We encourage all community members to participate actively in the open refinement and planning processes. Together, we can ensure that the project continually evolves and improves to meet the needs and expectations of its users.

Please see also [more](https://github.com/eclipse-tractusx/sig-release/blob/main/README.md) information related to sig-release repository and planning.

Follow our [news section](https://eclipse-tractusx.github.io/blog) and join our [Tractus-X mailing list](https://eclipse-tractusx.github.io/docs/oss/how-to-contribute/#dev-mailinglist)
and be part of our [Matrix Chat from Eclipse Tractus-X](https://chat.eclipse.org/#/room/#tools.tractus-x:matrix.eclipse.org).

In case of any questions, please create a discussion in the [sig-release](https://github.com/eclipse-tractusx/sig-release/discussions) repository.

## Feedback

We welcome feedback and suggestions on how we can improve our open refinement and planning processes. Please create a discussion in the [sig-release](https://github.com/eclipse-tractusx/sig-release/discussions) repository.

## NOTICE

This work is licensed under the [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0).

- SPDX-License-Identifier: Apache-2.0
- SPDX-FileCopyrightText: 2022,2023 Contributors to the Eclipse Foundation
- Source URL: https://github.com/eclipse-tractusx/sig-release