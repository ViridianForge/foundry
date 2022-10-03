# GRPC Reflecting Client Pytest

This is an example framework for the rapid development of pytest-based automated
testing of a GRPC server with reflection enabled, removing the need to maintain
and compile protos alongside test suites.

This framework builds around [GRPC for humans](https://github.com/spaceone-dev/grpc_requests) grpc client.

## Initial Development Environment Setup

* Create virtualenv: `python3 -m venv ./venv`
* Activate the virtual environment: `source ./venv/bin/activate`
* Make sure pip is up to date: `pip install --upgrade pip`
* Install pip dependencies: `pip install requirements.txt`

## Running example tests

### Start the Example Server

If the included compiled python proto files are up to date, start the server
by running:

```bash
python test_grpc_server/app.py
```

The server is configured to serve on port `50051`.

### Run Tests

With the server running, run the example test suite by running:

```bash
pytest
```

from the main test directory.

The framework comes with two marks predefined: `smoke` and `not-ready` defined
in `pytest.ini` to provide examples of using marks to differentiate tests.

## Additional Reading

* [pytest Docs](https://docs.pytest.org/en/7.1.x/contents.html)
* [Python Testing with pytest, Second Edition](https://pragprog.com/titles/bopytest2/python-testing-with-pytest-second-edition/)
* [pytest-html Docs](https://pytest-html.readthedocs.io/en/latest/)
