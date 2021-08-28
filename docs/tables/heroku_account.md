# Table: heroku_account

An account represents an individual signed up to use the Heroku platform.

## Examples

### Get account information

```sql
select
  *
from
  heroku_account
```

### Check two factor authentication

```sql
select
  email,
  two_factor_authentication
from
  heroku_account
```
