package now

import (
	"fmt"
	"testing"
	"time"

	"github.com/jinzhu/now"
)

func TestName(t *testing.T) {
	t.Run("BeginningOfWeek", func(t *testing.T) {
		now.WeekStartDay = time.Monday
		fmt.Println(now.BeginningOfWeek())
		fmt.Println(now.EndOfWeek())
		fmt.Println(now.BeginningOfMonth())
		fmt.Println(now.EndOfMonth())
	})
}
