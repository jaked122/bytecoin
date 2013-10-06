This is my attempt at making a readme while I develop, meaning that as I add
features and change things the documentation will keep up. For bytecoin I feel
that this is particularly necessary because there is a massive volume of new
cryptocurrency concepts.

Right now there is a simple client and server. The client connects to the
localhost server (or any server - it's a socket!) on port 8080 and prints
a basic statement.

The server takes whatever data the client sends and prints it to a file,
storage.ytc - storage.ytc will be responsible for holding all of the files that
back bytecoin. YTC is the ticker term for bytecoin.
