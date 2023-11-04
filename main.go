package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/ronaldr1985/StayAwake/icons/disabledicon"
	"github.com/ronaldr1985/StayAwake/icons/enabledicon"

	"github.com/gen2brain/beeep"
	"github.com/getlantern/systray"
	"github.com/ncruces/zenity"
	ge "github.com/ronaldr1985/graceful-exit"
	"gopkg.in/yaml.v3"
)

const (
	DEFAULT_TIME_INBETWEEN_NOTIFICATIONS = 20 * time.Minute
	ALERT_NOTIFICATION_ICON_URL          = ""
	DEFAULT_CONFIG_FILE                  = "Interval: 20\nDarkTheme: true"
)

type Config struct {
	Interval  int  `yaml:"Interval"`
	DarkTheme bool `yaml:"DarkTheme"`
}

var (
	DEFAULT_UNIX_DIRECTORY    = os.Getenv("HOME") + "/.config/StayAwake/" // Think of this as const
	DEFAULT_WINDOWS_DIRECTORY = os.Getenv("LOCALAPPDATA") + "/StayAwake/" // Think of this as const
	AlertNotificationIcon     = ""
	ConfigFile                = ""
	ProgramConfig             Config
)

func CheckIfFileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return false
	}
}

func WriteConfig(filename string, config Config) error {
	file, err := os.OpenFile(filename, os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	yamlEncoder := yaml.NewEncoder(file)

	err = yamlEncoder.Encode(config)
	if err != nil {
		return err
	}

	return nil
}

func ReadConfig(filename string) (Config, error) {
	config := &Config{}

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return *config, err
	}

	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return *config, fmt.Errorf("in file %q: %w", filename, err)
	}

	return *config, err
}

func onReady() {
	var enabled bool = true
	var disabledIcon []byte
	systray.SetIcon(enabledicon.Data)
	systray.SetTitle("Stay Awake")
	systray.SetTooltip("Stay Awake")
	mChangeGUI := systray.AddMenuItem(
		"Change interval", "Change how often a key is pressed",
	)
	mDarkTheme := systray.AddMenuItemCheckbox(
		"Dark Mode", "Use the dark theme for icons", ProgramConfig.DarkTheme,
	)
	systray.AddSeparator()
	mEnabled := systray.AddMenuItemCheckbox(
		"Enabled", "Whether we should keep the screen on", true,
	)
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()

	if ProgramConfig.DarkTheme {
		disabledIcon = disabledicon.LightIcon
	} else {
		disabledIcon = disabledicon.DarkIcon
	}

	go func() {
		nextTimeToPressKeys := time.Now().Add(time.Duration(-ProgramConfig.Interval) * time.Second)
		for {
			if enabled && time.Since(nextTimeToPressKeys) > time.Duration(ProgramConfig.Interval)*time.Second {
				PressAndReleaseF24Key()

				nextTimeToPressKeys = time.Now()
			}
		}
	}()

	go func() {
		for {
			select {
			case <-mEnabled.ClickedCh:
				if mEnabled.Checked() {
					mEnabled.Uncheck()
					systray.SetIcon(disabledIcon)
					enabled = false
				} else {
					systray.SetIcon(enabledicon.Data)
					mEnabled.Check()
					enabled = true
				}
			case <-mChangeGUI.ClickedCh:
				for {
					entered_seconds, err := zenity.Entry(
						"Enter a number of seconds",
						zenity.Title("StayAwake, enter a number of seconds"),
					)

					if err != nil {
						if strings.Contains(err.Error(), "not found") {
							err := beeep.Alert(
								"Error",
								"Error: Zenity is not installed, please install Zenity if you wish to use the change interval GUI",
								"assets/warning.png",
							)
							if err != nil {
								fmt.Fprintln(os.Stderr, "Failed to send alert failed with:", err)
							}

							fmt.Fprintln(os.Stderr, "Zenity not installed")
						}

						break
					}

					if ProgramConfig.Interval, err = strconv.Atoi(entered_seconds); err != nil {
						fmt.Fprintln(os.Stderr, "Failed to convert string to integer...")
						fmt.Fprintln(os.Stderr, "Got the following error:", err)
					} else {
						break
					}
				}

				WriteConfig(ConfigFile, ProgramConfig)
			case <-mDarkTheme.ClickedCh:
				if mDarkTheme.Checked() {
					mDarkTheme.Uncheck()
					disabledIcon = disabledicon.LightIcon
				} else {
					mDarkTheme.Check()
					disabledIcon = disabledicon.DarkIcon
				}

				if !enabled {
					systray.SetIcon(disabledIcon)
				}

				ProgramConfig.DarkTheme = !ProgramConfig.DarkTheme

				WriteConfig(ConfigFile, ProgramConfig)
			}
		}
	}()

}

func main() {
	ge.HandleSignals(false)

	ProgramConfig.Interval = int(DEFAULT_TIME_INBETWEEN_NOTIFICATIONS)

	var possibleAppDirectoryLocations []string
	possibleAppDirectoryLocations = append(possibleAppDirectoryLocations, os.Getenv("HOME")+"/.config/StayAwake/")
	possibleAppDirectoryLocations = append(possibleAppDirectoryLocations, os.Getenv("HOME")+"/.config/StayAwake/")
	possibleAppDirectoryLocations = append(possibleAppDirectoryLocations, os.Getenv("XDG_CONFIG_HOME")+"/StayAwake/")
	possibleAppDirectoryLocations = append(possibleAppDirectoryLocations, os.Getenv("XDG_CONFIG_HOME")+"/StayAwake/")
	possibleAppDirectoryLocations = append(possibleAppDirectoryLocations, os.Getenv("LOCALAPPDATA")+"/StayAwake/")
	possibleAppDirectoryLocations = append(possibleAppDirectoryLocations, os.Getenv("LOCALAPPDATA")+"/StayAwake/")

	appDirectory := "not found"
	for _, folder := range possibleAppDirectoryLocations {
		if _, err := os.Stat(folder); !os.IsNotExist(err) {
			appDirectory = folder
			break
		}
	}

	if appDirectory == "not found" {
		if runtime.GOOS == "windows" {
			appDirectory = DEFAULT_WINDOWS_DIRECTORY
		} else {
			appDirectory = DEFAULT_UNIX_DIRECTORY
		}

		fmt.Println("Creating config directory:", appDirectory)
		err := os.Mkdir(appDirectory, 0755)
		fmt.Println("Created folder:", appDirectory)
		if err != nil {
			panic(err)
		}
	}

	if fileExists := CheckIfFileExists(appDirectory + "config.yaml"); fileExists {
		ConfigFile = appDirectory + "config.yaml"
	} else {
		if fileExists := CheckIfFileExists(appDirectory + "config.yml"); fileExists {
			ConfigFile = appDirectory + "config.yml"
		}
	}

	if ConfigFile != "" {
		fmt.Println("Found config file at", ConfigFile)
	}

	if ConfigFile == "" {
		ConfigFile = appDirectory + "config.yaml"

		fmt.Println("Creating config file:", ConfigFile)
		f, err := os.Create(ConfigFile)
		if err != nil {
			panic("Failed to create config file")
		}
		fmt.Println("Created config file:", ConfigFile)

		fmt.Println("Writing default config to", ConfigFile)
		_, err = f.WriteString(DEFAULT_CONFIG_FILE)
		if err != nil {
			f.Close()
			panic("Failed to write to config file: " + ConfigFile)
		}
		f.Close()

		fmt.Println("Written default config to ", ConfigFile)
	}

	var err error
	ProgramConfig, err = ReadConfig(ConfigFile)
	if err != nil {
		panic("Failed to read config file")
	}

	onExit := func() {
		return
	}

	systray.Run(onReady, onExit)
}
