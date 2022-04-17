# Product Images

## Uploading 

Note: need to use `--data-binary` to ensure file is not converted to text

```
curl -vv localhost:9090/1/go.mod -X PUT --data-binary @test.png
```

To post a image to your server
```
curl localhost:9090/images/1/test.png -d @test.png
curl localhost:9090/images/1/girls.jpg -d @girls.jpg
```

## Getting 
To get a image to your server
```
curl localhost:9090/images/1/test.png
```