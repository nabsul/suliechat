# SulieChat

A very simple command-line messaging program, written in Go.
I wrote this client and the accompanying server to introduce my nieces to the glories of the 
command line. They both have Kano devices (Raspberry Pis packaged for kids), so I wrote the 
client in Go which can be easily cross-compiled and packaged as a stand-alone executable.

At the moment, this isn't really a chat application. It simply allows you to send messages
to individual users and check if anyone sent messages to you.

The server has only a few simply features at the moment, and is written as an Azure Functions application.
It can be found at (https://github.com/nabsul/suliechat-server)

## Usage

Once the program is installed, you'll need to run `suliechat [server] [username] [password]` to create a config 
file that stores these parameters. The program will use these settings when all the other commands are executed.

The `[server]` parameter is either `localhost` for local testing, or the ``

For a full list of commands type `suliechat help`.  

## Build Instructions

To check out the code and compile it yourself:

- Install Go
- `go get github.com/nabsul/suliechat`
- `go install github.com/nabsul/suliechat`

## Build for Kano (Raspberry Pi)

If not build on a Raspberry Pi , you'll need to have the following environment variables set:

```text
GOOS=linux
GOARCH=arm
GOARCH=5
```

On Windows I do this in PowerShell like so:

```bash
$env:GOOS = "linux"; $env:GOARCH = "arm"; $env:GOARM = "5";
```

Either way, you can then fetch and build a binary like so:

```bash
go get github.com/nabsul/suliechat
go build github.com/nabsul/suliechat
```

## Using `install_kano.sh`

I made a little script that can automatically install the chat client on a kano. 
It basically downloads the latest release binary to `/usr/local/bin` and sets permissions 
to executable.

This can be done on a Kano (or Raspberry Pi) line so:

```bash
curl https://raw.githubusercontent.com/nabsul/suliechat/master/install_kano.sh | sudo bash
```

You'll be asked to enter the admin password. The default on a Kano device is `kano`.
