# StayAwake

StayAwake is simple Go program that runs on Windows in the system tray and stops the screen from sleeping.

### Usage

When you run the application there will be a small eye icon that appears in the system tray, you can than right/left click to disable the application or exit.

## Compiling

 To compile StayAwake on Windows use the following command:
 
    go build -ldflags="-H windowsgui" 
