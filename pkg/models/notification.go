package models

import "sync"

type IndexRvNotifications map[string][]IndexRvModel

type NotificationStore struct {
	Data IndexRvNotifications
	mu   sync.RWMutex
}

func (ns *NotificationStore) Add(indexID string,
	notification IndexRvModel) {
	println("ADDING: " + indexID)
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.Data[indexID] = append(ns.Data[indexID], notification)
}

func (ns *NotificationStore) Get(indexID string) []IndexRvModel {
	ns.mu.RLock()
	defer ns.mu.RUnlock()
	return ns.Data[indexID]
}
