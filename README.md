# Active Directory Groups Guard (ADGG)

This light CLI utility can be used to detect unwanted changes to a list of AD groups.
By running this CLI every few minutes, you'll be the first one to know if a hostile AD takeover is taking place.

## Software settings
Copy and paste `.env.example` to `.env` and adapt your settings (AD address, service account, mail gateway).

## Default groups to watch:
- Account Operators
- Administrators
- Backup Operators
- Domain Admins
- Domain Controllers
- Enterprise Admins
- Enterprise Read-only Domain Controllers
- Group Policy Creator Owners
- Incoming Forest Trust Builders
- Microsoft Exchange Servers
- Network Configuration Operators
- Power Users
- Print Operators
- Read-only Domain Controllers
- Replicators
- Schema Admins
- Server Operators

## Add or edit monitored groups:
Edit `config/groups.txt` with one AD group per line