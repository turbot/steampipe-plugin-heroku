# Table: heroku_team_member

A team member is an individual with access to a team.

Notes:
* List queries require an `team_name`.

Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all team members

```sql
select
  email,
  role,
  created_at
from
  heroku_team_member;
```

### List all federated team members

```sql
select
  email,
  role,
  created_at,
  is_federated
from
  heroku_team_member
where
  is_federated;
```

### List all team members who haven't enabled two-factor authentication

```sql
select
  email,
  role,
  created_at,
  two_factor_authentication
from
  heroku_team_member
where
  not two_factor_authentication;
```
