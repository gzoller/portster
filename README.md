# portster
Inside-the-box Docker REST service to easily access externally-mapped ports.

GET /port/<num> - Returns a port # of externally-mapped port or 404 if not mapped.

There's also an error log forwarding feature:

POST /log/info
POST /log/warn
POST /log/error

Where the body of the post is any free-form string you want to log.  These get forwarded to wherever you specify in a /var/logfwd.json file like this:

{
  "info":"http://...",
  "warn":"http://...",
  "error":"http://..."
}

If one or more (or any) of these aren't specified logging defaults to STDERR.
