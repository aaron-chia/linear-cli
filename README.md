# Linear CLI (Fork)

Personal fork of [joa23/linear-cli](https://github.com/joa23/linear-cli) with agent-friendly enhancements. See [README.upstream.md](README.upstream.md) for full upstream documentation.

## Setup

### 1. Install

Build from source and install to your PATH:

```bash
make build
cp bin/linear ~/.local/bin/linear
```

### 2. Authenticate

The CLI checks for credentials in this order:

1. **Stored token** (`~/.config/linear/token`) -- set via `linear auth login`
2. **Environment variable** -- `LINEAR_API_KEY`
3. **Agent secrets** -- `LINEAR_API_KEY` in `~/.agents/secrets.json`

For agent use, the simplest path is adding your API key to agent secrets:

```bash
# Add to ~/.agents/secrets.json:
# { "LINEAR_API_KEY": "lin_api_your_key_here" }
```

Or authenticate interactively:

```bash
linear auth login    # Choose "Agent" mode for bots/automation
linear auth status   # Verify auth and source
```

### 3. Initialize

```bash
linear init          # Select default team, creates .linear.yaml
```

Optionally add a default project to `.linear.yaml`:

```yaml
team: CEN
project: my-project
```

## Skills

Copy the `linear-` prefixed skills to your global Claude Code skills directory:

```bash
cp -r plugin/skills/linear-* ~/.claude/skills/
```

| Skill | Purpose |
|-------|---------|
| `/linear-reference` | Complete CLI command reference |
| `/linear-prd` | Create agent-friendly tickets with PRDs and sub-issues |
| `/linear-triage` | Analyze and prioritize backlog |
| `/linear-cycle-plan` | Plan cycles using velocity analytics |
| `/linear-retro` | Generate sprint retrospectives |
| `/linear-deps` | Analyze dependency chains and blockers |
| `/linear-link-deps` | Discover and link related issues as dependencies |

## Fork Changes

- **Proxy support**: HTTP transport respects `HTTP_PROXY`/`HTTPS_PROXY`/`NO_PROXY`
- **Agent secrets**: Token fallback to `~/.agents/secrets.json`
- **Prefixed skills**: `linear-` prefixed skills for namespace grouping in `~/.claude/skills/`
- **JSON fix**: Detailed JSON output adds truncation hints to comment bodies

## Quick Reference

```bash
linear issues list --format full          # List issues
linear issues get CEN-123                 # Get issue details
linear issues create "Title" --priority 1 # Create issue
linear search "query" --team CEN         # Search
linear cycles analyze --team CEN         # Velocity analysis
linear deps --team CEN                   # Dependency graph
```

All commands support `--output json` for machine-readable output. See [README.upstream.md](README.upstream.md) for the full command reference.
