package packed

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"strings"
)

type pCode struct {
}

var Code pCode

// GetMsg 获取code码对应的msg
func (c *pCode) GetMsg(code int) string {
	var maps map[int]string
	maps = map[int]string{
		0:     "好耶",
		10100: "sorry,数据走丢了",
		10101: "你所在文章分类不存在",
		10501: "图片类型不正确",
		20100: "你的账号或者密码不对哩",
		20101: "不正确的或者过期的token，请重新获取",
		20102: "┗|｀O′|┛ 嗷~~，获取你信息失败了",
	}
	return maps[code]
}

// SetErr 设置一个业务码的err
func (c *pCode) SetErr(code int, msg ...string) (err error) {
	var text string
	if len(msg) == 0 {
		text = c.GetMsg(code)
	} else {
		text = strings.Join(msg, ",")
	}
	err = gerror.NewCode(gcode.New(code, "", nil), text)
	return
}
