# Get OIDC Identity Token

[![Step changelog](https://shields.io/github/v/release/bitrise-steplib/bitrise-step-get-identity-token?include_prereleases&label=changelog&color=blueviolet)](https://github.com/bitrise-steplib/bitrise-step-get-identity-token/releases)

The Step fetches an OIDC identity token.

<details>
<summary>Description</summary>

The "Prepare App Store Release" Step allows you to streamline the process of preparing a new release for your iOS app in the Release Management. This Step leverages the Bitrise Public API to facilitate the creation and configuration of an App Store release in the Release Management.

By utilizing this Step, you can automate the initial stages of the release process and ensure a consistent and efficient deployment experience. Instead of manually navigating through the Release Management interface to create a release, the Step empowers you to initiate the release setup programmatically, saving valuable time and effort.

It's important to note that this Step doesn't create a release directly in the App Store Connect. Instead, it streamlines the process by generating a release in the [Release Management](https://devcenter.bitrise.io/en/release-management.html).
</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `subject` | The subject for the identity token.  This could be the email or unique identifier of the user or service account for whom the token is being generated. | required |  |
| `audience` | The audience for the identity token.  This could be the URL of the service you want to access with the token or a specific identifier provided by the service. | required |  |
| `build_url` | Unique build URL of this build on Bitrise.io.  By default the step will use the Bitrise API. | required | `$BITRISE_BUILD_URL` |
| `build_api_token` | The build's API Token for the build on Bitrise.io  This will be used to communicate with the Bitrise API | required, sensitive | `$BITRISE_BUILD_API_TOKEN` |
| `verbose` | Enable logging additional information for debugging. | required | `false` |
</details>

<details>
<summary>Outputs</summary>

| Environment Variable | Description |
| --- | --- |
| `BITRISE_IDENTITY_TOKEN` | The newly generated identity token. |
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/bitrise-step-get-identity-token/pulls) and [issues](https://github.com/bitrise-steplib/bitrise-step-get-identity-token/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Note: this step's end-to-end tests (defined in e2e/bitrise.yml) are working with secrets which are intentionally not stored in this repo. External contributors won't be able to run those tests. Don't worry, if you open a PR with your contribution, we will help with running tests and make sure that they pass.

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)
