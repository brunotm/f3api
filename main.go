package main

/*
Copyright 2018 Bruno Moura <brunotm@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	api "github.com/brunotm/f3api/api.v1"

	"github.com/brunotm/f3api/server"

	"github.com/brunotm/f3api/store/badgerdb"
)

var (
	address  = flag.String("address", "localhost:8080", "address to listen")
	dataPath = flag.String("datapath", "./db", "database path to open or create")
	logLevel = flag.String("log", "info", "Log level")
)

func main() {
	flag.Parse()
	log.SetPrefix("f3api ")

	log.Println("openning db")
	store, err := badgerdb.Open(*dataPath)
	if err != nil {
		log.Fatalln("error openning db:", err)
	}

	config := server.Config{}
	config.Addr = *address
	srv := server.New(config, store)

	api.AddAllRoutes(srv)

	log.Println("starting f3api server")
	go srv.Start()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh

	log.Println("closing server", srv.Close(context.Background()))
	log.Println("closing store", store.Close())

}
