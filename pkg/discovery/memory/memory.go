package memory

import (
	"backend/pkg/discovery"
	"context"
	"errors"
	"sync"
	"time"
)

type serviceNameType string
type instanceIDType string

type Registry struct {
	sync.RWMutex
	serviceAddr map[serviceNameType]map[instanceIDType]*serviceInstance
}

type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

// NewRegistry creates a new in-memory registry instance.
func NewRegistry() *Registry {
	return &Registry{
		serviceAddr: make(map[serviceNameType]map[instanceIDType]*serviceInstance),
	}
}

// Register creates a service record in the registry
func (r *Registry) Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddr[serviceNameType(serviceName)]; !ok {
		r.serviceAddr[serviceNameType(serviceName)] = map[instanceIDType]*serviceInstance{}
	}
	r.serviceAddr[serviceNameType(serviceName)][instanceIDType(instanceID)] = &serviceInstance{
		hostPort:   hostPort,
		lastActive: time.Now(),
	}
	return nil
}

// Deregister removes a service record from the registry
func (r *Registry) Deregister(ctx context.Context, instanceID string, serviceName string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddr[serviceNameType(serviceName)]; !ok {
		return discovery.ErrNotFound
	}

	delete(r.serviceAddr[serviceNameType(serviceName)], instanceIDType(instanceID))
	return nil
}

// HealthCheck marks a service instance as active
func (r *Registry) HealthCheck(instanceID string, serviceName string) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.serviceAddr[serviceNameType(serviceName)]; !ok {
		return errors.New("service not registered yet")
	}

	if _, ok := r.serviceAddr[serviceNameType(serviceName)][instanceIDType(instanceID)]; !ok {
		return errors.New("service instance not registered yet")
	}

	r.serviceAddr[serviceNameType(serviceName)][instanceIDType(instanceID)].lastActive = time.Now()
	return nil
}

// Discover returns a list of service instances from the registry
func (r *Registry) Discover(ctx context.Context, serviceName string) ([]string, error) {
	r.RLock()
	defer r.RUnlock()

	if len(r.serviceAddr[serviceNameType(serviceName)]) == 0 {
		return nil, discovery.ErrNotFound
	}
	var res []string

	for _, v := range r.serviceAddr[serviceNameType(serviceName)] {
		if time.Since(v.lastActive) > 5*time.Second {
			continue
		}

		res = append(res, v.hostPort)
	}
	return res, nil
}
