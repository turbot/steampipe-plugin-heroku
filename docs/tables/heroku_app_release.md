---
title: "Steampipe Table: heroku_app_release - Query Heroku App Releases using SQL"
description: "Allows users to query Heroku App Releases, providing detailed information about the version, status, and description of each app release."
---

# Table: heroku_app_release - Query Heroku App Releases using SQL

Heroku is a platform as a service (PaaS) that enables developers to build, run, and operate applications entirely in the cloud. It supports several programming languages and allows developers to deploy, manage, and scale applications without the need for infrastructure. Heroku App Releases are a record of every deploy to an app, and each release has a unique version number that increments with each release.

## Table Usage Guide

The `heroku_app_release` table provides insights into each app release within Heroku. As a developer or a DevOps engineer, you can explore details about each app release, including its version number, status, and description. Utilize this table to track the history of app releases, identify any changes made, and monitor the overall progress of app development.

**Important Notes**
- List queries require an `app_name`.
- Get queries require an `app_name`, a release `id` or a release `version`.
- Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all app releases
Explore all releases of a specific application, identifying their versions, statuses, and creation dates. This can be used to track the app's development history and assess the implementation of current and past versions.

```sql+postgres
select
  id,
  status,
  version,
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe';
```

```sql+sqlite
select
  id,
  status,
  version,
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe';
```

### Get the current release version of an app
Explore the current release version of an app to stay updated on its status and creation date. This is useful for maintaining app version control and ensuring you're working with the most recent release.

```sql+postgres
select
  id,
  status,
  version,
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe' and is_current;
```

```sql+sqlite
select
  id,
  status,
  version,
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe' and is_current;
```

### Get the release information of an app by release ID
Explore the version history of a specific application, such as when the version was created and whether it's the current one, to understand its update timeline and status. This is particularly useful for tracking the app's development and ensuring it's up-to-date.

```sql+postgres
select
  id,
  status,
  version,
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe' and id = 'e8256596-5583-4df0-9a6d-cf0af5e11f02';
```

```sql+sqlite
select
  id,
  status,
  version,
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe' and id = 'e8256596-5583-4df0-9a6d-cf0af5e11f02';
```

### Get the release information of an app by release version
Explore the status of a specific version of an app to understand if it's the current version and when it was created. This can be useful to track the history and updates of the app.

```sql+postgres
select
  id,
  status,
  version,
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe' and version = 4;
```

```sql+sqlite
select
  id,
  status,
  version,
  is_current,
  created_at
from
  heroku_app_release
where
  app_name = 'steampipe' and version = 4;
```