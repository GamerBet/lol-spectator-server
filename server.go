package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/collinglass/mw"
	spec "github.com/gamerbet/lol-spectator-server/spectator"
	"github.com/gorilla/mux"
)

func LoggingMW(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Path: %s, Method: %s\n", r.URL.Path, r.Method)

		h.ServeHTTP(w, r)
	})
}

func main() {
	var port = flag.String("port", "8080", "Declare a port to listen on.")
	flag.Parse()

	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	// log.SetOutput(f)
	log.Printf("Listening on port %v...\n", *port)

	r := mux.NewRouter()
	mw.Decorate(
		r,
		LoggingMW,
	)

	r.HandleFunc("/observer-mode/rest/featured", spec.FeaturedHandler)
	r.HandleFunc("/observer-mode/rest/consumer/version", spec.VersionHandler)
	r.HandleFunc("/observer-mode/rest/consumer/getGameMetaData/{platformId}/{gameId}/{yolo}/token", spec.GetGameMetaDataHandler)
	r.HandleFunc("/observer-mode/rest/consumer/getLastChunkInfo/{platformId}/{gameId}/{param}/token", spec.GetLastChunkInfoHandler)
	r.HandleFunc("/observer-mode/rest/consumer/getLastChunkInfo/{platformId}/{gameId}/null", spec.EndOfGameStatsHandler)
	r.HandleFunc("/observer-mode/rest/consumer/getGameDataChunk/{platformId}/{gameId}/{chunkId}/token", spec.GetGameDataChunkHandler)
	r.HandleFunc("/observer-mode/rest/consumer/getKeyFrame/{platformId}/{gameId}/{keyFrameId}/token", spec.GetKeyFrameHandler)

	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		panic(err)
	}
}
