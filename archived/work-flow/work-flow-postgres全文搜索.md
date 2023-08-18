# 如何使用 postgres 的全文搜索功能

PostgreSQL 支持全文搜索功能，这可以让你搜索文本字段中的关键字。为了使用 PostgreSQL 的全文搜索功能，你需要：

创建一个带有文本字段的表：创建一个包含文本字段的表，这些文本字段包含要搜索的信息。

使用全文搜索解析器：PostgreSQL 支持多种全文搜索解析器，你可以根据需要选择一个解析器。比如，你可以使用 tsquery 解析器进行简单的全文搜索。

建立索引：创建一个全文索引，以提高搜索的速度。

使用 ILIKE 或者 tsvector 和 tsquery 运算符进行搜索：你可以使用 ILIKE 运算符进行简单的全文搜索，也可以使用 tsvector 和 tsquery 运算符进行更高级的全文搜索。

以下是一个例子：

```sql
-- 创建一个文档表
CREATE TABLE documents (
  id SERIAL PRIMARY KEY,
  title TEXT,
  content TEXT
);

-- 使用tsvector类型的字段存储文本信息
ALTER TABLE documents ADD COLUMN search_vector tsvector;

-- 为tsvector字段创建索引
CREATE INDEX documents_search_idx ON documents USING gin(search_vector);

-- 更新tsvector字段
UPDATE documents SET search_vector = to_tsvector('english', title || ' ' || content);

-- 查询文档表
SELECT id, title, content
FROM documents
WHERE search_vector @@ to_tsquery('english', '搜索关键字')


```

"@@" 是 PostgreSQL 中的全文搜索运算符，代表搜索匹配，它的意思是检查 search_vector 字段是否与 to_tsquery 函数返回的 tsquery 值匹配。

在这个查询中，"search_vector @@ to_tsquery('english', '搜索关键字')" 表示搜索文档表中的文本，看看是否包含英文关键字 "搜索关键字"。如果包含，则返回 true；否则返回 false。

第一个参数 'english' 是语言代码，用于指定使用哪个语言分词器进行文本分词。PostgreSQL 中提供了多种语言分词器，包括英语、法语、德语等。指定语言代码的目的是为了更好地处理该语言的文本，比如处理英文文本时分词的结果与中文文本不同，使用英文语言分词器可以得到更符合英文语法的分词结果。

如果你的数据表中存储的是英文文本，那么第一个参数应该是 'english'。如果是其他语言，则应使用相应的语言代码，如 'french'、'german' 等。

结合 WHERE 子句，你可以查询出包含特定关键字的文档记录。

MySQL 中的 LIKE 语句可以用于文本匹配，但它不是全文搜索功能。LIKE 可以使用通配符和表达式对文本进行匹配，但它不能对文本进行语言特定的分词和语法处理，也不能创建索引和搜索权重。

相比之下，PostgreSQL 的全文搜索功能更加强大，提供了更多的分词、语法处理和索引优化的功能，能够更好地处理文本搜索。

具体来说，PostgreSQL 中的全文搜索可以使用全文索引进行加速，这样搜索的效率比较高，并且可以支持语言特定的分词和语法处理，可以提供更精确的搜索结果。

所以，总的来说，PostgreSQL 的全文搜索相比 MySQL 中的文本匹配，更加强大和高效。

# to_tsvector

to_tsvector('english', title || ' ' || content) 这个语句的意思是：创建一个文本搜索向量，该向量表示由 title 和 content 字段拼接起来的文本。

to_tsvector 函数：该函数是用于创建文本搜索向量的。
'english' 参数：表示对文本使用英语分词器，并对文本进行语言特定的分词和语法处理。
title 和 content 字段：分别代表文本的两个字段。
title || ' ' || content：拼接 title 和 content 字段，中间插入一个空格，用于将这两个字段看作一个整体的文本。
通过这条语句，你可以创建一个文本搜索向量，这个向量表示了整个文本，它会在后面的搜索操作中被用到。

# 这背后是不是基于倒序索引？

是的，PostgreSQL 中的全文搜索功能基于倒排索引。倒排索引是一种特殊的索引，它存储了所有出现在文档中的词，以及这些词所在文档的列表。这样，当搜索包含一个词的文档时，搜索引擎可以直接查询该词的倒排索引，获取包含该词的文档列表。

在 PostgreSQL 中，tsvector 类型的字段可以被索引，并且它们存储了一个文本搜索向量，以支持快速全文搜索。当创建一个索引时，该索引会自动构建一个倒排索引，以便于更快地进行搜索。
