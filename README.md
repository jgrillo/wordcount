# Wordcount

It counts your words.

**POST** */words*

```
{
    "word" : ["word", "word", "word"]
}
```

Response:

```
{
    "counts" {
        "word": 3
    }
}
```