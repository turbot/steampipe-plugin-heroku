---
title: "Steampipe Table: heroku_key - Query Heroku API Keys using SQL"
description: "Allows users to query Heroku API Keys, providing a detailed view of each key’s id, fingerprint, public key, and other related information."
---

# Table: heroku_key - Query Heroku API Keys using SQL

Heroku API Keys are used to authenticate requests to the Heroku Platform API. These keys are tied to a Heroku user and provide a method for the user to interact with the Heroku platform programmatically. API Keys are a critical component of maintaining and managing applications on the Heroku platform.

## Table Usage Guide

The `heroku_key` table provides insights into API Keys within Heroku. As a developer or system administrator, explore key-specific details through this table, including key fingerprints, public keys, and associated metadata. Utilize it to manage and audit API Keys, such as those associated with specific users or applications, and to ensure the security and integrity of your Heroku applications.

**Important Notes**
- Get queries require a key `id`.
- Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all keys
Explore the full range of keys available within your Heroku application. This can help you manage and monitor access to your application effectively.

```sql+postgres
select
  *
from
  heroku_key;
```

```sql+sqlite
select
  *
from
  heroku_key;
```

### Keys older than 90 days
Explore which Heroku keys are older than 90 days. This is useful for maintaining security and ensuring keys are updated regularly.

```sql+postgres
select
  comment,
  created_at,
  date_part('day', now() - created_at) as age_in_days
from
  heroku_key
where
  created_at < now() - interval '90 days';
```

```sql+sqlite
select
  comment,
  created_at,
  julianday('now') - julianday(created_at) as age_in_days
from
  heroku_key
where
  julianday('now') - julianday(created_at) > 90;
```