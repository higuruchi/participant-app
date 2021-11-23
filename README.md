# participant-app

Server side of an application to record participants in a circle

[Router side](https://github.com/yassi-github/participant-app-router)

[Client side](https://github.com/yassi-github/participant-app-client)

# DEMO

# Requirement

# Installation

# Usage

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
    Content-Type:application/json; charset=UTF-8
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
    # user id
    id=19T999

    # name of user
    name=kagawa-taro

    # MAC address
    macaddress=aa:aa:aa:aa:aa:aa
    ```