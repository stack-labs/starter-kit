# Validate

Server及Client端自动参数验证

```shell script
service.Init(
    micro.WrapHandler(validate.NewHandlerWrapper()),
    micro.WrapCall(validate.NewCallWrapper()),
)
```
