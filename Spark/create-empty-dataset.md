## Create Empty Dataset with Schema

```scala
import org.apache.spark.sql.types._
import org.apache.spark.sql.Row

val schema = StructType(
    StructField("name", StringType, true)::
    StructField("age", IntegerType, true)::
    StructField("lastUpdateTime", LongType, true)::Nil
)

val emptyDs = spark.createDataFrame(spark.sparkContext.emptyRDD[Row], schema)
```

More examples refers to [sparkbyexample]

[sparkbyexample]: https://sparkbyexamples.com/spark/spark-how-to-create-an-empty-dataset/