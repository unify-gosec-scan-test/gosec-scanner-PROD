# vuln-go-api

A deliberately insecure Golang REST API with common API security flaws. Use for educational/testing purposes only.
 
## Key Vulnerabilities

- 🔓 Hardcoded JWT Secret
- 🔐 Plaintext password storage
- 📂 Directory traversal: `/files/../../etc/passwd`
- 🚫 No input validation or rate limiting
- 🌍 CORS misconfiguration (`*`)
- ⚠️ Verbose logging of secrets

## ⚠️ WARNING

**This code is insecure and should NOT be used in production.**
Use it only in isolated test environments (like local labs or CTFs).
