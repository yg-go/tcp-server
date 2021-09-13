package main

type dataType int
type DocType dataType

const (
	DataTypeTraces dataType = 1 + iota
	DataTypePlots
)

type modelA struct {
	DataSize int16 `json:"data_size" bson:"data_size"`
	OpCode   int16 `json:"op_code" bson:"op_code"`
}

type TrackInfo struct {
	ID     DocType `json:"doc_type" bson:"doc_type"`
	Header modelA  `json:"header" bson:"header"`
	NTrack int32   `json:"n_track" bson:"n_track"`
	Tracks []Track `json:"tracks" bson:"tracks"`
}

type PlotInfo struct {
	ID     DocType `json:"doc_type" bson:"doc_type"`
	Header modelA  `json:"header" bson:"header"`
	NPlot  int32   `json:"n_plot" bson:"n_plot"`
	Plots  []Plot  `json:"plots" bson:"plots"`
}

type Track struct {
	TrackID int32   `json:"track_id" bson:"track_id"`
	X       float32 `json:"x" bson:"x"`
	Y       float32 `json:"y" bson:"y"`
	Z       float32 `json:"z" bson:"z"`
	Power   float32 `json:"power" bson:"power"`
}

type Plot struct {
	PlotID int32   `json:"plot_id" bson:"plot_id"`
	X      float32 `json:"x" bson:"x"`
	Y      float32 `json:"y" bson:"y"`
	Z      float32 `json:"z" bson:"z"`
	Power  float32 `json:"power" bson:"power"`
}
