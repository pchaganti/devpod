package config

type Result struct {
	ContainerDetails    *ContainerDetails
	MergedConfig        *MergedDevContainerConfig
	SubstitutionContext *SubstitutionContext
}

func GetVSCodeConfiguration(mergedConfig *MergedDevContainerConfig) *VSCodeCustomizations {
	if mergedConfig.Customizations == nil || mergedConfig.Customizations["vscode"] == nil {
		return &VSCodeCustomizations{}
	}

	retVSCodeCustomizations := &VSCodeCustomizations{
		Settings:   map[string]interface{}{},
		Extensions: nil,
	}
	for _, customization := range mergedConfig.Customizations["vscode"] {
		vsCode := &VSCodeCustomizations{}
		err := Convert(customization, vsCode)
		if err != nil {
			continue
		}

		for _, extension := range vsCode.Extensions {
			if contains(retVSCodeCustomizations.Extensions, extension) {
				continue
			}

			retVSCodeCustomizations.Extensions = append(retVSCodeCustomizations.Extensions, extension)
		}

		for k, v := range vsCode.Settings {
			retVSCodeCustomizations.Settings[k] = v
		}
	}

	return retVSCodeCustomizations
}

func contains(stack []string, k string) bool {
	for _, s := range stack {
		if s == k {
			return true
		}
	}
	return false
}