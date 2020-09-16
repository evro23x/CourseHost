# Course - backend #
host - http://157.245.22.136 <br>
port - 80 <br>

## Methods:
1) user/new (POST) - Создание пользователя <br> 
    endpoint - host/user/new <br>
    request:  <br>
    ```json
   {
       "email": "test@test.test",
       "password": "1234567890"
   }
    ```
   response:
    ```json
   {
       "message": "User has been created",
       "status": true,
       "user": {
           "ID": 7,
           "CreatedAt": "2019-12-08T16:04:24.172681087Z",
           "UpdatedAt": "2019-12-08T16:04:24.172681087Z",
           "DeletedAt": null,
           "email": "test@test.test",
           "password": "",
           "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjd9.xBtr8h_ei9zcV21JlCLSKpqGxW2l2R0XZRV0Lr7AjXQ"
       }
   }
    ```
2) user/login (POST) - Авторизация(получение токена)<br> 
      endpoint - host/user/login <br>
      request:  <br>
      ```json
      {
          "email": "test@test.test",
          "password": "1234567890"
      }
      ```
     response:
      ```json
    {
        "message": "Logged In",
        "status": true,
        "user": {
            "ID": 7,
            "CreatedAt": "2019-12-08T16:04:24.172681Z",
            "UpdatedAt": "2019-12-08T16:04:24.172681Z",
            "DeletedAt": null,
            "email": "test@test.test",
            "password": "",
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjd9.xBtr8h_ei9zcV21JlCLSKpqGxW2l2R0XZRV0Lr7AjXQ"
        }
    }
      ```
3) page/new (POST) - Создание страницы <br> 
      endpoint - host/page/new <br> 
      request:  <br>
      ```json
    {
        "title": "main",
        "head": "script",
        "body": "html code"
    }
      ```
     response:
      ```json
    {
        "message": "success",
        "page": {
            "ID": 3,
            "CreatedAt": "2019-12-08T16:12:42.088694604Z",
            "UpdatedAt": "2019-12-08T16:12:42.088694604Z",
            "DeletedAt": null,
            "title": "main",
            "head": "script",
            "body": "html code"
        },
        "status": true
    }
      ```
4) page/{id} (GET) - Получение страницы <br> 
      endpoint - host/page/1<br> 
     response:
      ```json
    {
        "data": [
            {
                "ID": 1,
                "CreatedAt": "2019-12-08T16:12:34.500728Z",
                "UpdatedAt": "2019-12-08T16:12:34.500728Z",
                "DeletedAt": null,
                "title": "main",
                "head": "script",
                "body": "html code"
            }
        ],
        "message": "success",
        "status": true
    }
      ```
