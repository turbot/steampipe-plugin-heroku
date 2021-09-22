# Table: heroku_team

Teams allow you to manage access to a shared group of applications and other resources.

Notes:
* Get queries require a team `id`.

Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all teams

```sql
select
  name,
  team_type
from
  heroku_team
```

### Find the default team

```sql
select
  name,
  team_type
from
  heroku_team
where
  is_default
```
