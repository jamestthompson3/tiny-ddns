## DDNS

Tiny library for updating DNS records on digital ocean.

### Usage

ddns does only one thing, update a domain's root A record inside of DigitalOcean with the public IP address of the machine
upon which it runs:

`ddns -domain <domain>`

The only argument it takes is the domain which you want to modify. The domain must exist inside your
DigitalOcean dashboard, or else the API requests will fail. You also must have the nameservers in
your domain registrar pointing to DigitalOcean.

Your DigitalOcean API key _must_ be exported as an environment variable:

`export DIGITALOCEAN_TOKEN='<your-token-here>' ddns -domain <domain>`
