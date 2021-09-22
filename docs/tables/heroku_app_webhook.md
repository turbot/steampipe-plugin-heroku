# Table: heroku_app_webhook

Webhooks define what web routes should be routed to an app on Heroku.

Notes:
* List queries require an `app_name`.
* Get queries require an `app_name` and webhook `id`.

Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all webhooks

```sql
select
  id,
  url,
  level,
  created_at
from
  heroku_app_webhook
where
  app_name = 'steampipe';
```

### List all notify level webhooks

```sql
select
  id,
  url,
  level,
  created_at
from
  heroku_app_webhook
where
  app_name = 'steampipe' and level = 'notify';
```
