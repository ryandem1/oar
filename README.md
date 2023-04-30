![OAR Logo](static/oarLogo.png "OAR Logo") ![OAR Logo](static/oarLogo.png "OAR Logo") ![OAR Logo](static/oarLogo.png "OAR Logo") ![OAR Logo](static/oarLogo.png "OAR Logo") ![OAR Logo](static/oarLogo.png "OAR Logo")

# The OAR Framework for Software Test Reporting

## Outcome, Analysis, Resolution

See the GitHub wiki for more info on project background: https://github.com/ryandem1/oar/wiki/Background

## Links

### [OAR Service Docker Image](https://hub.docker.com/r/ryandem1/oar-service)
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
- **OAR enrichment UI**: A minimal UI for test engineers to maintain test result ownership, enrich test results, query 
for tests, and triage test/software issues.
- **OAR analytics UI**: A data visualization UI that connects to an instance of an OAR DB to provide fast, rich aggregate statistics
around software/test quality.
- More ingestion integrations (maybe some JS UI frameworks?)

### Getting Familiar with the OAR Service

Everything is built around the OAR service (or even more abstractly, the OAR data model). So getting familiar with the service
will get you familiar with any other OAR component.

The best place to get started with the `oar-service` api is to look at the ``oar-service-spec.json``.
This is the OpenAPI v3 schema for the service; copy and pasting it in swagger will provide an interface:
https://editor.swagger.io/.

#### Starting local instance for development/sandbox

To start a local instance of OAR, you could either pull and run the image (would need to configure environment variables),
or more simply, clone this repository and use a few make commands:

``make clean db build test-service service``

To break it down:
- ``clean``: Will delete orphan database volumes, teardown existing service, and perform other environment cleanups.
- ``db``: Starts a new instance of Postgres locally, waits for startup process, and will run the ``init-postgres.sql``
to create the OAR table.
- ``build``: Builds the ``oar-service`` image
- ``test-service``: Will run unit tests that do not do database cleanup, so it serves as a helpful way to seed the DB with
some data for experimenting
- ``service``: Starts a local ``oar-service`` container and will port-forward the service to your localhost.

After running these commands, use the Swagger UI or some other HTTP client to start!
