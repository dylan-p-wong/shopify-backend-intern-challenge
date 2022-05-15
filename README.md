# shopify-backend-intern-challenge

This is my submission for the Shopify Backend Intern Challenge. I implemented the "When deleting, allow deletion comments and undeletion" option along with the CRUD functionality. Below is a small documentation for the requests. You should be able to test these with the Repl.it link. Additionally I wrote a few tests and they can be ran by running go test in the test directory.  

Repl.it Link: https://YummyDualOpen64.dylanwong8.repl.co

Example: GET https://YummyDualOpen64.dylanwong8.repl.co/item

**Create Item**
----
  Creates a single item.

* **URL**

  /item

* **Method:**

  `POST`

* **JSON Body**

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

**Get Items**
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

* **JSON Body**

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

* **JSON Body**

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
