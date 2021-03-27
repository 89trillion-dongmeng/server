# api 说明

### /gift/create
1. 接口方式	http post
2. 路径	/gift/create
3. 参数	json 
    ```json
    {"userId":"1234","count": 10,"gifts": [{"coin": 100}]}
    ```
4. 返回  json 
    ```json
    {"code":"apgd48jk","message": "ok"}
    ```

type GiftGetRes struct {
	Changes map[string]string `changes`
	Message string         `message`
}
### /gift/get
1. 接口方式	http post
2. 路径	/gift/create
3. 参数	:userId,code 
4. 返回  json 
    ```json
    {"changes":{"coin": 100},"message":"ok"}
    ```

### 测试结果

![]("./images/create.png")

![]("./images/get.png")
 

