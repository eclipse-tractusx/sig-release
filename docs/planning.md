# Roadmap and Release Process

We follow a coordinated approach to plan improvements of Eclipse Tractus-X. See also [Open Planning and Refinement](open-refinement-and-planning.md).

While every repository in the [eclipse-tractusx](https://github.com/eclipse-tractusx) GitHub organization has its own issue management, the [Release Planning Board](https://github.com/orgs/eclipse-tractusx/projects/26) is used to align the overarching Tractus-X releases.


### How can I get involved

In case you experienced a bug, unexpected behaviour, or you want to propose enhancements to Eclipse Tractus-X,
feel free to use one of the provided [issue templates](https://github.com/eclipse-tractusx/sig-release/issues/new/choose) and describe your request.
Please be aware, that not every feature request can be integrated and that we also cannot treat every issue with the highest priority.

Every Release planning will be kicked off by two public alignment sessions. The dates and further details will be shared via [tractusx-dev](https://accounts.eclipse.org/mailing-list/tractusx-dev) mailing list.
In addition to that, you can also find public meetings and info about how to join on our [Open Meetings](https://eclipse-tractusx.github.io/community/open-meetings) community page.
Issues or bug reports, that should be discussed in these meetings, have to be opened prior to the meeting via our [issue templates](https://github.com/eclipse-tractusx/sig-release/issues/new/choose).


### What can I expect

We really welcome every contribution. Every Bug report and feature proposal takes time to prepare,
is valuable to our project, and we very much appreciate this input.
We are giving our best to give a first feedback in one week.
If we should miss that, please stick with us and just use the commenting function to remind us of the issue.


### Issue structure

Our issues do have important properties, that enable our planning process. These are:
- **Labels:** We use them to indicate the involved teams (kit or foss product). A label for each involved product is added to an issue
- **Issue Type:** To separate between bugs, feature requests and release criteria, we use a custom field `Issue Type`
- **Milestone:** Every Tractus-X release is represented by a `Milestone`.
- **Status:** The status field is used to integrate the progress of an issue


### Issue statuses

The following statuses are defined:
- **Inbox:** This is the initial status of all issues. It indicates, that involved products have to be identified and additional information gathered
- **Backlog:** If enough information is gathered, and we agreed to work on the issue, it is set from `Inbox` to `Backlog` to indicate it is ready for timeline planning
- **Work in Progress:** The issue is actively being worked on.
- **Done:** All relevant parts have been implemented and released


### Issue process

Every new feature proposal or bug report will be handled as issue in status `Inbox` initially. The alignment meetings are used to discuss the purpose and impact of the current issues. While in `Inbox` status, the involved products are discovered and respective `Labels` are added. Additionally, an `Assignee` is selected, who will coordinate efforts to solve the issue.
After these details are clarified, an issue is moved to `Backlog` to open it for detailed timeline planning. In this status, discussions about a fitting `Iteration` is held.
The fully refined status "backlog" is the entry point for our "open planning" ceremony. During "open planning" a milestone is set for each feature. This represents a commitment, given by the open source community to finalize the feature in due time.
As soon as actual work is started in the selected iteration, the issue is set to `Work in Progress`. This is especially helpful on our project milestone views to get an overview of the release progress.


### Issue hierarchy (WIP)

@stephanbcbauer to formulate details
=> detailled features are planned and executed within the repositories of the respective product. Summary features are planned and managed on "SIG release" level...


### Planning product changes

While the [Release Planning Board](https://github.com/orgs/eclipse-tractusx/projects/26) is used to coordinate overarching feature and bug requests, we encourage every product team to break these issues down to their product repositories/projects.
When doing so, make sure you link to the overarching issue in your product issue description.

See [product release](product_release.md), for information about singular product releases.

See [Tractus-X release](tractus-x-release.md), for our overarching release strategy.
