Proxy for Policy bypass on Imperva WAF
=======================

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/f821affc94304b018f556ebd856e8404)](https://app.codacy.com/app/MonaxGT/ProxyImperva?utm_source=github.com&utm_medium=referral&utm_content=MonaxGT/ProxyImperva&utm_campaign=Badge_Grade_Dashboard)

Proxy with removing header Content-Type.

Run:
```
go run proxy.go -p 127.0.0.1:8080
```

Flags:
```
-p Proxy IP Address (default: ":8080")
```

Docker:
```
docker build . -t proxyimperva
docker run -d -p 8080:8080 proxyimperva
```

Details:
```
Exploit Title: Policy Bypass on Imperva SecureSphere Web Application
Firewall
Date: 08/05/2018
Author: Damien Cabrié
Contact: https://twitter.com/nawhack
Vendor Homepage: http://www.imperva.com
Version: Imperva SecureSphere WAF 11.5 in all deployment options
Tested on: Imperva SecureSphere WAF 11.5.0.95_0 (bridge and reverse proxy
mode)
Class: Policy Bypass
```

[VULNERABILITY DETAILS]

The Imperva WAF provides solutions to protect websites against attacks (SQL
injections, cross site scripting, illegal resource access). The protect is
base with policies building from Signatures (network, generic attack, known
web application vulnerabilities), Application profiling and Threatradar
Reputation Service.

There is a bug in the Web Correlation Policy engine which protect against
SQLi and XSS.

The WAF is not able to detect malicious SQLi or XSS content in the body of
POST requests without the "Content-Type" header.

An attacker can easily craft a POST request method without the Content-Type
header to bypass firewall protections.

Applications protected by the WAF could be compromised with this bug.

[REMEDIATION]
If possible, block POST requests without Content-Type header.

[DISCLOSURE TIME-LINE]
    * 08/05/2018 - Initial vendor contact.

    * 27/06/2018 - Imperva confirmed the issue. Fix targeted for Q4 2018.

    * 12/07/2018 - Ticket closed.

    * 18/08/2018 - Public disclosure.

[Source](https://seclists.org/fulldisclosure/2018/Sep/16?utm_source=feedburner&utm_medium=feed&utm_campaign=Feed%3A+seclists%2FFullDisclosure+%28Full+Disclosure%29)
