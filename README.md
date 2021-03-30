# Process Launcher
Start a process when another is detected to be running. 

## Instructions

1. Install [Go](https://golang.org/)
2. Download the source code for Process Launcher
3. Open `Main.go` in a text editor
4. Modify the `application` variable to be the name of the process you wish to wait until is launched. _(Make sure to include the extension, such as `.exe`)_
5. Modify the `directory` variable to be the directory to the file you wish to launch when `application` is launched. _(Make sure to use double backslash)_
6. Save the file.
7. Open your terminal in the directory that contains `Main.go` and run the command `go build -ldflags "-H=windowsgui" Main.go`.
8. Run the `Main.go` executable to start Process Launcher

To stop Process Launcher just navigate to your system tray, look for the running person icon, click on it, then click "Quit"

### Example

Launch "[ProtoSmasher.exe](https://protosmasher.net/)" when "[RobloxBetaPlayer.exe](https://www.roblox.com/download/client)" starts.

```go
const (
	application string = "RobloxPlayerBeta.exe"
	directory string = "C:\\Users\\Andrew\\Desktop\\ProtoSmasher v3 Release\\ProtoSmasher.exe"
)
```

## Dependancies
* https://github.com/mitchellh/go-ps
* https://github.com/getlantern/systray
