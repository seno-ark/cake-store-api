
## API Endpoints

## Get Cake list
__GET /cakes__

Query Params:
- page: default 1
- count: default 10, max 10

200
400
500

## Get Cake by ID
__GET /cakes/:cake_id__

200
400
404
500

## Add new Cake
__POST /cakes__

Json Body:

{
  "title": "Lemon cheesecake",
  "description": "A cheesecake made of lemon",
  "rating": 7,
  "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
}

201
400
500

## Update Cake
__PATCH /cakes/:cake_id__

Json Body:

{
  "title": "Lemon cheesecake",
  "description": "A cheesecake made of lemon",
  "rating": 7,
  "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
}

200
400
404
500

## Delete Cake
__DELETE /cakes/:cake_id__

200
400
500
