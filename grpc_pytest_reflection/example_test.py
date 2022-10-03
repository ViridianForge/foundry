import pytest

from grpc_requests import Client

"""

"""

@pytest.fixture
# TODO: Should be a class-wide method auto-run at test start. Don't need to recreate
# the client connection for every test.
def service_client():
  # TODO: Host definitions should be coming from some central configuration
  # Insecure connections for local testing
  
  try:
    client=Client.get_by_endpoint("localhost:50051")
    # To do TLS try the following line
    # client = Client.get_by_endpoint("localhost:443", ssl=True)
    yield client
    # Place any clean up code after here
  except Exception:
    pytest.fail("test server unavailable")

@pytest.fixture
def service_methods_map():
  yield {
    "helloworld.Greeter":['SayHello','SayHelloGroup','HelloEveryone','SayHelloOneByOne'],
    "grpc.health.v1.Health":['Check','Watch'],
    "grpc.reflection.v1alpha.ServerReflection":['ServerReflectionInfo'],
  }


@pytest.mark.smoke
def test_available_services(service_client: Client, service_methods_map: dict):
  expected_service_names = list(service_methods_map.keys())

  for service_name in service_client.service_names:
    assert service_name in expected_service_names


@pytest.mark.smoke
def test_available_methods(service_client: Client, service_methods_map: dict):

  for service_name in service_methods_map.keys():
    service = service_client.service(service_name)

    for method_name in service.method_names:
      assert method_name in service_methods_map[service_name]


@pytest.mark.smoke
@pytest.mark.parametrize("method,grpc_request,grpc_response",
  [
    ("SayHello", {"name":"Gimli Son of Gloin"},{"message":"Hello Gimli Son of Gloin!"}),
  ]
)
def test_greeter_sayhello(service_client: Client, method: str, grpc_request: dict, grpc_response: dict):
  response = service_client.request("helloworld.Greeter", method, grpc_request)
  # TODO: Add response validation somehow
  assert response == grpc_response


@pytest.mark.skip("Test Not Ready")
def test_greeter_sayhellogroup(service_client:Client):
  raise NotImplementedError


@pytest.mark.skip("Test Not Ready")
def test_greeter_helloeveryone(service_client:Client):
  raise NotImplementedError


@pytest.mark.skip("Test Not Ready")
def test_greeter_sayhelloonebyone(service_client:Client):
  raise NotImplementedError


@pytest.mark.smoke
def test_health_service(service_client: Client):
  health_service = service_client.service('grpc.health.v1.Health')
  health_response = health_service.Check()
  assert health_response == {'status': 'SERVING'}