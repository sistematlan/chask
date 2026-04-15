# Chask — Development Backlog

## Phase 1 — Skeleton (current)

- [ ] **Daemon:** Go binary with Unix socket listener, healthcheck endpoint
- [ ] **Daemon:** Request validation and JSON protocol
- [ ] **Panel:** CodeIgniter 4 app with auth (login/session)
- [ ] **Panel:** ChaskClient library for daemon communication via socket
- [ ] **Panel:** Dashboard view with server metrics (CPU, RAM, disk)
- [ ] **Installer:** Basic install script (Caddy, PHP 8.4, MySQL 8, Go, CI4)
- [ ] **Installer:** systemd service files for chaskd
- [ ] **CI:** GitHub Actions — Go build + PHP lint

## Phase 2 — Sites + PHP

- [ ] **Daemon:** sites.create — generate Caddyfile + FPM pool + directory structure
- [ ] **Daemon:** sites.delete — remove vhost, FPM pool, optionally delete files
- [ ] **Daemon:** sites.list — return all configured sites with status
- [ ] **Daemon:** sites.update — change PHP version, document root
- [ ] **Daemon:** php.versions — list installed PHP versions
- [ ] **Daemon:** php.install — install new PHP version from Ondřej PPA
- [ ] **Panel:** Sites CRUD UI
- [ ] **Panel:** PHP version selector per site

## Phase 3 — Databases

- [ ] **Daemon:** databases.create — MySQL database + user with permissions
- [ ] **Daemon:** databases.drop — remove database and user
- [ ] **Daemon:** databases.list — all databases with sizes
- [ ] **Daemon:** databases.import — import SQL file
- [ ] **Daemon:** databases.export — mysqldump to file
- [ ] **Panel:** Database management UI
- [ ] **Panel:** SQL import/export from browser

## Phase 4 — Firewall + SSL

- [ ] **Daemon:** firewall.list — current UFW rules
- [ ] **Daemon:** firewall.allow / firewall.deny — add rules (port, IP, protocol)
- [ ] **Daemon:** firewall.delete — remove rule
- [ ] **Daemon:** ssl.status — per-site certificate status from Caddy
- [ ] **Daemon:** ssh.keys — list/add/remove authorized keys
- [ ] **Panel:** Firewall rules UI
- [ ] **Panel:** SSL status per site
- [ ] **Panel:** SSH key management UI

## Phase 5 — DNS

- [ ] **Daemon:** dns.zones — list DNS zones
- [ ] **Daemon:** dns.records — CRUD for A, AAAA, CNAME, MX, TXT, NS records
- [ ] **Daemon:** DNS server setup (BIND9 or PowerDNS)
- [ ] **Panel:** DNS zone editor UI

## Phase 6 — Backups

- [ ] **Daemon:** backups.create — tar files + mysqldump per site
- [ ] **Daemon:** backups.restore — restore files + database from backup
- [ ] **Daemon:** backups.schedule — cron-based scheduled backups
- [ ] **Daemon:** backups.list — available backups with sizes/dates
- [ ] **Daemon:** backups.retention — auto-delete old backups
- [ ] **Panel:** Backup management UI (create, restore, schedule, download)

## Future (v2+)

- [ ] Multi-user: admin creates Linux users, each with isolated sites/DBs
- [ ] Resellers: admin → resellers → users hierarchy
- [ ] Email: Postfix + Dovecot integration
- [ ] File Manager: web-based file browser/editor
- [ ] FTP/SFTP: user access to site files
- [ ] Cron Jobs: manage cron from the panel
- [ ] Public API: REST API for automation
- [ ] Nginx: optional web server alternative to Caddy
- [ ] Node.js / Python: support for non-PHP apps
