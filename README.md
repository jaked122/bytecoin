This is my attempt at making a readme while I develop, meaning that as I add
features and change things the documentation will keep up. For bytecoin I feel
that this is particularly necessary because there is a massive volume of new
cryptocurrency concepts.

Right now there is a client and a server.

The client takes the argument 'store' followed by a filename, opens a
connection to the server on port 8080 and sends the file to the server.

The server takes the file and prints it to storage.ytc

In order to store the block chain, it was decided that a full database would
be necessary. The database chosen is sqlite. For that reason, you will need to
run the command: 'go get github.com/mattn/go-sqlite3'

the database currently contains a single table of hosts and files. each file
must be unique. each file has an associated host, multiple files can have
the same host. The database is saved to a file for persistence. Currently hosts
and files are always stored with the same key.
