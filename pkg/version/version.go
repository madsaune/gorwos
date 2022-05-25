package version

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	Version   string
	BuildDate string
}

// Version of gorwos
var Version = "dev"

// BuildDate of the release
var BuildDate = "unknown"

func GetVersionInfo() Info {
	return Info{
		Version:   Version,
		BuildDate: BuildDate,
	}
}

func AsJSON() string {
	if data, err := json.MarshalIndent(GetVersionInfo(), "", " "); err == nil {
		return string(data)
	}

	return ""
}

func String() string {
	v := GetVersionInfo()
	return fmt.Sprintf("gorwos version %s (%s)\n", v.Version, v.BuildDate)
}
