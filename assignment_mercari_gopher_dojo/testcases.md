# Test cases

## Test case 1
### Input

```
echo "dog dog dog. cat. fish fish. go go go go." | ./wordcount
```

### Expected

```
go
dog
fish
```


## Test case 2
### Input

```
echo "bat bat go go go dog dog dog. bat. fish fish fish" | ./wordcount
```

### Expected

```
bat
dog
go
```
