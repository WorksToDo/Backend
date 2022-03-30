<div align="right">
  <img src="https://eneskzlcn.github.io/my-published-images/todo-icon.png" width="50" height="50">
</div>
<h3 align="center">About Todo Frontend Project</h3>
<br>

<hr>

#### Project Information And Purpose

This project is the backend part of todo project which provides a REST api to handle
getTodos and addTodo request. When a getTodos request arrives to the related endpoint
it returns all todos in database. When an addTodo request arrives with a new todo to the
related endpoint, it adds the todo the database and returns the added todo.

<hr>

#### Architectural Decisions

This project implemented with domain driven design(DDD) approach. The architecture is shaped 
by DDD. Started with developing todo domain and all the other external library or architectures like server or memory layer,
stands out of domain. Because a server or a memory part, do not related with todo domain. If you decide to use
it in the todo domain, you can put the server in another package and put a reference to the todo domain.
``` html
 Directory Strutcure On DDD Will Be Like
 
    <project-dir>
        <todo>
            <repository>
            <service>
            <handler>
            <models-entities>
            <transactions>
        <server>
        <memory-layer> 
        <main.go>
```

The handler layer responsible for catching sending request on related endpoints. The
query or params or another needed data obtained here and then the service layer called
to execute given query.

The service layer responsible for encapsulating repository layer and being a bride between handler and repository layer.
called by handler layer to execute given query like getTodos or
addTodo. In the service layer, you can put your business logic to control the given, context creating or etc. This layer
calls repository layer when decided to add or get data from database.

The repository layer responsible for database transactions. When a database transaction is
needed, service layer calls related functionality of repository layer and gets the data.

```
Client --> Get Todo Request --> Handler <--> Service <--> Repository
```
<img src="https://eneskzlcn.github.io/my-published-images/todo-backend-schema.png" width="600" height="600">

#### Verify Provider Test

Suppose that you have a published contract from frontend to post todos. Pact like:

```json
{
  "description": "POST todos request",
  "providerState": "posted todo successfully",
  "request": {
    "method": "POST",
    "path": "/todos",
    "body": {
      "task": "get some bread"
    },
    "matchingRules": {
      "$.body": {
        "match": "type"
      }
    }
  },
  "response": {
    "status": 200,
    "headers": {
      "Content-Type": "application/json"
    },
    "body": {
      "id": 0,
      "task": "buy some milk"
    },
    "matchingRules": {
      "$.body": {
        "match": "type"
      }
    }
  }
```
So, start to write the provider test
```go
import (
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
)
func createPact() (pact *dsl.Pact, cleanUp func()) {
	pact = &dsl.Pact{
		Host:                     "localhost",
		Consumer:                 "todo-frontend",
		Provider:                 "todo-backend",
		DisableToolValidityCheck: true,
	}

	cleanUp = func() { pact.Teardown() }

	return pact, cleanUp
}


```

Here we create the pact and cleanUp function which will called by provider test.

```go
pact, cleanUp := createPact()
defer cleanUp()
port,_ := utils.GetFreePort()
app := fiber.New()
go func() {
	err := app.Listen(fmt.Sprintf(":%d",port))
}
_, err := pact.VerifyProvider(t,  types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", port),
		PactURLs:                   []string{"pact-url-in-pactflow-"},
		BrokerURL:                  "https://any.pactflow.io",
		BrokerToken:                "ASZ12415AS",
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		StateHandlers:              map[string]types.StateHandler{
			"posted todo successfully": func() error {
				return nil
			},
		},

	})
```
You expect the test will fail because there is no handler listens for POST request of endpoint '/todos'
. So add the code in below directly after the new fiber app created.
```go
type todo struct {
	Id int `json:"id"`
	Task string `json:"task"`
}
app.Get("todos", func(ctx *fiber.Ctx) error {
	return ctx.JSON(todo{Id:2,Task:"buy some milk"})
})
```
After you add that you expect the test is pass.

<hr>

#### Unit Tests And Mocking 

Because of the architecture, if you want to test handler, service, repository as unit.
You need to mock the encapsulated strucutres. Example;
```go
type Handler struct {
	service IService
}
```
As you see your handler structure has a service inside. So if you want to test handler
as unit, you need to mock the service and then test the handler.

An example with of getTodos scenario;

```go
// handler.go

type Handler struct {
	service IService
}
func (h *Handler) GetTodos(c *fiber.Ctx) error {
    todos, err := h.service.GetTodos()
    if err != nil {
    return c.SendStatus(fiber.StatusBadRequest)
    }
return c.Status(fiber.StatusOK).JSON(todos)
}

func (h *Handler) RegisterRoutes(app *fiber.App){
    app.Get("/todos", h.GetTodos)
}

```
```go
// service.go
type IService interface {
    GetTodos() ([]Todo, error)
}

type Service struct {
    repository IRepository
}
```
You need to generate mock service to test handler independent from service.

```bash
$ mockgen -destination=mocks/mock_service.go -package=mocks todo-backend/todo IService
```
After that command your mock_service will be in the destination mocks/mock_service.go
Use it on your handler test

```go
// handler_test.go
func TestWhenGetTodosRequestArrivesWithValidRequestItReturnsTodosWithoutError(t *testing.T){
    request := struct {
        expectedTodos []todo.Todo
    }{
        expectedTodos: []todo.Todo{
        {
            ID:   0,
            Task: "buy some milk",
        },
        {
            ID:   1,
            Task: "buy some water",
        },
    },
    }
    mockController := gomock.NewController(t)
    defer mockController.Finish()
    mockService := mocks.NewMockIService(mockController)
    mockService.EXPECT().GetTodos().Return(request.expectedTodos, nil)
    
    handler, _ := todo.NewHandler(mockService)
    app := fiber.New()
    handler.RegisterRoutes(app)
    
    // make request without body is just a utility to create a http request without body
    testRequest := MakeRequestWithoutBody(http.MethodGet, "/todos")
    response, err := app.Test(testRequest)
    assert.Nil(t, err)
    
    AssertBodyEqual(t, response.Body, request.expectedTodos)
}

```
After that, you expect your test to finish. So you need to mock the dependency injected
structure before writing test of the related structure.

<hr>

#### CI-CD Pipeline Of Project

The CI-CD pipeline has the following stages:

stages:
- build
- test
- package
- prepare-test-artifacts
- deploy-to-test-env
- pact-provider-verify
- prepare-prod-artifacts
- deploy-to-prod-env

**build:**
In the build stage, the project is building with a go docker image and
we expect the project build successfully in this stage.

**test:**
In the test stage, all unit tests of the project runs in a go
docker image. Expect the test pass to complete successfully this stage.

**package:**
In the package stage, the project is containerized with docker using a docker image and dind(docker in docker)
service. The reason for using the dind service is to make the docker image run in the runner docker image successfully.
Purpose of this stage is running the docker commands to build, tag and publish the project as docker image to the [docker.hub](https://hub.docker.com/)
The packaged project image is Dockerfile which stands on project directory. This
dockerfile will be used in test and production environment.

**prepare-test-artifacts:**
In the prepare-test-artifacts stage, helm charts of todo-backend project
which in [todo-helm](https://gitlab.com/todo32/helm) repository downloaded and used to render the newest todo-backend
test environment deployment artifacts including the newest version of todo-backend image which is built in package stage.
After all needed k8s deployment files for test environment is rendered with specific helm values, all rendered
deployment files are pushed to the [todo-deployment-artifacts](https://gitlab.com/todo32/deployment-artifacts) repository to use later to deploy.

**deploy-to-test-env:**
In the deploy-to-test-env, all the newest deployment files just added to the deployment-artifacts repo in previous stage,
is deployed to the test environment with argocd which already inside the test k8s cluster. For needed
configs, gcloud cli is used. Also the related argocd is inside the .cd directory to apply before use it if not exist in k8s cluster.

**pact-provider-verify:**
In this stage, the provider tests runs in go image and expect the test passed.

**prepare-prod-artifacts:**
In the prepare-prod-artifacts stage, helm charts of todo-backend project
which in [todo-helm](https://gitlab.com/todo32/helm) repository downloaded and used to render the newest todo-backend
production environment deployment artifacts including the newest version of todo-backend image which is built in package stage.
After all needed k8s deployment files for prod environment is rendered with specific helm values, all rendered
deployment files are pushed to the [todo-deployment-artifacts](https://gitlab.com/todo32/deployment-artifacts) repository to use later to deploy.

**deploy-to-prod-env:**
In the deploy-to-prod-env, all the newest deployment files just added to the deployment-artifacts repo in previous stage,
is deployed to the prod environment with argocd which already inside the prod k8s cluster. For needed
configs, gcloud cli is used. Also the related argocd is inside the .cd directory to apply before use it if not exist in k8s cluster.

<hr>

#### Build Setup
```bash
# to install all dependencies
$ go mod tidy -go=1.16 && go mod tidy -go=1.17
# to start

# directly with 
$ go run main.go

# or 

$ go build -o bin/main main.go && ./bin/main
```
#### Docker Compose Setup
```bash

$ docker build -t todo-backend:latest
$ docker-compose up
```
<hr>

#### Reachable Project Parts
Todo assignment project consist of 5 different projects including this project. You can reach;
- Todo project acceptance repository <a href="https://gitlab.com/todo32/acceptance"> here</a>,
- Todo project frontend repository <a href="https://gitlab.com/todo32/frontend"> here</a>,
- Todo project helm repository <a href="https://gitlab.com/todo32/helm"> here</a>,
- Todo project deployment artifacts repository <a href="https://gitlab.com/todo32/deployment-artifacts"> here </a>

#### Running Project Environments

The all parts of project getting deployed into to environments; test and production environments. The project is live and available in;
- Test environment <a href="http://34.116.156.27:8090/">here</a>,
- Production environment <a href="http://34.116.223.97:8090/">here</a>.

<hr>

#### References
- **Self Made**
    - Kubernetes [docs](https://eneskzlcn.github.io/my-documentations/CI-CD/Kubernetes/Index)
    - ArgoCD [docs](https://eneskzlcn.github.io/my-documentations/CI-CD/ArgoCD/Index)
    - Google Cloud K8s Engine [docs](https://eneskzlcn.github.io/my-documentations/CI-CD/GoogleK8S/Index)
- **External**
    - Kubernetes [docs](https://kubernetes.io/docs/home/), [playground](https://www.katacoda.com/courses/kubernetes/playground)
    - ArgoCD [docs](https://argo-cd.readthedocs.io/en/stable/)
    - Helm [docs](https://helm.sh/docs/)
    - Very beneficial video source [TechWorld With Nana](https://www.youtube.com/c/TechWorldwithNana) from youtube for all devops technologies including k8s, argocd, helm or etc.
    - Fiber [docs](https://docs.gofiber.io/)
    - Gorm [docs](https://gorm.io/docs/)
    - Test with mocks 
    - DDD with example from programmingmercy.tech [here](https://programmingpercy.tech/blog/how-to-domain-driven-design-ddd-golang/)
    - Gomock an excellent getting started [source](https://blog.codecentric.de/en/2017/08/gomock-tutorial/)
  