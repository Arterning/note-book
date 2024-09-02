Discovered something neat with the new version of MySQL and thought it warranted a mention. Storing tree structures in a relational database is a common use case across many different areas of tech. The problem comes when you need to construct a query based on a subset of that tree.

But MySQL 8 has some nice new features that makes doing this a breeze.

For example, let’s assume you have a set of tables that look like this:


```sql
CREATE TABLE `files` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `parent_id` bigint(20) NOT NULL,
  `kind` enum('file','folder') NOT NULL,
  PRIMARY KEY (`id`),
  KEY `parent_id` (`parent_id`)
);

CREATE TABLE `tags` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `tag_file` (
  `tag_id` bigint(20) NOT NULL,
  `file_id` bigint(20) NOT NULL,
  PRIMARY KEY (`tag_id`,`file_id`)
);
```

This is a pretty standard setup for storing a tree setup in a relational table, with the `parent_id` key referencing `id` of the parent. And this all works well … right up until you need to query a parent and all of it’s children. So let’s say you want to find all children that have a tag of ‘Important’.

MySQL 8 includes support for [recursive common table expressions](https://dev.mysql.com/doc/refman/8.0/en/with.html). Using this, this becomes a pretty easy query. You can create a CTE query and recursively call it! You could so something like this:

```sql
with recursive cte (id, parent_id) as (
    select id, parent_id from files where parent_id = ? and kind = 'folder'
    union all
    select p.id, p.parent_id from files p inner join cte on p.parent_id = cte.id where kind = 'folder'
)
select * from cte;

select * from files inner join tag_file on (tag_file.file_id = file.id)
inner join tags on (tags.id = tag_file.tag_id) where kind = 'file'
and tag = 'Important' and parent_id in (select id from cte);

```

(That is a prepared statement, replace the `?` with the ID of the parent.)

What surprised me was how _fast_ that query is with the right keys. On a table that has nearly 600,000 items in it, that query completes in about 0.3 seconds. Slow, but considering the number of rows in the table quite fast.

Thanks to [this post](https://stackoverflow.com/a/33737203/9348119) on Stack Overflow for the heads up.















