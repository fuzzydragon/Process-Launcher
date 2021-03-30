# Process Launcher
Start a process when another is detected to be running. 

## Instructions

1. Modify the `application` variable to be the name of the process you wish to wait until is launched. _(Make sure to include the extension, such as `.exe`)_
2. Modify the `directory` variable to be the directory to the file you wish to launch when `application` is launched. _(Make sure to use double backslash)_

To stop Process Launcher just navigate to your system tray, look for the running person icon, click on it, then click "Quit". 

### Example

Launch "[ProtoSmasher.exe](https://protosmasher.net/)" when "[RobloxBetaPlayer.exe](https://www.roblox.com/download/client)" starts

```go
const (
	application string = "RobloxPlayerBeta.exe"
	directory string = "C:\\Users\\Andrew\\Desktop\\ProtoSmasher v3 Release\\ProtoSmasher.exe"
)
```

## Dependancies
* https://github.com/mitchellh/go-ps
* https://github.com/getlantern/systray
