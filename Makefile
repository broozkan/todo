generate-mocks:
		mockgen -destination=internal/mocks/mock_stock_repository.go -package mocks  playground/configuration_manager/internal/client Todo

