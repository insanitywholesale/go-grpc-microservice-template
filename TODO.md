# TODOs
Things left to be done before the project is in an okay state

## Structure
Consider modularizing the project further

### Service/Logic layer
A layer that handles the logic so it's decoupled from the grpc implementation

#### Options
The ways to achieve it that I can think of are
- a struct with methods that implement the logic and will also hold the db which gets passed to it from the grpc server struct
- a set of functions that implement parts of the logic and the grpc layer will handle db interaction
- a set of functions that take the db as an argument and will be called from the grpc layer

## Documentation
Since this is a showcase of a template there should be good documentation

### Code comments
Add more comments in code to explain the rationale behind decisions as well as the not immediately-obvious parts

### Architecture
By far the most complicated thing here is the internal structure
- add `ARCHITECTURE.md` to explain it
- add graphviz graph to explain it

### Environment variables
Create a table of environment variables used throughout the project with a description and default values

## CI-type files
The more CI, the better
- add `Jenkinsfile`
- add `.travis.yml`
- add whatever Tekton has

## Deployment files
Need a way to run the application
- add `docker-compose.yml` for easier local testing with a real db
- add raw kubernetes yaml
- add kustomize yaml to demonstrate testing/prod

## Tests
A lot of things don't have tests
- real databases (currently the postgres implementations)
- models/interfaces (no idea how to test these)

## Errors
Wrap errors to add details at each step instead of doing `return nil, err`

## Functionality
Functionality is not yet complete

### Replace Empty gRPC message
If we want to accept something other than Empty there will be problems in the future so to avoid this replace it with a HelloFilter message or something

### Additional version demo
Add later API versions to demonstrate that fields can be added but breaking changes need a full API version change

### Rethink REST api versioning
Might want to actually namespace this but it works somewhat okay

### Repo
Adjust repos to store and retrieve all info for HelloRequest/HelloResponse

### Multiplexing
Add the option to run both grpc and rest on the same port but explain why it's dumb

### Standalone server
Add function CreateListenerFromPortAndAddress in `utils` to be able to create a listener from a user-supplied address.
Also add an env var or two (`HELLO_GRPC_ADDRESS`, `HELLO_REST_ADDRESS`) and maybe a CLI too to run with that user-supplied address
