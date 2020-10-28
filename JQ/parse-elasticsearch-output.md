## Parse ElasticSearch Output

*Refer to [jq]*

### Output Example
```json
{
  "hits": {
    "hits": [
      {
        "_source": {
          "name": "name",
          "age": 10,
          "pets": ["dog", "cat"]
        }
      }
    ]
  }
}
```

#### Read Name and Age

```
cat output.json | jq '.hits.hits[] | (._source.name + "," + (._source.age|tostring))'
```

```
"name,10"
```

#### Read Name, Age and Pet

```
cat output.json | jq '.hits.hits[] | (._source.name + "," + (._source.age|tostring) + "," + (._source.pets[]))'
```

```
"name,10,dog"
"name,10,cat"
```

[jq]: https://stedolan.github.io/jq/