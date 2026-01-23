查询所有单词
GET /word_desc/_doc/_search
{
  "query": {
    "match_all": {}
  }
}

按照单词查询
GET /word_desc/_search
{
  "query": {
    "match": {
      "word": "wish"
    }
  }
}

按照id查询
GET /word_desc/_search
{
  "query": {
    "match": {
      "_id":"0"
    }
  }
}

kibana: http://localhost:5601
elasticsearch: http://localhost:9200
rabbitmq: http://localhost:15672


