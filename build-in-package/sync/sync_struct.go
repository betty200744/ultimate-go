package sync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"ultimate-go/utils"
)

type UserInfo struct {
	Id   int64
	Name string
}

type UserService struct {
	stopChan        chan struct{}
	stopOnce        sync.Once
	userInfoMap     map[int64]UserInfo
	userInfoMapLock sync.RWMutex
	userList        atomic.Value // type is []UserInfo
	userLock        sync.Mutex
	userNumber      int
	userNumberCond  *sync.Cond
}

func NewUserService() *UserService {
	userService := &UserService{
		stopChan:        make(chan struct{}),
		userInfoMap:     map[int64]UserInfo{},
		userInfoMapLock: sync.RWMutex{},
		userList:        atomic.Value{},
		userLock:        sync.Mutex{},
	}
	userService.userNumberCond = sync.NewCond(&userService.userLock)
	return userService
}

func (s *UserService) Start() {
	//	go1
	//  go2
}

func (s *UserService) Stop() {
	s.stopOnce.Do(func() {
		close(s.stopChan)
	})
}

func (s *UserService) refreshRoutine() {
	tick := time.NewTicker(time.Second * 2)
	for {
		select {
		case <-tick.C:
			s.addRandomUserInfo()
		case <-s.stopChan:
			return
		}
		// user lock
		s.userLock.Lock()
		// user Cond signal
		s.userNumberCond.Signal()
		// user unlock
		s.userLock.Unlock()
	}
}
func (s *UserService) refreshUserListRoutine() {
	for {
		// user obtain lock
		s.userLock.Lock()
		// user Cond wait
		s.userNumberCond.Wait()
		v := s.userList.Load()
		if v == nil {
			s.userNumber = 0
		} else {
			s.userNumber = len(v.([]UserInfo))
		}
		// user unlock
		s.userLock.Unlock()
	}
}
func (s *UserService) addRandomUserInfo() {
	s.userInfoMapLock.Lock()
	defer s.userInfoMapLock.Unlock()
	id := utils.RandInt64(10000, 99999)
	s.userInfoMap[id] = UserInfo{
		Id:   id,
		Name: fmt.Sprintf("name_%v", id),
	}
	s.userLock.Lock()
	defer s.userLock.Unlock()
	var userList []UserInfo
	v := s.userList.Load()
	if v == nil {
		userList = make([]UserInfo, 0)
	} else {
		userList = v.([]UserInfo)
	}
	userList = append(userList, s.userInfoMap[id])
	s.userList.Store(userList)
	fmt.Println(s.userInfoMap)
	v1 := s.userList.Load()
	fmt.Println(v1.([]UserInfo))
}
