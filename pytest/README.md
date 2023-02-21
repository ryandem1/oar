# OAR PyTest Plugin

## Summary

The OAR PyTest plugin allows for an intuitive declarative interface for OAR with configurations to connect to any 
OAR instance. This will allow for both offline OAR results or uploading OAR results to an OAR instance

## Quickstart

Here is a walkthrough of a new PyTest project with OAR:

### Installation

To start, install the OAR PyTest plugin in your Python interpreter:

```
pip install pytest-oar
```

Next we need to install the plugin in PyTest, this guide assumes a fresh project.

Create a new file called ``conftest.py`` and add the following line:

```
pytest_plugins = ["oar.plugin"]
```
OAR is now installed!

### Creating a test
OAR is designed to be as much or as little as you want it to be. Here is a minimal test that will be reported 
to OAR:

```
class TestExample:

    def test_one_equals_one(self, oar_test):
        assert 1 == 1
```

The inclusion of the ``oar_test`` fixture is the key here, this will both mark this test to 
have its results reported, and yield a test object to enrich during test runtime.

Let's run the test and see what we get:
```
pytest -s
```

Result:
```
collected 1 item                                                                                                                                               

test_example.py .
============OAR SUMMARY===============
Passed IDs: [0]
Failed IDs: []
Tests that need analysis: []
Tests that need resolution: []
======================================


====================================================================== 1 passed in 0.00s =======================================================================
```
Not necessarily the most helpful on its own.

In the current working directory, there is a new folder, by default titled **oar-results**:

```
.
├── __pycache__
├── conftest.py
├── oar-results
├── test_example.py
└── venv
```

Each time OAR runs, by default, it will store an offline JSON result. Ours is:

```
{
    "start_time": "2023-02-21 04:03:26.406457",
    "completed_time": "2023-02-21 04:03:26.441054",
    "tests": [
        {
            "id": 0,
            "summary": "test_one_equals_one",
            "outcome": "Passed",
            "analysis": "TrueNegative",
            "resolution": "NotNeeded"
        }
    ],
    "failed_ids": [],
    "passed_ids": [
        0
    ],
    "need_analysis_ids": [],
    "need_resolution_ids": [],
    "all_ids": [
        0
    ]
}
```
Notice how the "id" is set to 0? When running OAR in offline mode, unless you manually set an ID,
the ID will be 0. If running in online mode, a new ID will be returned corresponding to the ID in the
Postgres DB.

### Configuration
Let's make OAR a bit more useful and connect our test to an instance of OAR. To do this, we must
configure OAR. 

Configuration is handled through Pydantic if you are familiar with that, if not, no worries.

Configuration can be done through a ``.toml`` file, a ``.json`` file, or environment variables.

By default, it will look for a ``.toml`` or ``.json`` file located in: ``CWD/<OAR_CONFIG_PATH>`` 
(defaults to look for ``oar-config.toml``)

Let's configure OAR with a ``.toml`` file:

Here is our tree now:
```
.
├── __pycache__
├── conftest.py
├── oar-config.toml
├── oar-results
├── test_example.py
└── venv
```

To connect OAR to an instance, we need to set two variables: ``host`` and ``send_results``.

Here is our ```config.toml```:

```
host = "http://localhost:8080"
send_results = true
```
> **Note**: In this example, I have a local OAR instance running, accessible on my localhost's port 8080

The ```send_results``` environment variable is by default False because you might not want to just upload 
any local test results, makes it a bit more deliberate.

When running the test again:

```
collected 1 item                                                                                                                                               

test_example.py .OAR Test Result ID: 59

============OAR SUMMARY===============
Passed IDs: [59]
Failed IDs: []
Tests that need analysis: []
Tests that need resolution: []
======================================


====================================================================== 1 passed in 0.02s =======================================================================
```

Notice, in this result, we have result IDs > 0, this is because it is returning the DB ID now. After each OAR test uploaded,
the ID will also be logged.

> **Note**: Because configuration is done with Pydantic, environment variable configuration can be done with the same 
> config variable names with the "OAR_" prefix prepended.
> 
> So for example, if I wanted to have an equivalent configuration via environment variables, I could do the following:
> ``OAR_HOST='http://localhost:8080' OAR_SEND_RESULTS=True pytest -s'``

Here is a current list of environment variables (with their defaults):

```
host: str = "oar-service:8080"  # Base URL of the OAR instance to send results to
send_results: bool = False  # This is what will control sending the results to the OAR instance
store_results: bool = True  # This will enable the `oar_results` fixture, will not prevent sending results to OAR
log_summary: bool = True  # This will control the logging of summary statistics in a run
output_file: bool = True  # This will output a JSON results file with name `oar-results-<utc-timestamp>.json`
output_dir: str = "oar-results"  # Controls where JSON results files will be stored. Relative to CWD
```

### Enriching Tests
Let's go back to our tests to see how we can make the result more useful:

Let's say that we have to test an API endpoint ``/user``, and our simplified test looks like this:

```
class TestExample:

    def test_users_endpoint_returns_user(self, oar_test):
        request_body = {"id": 1}
        response = simulate_users_endpoint(request_body)
        assert response["username"]
        assert "@" in response["email"]
```

The ``oar_test`` fixture also serves as an object that can be enriched with runtime data.

Let's start by adding some test metadata, like an owner and app name of this test case:

```
class TestExample:

    def test_users_endpoint_returns_user(self, oar_test):
        oar_test.owner = "Ryan"
        oar_test.app = "user-service"

        request_body = {"id": 1}
        response = simulate_users_endpoint(request_body)
        assert response["username"]
        assert "@" in response["email"]
```

OAR tests can be enriched like normal Python objects. All fields added to the ``oar_test`` will be placed 
in the dynamic ```doc``` column of the OAR table.

Great! Later we want to analyze results, we can now query by these fields.

Let's add some runtime data too, that can help us potentially debug the test:

```
class TestExample:

    def test_users_endpoint_returns_user(self, oar_test):
        oar_test.owner = "Ryan"
        oar_test.app = "user-service"

        request_body = {"id": 1}
        oar_test.request_body = request_body

        response = simulate_users_endpoint(request_body)
        oar_test.response_body = response

        assert response["username"]
        assert "@" in response["email"]
```
Even better! Now when we go to debug this test, all information will be included onto it. When we analyze results later, 
we can look for patterns in request/response bodies that can help us debug and find weaknesses in tests and apps.

Let's take a look at our generated JSON test report now:

```
{
    "start_time": "2023-02-21 04:41:21.064272",
    "completed_time": "2023-02-21 04:41:21.119699",
    "tests": [
        {
            "id": 62,
            "summary": "test_users_endpoint_returns_user",
            "outcome": "Passed",
            "analysis": "TrueNegative",
            "resolution": "NotNeeded",
            "app": "user-service",
            "request_body": {
                "id": 1
            },
            "response_body": {
                "username": "aUser",
                "email": "user@mail.com"
            },
            "owner": "Ryan"
        }
    ],
    "failed_ids": [],
    "passed_ids": [
        62
    ],
    "need_analysis_ids": [],
    "need_resolution_ids": [],
    "all_ids": [
        62
    ]
}
```
Our test is also available in the OAR DB now:

```
62,test_users_endpoint_returns_user,Passed,TrueNegative,NotNeeded,2023-02-21 04:41:21.113327,2023-02-21 04:41:21.113327,"{""app"": ""user-service"", ""owner"": ""Ryan"", ""request_body"": {""id"": 1}, ""response_body"": {""email"": ""user@mail.com"", ""username"": ""aUser""}}"
```

If you see by default, our test Summary takes the name of the PyTest test, we can overwrite any default OAR attribute by
just defining it in our test:

```
oar_test.summary = "Ensures the users endpoint returns a valid username/email"
```
You can manually set Outcome, Analysis, and Resolution within a test too to implement your own custom pass/fail conditions,
automatic analysis, and even resolution.

### Custom Test Types
This previous example demonstrated how we can add information to a test dynamically, but, as tests scale, 
there might be a desire to re-use custom types.

With the previous test as an example, let's say we wanted to make a formal ``APITest``type to ensure certain
fields get set as a way to harden data and reduce reporting errors.

The ``oar_test`` fixture yields an object of the type ``oar.Test``, this is a Pydantic model. We can inherit
and extend this model. Here is an example:

```
class APITest(oar.Test):
    owner: str
    app: str
    request_body: dict[str, typing.Any]
    response_body: dict[str, typing.Any]
```

With this created, to enforce this type on our test, all we have to do is add the type hint:

```
class TestExample:

    def test_users_endpoint_returns_user(self, oar_test: APITest):
        oar_test.summary = "Ensures the users endpoint returns a valid username/email"
        oar_test.owner = "Ryan"
        oar_test.app = "user-service"

        request_body = {"id": 1}
        oar_test.request_body = request_body

        response = simulate_users_endpoint(request_body)
        oar_test.response_body = response

        assert response["username"]
        assert "@" in response["email"]
```

Now if we don't define a field for this testcase, an error will be raised, and the test will NOT be reported until it is 
fixed.

> **Note**: It is not advised to add methods to custom types, they will not be available at a test's runtime, as all 
> tests start as a ```oar.Test``` and will only be converted to the specific type after the test is over.

Let's see the result of this test:

```
{
    "start_time": "2023-02-21 04:55:50.615502",
    "completed_time": "2023-02-21 04:55:50.668846",
    "tests": [
        {
            "id": 65,
            "summary": "Ensures the users endpoint returns a valid username/email",
            "outcome": "Passed",
            "analysis": "TrueNegative",
            "resolution": "NotNeeded",
            "owner": "Ryan",
            "app": "user-service",
            "request_body": {
                "id": 1
            },
            "response_body": {
                "username": "aUser",
                "email": "user@mail.com"
            },
            "type": "APITest"
        }
    ],
    "failed_ids": [],
    "passed_ids": [
        65
    ],
    "need_analysis_ids": [],
    "need_resolution_ids": [],
    "all_ids": [
        65
    ]
}
```
Notice that, for tests of a specific type, a "type" field will automatically be added.

## Conclusion

That's about it! If you made it this far, you know how to use OAR in PyTest. This is only the 
beginning of the process, pay attention to tests that need analysis/resolution after each run if you 
would like to follow the process.

Also enrich what is valuable, if it would be helpful for debugging or querying, add it! If it is not,
it is probably best to leave it out.
