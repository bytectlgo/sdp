package sdp

import (
	"fmt"
	"strings"
)

// CodecGB 媒体参数描述
type CodecGB struct {
	// v：后续参数为视频的参数；各参数间以 “/”分割；
	// 编码格式（十进制整数字符串表示）：1——MPEG-4  2——H.264  3——SVAC  4——3GP	5——H.26
	Codec string
	// 分辨率（字符串表示）：1——QCIF	 2——CIF  3——4CIF  4——D1  5——720P		6——1080P/I  其余分辨率用WxH表示（W表示宽，H表示高
	Resolution string
	// 帧率(十进制整数字符串表示)：0～99
	Framerate string
	// 码率类型: 1——固定码率（CBR）  2——可变码率（VBR）
	BitrateType string
	// 码率大小(十进制整数字符串表示)： 0～100000 （如 1表示1kbps
	BitrateSize string
	// a：后续参数为音频的参数；各参数间以 “/”分割；
	// 音频编编码格式(十进制整数字符串表示)： 1——G.711  2——G.723.1  3——G.729	     4——G.722.1  5——SVAC  6——AAC
	AudioCodec string
	// 音频编码码率类型
	// 注1：G.723.1中使用1、2。
	// 注2：G.722.1中使用4、5、6、7。
	// 注3：G.729中使用3。
	// 注4：G.711中使用8。
	// 注5：SVAC中使用5、9、20、21、22、23、24、25、26、27、28、29、30。
	// 注6：AAC中使用4、5、6、7、8、9、10、11、12、13、14、15、16、17、18、19。
	AudioBitrateType string
	// 采样率类型
	// G.711中使用1。
	// G.722.1中使用2、3、4。
	// G.723.1中使用1。
	// G.729中使用1。
	// SVAC中使用3、4、9、11，15、16、17。
	// AAC中使用1、3、4、5、6、7、8、9、10、11、12、13、14。
	SampleRateType string
}

func (c *CodecGB) String() string {
	return fmt.Sprintf("v/%s/%s/%s/%s/%sa/%s/%s/%s", c.Codec, c.Resolution, c.Framerate, c.BitrateType, c.BitrateSize, c.AudioCodec, c.AudioBitrateType, c.SampleRateType)
}

// 解析音视频参数描述字符串
func ParseCodecGB(str string) *CodecGB {
	if str == "" {
		return nil
	}
	arr := strings.Split(str, "/")
	if len(arr) != 9 {
		return nil
	}
	if arr[0] != "v" {
		return nil
	}
	if !strings.HasSuffix(arr[5], "a") {
		return nil
	}
	return &CodecGB{
		Codec:            arr[1],
		Resolution:       arr[2],
		Framerate:        arr[3],
		BitrateType:      arr[4],
		BitrateSize:      strings.TrimSuffix(arr[5], "a"),
		AudioCodec:       arr[6],
		AudioBitrateType: arr[7],
		SampleRateType:   arr[8],
	}
}
