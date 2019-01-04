# BeegoAccessLog
Baseed on beego/logs implete Nginx-like access log format
# How to use?
- If you use beego, and you want to log some params like request host,request ip, user-agent and request method,statusCode,etc..
this package is for you.
- In functions of controller , use the folowing code:
```
type TestController struct {
	beego.Controller
}
func (this *TestController) GetAll() {
beginTime := h.GetString("beginTime")

beegoInput := this.Ctx.Request
ss,err := models.GetAllThings(limit,offset,beginTime)
if err != nil {
this.Data["json"] = err
ss.Base.StatusCode = 404
ss.Base.Success = false
accessLog := BeegoAccessLog.FormatAccessLog(int64(ss.Base.StatusCode),beegoInput)
beego.Debug(accessLog)
}
accessLog := BeegoAccessLog.FormatAccessLog(int64(ss.Base.StatusCode),beegoInput)
beego.Debug(accessLog)

}
```
