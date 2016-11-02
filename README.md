# SAGE Assessment Server

## Dependencies
- [Go](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)

## Installation
Installation instructions are currently only available for Macs.

### MongoDB
1. Install [Homebrew](http://brew.sh/).
2. Use Homebrew to install MongoDB: `brew install mongodb`.

### Go
1. Install [Homebrew](http://brew.sh/).
2. Use Homebrew to install Go: `brew install go`.
3. Set up the Go workspace

  ```
  mkdir ~/Go
  cd ~/Go
  mkdir bin pkg src
  ```
  
4. Set up Go environmental variables. Typically you want to copy and paste this into `~/.bash_profile`.

  ```
  export GOPATH=$HOME/Go
  export GOROOT=/usr/local/opt/go/libexec
  export PATH=$PATH:$GOPATH/bin
  export PATH=$PATH:$GOROOT/bin
  ```
    
4. Clone this repository into ~/Go/src/github.com/

  ```
  mkdir -p ~/Go/src/github.com/cu-sage
  cd ~/Go/src/github.com/cu-sage
  git clone git@github.com:cu-sage/sage.git
  ```
5. Install the package dependencies

  ```
  cd ~/Go/src/github.com/cu-sage/sage
  go get
  go install
  ```
  
## Starting the Server
First follow the instructions to start the MongoDB database. Then follow the instructions to start the Go server.

### MongoDB
1. Open a new window in Terminal.
2. Start the daemon process: `mongod --config /usr/local/etc/mongod.conf`.

### Go
1. Open a window in Terminal.
2. Go to `~/Go/github.com/cu-sage-sage`.
3. Pull the latest version of the server: `git pull`.
4. Compile the code: `go build`.
5. Start the server: `go run main.go`.
