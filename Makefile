generate-mocks:
	mockgen -destination=internal/mocks/mock_todo_client.go -package mocks github.com/broozkan/todo/internal/client Todo


