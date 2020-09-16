package models

type Versions []VersionString

func (v Versions) Latest() string {
	var latest VersionString
	for _, ver := range v {
		if ver.Major >= latest.Major && ver.Minor >= latest.Minor && ver.Hotfix >= latest.Hotfix {
			latest = ver
		}
	}

	return latest.String()
}
