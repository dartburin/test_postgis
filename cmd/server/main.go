package main

import (
	"flag"
	"os"
	rest "test_postgis/internal/api/server"
	pdb "test_postgis/internal/gis"
	lg "test_postgis/internal/logger"
)

func main() {
	// Load init parameters
	dbHostName := flag.String("dbhost", "localhost", "host name")
	dbUser := flag.String("dbuser", "postgres", "user db name")
	dbPass := flag.String("dbpass", "", "user db pass")
	dbBase := flag.String("dbbase", "", "database name")
	dbPort := flag.String("dbport", "5432", "port for database connect")

	logLevel := flag.String("loglvl", "", "logging message level")
	logFile := flag.String("logfile", "", "logging message to file")

	restHostName := flag.String("httphost", "", "host name")
	restPort := flag.String("httpport", "", "port for http connect")
	flag.Parse()

	//Init log system
	log := lg.LogInit(*logLevel, *logFile)

	log.Println("")
	log.Println("Start application.")
	log.Println("Server log system init.")
	lg.PrintOsArgs(log)

	// Check existing obligatory http and db parameters
	if *restHostName == "" || *restPort == "" || *dbPass == "" || *dbBase == "" {
		flag.PrintDefaults()
		log.Println("")
		log.Fatalln("Init server error: set not all obligatory parameters.")
		os.Exit(1)
	}

	// Set db parameters
	var configDB pdb.Config
	configDB.Port = *dbPort
	configDB.Host = *dbHostName
	configDB.User = *dbUser
	configDB.Pass = *dbPass
	configDB.Db = *dbBase

	// Connect to DB
	parDB, err := pdb.ConnectToDB(configDB, log)
	if err != nil {
		log.Fatalf("Don`t connect for database (%s).", err.Error())
		os.Exit(1)
	}
	defer parDB.Base.Close()

	//Start REST server
	log.Println("Start REST server.")
	rr := rest.New(parDB, *restHostName, *restPort)
	err = rr.Start()
	if err != nil {
		log.Fatalf("Server REST error: %s.", err.Error())
	}
}
