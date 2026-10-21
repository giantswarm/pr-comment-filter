# PR-comment-filter

Designed to be an entrypoint task for GitHub webhook triggers based on PR comments. `pr-comment-filter` will check the comment text for recognised triggers (e.g. `/run my-pipeline`) and generate individual `PipelineRun` resources for each.

The generated `PipelineRun` will have details about the PR passed through as params as well as any arguments added to the trigger comment.

Details about files updated in the PR will also be included as parameters on the `PipelineRun` with the following params being defined with a comma separated string value:

* `PR_FILES` - Includes all files contained in the PR
* `PR_FILES_ADDED` - Includes all new files
* `PR_FILES_CHANGED` - Includes all modified/changed/renamed files
* `PR_FILES_REMOVED` - Includes all files removed in the PR

## Trigger format

All triggers must follow the following format:

```
/run ${PIPELINE_NAME} [KEY=val ...]
```

Examples:

```
/run build-and-publish
/run test-cluster-create PRIVATE_NETWORK=true
/run test-cluster-create PREVIOUS_VERSION=1.2.6
/run test-cluster-upgrade PRIVATE_NETWORK=false PREVIOUS_VERSION=1.2.6
/run hold wait-for-tests
/run help NAMESPACE=foo-bar test-cluster-create
```

Some notes:

* Multiple triggers can be defined in a single comment but must each be on their own line
* Triggers must start the line with `/run ` and cannot be placed mid-sentance
* Arguments are optional and in the format of either:
  * `KEY=value` where the key must be all uppercase and is used as the key/val pair of environment variables
  * space seperated words - these are treated as "positional arguments" and added as the `POS_ARGS` env var with the values comma seperated
* Multiple arguments can be provided as long as they all appear on the same line as the trigger
* Argument values with spaces in them is not currently supported
* A Pipeline from a specific namespace can be run by specifying a `NAMESPACE=xxx` argument along with the `/run` trigger line
* If a user provided namespace isn't provided the pipeline will first be looked for in a namespace matching the repo name and if not found then default back to the `tekton-pipelines` namespace.

## Who can trigger pipelines

Comments are only acted on when they come from one of:

* an active member of the `giantswarm` GitHub organization (account type `User`), or
* an allowed bot (account type `Bot`), matched on its numeric user ID, never on its login.

The allowed bots are configured with environment variables:

| Variable | Default | Description |
| --- | --- | --- |
| `RENOVATE_BOT_USER_ID` | `29139614` | User ID of the Renovate bot. Kept for backwards compatibility. |
| `ALLOWED_BOT_USER_IDS` | `342014845` | Comma-separated user IDs of further allowed bots. The default is the `giantswarm-release` GitHub App (`giantswarm-release[bot]`). |

Both variables are combined. Setting `ALLOWED_BOT_USER_IDS` replaces its default, so include `342014845` if the `giantswarm-release` bot should stay allowed. Other bots, such as `github-actions[bot]` (ID `41898282`), are rejected.
