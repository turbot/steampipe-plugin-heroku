---
title: "Steampipe Table: heroku_region - Query Heroku Regions using SQL"
description: "Allows users to query Heroku Regions, specifically information regarding the unique identifier, name, and description of each region."
---

# Table: heroku_region - Query Heroku Regions using SQL

Heroku Regions are geographical areas in which Heroku infrastructure is located. Each region is completely isolated from other regions and has its own set of resources. Regions are designed to allow developers to deploy and run applications closer to their user base for improved latency.

## Table Usage Guide

The `heroku_region` table provides insights into the geographical areas where Heroku infrastructure is located. As a developer or system administrator, you can explore region-specific details through this table, including unique identifiers, names, and descriptions. Utilize it to understand the infrastructure distribution and to plan for application deployment for improved latency.

**Important Notes**
- Get queries require a region `id`.
- Pagination is not currently supported for this resource type in the SDK.

## Examples

### List all regions
Explore the different regions where Heroku services are available, allowing you to understand the geographical distribution and plan your deployments accordingly.

```sql+postgres
select
  *
from
  heroku_region
order by
  name;
```

```sql+sqlite
select
  *
from
  heroku_region
order by
  name;
```

### List all US regions
Explore which Heroku regions are located within the United States. This is useful for understanding the geographical distribution of your Heroku resources.

```sql+postgres
select
  *
from
  heroku_region
where
  country = 'United States'
order by
  name;
```

```sql+sqlite
select
  *
from
  heroku_region
where
  country = 'United States'
order by
  name;
```