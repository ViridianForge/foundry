# GRPC Reflecting Client Pytest

This is a test framework to run pytest based testing of a GRPC endpoint without
needing to have access to, or compile, that service's protos.

## Initial Development Setup

* Install Python 3.8
  * `pyenv install 3.8.13`
* Create virtualenv
  * `python3.8 venv ./grpcPytest`
* Install pip dependencies
  * `pip install requirements.txt`

## Running example tests

`pytest test_suites`
