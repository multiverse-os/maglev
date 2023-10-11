# Starship Yard Templating 
Template caching using a memory key/value store needs to be implemented soon as
possible.

Should use TTL in addition to chunksum sections that can be checked after data
is pulled ajnd rendered to to see if the next load needs to have the cache
refreshed.

### Defaults
This sub-library should define reasonable and simple CSS-only defaults that are
overridden in views folder. This functionality should be very similar to Rails,
including providing the mail template in views (but in Go instead of a
templatinng file like *.erb).
