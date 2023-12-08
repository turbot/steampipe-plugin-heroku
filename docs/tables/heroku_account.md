---
title: "Steampipe Table: heroku_account - Query Heroku Accounts using SQL"
description: "Allows users to query Heroku Accounts, providing insights into account details like email, identity provider, last login, created and updated times."
---

# Table: heroku_account - Query Heroku Accounts using SQL

Heroku is a cloud platform as a service (PaaS) that lets companies build, deliver, monitor, and scale apps. The platform is flexible for developers, intuitive for teams, and built with love for the users. It is designed for developers to build and run applications entirely in the cloud.

## Table Usage Guide

The `heroku_account` table provides insights into Heroku Accounts within the Heroku platform. As a DevOps engineer, explore account-specific details through this table, including email, identity provider, last login, created and updated times. Utilize it to uncover information about accounts, such as those with specific identity providers, the last login times, and the creation and update times.

## Examples

### Get account information
Explore the details of your Heroku account to gain insights into various aspects such as billing, usage, and settings. This can be useful for auditing, optimizing costs, and ensuring the account is configured correctly.

```sql+postgres
select
  *
from
  heroku_account;
```

```sql+sqlite
select
  *
from
  heroku_account;
```

### Check two factor authentication
Analyze the settings to understand which Heroku accounts have two-factor authentication enabled. This is useful for ensuring security best practices are being followed within your organization.

```sql+postgres
select
  email,
  two_factor_authentication
from
  heroku_account;
```

```sql+sqlite
select
  email,
  two_factor_authentication
from
  heroku_account;
```