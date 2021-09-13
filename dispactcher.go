package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"math"
)

var (
	COM_CODE_TRACK_INFO int16 = 0x07D6
	COM_CODE_PLOT_INFO  int16 = 0x07D5
)

type Payload []byte

func parseHeader(payload Payload) (*modelA, int) {
	message := modelA{}
	offset := 0

	message.DataSize = int16(binary.LittleEndian.Uint16(payload[offset : offset+2]))
	offset += 2
	message.OpCode = int16(binary.LittleEndian.Uint16(payload[offset : offset+2]))
	offset += 2

	return &message, offset
}

func parseTrackInfo(header *modelA, payload Payload) *TrackInfo {
	info := TrackInfo{}
	info.ID = DocType(DataTypeTraces)
	info.Header = *header

	offset := 0
	info.NTrack = int32(binary.LittleEndian.Uint32(payload[offset : offset+4]))
	offset += 4
	for i := 0; i < int(info.NTrack); i++ {
		track := Track{}
		track.TrackID = int32(binary.LittleEndian.Uint32(payload[offset : offset+4]))
		offset += 4
		track.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[offset : offset+4]))
		offset += 4
		track.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[offset : offset+4]))
		offset += 4
		track.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[offset : offset+4]))
		offset += 4
		track.Power = math.Float32frombits(binary.LittleEndian.Uint32(payload[offset : offset+4]))
		offset += 4

		info.Tracks = append(info.Tracks, track)
	}
	return &info
}

func parsePlotInfo(header *modelA, payload Payload) *PlotInfo {
	info := PlotInfo{}
	info.ID = DocType(DataTypePlots)
	info.Header = *header

	offset := 0
	info.NPlot = int32(binary.LittleEndian.Uint32(payload[offset : offset+4]))
	offset += 4

	info.Plots = make([]Plot, 0)
	for i := 0; i < int(info.NPlot); i++ {
		plot := Plot{}
		plot.PlotID = int32(binary.LittleEndian.Uint32(payload[offset : offset+4]))
		offset += 4
		plot.X = math.Float32frombits(binary.LittleEndian.Uint32(payload[offset : offset+4]))
		offset += 4
		plot.Y = math.Float32frombits(binary.LittleEndian.Uint32(payload[offset : offset+4]))
		offset += 4
		plot.Z = math.Float32frombits(binary.LittleEndian.Uint32(payload[offset : offset+4]))
		offset += 4
		plot.Power = math.Float32frombits(binary.LittleEndian.Uint32(payload[offset : offset+4]))
		offset += 4
		info.Plots = append(info.Plots, plot)
	}

	return &info
}

func DispatcherHeader(in <-chan Payload, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case payload := <-in:
			head, dataOffset := parseHeader(payload)
			fmt.Println(payload, head, dataOffset)
			switch head.OpCode {
			case COM_CODE_TRACK_INFO:
				tracks := parseTrackInfo(head, payload[dataOffset:])
				fmt.Println(tracks)
				break
			case COM_CODE_PLOT_INFO:
				plots := parsePlotInfo(head, payload[dataOffset:])
				fmt.Println(plots)
				break
			}
		}
	}
}
