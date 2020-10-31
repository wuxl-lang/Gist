## Aggs


### Aggs with Filter, Group By and Distinct Count


```
{
    "query": {
        "bool": {
            "filter": {
                "bool": {
                    "must":[{
                        "range": {
                            "age": {
                                "gt": 18
                            }
                        }
                    }]
                }
            }
        }
    },
    "aggs": {
        "With Male" : {
            "filter": {
                "term": {
                    "gender": "Male"
                }
            },
            "aggs": {
                "GROUP BY": {
                    "terms": {
                        "field": "visitedCity"
                    },
                    "aggs": {
                        "DISTINCT COUNT": {
                            "cardinality" : {
                                "field": "id"
                            }
                        }
                    }                    
                }
            }
        }
    }
}
```