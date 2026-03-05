// Package skills provides embedded skill templates for Claude Code.
package skills

import "embed"

//go:embed linear-reference/* linear-prd/* linear-triage/* linear-cycle-plan/* linear-retro/* linear-deps/* linear-link-deps/*
var SkillFiles embed.FS

// SkillInfo describes an available skill
type SkillInfo struct {
	Name        string
	Description string
	Dir         string
}

// AvailableSkills lists all skills that can be installed
var AvailableSkills = []SkillInfo{
	{
		Name:        "linear-reference",
		Description: "Linear issue tracking - MUST READ before using Linear commands",
		Dir:         "linear-reference",
	},
	{
		Name:        "linear-prd",
		Description: "Create agent-friendly tickets with PRDs, sub-issues, and success criteria",
		Dir:         "linear-prd",
	},
	{
		Name:        "linear-triage",
		Description: "Analyze and prioritize Linear backlog based on staleness and blockers",
		Dir:         "linear-triage",
	},
	{
		Name:        "linear-cycle-plan",
		Description: "Plan cycles using velocity analytics and capacity",
		Dir:         "linear-cycle-plan",
	},
	{
		Name:        "linear-retro",
		Description: "Generate sprint retrospective analysis from completed cycles",
		Dir:         "linear-retro",
	},
	{
		Name:        "linear-deps",
		Description: "Analyze dependency chains, find blockers and circular dependencies",
		Dir:         "linear-deps",
	},
	{
		Name:        "linear-link-deps",
		Description: "Discover and link related issues as dependencies across your backlog",
		Dir:         "linear-link-deps",
	},
}

// GetSkillByName returns a skill by name, or nil if not found
func GetSkillByName(name string) *SkillInfo {
	for _, s := range AvailableSkills {
		if s.Name == name {
			return &s
		}
	}
	return nil
}
