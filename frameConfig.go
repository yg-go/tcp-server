package main

import (
	"encoding/binary"
	"github.com/rolancia/goframe"
	server "github.com/rolancia/thing"
)

func TcpFrameConfig() server.FrameConfig {
	decoder := goframe.DecoderConfig{
		ByteOrder:           binary.LittleEndian,
		LengthFieldOffset:   0,  /*Length field started at 12bytes after*/
		LengthFieldLength:   2,  /*2bytes*/
		LengthAdjustment:    -2, /*Header length appended to length field*/
		InitialBytesToStrip: 0,  /*No Skip Header*/
	}
	encoder := goframe.EncoderConfig{
		ByteOrder:                       binary.LittleEndian,
		LengthFieldLength:               4,
		LengthAdjustment:                0,
		LengthIncludesLengthFieldLength: false,
	}
	return server.FrameConfig{encoder, decoder}
}
