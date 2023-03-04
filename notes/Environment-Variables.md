# Environment variables

# Helpful commands

## System wide Environment Variables
You want to be sure that your `$GOPATH/bin` and `$GOROOT/bin` existing in your `PATH` environment variables. In linux your 

### Show your system wide paths
In Linux (`PATH` = string of `:` delimted paths)
```bash
env | grep PATH
```
In Windows CMD (`Path` = string of `;` delimted paths)
```bash
SET | findstr Path
```

### Adding your Go paths to system wide paths
In Linux
```bash
export GOROOT = /ussr/local/go
export GOPATH = $HOME/workspace
PATH = PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

In Windows CMD      
```bash
TODO
```

# Go related variables
List out all your go related environment variables
```bash
go env
```

# Paths
There are two main paths to worry about when using go
 - `GOROOT`, this is the path to your go binary typical location is `/usr/local/go`
 - `GOPATH` this is the path to your "workspace" (where your go projects live)
 