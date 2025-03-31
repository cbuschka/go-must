# Generic helper for errors, that cannot happen

## Example

Instead of ...

```golang
var resource *Resource
if resource, err := acquire(); err != nil {
	return err
}
```

... you can then write ...

```golang
resource := must.Be(acquire())
```

... and it panics in case of error. 

## License

[MIT-0](./license.txt)