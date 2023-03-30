### 说明
可以不用搭环境，自己有境外的vps就行，最好是openai支持的地区的vps，下载bin里面的[执行文件](./bin/api_proxy)直接就能跑，最简单的api proxy方式，最重要的是支持SSE，让客户端请求时响应得更加迅速，也提供了golang的源码，需要定制的可以自行完善。
```sh
./api_proxy -daemon -port 9000 # 最好开启daemon守护进程模式
```

### 客户端使用方法
python使用案例：
```python
import os
import openai

openai.api_key = YOUR-API-KEY
openai.api_base = "http://host:port/v1" # 一定要加v1

for resp in openai.ChatCompletion.create(
                                    model="gpt-3.5-turbo",
                                    messages=[
                                      {"role": "user", "content": "冒泡排序"}
                                    ],
                                    stream = True): # 流式输出，支持SSE
    if 'content' in resp.choices[0].delta:
        print(resp.choices[0].delta.content, end="", flush=True) # flush及时打印
```

js使用案例，以 https://www.npmjs.com/package/chatgpt 为例：
```js
chatApi= new gpt.ChatGPTAPI({
    apiKey: 'sk.....:<proxy_key写这里>',
    apiBaseUrl: "http://host:port", // 传递代理地址
});
```

（推荐：）服务器使用[ChatGPT-Next-Web](https://github.com/Yidadaa/ChatGPT-Next-Web) 的例子，设置key后，可以设置code密码来访问，api则使用当前的代理，非常好用，参考网页https://gpt.nephen.cn/。
```sh
docker pull nephen2023/chatgpt-next-web:v1.7.1
docker run -d -p 3000:3000 -e OPENAI_API_KEY="" -e CODE="" -e BASE_URL="ip:port" -e PROTOCOL="http" nephen2023/chatgpt-next-web:v1.7.1
```

### 支持
![](https://nephen-blog.oss-cn-beijing.aliyuncs.com/post/20230315130826.png)
