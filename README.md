# participant-app

Server side of an application to record participants in a circle

[Router side](https://github.com/yassi-github/participant-app-router)

[Client side](https://github.com/yassi-github/participant-app-client)

# Requirement

- docker
- docker-compose

# Installation

```
$ git clone https://github.com/higuruchi/participant-app.git
```

# Usage

- Write config file in yaml and put in the directory under the participant-app 

    ```yaml
    server:
        # service port number
        port: 1323
    db:
        # dagbase user name
        user: user
        # database user password
        password: password
        # database ip address
        ip: 192.168.0.104
        #database port number
        port: 3306
        # database name
        name: participant-app
    ```

- Run the following command

    ```bash
    $ cd ./participant-app/deployments
    $ docker-compose up --build
    ```

## Get JSON of participants

- Method

    **GET**

- HTTP request path

    ```
    /participants/:year/:month/:date
    ```

- Return value

    ```JSON
    {
        "B1": [
            {
                "Id": "21T999",
                "Name": "kagawa-taro"
            },
            {
                "Id": "21T998",
                "Name": "kagawa-jiro"
            }
        ],
        "B2": [
            {
                "id": "20T999",
                "Name": "kagawa-sabro"
            }
        ],
        "B3": null,
        "B4": null
    }
    ```

## Save participants

- Method
    
    **POST**

- HTTP request path

    ```
    /participants
    ```

- HTTP request header

    ```http
    Content-Type:application/json; 
    ```

- HTTP request body

    ```JSON
    {
        "year": 2021,
        "month": 11,
        "day": 16,
        "hour": 14,
        "minute": 11,
        "second": 54,
        "macaddresses": [
            "c4:3c:ea:85:f0:08",
            "2c:d0:5a:27:82:3e"
        ]
    }
    ```

## Register user

- Method

    **POST**

- HTTP request path

    ```
    /user
    ```

- HTTP request header

    ```http
    Content-Type:application/x-www-form-urlencoded
    ```

- HTTP request body

    ```
    id=19T999&name=kagawataro&macaddress=aa%3Aaa%3Aaa%3Aaa%3Aaa%3Aaa
    ```

### Update user macaddr

- Method

    **PUT**

- HTTP request path

    ```
    /macaddr/:id
    ```

- HTTP request header

    ```http
    Content-Type:application/x-www-form-urlencoded
    ```

- HTTP request body

    ```
    # macaddress
    macaddress=aa%3Aaa%3Aaa%3Aaa%3Aaa%3Aaa
    ```
