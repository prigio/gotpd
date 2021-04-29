# gotpd
A simple golang-based web server serving only one static folder.

This is a dead-simple webserver serving (static) contents of the specified folder on the specified (http-only) address.

This program is maintained at https://github.com/prigio/gotpd

Command-line parameters:

```
  -addr string
    	IP:Port to use for the webserver (default ":8080")
  -folder string
    	Folder to serve static content from (default "/var/www/")
 
```

## Examples

```bash
    gotpd -addr :8080 -folder /var/www &
```

## Building

A makefile is provided