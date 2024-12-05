package sdp

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		CodecGB string
		want    *CodecGB
	}{
		{
			name:    "test1",
			CodecGB: "v/1/1/1/1/1a/1/1/1",
			want: &CodecGB{
				Codec:            "1",
				Resolution:       "1",
				Framerate:        "1",
				BitrateType:      "1",
				BitrateSize:      "1",
				AudioCodec:       "1",
				AudioBitrateType: "1",
				SampleRateType:   "1",
			},
		},
		{
			name:    "test2",
			CodecGB: "v/2/2/2/2/2a/2/2/2",
			want: &CodecGB{
				Codec:            "2",
				Resolution:       "2",
				Framerate:        "2",
				BitrateType:      "2",
				BitrateSize:      "2",
				AudioCodec:       "2",
				AudioBitrateType: "2",
				SampleRateType:   "2",
			},
		},
		{
			name:    "test3",
			CodecGB: "v/////a/1/8/1",
			want: &CodecGB{
				Codec:            "",
				Resolution:       "",
				Framerate:        "",
				BitrateType:      "",
				BitrateSize:      "",
				AudioCodec:       "1",
				AudioBitrateType: "8",
				SampleRateType:   "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCodecGB(tt.CodecGB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCodecGB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodecGB_String(t *testing.T) {
	tests := []struct {
		name string
		cf   *CodecGB
		want string
	}{
		{
			name: "test1",
			cf: &CodecGB{
				Codec:            "1",
				Resolution:       "1",
				Framerate:        "1",
				BitrateType:      "1",
				BitrateSize:      "1",
				AudioCodec:       "1",
				AudioBitrateType: "1",
				SampleRateType:   "1",
			},
			want: "v/1/1/1/1/1a/1/1/1",
		},
		{
			name: "test2",
			cf: &CodecGB{
				Codec:            "2",
				Resolution:       "2",
				Framerate:        "2",
				BitrateType:      "2",
				BitrateSize:      "2",
				AudioCodec:       "2",
				AudioBitrateType: "2",
				SampleRateType:   "2",
			},
			want: "v/2/2/2/2/2a/2/2/2",
		},
		{
			name: "test3",
			cf: &CodecGB{
				Codec:            "",
				Resolution:       "",
				Framerate:        "",
				BitrateType:      "",
				BitrateSize:      "",
				AudioCodec:       "1",
				AudioBitrateType: "8",
				SampleRateType:   "1",
			},
			want: "v/////a/1/8/1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cf.String(); got != tt.want {
				t.Errorf("CodecGB.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
