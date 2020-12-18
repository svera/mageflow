# Mageflow

A CLI application to help with the creation and management of pull requests against Magento CE, EE , B2B and Infrastructure repositories.

**WARNING**: *This application is just a proof of concept at the moment. Output is faked, with the intention to show how the real application would work.*

## Motivation

The flow for merging code in Magento is fairly complex at the moment, requiring a minimum of two pull requests, with tests having to pass twice as a result and different requirements depending of which branch we want to merge against.

There are several tools helping with that process, mainly in the form of Jenkins jobs. Unfortunately, they require quite a lot of information to fill, most of which is the same after each run or can be easily inferred, making this process slightly cumbersome.

Here's where mageflow comes into play. Filling a simple configuration file once, and making use of remote access APIs provided by both Github and Jenkins, we can create and manage pull requests from a good old terminal.

## Installation

No binaries are provided at the moment, so you'll need to build it from source. Go 1.11 or greater is required.

To install, just clone the repository and type `go install` from there.

## Configuration

Before starting to use `mageflow`, you need to create a configuration file called `.mageflow.json` in your home directory, and fill it with the required data.

* `repos` holds the location of the repositories for the different Magento editions in your local environment.
* `organisation` is the URL of your team organisation in Github.
* `github.token` is used to store your personal Github API token, required to do requests to Github.
* `jenkins.credentials` is also used to store Jenkins credentials required to to requests to Jenkins.

An example file is provided.

## Usage

`mageflow` is composed of 4 commands: `create`, `status`, `retest` and `promote`.

`create`

Creates one or several pull requests in your team's selected repositories. A form is shown, where we must select:

* Against which Magento version this PR will be created.
* Which branches from our local environment we want to push.
* Wheter we want to bypass PR validation or not.

Note that information related to the delivery profile is missing from this form. The idea is to provide a remote API that, given a Magento version, returns the correspondent delivery profile.

A job ID is returned after a successful creation, which must be later used to retrieve information by other mageflow commands.

`status <jobID>`

Show information about all pull requests belonging to a job ID as a table, where columns show pull requests and rows show tests. Table is color coded, so succesful tests are in green, while failed ones are in red.

`retest <jobID>`

Triggers again failed tests in all PR belonging to the passed job ID. If you want to rerun all tests, use the flag `--all`.

`promote <jobID>`

Copies the pull requests of the passed job in the mainline. By default, PR's are not copied if there are incomplete or failed tests. This behaviour can be omitted using the ` --force` flag.
