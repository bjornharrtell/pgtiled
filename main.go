package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/paulmach/go.vector_tile"
	"github.com/valyala/fasthttp"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	const vertice uint32 = 1
	feature := vector_tile.Tile_Feature{
		Geometry: []uint32{vertice},
	}
	layer := vector_tile.Tile_Layer{
		Features: []*vector_tile.Tile_Feature{&feature},
		Name:     proto.String("test"),
		Version:  proto.Uint32(2),
	}
	t := &vector_tile.Tile{
		Layers: []*vector_tile.Tile_Layer{&layer},
	}

	bytes, err := vector_tile.Encode(t)
	fmt.Printf("Tile bytes: %d\n", len(bytes))
	if err == nil {
		ctx.SetContentType("application/octet-stream")
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.Write(bytes)
	} else {
		panic(err)
	}
}

func main() {
	fasthttp.ListenAndServe(":8080", fastHTTPHandler)
}
