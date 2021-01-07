# StayAwake

StayAwake is simple Go program that runs on Windows in the system tray and stops the screen from sleeping.

### Usage

When you run the application there will be a small eye icon that appears in the system tray, you can than right/left click to disable the application or exit. There's a couple of screenshots and some more details about StayAwake on my personal website and blog found [here](https://reganm.xyz/blog/stayawake.html)

## Compiling

 To compile stayawake use the following command:
 
    go build -ldflags="-H windowsgui" 

Note that compiling requires GCC in the PATH variable.