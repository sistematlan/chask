# Chask — Agent Instructions

## What is Chask?

Open-source server panel (cPanel alternative). Runs monolithic on Ubuntu VPS (no Docker).
Two components: a **Go daemon** (chaskd) that runs as root, and a **CodeIgniter 4 panel** served by PHP-FPM.

## Architecture

```
Browser (:2083) → Panel CI4 (PHP-FPM) → Unix socket → chaskd (Go, root) → system operations
```

- Panel NEVER executes system commands directly
- All privileged operations go through the daemon via `/var/run/chaskd.sock`
- Daemon validates every request before executing

## Project Structure

- `daemon/` — Go daemon (chaskd). Entry point: `cmd/chaskd/main.go`
- `panel/` — CodeIgniter 4 app. Communication with daemon via `app/Libraries/ChaskClient.php`
- `scripts/` — Installation scripts. Target: Ubuntu 22.04/24.04
- `docs/` — Architecture and development documentation

## Coding Standards

### Go (daemon)
- Standard Go project layout
- Use `internal/` for non-exported packages
- Error handling: return errors, don't panic
- Logging: use `slog` (structured logging)
- No external dependencies unless absolutely necessary

### PHP (panel)
- CodeIgniter 4 conventions
- Controllers: thin, delegate to models/libraries
- Views: minimal logic, use CI4 view cells for components
- No `console.log` or `var_dump` in committed code
- PSR-12 coding style

### General
- Conventional commits: `feat(sites):`, `fix(daemon):`, `docs:`, etc.
- Scopes: `daemon`, `panel`, `sites`, `databases`, `dns`, `firewall`, `backups`, `installer`
- Keep files under 400 lines
- No Docker — this project runs directly on the host OS
- Security first: validate all inputs, sanitize all outputs, no shell injection

## Communication Protocol (Panel ↔ Daemon)

Unix socket at `/var/run/chaskd.sock`. JSON request/response.

```json
// Request
{
  "action": "sites.create",
  "data": {
    "domain": "example.com",
    "php_version": "8.4"
  }
}

// Response
{
  "status": "ok",
  "data": { ... }
}

// Error
{
  "status": "error",
  "error": "domain already exists"
}
```

## Security Rules

- The daemon runs as root — every input MUST be validated
- Domain names: validate format, no path traversal
- Database names: alphanumeric + underscores only
- Never pass user input directly to shell commands — use Go exec.Command with args
- Never use `shell_exec()` or `exec()` in PHP — always go through the daemon
- UFW rules: validate ports (1-65535), IPs (valid CIDR)

## Development Phases

Current: **Phase 1 — Skeleton**

1. Skeleton — daemon socket + panel auth + communication working
2. Sites + PHP — Caddy vhosts, PHP-FPM pools
3. Databases — MySQL management
4. Firewall + SSL — UFW UI, SSL status
5. DNS — Zone management
6. Backups — Scheduled backups and restore
