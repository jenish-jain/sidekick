package setupMySystem

var toolBelt []*Tool

type toolCategory string

type userProfile string

const (
	Utility toolCategory = "utility"
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

func NewTool(name string, isDefault bool, category toolCategory, userProfiles []userProfile) *Tool {
	tool := &Tool{
		name:         name,
		isDefault:    isDefault,
		category:     category,
		userProfiles: userProfiles,
	}
	toolBelt = append(toolBelt, tool)
	return tool
}

func GetToolBelt() []*Tool {
	return toolBelt
}

var the_fuck = NewTool("thefuck", true, Utility, []userProfile{Developer, Devops})
