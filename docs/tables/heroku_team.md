---
title: "Steampipe Table: heroku_team - Query Heroku Teams using SQL"
description: "Allows users to query Heroku Teams, providing information about each team's id, name, membership limit, and more."
---

# Table: heroku_team - Query Heroku Teams using SQL

Heroku Teams is a feature within the Heroku platform that allows users to collaborate on applications. Teams provide a shared, collaborative space where members can access, manage, and deploy applications. It offers a central point for managing access to apps, pipelines, and resources.

## Table Usage Guide

The `heroku_team` table provides comprehensive insights into the teams within Heroku. As an application developer or DevOps engineer, you can use this table to explore details about each team, including its members, associated applications, and limits. This can be particularly useful for managing access to applications, tracking team usage, and ensuring compliance with organizational policies.

**Important Notes**
- Get queries require a team `id`.
- Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all teams
Explore which teams are available and their types to understand the structure and organization within your Heroku environment. This can assist in managing team resources and ensuring appropriate access and permissions.

```sql+postgres
select
  name,
  team_type
from
  heroku_team;
```

```sql+sqlite
select
  name,
  team_type
from
  heroku_team;
```

### Find the default team
Analyze the settings to understand which team has been set as the default in your Heroku application. This can help streamline team management and ensure proper access controls are in place.

```sql+postgres
select
  name,
  team_type
from
  heroku_team
where
  is_default;
```

```sql+sqlite
select
  name,
  team_type
from
  heroku_team
where
  is_default = 1;
```