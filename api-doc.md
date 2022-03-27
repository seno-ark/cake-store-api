# API Documentation

# Endpoints

## Get Cake list
__GET /cakes__

Query Params:
- page: default 1
- count: default 10, max 10

Success Response:

__200__ OK
```
{
    "message": "",
    "data": {
        "cakes": [
            {
                "id": 1,
                "title": "Lemon cheesecake v2",
                "description": "A cheesecake made of lemon",
                "rating": 5,
                "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
                "created_at": "2022-03-27T01:58:11Z",
                "updated_at": "2022-03-27T02:00:02Z"
            }
        ],
        "count": 10,
        "page": 1,
        "total_data": 1
    }
}
```

## Get Cake by ID
__GET /cakes/:cake_id__

Success Response:

__200__ OK
```
{
    "message": "",
    "data": {
        "id": 1,
        "title": "Lemon cheesecake v2",
        "description": "A cheesecake made of lemon",
        "rating": 5,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
        "created_at": "2022-03-27T01:58:11Z",
        "updated_at": "2022-03-27T02:00:02Z"
    }
}
```

## Add new Cake
__POST /cakes__

Json Body:
```
{
  "title": "Lemon cheesecake",
  "description": "A cheesecake made of lemon",
  "rating": 4.5,
  "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
}
```

Success Response:

__201__ Created
```
{
    "message": "",
    "data": {
        "id": 1,
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 4.5,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
        "created_at": "2022-03-27T01:58:11Z",
        "updated_at": "2022-03-27T01:58:11Z"
    }
}
```

## Update Cake
__PATCH /cakes/:cake_id__

Json Body:
```
{
    "title": "Lemon cheesecake v2",
    "description": "A cheesecake made of lemon",
    "rating": 5,
    "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
}
```

Success Response:

__200__ OK
```
{
    "message": "",
    "data": {
        "id": 1,
        "title": "Lemon cheesecake v2",
        "description": "A cheesecake made of lemon",
        "rating": 5,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
        "created_at": "2022-03-27T01:58:11Z",
        "updated_at": "2022-03-27T02:00:02Z"
    }
}
```

## Delete Cake
__DELETE /cakes/:cake_id__

Success Response:

__200__ OK
```
{
    "message": "Success",
    "data": null
}
```

# Error Response

__400__ Bad Request
```
{
    "message": "Invalid Data: cake_id",
    "data": null
}
```

__404__ Not Found
```
{
    "message": "Cake not found",
    "data": null
}
```

__500__ Internal Server Error
```
{
    "message": "Database Error",
    "data": null
}
```
