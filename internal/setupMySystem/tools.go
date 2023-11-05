package setupMySystem

var toolBelt []*Tool

type toolCategory string

type userProfile string

const (
	Utility  toolCategory = "utility"
	Services toolCategory = "services"
	Language toolCategory = "language"
)

const (
	Developer userProfile = "developer"
	Devops    userProfile = "devops"
)

type Tool struct {
	name         string
	isDefault    bool
	category     toolCategory
	userProfiles []userProfile
}

func addToolToBelt(name string, isDefault bool, category toolCategory, userProfiles []userProfile) {
	tool := &Tool{
		name:         name,
		isDefault:    isDefault,
		category:     category,
		userProfiles: userProfiles,
	}
	toolBelt = append(toolBelt, tool)
}

func GetToolBelt() []*Tool {
	return toolBelt
}

func InitToolBelt() {
	addToolToBelt("watch", true, Utility, []userProfile{Developer, Devops})
	addToolToBelt("stern", true, Utility, []userProfile{Developer, Devops})
	addToolToBelt("wget", true, Utility, []userProfile{Developer, Devops})
	addToolToBelt("zsh", false, Utility, []userProfile{Developer, Devops})
	addToolToBelt("git", true, Utility, []userProfile{Developer, Devops})
	addToolToBelt("mitmproxy", false, Utility, []userProfile{Developer})
	addToolToBelt("kubectl", true, Utility, []userProfile{Developer, Devops})
	addToolToBelt("thefuck", true, Utility, []userProfile{Developer, Devops})
	addToolToBelt("h3", false, Utility, []userProfile{Developer})
	addToolToBelt("nvm", true, Utility, []userProfile{Developer})

	addToolToBelt("kafka", false, Services, []userProfile{Developer})

	addToolToBelt("node", true, Language, []userProfile{Developer})
}
