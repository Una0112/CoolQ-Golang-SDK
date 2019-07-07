package util

import (
	"testing"
)

func TestUnpackGroupList(t *testing.T) {
	mems, err := UnpackGroupList("AAAABABEAAAAADOGQH8AAAAAWHjaagAKqVnI9NPqx6fStgAAAAAAAQAAABMAAFwnP4pcoWM7AAAAAAACAAAAAAAAAAAAAAAAAAAARgAAAAAzhkB/AAAAAGXEFuoADKhzhvnTx6erlZnYyAAAAAAAAQAAABQAAFxQNkBcdT/VAAAAAAACAAAAAAAAAAAAAAAAAAAARAAAAAAzhkB/AAAAAGpre0AACs3svv314s/QzaUAAAAAAAEAAAAUAABcStJZXLBr7gAAAAAAAQAAAAAAAAAAAAAAAAAAAEIAAAAAM4ZAfwAAAACQL1MEAAiBN6I4ucjT6gAAAAAAAAAAABQAAFwnP4Vc3ASjAAAAAAADAAAAAAAAAAAAAAAAAAA=") //测试数据由晓影提供
	if err != nil {
		t.Log("解码出错 ", err)
	}

	if len(mems) != 4 {
		t.Errorf("成员数量解析错误，应为%d而不是%d", 4, len(mems))
	}

	t.Log(mems)
	//TODO(Tnze): 测试其余的值是否正确
}
