package cdc

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
	"todo-backend/config"
	"todo-backend/server"
	"todo-backend/todo"
	"todo-backend/todo/mocks"
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
func TestProvider(t *testing.T) {
	pact, cleanUp := createPact()
	defer cleanUp()

	port,_ := utils.GetFreePort()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockIRepository(mockCtrl)
	service,_  := todo.NewService(mockRepo)
	handler,_  := todo.NewHandler(service)
	server := server.NewServer(config.Server{Port: fmt.Sprintf("%d",port)},[]todo.IHandler{ handler})

	go server.Run()

	_, err := pact.VerifyProvider(t,  types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", port),
		PactURLs:                   []string{"https://zumber.pactflow.io/pacts/provider/todo-backend/consumer/todo-frontend/version/fac1d5f051c9cf31f31b2dc60bbde365f86c1cb0"},
		BrokerURL:                  "https://zumber.pactflow.io",
		BrokerToken:                "AhT3yO3pFlB6lZpWskVDcA",
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		StateHandlers:              map[string]types.StateHandler{
			"all todos successfully": func() error {
				mockData := []todo.Todo{
					{
						ID: 1,
						Task: "drink some milk Ahmet",
					},{
						ID: 2,
						Task: "drink some milk Enes",
					},
				}
				mockRepo.EXPECT().GetTodos().Return(mockData, nil)
				return nil
			},

			"posted todo successfully": func() error {
				mockRepo.EXPECT().AddTodo(gomock.Any()).Return(todo.Todo{
					ID:   0,
					Task: "buy some milk",
				}, nil)
				return nil
			},
		},

	})
	if err != nil {
		t.Fatal(err)
	}
}