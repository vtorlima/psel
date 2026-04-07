package main

import (
	"fmt"
	"sync"
)

// Backend representa um servidor backend com informações de carga
type Backend struct {
	Address     string
	ActiveConns int64
	mu          sync.Mutex
}

// Incrementa o contador de conexões ativas
func (b *Backend) IncrementConns() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.ActiveConns++
}

// Decrementa o contador de conexões ativas
func (b *Backend) DecrementConns() {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.ActiveConns--
}

// Retorna o número atual de conexões ativas
func (b *Backend) GetActiveConns() int64 {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.ActiveConns
}

// Escolhe o backend com menos conexões ativas (least connections)
func pickBackend() *Backend {
	minIdx := 0
	minConns := backendList[0].GetActiveConns()

	for i := 1; i < len(backendList); i++ {
		conns := backendList[i].GetActiveConns()
		if conns < minConns {
			minConns = conns
			minIdx = i
		}
	}

	selected := backendList[minIdx]
	fmt.Printf("[LB] Selected backend %s (active: %d) | Load: %d | %d | %d\n",
		selected.Address, selected.GetActiveConns(),
		backendList[0].GetActiveConns(),
		backendList[1].GetActiveConns(),
		backendList[2].GetActiveConns())

	return selected
}
