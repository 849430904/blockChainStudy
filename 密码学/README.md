#### 大纲
* ![](img/001.png) ![](img/002.png) ![](img/003.png)![](img/004.png)![](img/005.png)![](img/006.png)![](img/007.png)![](img/008.png)![](img/009.png)![](img/010.png)![](img/011.png)![](img/012.png)
*

* 对称加密：
* ![](img/013.png)![](img/014.png)
* DES
* ![](img/015.png)现在DES也不是那么安全了，因为计算机越来越快了；![](img/017.png)
* DES细节：![](img/018.png)![](img/022.png)
* ![](img/019.png)![](img/020.png)

![](img/021.png) ![](img/022.png)![](img/023.png)
* [DES加解密代码](code/Ctypto)
* ![](img/024.png)![](img/025.png)![](img/026.png)![](img/027.png)![](img/028.png)

### AES
*  ![](img/029.png),对称加密中，现阶段安全级别最高的加密算法 ![](img/030.png)![](img/031.png)
*  ![](img/032.png),新项目如果要用对称加密，建议使用AES加密算法

### 非对称加密
* 对称加密存在密钥分发困难，分发密钥的时候，可能被对方监听到。因而有了非对称加密
* ![](img/033.png)![](img/034.png)![](img/035.png)![](img/036.png)![](img/037.png)![](img/038.png) ![](img/039.png),只有Bob能解密数据，Alice可以用公钥向Bob发送数据.非对称加密只需要向对方分发公钥就OK了
* ![](img/040.png)![](img/041.png)![](img/042.png)![](img/043.png)![](img/044.png)![](img/045.png)![](img/046.png)![](img/047.png)![](img/048.png)![](img/049.png)

* [非对称加解密demo](code/Ctypto/MyRsa.go)
* ![](img/050.png)![](img/051.png),密钥越长，程序运行的时间越长，效率越低
* ![](img/052.png)![](img/053.png)![](img/054.png)![](img/055.png)![](img/056.png)![](img/057.png)![](img/058.png)![](img/059.png)![](img/060.png)![](img/061.png)![](img/062.png)![](img/063.png)![](img/064.png)