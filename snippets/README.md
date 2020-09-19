# Decrypting

Authorized users only!

```bash
# Example
curl "https://gitlab.com/nolim1t/documents/-/raw/master/snippets/<FILENAME>" 2>/dev/null | gpg --decrypt

# Get TOR ssh address for pinephone 
curl "https://gitlab.com/nolim1t/documents/-/raw/master/snippets/pinephone-ssh-details.txt" 2>/dev/null | gpg --decrypt
```


# Signing stuff

```bash
# signing
gpg -u <user> --clear-sign <filename>

# verify
gpg --verify <outputfilename>.asc
```

