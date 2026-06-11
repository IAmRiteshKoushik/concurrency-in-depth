package main

import (
	"errors"
	"log/slog"
	"sync"
	"time"
)

const (
	Closed   = "closed"
	Open     = "open"
	HalfOpen = "half-open"
)

type circuitBreaker struct {
	mu                   sync.Mutex // Guards the circuit breaker state
	state                string     // Current state of the circuit breaker
	failureCount         int        // Number of consecutive failures
	lastFailureTime      time.Time  // Time of the last failure
	halfOpenSuccessCount int        // Successful requests in half-open state

	failureThreshold    int           // Failure to trigger open state
	recoveryTime        time.Duration // Wait time before half-open
	halfOpenMaxRequests int           // Requests allowed in half-open state
	timeout             time.Duration // Timeout for requests
}

func NewCircuitBreaker(
	failureThreshold int,
	recoveryTime time.Duration,
	halfOpenMaxRequests int,
	timeout time.Duration,
) *circuitBreaker {
	return &circuitBreaker{
		state:               Closed,
		failureThreshold:    failureThreshold,
		recoveryTime:        recoveryTime,
		halfOpenMaxRequests: halfOpenMaxRequests,
		timeout:             timeout,
	}
}

func (cb *circuitBreaker) handleOpenState() (any, error) {

}

func (cb *circuitBreaker) handleClosedState(fn func() (any, error)) (any, error) {

}

func (cb *circuitBreaker) handleHalfOpenState(fn func() (any, error)) (any, error) {

}

func (cb *circuitBreaker) Call(fn func() (any, error)) (any, error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	slog.Info("Making a request", "state", cb.state)

	switch cb.state {
	case Closed:
		return cb.handleClosedState(fn)
	case Open:
		return cb.handleOpenState(fn)
	case HalfOpen:
		return cb.handleHalfOpenState(fn)
	default:
		return nil, errors.New("unknown circuit state")
	}
}

func (cb *circuitBreaker) resetCircuit() {
	cb.failureCount = 0
	cb.state = Closed
	slog.Info("Circuit reset to closed state")
}

func (cb *circuitBreaker) runWithTimeout(
	fn func() (any, error),
) (any, error) {

}
