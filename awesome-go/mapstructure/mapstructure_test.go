package mapstructure

import (
	"testing"
	"time"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/assert.v1"
)

type Person struct {
	Id       int               `json:"id"`
	Name     string            `json:"name"`
	Channels []string          `json:"channels"`
	UserCode string            `json:"user_code"`
	Extra    map[string]string `json:"extra"`
}

// 游戏列表
type GameInfo struct {
	Id            int32     `gorm:"column:id;primary_key;AUTO_INCREMENT"`            // 自增id
	Name          string    `gorm:"column:name;NOT NULL"`                            // 大写英文名字缩写，未用
	DisplayName   string    `gorm:"column:display_name;NOT NULL"`                    // 中文名字
	Summary       string    `gorm:"column:summary;NOT NULL"`                         // 简介
	Platform      string    `gorm:"column:platform;NOT NULL"`                        // 平台
	LogoPath      string    `gorm:"column:logo_path;NOT NULL"`                       // logo文件名
	WebCoverPath  string    `gorm:"column:web_cover_path;NOT NULL"`                  // web_cover封面文件名
	PcCoverPath   string    `gorm:"column:pc_cover_path;NOT NULL"`                   // pc_cover封面文件名
	AppCoverPath  string    `gorm:"column:app_cover_path;NOT NULL"`                  // app_cover封面文件名
	InstallPath   string    `gorm:"column:install_path;NOT NULL"`                    // 游戏启动路径/启动文件名
	Argument      string    `gorm:"column:argument;NOT NULL"`                        // 游戏启动参数
	StartType     int       `gorm:"column:start_type;default:0;NOT NULL"`            // 启动分类
	StartMode     int       `gorm:"column:start_mode;default:0;NOT NULL"`            // 启动模式
	ClickCountPc  int       `gorm:"column:click_count_pc;default:0;NOT NULL"`        // 点击次数
	ClickCountApp int       `gorm:"column:click_count_app;default:0;NOT NULL"`       // 点击次数
	PlayCountPc   int       `gorm:"column:play_count_pc;default:0;NOT NULL"`         // 玩过
	PlayCountApp  int       `gorm:"column:play_count_app;default:0;NOT NULL"`        // 玩过
	PlayHeat      int       `gorm:"column:play_heat;default:0;NOT NULL"`             // 热度，未用
	StatusPc      int32     `gorm:"column:status_pc;default:0;NOT NULL"`             // pc状态：1-上架，2-下架
	StatusApp     int32     `gorm:"column:status_app;default:0;NOT NULL"`            // app状态：1-上架，2-下架
	Permit        int       `gorm:"column:permit;default:0;NOT NULL"`                // 权限：0-都可见，1-有权限可见
	Ctime         time.Time `gorm:"column:ctime;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	Mtime         time.Time `gorm:"column:mtime;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func Test_Decode(t *testing.T) {
	input1 := map[string]interface{}{"id": 1, "name": "betty", "channels": []string{"c1", "c2"}, "extra": map[string]string{"a": "a"}}
	output := &Person{}
	mapstructure.Decode(input1, output)
	assert.Equal(t, input1["id"].(int), output.Id)
	assert.Equal(t, input1["extra"].(map[string]string)["a"], output.Extra["a"])
}
func Test_DecodeArray(t *testing.T) {
	input1 := map[string]interface{}{"id": 1, "name": "betty", "channels": []string{"c1", "c2"}}
	input2 := map[string]interface{}{"id": 1, "name": "betty", "channels": []string{"c1", "c2"}}
	inputs := []map[string]interface{}{input1, input2}
	output := make([]*Person, 0)
	mapstructure.Decode(inputs, &output)
	assert.Equal(t, input1["id"].(int), output[0].Id)
}
func TestMapStructureJsonTagDecode(t *testing.T) {
	input1 := map[string]interface{}{"id": 1, "name": "betty", "channels": []string{"c1", "c2"}, "user_code": "code", "extra": map[string]string{"a": "a"}}
	output := &Person{}
	MapStructureJsonTagDecode(input1, output)
	assert.Equal(t, input1["user_code"].(string), output.UserCode)
	assert.Equal(t, input1["extra"].(map[string]string)["a"], output.Extra["a"])
}
