# shopify-backend-intern-challenge

**Create Item**
----
  Creates a single item.

* **URL**

  /item

* **Method:**

  `POST`

* **Data Params**

  ```
  {
    "title": "afasdfasdfas2",
    "description": "",
    "count": 122
  }
  ```

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```
    {
      "id": "35362ee6-d468-11ec-ac0d-0242ac12002c",
      "title": "afasdfasdfas2",
      "description": "",
      "count": 122,
      "created_at": "2022-05-15T16:00:57.608310752Z",
      "updated_at": "2022-05-15T16:00:57.608310752Z",
      "deleted": false,
      "deletion_comments": ""
    }
    ```
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** 
    ```
      { "errors" : ["Count must be at least 0"] }
    ```

**Create Item**
----
  Get all not deleted items.

* **URL**

  /item

* **Method:**

  `GET`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```
    [
      {
        "id": "35362ee6-d468-11ec-ac0d-0242ac12002f",
        "title": "test2",
        "description": "",
        "count": 122,
        "created_at": "2022-05-15T16:00:57.608310752Z",
        "updated_at": "2022-05-15T16:00:57.608310752Z",
        "deleted": false,
        "deletion_comments": ""
      },
      {
        "id": "35362ee6-d468-11ec-ac0d-0242ac12002c",
        "title": "test",
        "description": "",
        "count": 1,
        "created_at": "2022-04-15T16:00:57.608310752Z",
        "updated_at": "2022-04-15T16:00:57.608310752Z",
        "deleted": false,
        "deletion_comments": ""
      }
    ]
    ```
 
* **Error Response:**

  * **Code:** 500 <br />


**Edit Item**
----
  Edits a single item.

* **URL**

  /item/:id

* **Method:**

  `PATCH`

* **Data Params**

  ```
  {
    "title": "afasdfasdfas2",
    "description": "",
    "count": 122
  }
  ```

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```
    {
      "id": "35362ee6-d468-11ec-ac0d-0242ac12002c",
      "title": "afasdfasdfas2",
      "description": "",
      "count": 122,
      "created_at": "2022-05-15T16:00:57.608310752Z",
      "updated_at": "2022-05-15T16:00:57.608310752Z",
      "deleted": false,
      "deletion_comments": ""
    }
    ```
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** 
    ```
      { "errors" : ["Count must be at least 0"] }
    ```

**Delete Item**
----
  Deletes a single item.

* **URL**

  /item/:id

* **Method:**

  `DELETE`

* **Data Params**

  ```
  {
    "deletion_comments": "Bad item."
  }
  ```

* **Success Response:**

  * **Code:** 204 <br />
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** 
    ```
      { "errors" : ["Not found."] }
    ```

**Un Delete Item**
----
  Un Deletes a single item.

* **URL**

  /item/:id/undelete

* **Method:**

  `PATCH`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** 
    ```
    {
      "id": "35362ee6-d468-11ec-ac0d-0242ac12002c",
      "title": "afasdfasdfas2",
      "description": "",
      "count": 122,
      "created_at": "2022-05-15T16:00:57.608310752Z",
      "updated_at": "2022-05-15T16:00:57.608310752Z",
      "deleted": false,
      "deletion_comments": ""
    }
    ```
 
* **Error Response:**

  * **Code:** 400 <br />
    **Content:** 
    ```
      { "errors" : ["Not found."] }
    ```
