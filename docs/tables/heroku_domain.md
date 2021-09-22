# Table: heroku_domain

Domains define what web routes should be routed to an app on Heroku.

Notes:
* List queries require an `app_name`.
* Get queries require an `app_name` and a domain `id`.

Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all domains

```sql
select
  id,
  status,
  kind,
  hostname
from
  heroku_domain
where
  app_name = 'steampipe';
```

### List all custom domains

```sql
select
  id,
  status,
  kind,
  hostname
from
  heroku_domain
where
  app_name = 'steampipe' and kind = 'custom';
```
