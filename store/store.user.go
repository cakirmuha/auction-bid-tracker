package store

import "fmt"

func (db *DB) GetUserNameByID(userID uint32) (*string, error) {
	db.cache.userMu.RLock()
	user, ok := db.cache.userCache[userID]
	db.cache.userMu.RUnlock()
	if ok {
		return &user.Name, nil
	}
	return nil, fmt.Errorf("user(%v) not available on the system", userID)
}
