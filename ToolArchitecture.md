---
markmap:
  colorFreezeLevel: 3
---

# Tool Architecture
## Command
### Common
- u
    - Show HTTP response status code.
- s
    - Subdomains Enumeration
- o
    - Operation Type
        - A for aggressive operation.
        - P for passive operation.
- h
    - Show flags useages.
- d
    - Show futher details.
- i
    - Proceed when HTTP response is not 200.

### Red Team & Peneration Test
- xss
    - Use XSS attack.
- sql
    - Use SQL Injection attack.

## Function
### General
- TargetStatus
- ReadPayload
- Login
- DomainPath
    - not used
- FindCred
    - not used
- SubdomainEnumeration
    - not used

### Tool
- NmapExecution

### CrossSiteScript
- ReflectedXSS

### SQLInjection
- SQLdbInfo
- SQLdbVersion
- GenericSQLInjection
- AuthBypassSQLInjection
- UnionSQLInjection
- BlindSQLInjection


