Map<String, String> requestData = new HashMap<String, String>();


/***银联全渠道系统，产品参数，除了encoding自行选择外其他不需修改***/

//版本号，全渠道默认值

requestData.put("version", DemoBase.version);

//字符集编码，可以使用UTF-8,GBK两种方式

requestData.put("encoding", DemoBase.encoding_UTF8);

//签名方法

requestData.put("signMethod", SDKConfig.getConfig().getSignMethod());

//交易类型 ，01：消费

requestData.put("txnType", "01");

//交易子类型， 01：自助消费

requestData.put("txnSubType", "01");

//业务类型，B2C网关支付，手机wap支付

requestData.put("bizType", "000201");

//渠道类型，这个字段区分B2C网关支付和手机wap支付；07：PC,平板 08：手机

requestData.put("channelType", "07");


/***商户接入参数***/

//商户号码，请改成自己申请的正式商户号或者open上注册得来的777测试商户号

requestData.put("merId", merId);

//接入类型，0：直连商户

requestData.put("accessType", "0");

//商户订单号，8-40位数字字母，不能含“-”或“_”，可以自行定制规则

requestData.put("orderId",DemoBase.getOrderId());

//订单发送时间，取系统时间，格式为YYYYMMDDhhmmss，必须取当前时间，否则会报txnTime无效

requestData.put("txnTime", DemoBase.getCurrentTime());

//交易币种（境内商户一般是156 人民币）

requestData.put("currencyCode", "156");

//交易金额，单位分，不要带小数点

requestData.put("txnAmt", txnAmt);



//前台通知地址 （需设置为外网能访问 http https均可），支付成功后的页面 点击“返回商户”按钮的时候将异步通知报文post到该地址

//如果想要实现过几秒中自动跳转回商户页面权限，需联系银联业务申请开通自动返回商户权限

//异步通知参数详见open.unionpay.com帮助中心 下载 产品接口规范 网关支付产品接口规范 消费交易 商户通知

requestData.put("frontUrl", DemoBase.frontUrl);


//后台通知地址（需设置为【外网】能访问 http https均可），支付成功后银联会自动将异步通知报文post到商户上送的该地址，失败的交易银联不会发送后台通知

//后台通知参数详见open.unionpay.com帮助中心 下载 产品接口规范 网关支付产品接口规范 消费交易 商户通知

//注意:1.需设置为外网能访问，否则收不到通知 2.http https均可 3.收单后台通知后需要10秒内返回http200或302状态码

// 4.如果银联通知服务器发送通知后10秒内未收到返回状态码或者应答码非http200，那么银联会间隔一段时间再次发送。总共发送5次，每次的间隔时间为0,1,2,4分钟。

// 5.后台通知地址如果上送了带有？的参数，例如：http://abc/web?a=b&c=d 在后台通知处理程序验证签名之前需要编写逻辑将这些字段去掉再验签，否则将会验签失败

requestData.put("backUrl", DemoBase.backUrl);


//////////////////////////////////////////////////

//

// 报文中特殊用法请查看 PCwap网关跳转支付特殊用法.txt

//

//////////////////////////////////////////////////


/**请求参数设置完毕，以下对请求参数进行签名并生成html表单，将表单写入浏览器跳转打开银联页面**/

//报文中certId,signature的值是在signData方法中获取并自动赋值的，只要证书配置正确即可。

Map<String, String> submitFromData = AcpService.sign(requestData,DemoBase.encoding_UTF8);


//获取请求银联的前台地址：对应属性文件acp_sdk.properties文件中的acpsdk.frontTransUrl

String requestFrontUrl = SDKConfig.getConfig().getFrontRequestUrl();

//生成自动跳转的Html表单

String html = AcpService.createAutoFormHtml(requestFrontUrl, submitFromData,DemoBase.encoding_UTF8);


LogUtil.writeLog("打印请求HTML，此为请求报文，为联调排查问题的依据："+html);

//将生成的html写到浏览器中完成自动跳转打开银联支付页面；这里调用signData之后，将html写到浏览器跳转到银联页面之前均不能对html中的表单项的名称和值进行修改，如果修改会导致验签不通过

resp.getWriter().write(html);