package sync

import (
	"fmt"
	"testing"
	"time"
)

/*
struct {map, list,  string }  , Mutex, RWMutex , Lock, UnLock
struct{stopChan, stopOnce } , sync.Once, stopOnce.Do()
*sync.Cond, cond.Signal(), cond.Broadcast, cond.Wait(), routine a 通知routine b，且routine b用routine的lock
*/

func Test_Cond(t *testing.T) {
	userService := NewUserService()
	go userService.refreshRoutine()
	go userService.refreshUserListRoutine()
	fmt.Println(userService.userNumber)
	fmt.Println(userService.userInfoMap)
	time.Sleep(time.Second * 10)
	fmt.Println(userService.userNumber)
	fmt.Println(userService.userInfoMap)
	fmt.Println("Done")
	fmt.Println(userService.userNumber)
	userService.Stop()
}
