# Table: heroku_key

Keys represent public SSH keys associated with an account and are used to authorize accounts as they are performing git operations.

Notes:
* Get queries require a key `id`.

Pagination is not currently supported for this table.

## Examples

### List all keys

```sql
select
  *
from
  heroku_key
```

### Keys older than 90 days

```sql
select
  comment,
  created_at,
  date_part('day', now() - created_at) as age_in_days
from
  heroku_key
where
  created_at < now() - interval '90 days'
```
