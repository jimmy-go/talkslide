package main

var (
	port       = flag.String("port", "2323", "HTTP port server.")
	dbHost     = flag.String("db-host", "", "MongoDB replicas hosts separated by comma (,). e.g.: host1:27017,host2:10001")
	dbDatabase = flag.String("db-database", "", "MongoDB database.")
)

func main() {
	... start all the things.
