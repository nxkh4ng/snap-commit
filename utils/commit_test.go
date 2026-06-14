package utils

import "testing"

func TestFormatCommitMsg(t *testing.T) {
	tests := []struct {
		name           string
		commitType     string
		scope          string
		summary        string
		description    string
		breakingChange string
		want           string
	}{
		{
			name:       "minimal - type and summary only",
			commitType: "feat",
			summary:    "add new feature",
			want:       "feat: add new feature",
		},
		{
			name:       "with scope",
			commitType: "fix",
			scope:      "api",
			summary:    "handle timeout",
			want:       "fix(api): handle timeout",
		},
		{
			name:        "with description",
			commitType:  "feat",
			summary:     "add login",
			description: "implement OAuth2 login flow",
			want:        "feat: add login\n\nimplement OAuth2 login flow",
		},
		{
			name:           "with breaking change",
			commitType:     "feat!",
			summary:        "rewrite auth",
			breakingChange: "drops v1 API",
			want:           "feat!: rewrite auth\n\nBREAKING CHANGE: drops v1 API",
		},
		{
			name:           "full - all fields",
			commitType:     "feat!",
			scope:          "api",
			summary:        "add v2 endpoints",
			description:    "migrate to new router",
			breakingChange: "removes v1 endpoints",
			want:           "feat(api)!: add v2 endpoints\n\nmigrate to new router\n\nBREAKING CHANGE: removes v1 endpoints",
		},
		{
			name:       "trims whitespace from inputs",
			commitType: "  chore  ",
			scope:      "  deps  ",
			summary:    "  update deps  ",
			want:       "chore(deps): update deps",
		},
		{
			name:       "scope is lowercased",
			commitType: "docs",
			scope:      "README",
			summary:    "fix typo",
			want:       "docs(readme): fix typo",
		},
		{
			name:       "empty scope omitted",
			commitType: "refactor",
			scope:      "",
			summary:    "simplify loop",
			want:       "refactor: simplify loop",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatCommitMsg(tt.commitType, tt.scope, tt.summary, tt.description, tt.breakingChange)
			if got != tt.want {
				t.Errorf("FormatCommitMsg() =\n%q\nwant:\n%q", got, tt.want)
			}
		})
	}
}

func TestParseCommitMsg(t *testing.T) {
	tests := []struct {
		name               string
		msg                string
		wantCommitType     string
		wantScope          string
		wantSummary        string
		wantDescription    string
		wantBreakingChange string
		wantHasBreaking    bool
	}{
		{
			name:               "full conventional commit",
			msg:                "feat(api): add endpoint\n\nimplement new route\n\nBREAKING CHANGE: drops v1",
			wantCommitType:     "feat",
			wantScope:          "api",
			wantSummary:        "add endpoint",
			wantDescription:    "implement new route",
			wantBreakingChange: "drops v1",
			wantHasBreaking:    false,
		},
		{
			name:               "breaking change with !",
			msg:                "feat!: rewrite\n\ndescription\n\nBREAKING CHANGE: major refactor",
			wantCommitType:     "feat!",
			wantScope:          "",
			wantSummary:        "rewrite",
			wantDescription:    "description",
			wantBreakingChange: "major refactor",
			wantHasBreaking:    true,
		},
		{
			name:               "no scope, no body",
			msg:                "fix: typo",
			wantCommitType:     "fix",
			wantScope:          "",
			wantSummary:        "typo",
			wantDescription:    "",
			wantBreakingChange: "",
			wantHasBreaking:    false,
		},
		{
			name:               "with scope, no body",
			msg:                "docs(readme): update install",
			wantCommitType:     "docs",
			wantScope:          "readme",
			wantSummary:        "update install",
			wantDescription:    "",
			wantBreakingChange: "",
			wantHasBreaking:    false,
		},
		{
			name:               "breaking with scope and !",
			msg:                "feat(api)!: add v2",
			wantCommitType:     "feat!",
			wantScope:          "api",
			wantSummary:        "add v2",
			wantDescription:    "",
			wantBreakingChange: "",
			wantHasBreaking:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotType, gotScope, gotSummary, gotDesc, gotBC, gotHasBreaking := ParseCommitMsg(tt.msg)

			if gotType != tt.wantCommitType {
				t.Errorf("commitType = %q, want %q", gotType, tt.wantCommitType)
			}
			if gotScope != tt.wantScope {
				t.Errorf("scope = %q, want %q", gotScope, tt.wantScope)
			}
			if gotSummary != tt.wantSummary {
				t.Errorf("summary = %q, want %q", gotSummary, tt.wantSummary)
			}
			if gotDesc != tt.wantDescription {
				t.Errorf("description = %q, want %q", gotDesc, tt.wantDescription)
			}
			if gotBC != tt.wantBreakingChange {
				t.Errorf("breakingChange = %q, want %q", gotBC, tt.wantBreakingChange)
			}
			if gotHasBreaking != tt.wantHasBreaking {
				t.Errorf("hasBreaking = %v, want %v", gotHasBreaking, tt.wantHasBreaking)
			}
		})
	}
}
