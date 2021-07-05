# Release Notes Action

This action ensures that the PR description contains a release-note code block. This replicates the functionality of the Prow plugin used by the Kubernetes / Kubernetes-Sigs projects.

## How to use the action

Create a workflow that uses the GitHub action that is triggered via a pull request. For example:

```yaml
name: releasenote

on: [pull_request]

jobs:
  labeler:
    runs-on: ubuntu-latest
    name: Release notes check
    steps:
      - name: Check for release note
        id: rn
        uses: weaveworks/release-note-action@v0.1.0
        with:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

Once this is enabled it will check that the PR description contains a `release-notes` code block. For example:
````
```release-note
Added feature Y to accomplish X
```
````

The above example is for a **normal** release note where you feel the PR description isn't descriptive enough. If the PR title is enough of a description the following release note can be used:

````
```release-note
NONE
```
````

If a change requires special hightlighting in the release notes because it needs action from the user then a release note block can be marked as requiring action:

````
```release-note
action required
Added feature Y to accomplish X. This is a breaking change and requires you do Z
```
````

## How to use the release notes

You can generate a markdown file containing the release notes by using Kubernetes **release-notes** cli

First build the CLI:

```bash
GO111MODULE=on go get k8s.io/release/cmd/release-notes
```

Create a template file for use by the cli. You can use [this sample](sample/changelog.tpl). 

The easiest way is to tag your release and then run the CLI passing in the SHA of the current release, SHA of the previous release and the path to the template. For example:

```bash
release-notes --debug --org myorg --repo myrepo --start-sha $(shell git rev-list -n 1 ${PREVIOUS_VERSION}) --end-sha $(shell git rev-list -n 1 ${RELEASE_TAG}) --output out/CHANGELOG.md --go-template go-template:sample/changelog.tpl --dependencies=false
```
