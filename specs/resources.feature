# language: zh-CN  
功能: 用户登录  
  为了能够浏览网站只对在线会员可见的那些内容  
  作为一名访客  
  我希望能够登录  
  
  场景: 用户登录功能  
    假如 没有<somebody@somedomain.com>这个用户  
    当 我以<somebody@somedomain.com/password>这个身份登录  
    那么 我应该看到<用户名或密码错误>的提示信息  
    而且 我应该尚未登录  
    
功能：支持url路由    
    场景：增加一个路由
        如果 没有<get users>这个路由
        当 我以<get "/users" UserController.Index>增加一个路由
        那么 我应该得到如下路由
            |GET    "/users"    UserController.Index|

功能:  支持REST样式的资源
    为了 简化REST资源的路径映射
    作为 API服务器
    我希望能够支持REST资源

    场景：增加一个资源
        如果 没有<users>这个资源
        当 我以<users>这个资源名
        和 <UserController>这个控制器名字增加一个资源
        那么 我应该得到如下的url路由
            |GET        "/users"    UserController.Index|                        
            |POST       "/users"    UserController.Create|
            |GET        "/users/:id"  UserController.Show|
            |PUT        "/users/:id"    UserController.Update|
            |DELETE    "/user/:id"  UserController.Delete|

    场景：增加一个只读资源
        如果 没有<user>这个资源
        当 我以<users>这个资源名
        和 <UserController>这个控制器名字
        和 指定只读属性（read_only=True）增加一个资源
        那么 我应该得到如下的url路由
            |GET        "/users"    UserController.Index|                        
            |GET        "/users/:id"  UserController.Show|
        而且 不会有多余的其他路由
        


