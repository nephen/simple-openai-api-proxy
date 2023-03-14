> [中文说明](./README.zh.md)

### description
You don’t need to set up an environment, just have an overseas vps, preferably a vps in an area supported by openai, download the [executive file] (./bin/api_proxy) in the bin and run it directly, the simplest api proxy method, the most important The most important thing is to support SSE, so that the client can respond more quickly when requesting, and also provides the source code of golang, which can be improved by itself if it needs to be customized.
```sh
./api_proxy -daemon -port 9000 # It is best to open the daemon process mode
```

### How to use the client
Python use cases：
```python
import os
import openai

openai.api_key = YOUR-API-KEY
openai.api_base = "http://host:port/v1" # Be sure to add v1

for resp in openai.ChatCompletion.create(
                                    model="gpt-3.5-turbo",
                                    messages=[
                                      {"role": "user", "content": "Bubble Sort"}
                                    ],
                                    stream = True): # Streaming output, support SSE
    if 'content' in resp.choices[0].delta:
        print(resp.choices[0].delta.content, end="", flush=True) # flush prints in time
```
JS use case, Take https://www.npmjs.com/package/chatgpt as an example
```js
chatApi= new gpt.ChatGPTAPI({
    apiKey: 'sk.....:<proxy_key write here>',
    apiBaseUrl: "http://host:port", // delivery proxy address
});
```