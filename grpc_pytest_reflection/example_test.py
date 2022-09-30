import pytest

from grpc_requests import Client

"""

"""

@pytest.fixture
def service_client():
  # TODO: Host definitions should be coming from some central configuration
  # Insecure connections for local testing
  client=Client.get_by_endpoint("localhost:8000")
  # To do TLS try the following line
  # client = Client.get_by_endpoint("localhost:443", ssl=True)
  yield client


@pytest.testcasebullshit
@pytest.mark.smoke
def test_available_services(service_client: Client, service_names: list[str], method_names: list[str]):
  assert service_client.service_names == service_names
  for service_name in service_names:
    service = service_client.service(service_name)
    assert service.method_names == method_names[service_name]


@pytest.testcasebullshit
@pytest.mark.smoke
def test_example_service(service_client: Client, service_name: str, method_name: str, request: dict, expected_response: dict):
  test_service = service_client.service(service_name)
  actual_response = test_service.method_name(request)
  # TODO: Add response validation somehow
  assert actual_response == expected_response

