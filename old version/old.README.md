# My Tool
A command line tool to execute the penration testing and OS exploit tool. Adding some base knowledge for me.
The vision would be using this tool to do all the work, when in need to penetration test or OS exploit, and if find the bug, return the result, and is able to request more information related to the vulnerability, such as vulnerbility category name.

## Bacis
- [x] MacOS
- [x] Windows
    - https://appuals.com/fix-account-restrictions-are-preventing-this-user-from-signing-in/
- [ ] Linux
- [ ] download tools
    - [ ] MacOS
        - [x] nmap
        - [x] sqlmap
    - [ ] Linux
        - [x] nmap, not tested
        - [x] sqlmap, not tested

## Peneration Test & Red Team
### Basic Command
- [x] help
- [x] ignore, proceed if HTTP response code is not 200
- [x] ignore additional URLs, one URL at a times
- [ ] a list of URLs

### Basic Feature
- [ ] output format
- [ ] loop until exit
- [x] parse flag first then deal with the output
- [x] flag can be placed anywhere in command line
    - [x] deal with command conflict
        -  [x] generate output if multiple flags contains have overlap
    - [x] format the output of multiple flags
- [x] handle error flag
- [x] use mytool globally
- [x] loading icon
- [x] check if webpage context contains "error"
- [ ] list of URLs
- [ ] list of admin names, usernames
- [ ] find username and password input boxes in websites

### Reconnaissance
- [x] URL response (200, 404 etc.)
    - [x] URL
    - [x] response code
    - [x] other details, protocal version, header etc.
- [x] subdomains enumeration
- [ ] directory search
- [ ] nmap
    - [ ] list
        - [ ] -iL <inputfilename>: Input from list of hosts/networks
    - [ ] scan user input command-line for nmap command 
    - [x] scan subdomain enumeration results
        - 1: -A
    - [ ] DNS
        - [ ] -n (No DNS resolution)
        - [ ] -R (DNS resolution for all targets)
        - [ ] --system-dns (Use system DNS resolver)
        - [ ] --dns-servers <server1>[,<server2>[,...]] (Servers to use for reverse DNS queries)
    - [ ] WHOIS, nmap --script whois-domain.nse <target>
    - [ ] scan directory search results
    - [ ] exception
        - [ ] --exclude <host1[,host2][,host3],...>: Exclude hosts/networks
- [ ] brute force
- [ ] domain enumeration
    - [ ] OWASP Amass

### Resource Development
- [ ] web shell
- [ ] server-side DNS C2 traffic

### Initial Access
- [ ] SSH
- [ ] SMB
- [ ] SQL Injection
    - [ ] url encode
    - [x] find database version
    - [x] Generic SQL Injection
    - [x] Auth Bypass SQL Injection
        - [x] login
    - [ ] UNION SQL Injection
        - [ ] determine the database (Oracle, Microsoft, PostgreSQL, MySQL)
        - [ ] determind "SELECT" in database
        - [ ] count colums (1, 2, 3, ..., aoumnt of NULL)
        - [ ] swap NULL with disire things, such as "BANNER" and "FROM+v$version"
        - [x] UNION SELECT (NULL, ... ,NULL--)
        - [x] find credential in HTML
        - [x] login with credential
        - [x] extract domain
    - [ ] blind SQL Injection
        - [X] extract cookie
        - [ ] Inferential SQLi (Blind SQLi)
        - [ ] Boolean-based (content-based) Blind SQLi
        - [ ] Time-based Blind SQLi
- [ ] XSS
    - [ ] Stored XSS
    - [x] Reflected XSS
    - [ ] DOM-based XSS
- [ ] SSRF
    - [ ] PHP
    - [ ] redis (RESP)
    - [ ] gopher
- [ ] SQLmap


### Execution


### Persistence


### Privilege Escalation


### Defense Evasion


### Credential Access


### Discovery


### Lateral Movement


### Collection


### Command and Control


### Exfiltration


### Impact

### Feature
- [ ] FFUF commands lookup
- [ ] netcat commands lookup
- [ ] execute CVE PoC
- [ ] matching with OWASP Top 10
    - [ ] detect the target version
    - [ ] the target version is vulnerable or not
        - [ ] CVE api
    - [ ] exploitation/script

### Attack
- [ ] LDAP
- [ ] File Upload
    - [ ] PHP
    - [ ] jpg


## OS Exploit
### Feature
- [ ] Windows stack overflow
- [ ] Linux stack overflow
- [ ] Windows Heap-based Buffer Overflow
- [ ] Linux Heap-based Buffer Overflow
- [ ] Windows Integer Overflow
- [ ] Linux Integer Overflow
- [ ] Windows ROP
- [ ] Linux ROP
- [ ] Windows ASLR
- [ ] Linux ASLR
- [ ] DEP
- [ ] NX
- [ ] CFG
- [ ] CFI
- [ ] SMEP
- [ ] SMAP

## Documents
- [flag](https://pkg.go.dev/flag)
