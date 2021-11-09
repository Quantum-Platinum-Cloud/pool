package pool

// This program generates mocks. It can be invoked by running:
//
// make mock
//

//go:generate mockgen -source=sidecar/interfaces.go -package=sidecar -destination=sidecar/mock_interfaces.go
//go:generate mockgen -source=internal/test/interfaces.go -package=test -destination=internal/test/mock_interfaces.go
