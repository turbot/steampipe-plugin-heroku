---
title: "Steampipe Table: heroku_team_member - Query Heroku Team Members using SQL"
description: "Allows users to query Heroku Team Members, specifically the user roles and permissions within a team, providing insights into team structure and access control."
---

# Table: heroku_team_member - Query Heroku Team Members using SQL

Heroku Team Members are individual user accounts associated with a particular team within the Heroku platform. Each team member has a defined role that determines their permissions and access levels within the team. These roles include admin, member, and owner, each with varying levels of control over team resources and settings.

## Table Usage Guide

The `heroku_team_member` table provides insights into user roles within Heroku teams. As a team manager or DevOps engineer, explore user-specific details through this table, including roles, permissions, and associated metadata. Utilize it to uncover information about team members, such as those with admin permissions, the roles distribution within the team, and the verification of access controls.

**Important Notes**
- List queries require an `team_name`.
- Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all team members
Explore which roles are assigned to each team member and when they were added to the team. This can help in understanding the team's composition and its evolution over time.

```sql+postgres
select
  email,
  role,
  created_at
from
  heroku_team_member;
```

```sql+sqlite
select
  email,
  role,
  created_at
from
  heroku_team_member;
```

### List all federated team members
Explore which team members have federated roles in Heroku. This can be useful for understanding the distribution of roles and identifying potential security implications.

```sql+postgres
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

```sql+sqlite
select
  email,
  role,
  created_at,
  is_federated
from
  heroku_team_member
where
  is_federated = 1;
```

### List all team members who haven't enabled two-factor authentication
Explore which team members have not yet activated two-factor authentication. This is useful for identifying potential security risks within your team and ensuring all members are adhering to best practices for account security.

```sql+postgres
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

```sql+sqlite
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