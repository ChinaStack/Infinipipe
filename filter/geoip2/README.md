gogstash geoip2 filter module
=============================

You need to download geoip2 database manually and setup file path in config.

* http://dev.maxmind.com/geoip/geoip2/downloadable/
* http://dev.maxmind.com/geoip/geoip2/geolite2/

The database is automatically reloaded on changes.

## Synopsis

```yaml
filter:
  - type: geoip2
    # (required) geoip2 database file path, default: GeoLite2-City.mmdb
    db_path: "geoip2/database/file/path/GeoLite2-City.mmdb"

    # (required) ip address field to parse
    ip_field: remote_addr

    # (optional) does not try to process private IP networks as they will fail, default: false
    skip_private: true

    # (optional) lets you specify your own definition for private IP addresses, both IPv4 and IPv6, default is private IP addresses
    private_net:
      - 10.0.0.0/8
      - 192.168.0.0/16

    # (optional) size of cache entries on IP addresses, so lookups don't go through the database, default is 100000
    cache_size: 100000

    # (optional) parsed geoip info should saved to field, default: geoip
    key: geoip

    # (optional) parsed geoip info into flat format, default: false
    # `city_name`, `continent_code`, `country_code`, `country_name`,
    # `ip`, `latitude`, `longitude`, `postal_code`, `region_code`, `region_name` and `timezone`.
    flat_format: false

    # (optional) if true does not log lookup failures from the database, default is false
    quiet: true
```
