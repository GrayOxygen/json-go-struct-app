<b>[English Doc/英文文档](https://github.com/GrayOxygen/json-go-struct-app/blob/master/README.md "英文文档")</b>
# 介绍
这个工程只是做一个GUI图形界面app，拷贝json到最左边的文本框点击一下别处就可以得到grpc message的定义了 <br>
用到的开源项目：https://github.com/asticode/go-astilectron<br>
实际的json解析转换实现是在我的另一个工程里 https://github.com/GrayOxygen/json-to-grpc

</br> 如果你要做一些定制化的东西，记得按照如下步骤进行打包生成app（下面就不翻译了，你可以看懂）


# Step 1: install the app

Run the following commands:

    $ go get -u github.com/GrayOxygen/json-go-struct-app/...
    $ rm $GOPATH/src/github.com/GrayOxygen/json-go-struct-app/bind.go

# Step 2: install the bundler

Run the following command:

    $ go get -u github.com/asticode/go-astilectron-bundler/...
    
And don't forget to add `$GOPATH/bin` to your `$PATH`.
    
# Step 3: bundle the app for your current environment

Run the following commands:

    $ cd $GOPATH/src/github.com/GrayOxygen/json-go-struct-app
    $ astilectron-bundler -v
    
# Step 4: use the app

The result is in the `output/<your os>-<your arch>` folder and is waiting for you to use it!

# Step 5: bundle the app for more environments

To bundle the app for more environments, add an `environments` key to the bundler configuration (`bundler.json`):

```json
"environments": [
  {"arch": "amd64", "os": "linux"},
  {"arch": "amd64", "os": "windows"},
  {"arch": "386", "os": "windows"},
  {"arch": "amd64", "os": "darwin"},
]
```

and repeat **step 3**.
