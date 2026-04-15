<p align="center">
  <h1 align="center">Chask</h1>
</p>

<p align="center"><strong>Open-source server panel. Fast, lightweight, no Docker.</strong></p>

<p align="center">
A modern alternative to cPanel, built for small VPS.<br>
Manage sites, databases, DNS, SSL, firewall, and backups from a single panel.
</p>

> **Status:** Early development. Not ready for production.

---

## Why Chask?

- **cPanel** costs ~$45/month per VPS
- **Docker-based panels** (Coolify, Dokku) eat too much RAM on small servers
- **Chask** runs monolithic — no containers, no overhead. Panel + daemon use ~20MB RAM total

Named after the **Chasqui** — Inca empire messengers who ran across the Andes delivering orders. Fast, reliable, efficient.

## Architecture

```
Browser (:2083)
    │
    ▼
Panel (CodeIgniter 4 + PHP-FPM)
    │
    │ Unix socket (/var/run/chaskd.sock)
    │ JSON API
    │
    ▼
chaskd (Go daemon, root)
    │
    ├─ Sites      → Caddy vhosts, HTTPS automatic
    ├─ Databases  → MySQL create/drop/import/export
    ├─ PHP        → Multiple versions via FPM pools
    ├─ DNS        → Zone management
    ├─ Firewall   → UFW rules
    ├─ Backups    → Scheduled, per-site, restore
    └─ System     → CPU, RAM, disk metrics
```

## Stack

| Component | Technology | Why |
|-----------|-----------|-----|
| Web Server | Caddy | Auto HTTPS, simple config, PHP FastCGI, hot reload |
| Panel | CodeIgniter 4 | ~15MB RAM, fast boot, lightweight |
| Daemon | Go | Single binary, ~5MB RAM, no dependencies |
| Database | MySQL 8 | Standard, reliable |
| Firewall | UFW | Simple, effective |
| Target OS | Ubuntu 22.04 / 24.04 | Most popular VPS OS |

## Modules (v1)

- **Sites** — Create/edit/delete Caddy vhosts, per-site PHP version
- **Databases** — MySQL databases and users, import/export SQL
- **SSL** — Automatic via Caddy (Let's Encrypt), zero config
- **PHP** — Multiple versions (8.4, 8.1, 7.4) with per-site FPM pools
- **DNS** — Zone management (A, AAAA, CNAME, MX, TXT)
- **Firewall** — UFW rules from the panel UI
- **Backups** — Scheduled backups (files + DB), restore from panel
- **Dashboard** — Server overview (CPU, RAM, disk, uptime)

## Installation

```bash
curl -sSL https://get.chask.dev | bash
```

> Not available yet. See [Development](#development) to run locally.

## Development

### Requirements

- Go 1.22+
- PHP 8.4 + Composer
- MySQL 8
- Caddy 2

### Setup

```bash
git clone https://github.com/sistematlan/chask.git
cd chask

# Build daemon
cd daemon && go build -o chaskd ./cmd/chaskd

# Install panel
cd ../panel && composer install
cp env .env
# Edit .env with your settings
php spark migrate
php spark db:seed AdminSeeder
```

## Project Structure

```
chask/
├── daemon/                  ← Go daemon (chaskd)
│   ├── cmd/chaskd/          ← main entry point
│   └── internal/
│       ├── api/             ← Unix socket JSON API handlers
│       └── system/          ← System operations (sites, db, firewall...)
├── panel/                   ← CodeIgniter 4 web panel
│   ├── app/
│   │   ├── Controllers/     ← Panel controllers
│   │   ├── Models/          ← Data models
│   │   ├── Views/           ← UI templates
│   │   └── Libraries/       ← ChaskClient (socket communication)
│   └── public/              ← Web root
├── scripts/                 ← Installation and helper scripts
│   └── install.sh           ← Main installer
├── docs/                    ← Documentation
│   └── ARCHITECTURE.md      ← Detailed architecture doc
├── CLAUDE.md                ← Agent instructions
├── BACKLOG.md               ← Development roadmap
└── VERSION                  ← Current version
```

## Roadmap

See [BACKLOG.md](BACKLOG.md) for the full roadmap.

**v1 (current):** Sites, Databases, SSL, PHP, DNS, Firewall, Backups, Dashboard
**v2:** Multi-user (admin → Linux users), File Manager, FTP/SFTP, Cron jobs
**v3:** Resellers, Email (Postfix + Dovecot), API, Nginx support

## License

MIT — see [LICENSE](LICENSE).
