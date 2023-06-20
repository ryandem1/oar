![OAR Logo](static/oarLogo.png "OAR Logo") ![OAR Logo](static/oarLogo.png "OAR Logo") ![OAR Logo](static/oarLogo.png "OAR Logo") ![OAR Logo](static/oarLogo.png "OAR Logo") ![OAR Logo](static/oarLogo.png "OAR Logo")

# The OAR Framework for Software Test Reporting

## Outcome, Analysis, Resolution

See the GitHub wiki for more info on project background: https://github.com/ryandem1/oar/wiki/Background

## Links

### [OAR Service Docker Image](https://hub.docker.com/r/ryandem1/oar-service)
### [OAR Enrich UI](https://hub.docker.com/r/ryandem1/oar-enrich-ui)
### [OAR PyTest Plugin](https://pypi.org/project/pytest-oar/)
### [GitHub Wiki](https://github.com/ryandem1/oar/wiki)


### Overview

The core of OAR is the oar-service. It serves as the backend for test ingestion/enrichment/querying.

Currently, the OAR service is provided in a minimal Docker image. All parts of OAR are open-source and are assumed to be 
self-hosted. The oar-service image strives to be compatible with a wide variety of hosting environments, although exact
steps for hosting OAR will vary and not be covered in this documentation.

#### 🔌🔌 OAR Integrations 🔌🔌

There are the following packages to easily integrate OAR ingestion into common testing frameworks:

- [pytest-oar](https://pypi.org/project/pytest-oar/): OAR ingestion PyTest plugin and general oar-service Python interface.


#### 🚀🚀 The following components are coming in the future 🚀🚀:
- **OAR analytics UI**: A data visualization UI that connects to an instance of an OAR DB to provide fast, rich aggregate statistics
around software/test quality.
- More ingestion integrations (maybe some JS UI frameworks?)

### Getting Familiar with the OAR Service

Everything is built around the OAR service (or even more abstractly, the OAR data model). So getting familiar with the service
will get you familiar with any other OAR component.

The best place to get started with the `oar-service` api is to look at the ``oar-service-spec.json``.
This is the OpenAPI v3 schema for the service; copy and pasting it in swagger will provide an interface:
https://editor.swagger.io/.

### Getting Familiar with the OAR Enrich UI

The OAR process encourages test enrichment; meaning to analyze test results (specifically failures) and adding addtional 
metadata regarding the outcome of the test. This includes indicating if it was a true/false positive, if a ticket came out of it,
ticket reference, notes, diagnostic steps to make it easier to debug next time, etc.

All data will go back into the OAR DB, and can be captured in any sort of reporting layer. This is only as valuable as 
the data you add. Enrichment will cost time, but the idea is that the extra insights are vital and having them
will save time in the end.

The UI is one way to perform this enrichment, but there are potentially some lighter-weight alternatives like some sort
of Slack channel/plugin combo to enrich results.

#### Starting local instance for development/sandbox

To start a local instance of OAR, you could either pull and run the image (would need to configure environment variables),
or more simply, clone this repository and use a few make commands:

``make clean db build-service test-service service enrich-ui-dev``

To break it down:
- ``clean``: Will delete orphan database volumes, teardown existing service, and perform other environment cleanups.
- ``db``: Starts a new instance of Postgres locally, waits for startup process, and will run the ``init-postgres.sql``
to create the OAR table.
- ``build-service``: Builds the ``oar-service`` image
- ``test-service``: Will run unit tests that do not do database cleanup, so it serves as a helpful way to seed the DB with
some data for experimenting
- ``service``: Starts a local ``oar-service`` container and will port-forward the service to your localhost.
- ``enrich-ui-dev``: Starts a dev instance of the ``oar-enrich-ui``, which will run on ``http://localhost:5173``. Configure
the base URL of the service in the settings.

After running these commands, use the Swagger UI or some other HTTP client to start adding tests!
