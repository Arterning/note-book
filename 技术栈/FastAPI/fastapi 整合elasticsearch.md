
## 安装elasticsearch

```
docker run -d \
	--name elasticsearch \
    -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" \
    -e "discovery.type=single-node" \
    -p 9200:9200 \
    -p 9300:9300 \
elasticsearch:7.12.1
```

安装完成后，在浏览器中输入：http://192.168.211.130:9200/ 即可看到elasticsearch的响应结果



## 创建索引

```python

from elasticsearch import Elasticsearch

# 连接到 Elasticsearch
es = Elasticsearch([{'host': 'localhost', 'port': 9200}])

# 创建一个索引（如果尚未创建）
es.indices.create(
    index='docs',
    body={
        'mappings': {
            'properties': {
                'title': {'type': 'text'},
                'content': {'type': 'text'}
            }
        }
    },
    ignore=400  # 如果索引已存在，忽略错误
)

# 索引示例文档
documents = [
    {"id": 1, "title": "First Document", "content": "This is the content of the first document."},
    {"id": 2, "title": "Second Document", "content": "This is the content of the second document."},
    {"id": 3, "title": "Third Document", "content": "ning is 天才,This is the content of the third document."}
]

for doc in documents:
    es.index(index='docs', id=doc['id'], body=doc)

```


可以更新索引映射，添加 'time' 字段

```python
# 更新索引映射，添加 'time' 字段
es.indices.put_mapping(
    index='docs',
    body={
        'properties': {
            'time': {
                'type': 'date',  # 或者其他类型，例如 'keyword'
                'index': False   # 不索引该字段
            }
        }
    }
)
```

在这个例子中，`'time'` 字段被设置为类型 `date`（你可以根据实际需要使用其他类型），并且通过 `'index': False` 指定该字段不被索引。


## mapping 的其他属性

在 Elasticsearch 中，`mapping` 用于定义索引中字段的类型和属性。除了 `type` 和 `index`，还有许多其他属性可以用来控制字段的行为和存储方式。以下是一些常见的字段属性及其说明：

### 常见字段属性

1. **`type`**: 字段的数据类型（如 `text`, `keyword`, `integer`, `date` 等）。决定了字段如何存储和索引。

2. **`index`**: 控制字段是否被索引（默认值为 `true`）。如果设置为 `false`，字段不会被索引，因此不能用于搜索。

3. **`store`**: 控制是否将字段存储在索引中（默认值为 `false`）。如果设置为 `true`，字段将被存储，尽管可能未被索引，仍可以通过 `get` API 访问。

4. **`analyzer`**: 指定用于分析该字段的分析器。适用于 `text` 类型字段，用于分词和标准化。

5. **`search_analyzer`**: 用于搜索时的分析器。如果设置了 `analyzer`，可以指定不同的分析器在搜索时使用。

6. **`norms`**: 控制是否存储字段的标准化信息（如文本的长度和重要性）。对 `text` 类型字段有效（默认值为 `true`）。

7. **`boost`**: 设置字段的权重。对查询结果的评分进行加权。通常用于 `query` 阶段。

8. **`doc_values`**: 控制字段是否启用 doc values。对于 `keyword`, `numeric`, `date` 类型字段，`doc_values` 用于高效的排序和聚合操作（默认值为 `true`）。

9. **`format`**: 控制如何格式化字段的值。适用于日期字段，指定日期的格式。

10. **`null_value`**: 指定在索引中用作 `null` 值的替代值。例如，可以将 `null` 值替换为某个默认值进行索引。

11. **`properties`**: 用于定义嵌套字段的映射。对于对象字段，`properties` 属性定义了对象中各个子字段的映射。

12. **`copy_to`**: 将字段的值复制到一个或多个其他字段，以便进行复合查询。

13. **`dynamic`**: 控制对象字段如何处理动态添加的字段。可以设置为 `true`（允许动态字段）、`false`（禁止动态字段）、`strict`（拒绝动态字段）。

14. **`fields`**: 为字段定义多个子字段，以便对同一字段使用不同的索引和分析策略。例如，可以为 `text` 字段添加一个 `keyword` 类型的子字段，用于精确匹配。

### 示例

以下是一个包含多个属性的字段映射示例：

```json
PUT /my_index
{
  "mappings": {
    "properties": {
      "title": {
        "type": "text",
        "analyzer": "standard",
        "search_analyzer": "whitespace",
        "boost": 2.0
      },
      "content": {
        "type": "text",
        "analyzer": "english",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "publish_date": {
        "type": "date",
        "format": "yyyy-MM-dd",
        "null_value": "1970-01-01"
      },
      "views": {
        "type": "integer",
        "doc_values": false
      },
      "tags": {
        "type": "keyword",
        "ignore_above": 1024
      }
    }
  }
}
```

### 总结

- **`type`**: 数据类型（如 `text`, `keyword` 等）。
- **`index`**: 是否索引字段（默认 `true`）。
- **`store`**: 是否存储字段（默认 `false`）。
- **`analyzer`**: 字段的分析器。
- **`search_analyzer`**: 搜索时的分析器。
- **`norms`**: 是否启用标准化（默认 `true`）。
- **`boost`**: 权重提升。
- **`doc_values`**: 是否启用 doc values（默认 `true`）。
- **`format`**: 日期格式。
- **`null_value`**: `null` 值替代。
- **`properties`**: 嵌套字段的映射。
- **`copy_to`**: 字段值复制。
- **`dynamic`**: 对动态字段的处理。
- **`fields`**: 多字段定义。

通过这些属性，你可以精确控制字段的存储、索引和检索行为。希望这些信息对你有帮助！如果你有更多问题或需要进一步的帮助，请随时告诉我。


## 脚本搜索

```python
from elasticsearch import Elasticsearch

es = Elasticsearch([{'host': 'localhost', 'port': 9200}])

# 执行 Elasticsearch 搜索
response = es.search(
    index='docs',
    body={
        'query': {
            'match': {
                'content': "天才"
            }
        }
    }
)

print(response)

# 处理 Elasticsearch 搜索结果
hits = response['hits']['hits']
results = [
    {
        'id': hit['_id'],
        'title': hit['_source']['title'],
        'content': hit['_source']['content']
    }
    for hit in hits
]

print(results)
```


## fastapi整合

```python
from fastapi import FastAPI, Query
from pydantic import BaseModel
from typing import List, Optional
from elasticsearch import Elasticsearch

app = FastAPI()
es = Elasticsearch([{'host': 'localhost', 'port': 9200}])

class SearchResponse(BaseModel):
    id: int
    title: str
    content: str

@app.get("/search", response_model=List[SearchResponse])
async def search(
    query: str = Query(..., description="Search query")
):
    # 执行 Elasticsearch 搜索
    response = es.search(
        index='docs',
        body={
            'query': {
                'match': {
                    'content': query
                }
            }
        }
    )
    
    # 处理 Elasticsearch 搜索结果
    hits = response['hits']['hits']
    results = [
        {
            'id': hit['_id'],
            'title': hit['_source']['title'],
            'content': hit['_source']['content']
        }
        for hit in hits
    ]
    
    return results

```
