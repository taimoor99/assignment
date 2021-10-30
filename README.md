# Assignment
Project have three layer the controller layer, model layer(for business logic) and the repository layer. By following clean architecture golang. 
Mongodb is used for message storage.

To run project locally make sure have docker installed in your system. Or you need to have and mongodb installed on system add mongodb url in environment variable.
Then run go run cmd/assigment/main.go.

The default port of server is 8080.

### Make commands to run project
* make build_containers (build containers)
* make run_containers (run containers)

These two commands will start go server and mongodb locally. Now you can try out the bellow api operations.

This assignment consist of the 4 endpoints post message, get detail of specific message,
get list of all messages, and delete message.

* Post message  
    * Method: Post
    * Url: /message
    * Body: { "message": "test" }
    * Response: {
          "body": {
              "_id": "617d52f9ac41968e55221df6",
              "message": "test",
              "created_at": "2021-10-30T14:13:13.0121858Z",
              "updated_at": "2021-10-30T14:13:13.0121859Z"
          }
      }
* Get message
    * Method: Get
    * Url: /message/{id}
    * Param: 617d52f9ac41968e55221df6
    * Response: {
      "message": "",
          "body": {
              "_id": "617d52f9ac41968e55221df6",
              "message": "test",
              "created_at": "2021-10-30T14:13:13.0121858Z",
              "updated_at": "2021-10-30T14:13:13.0121859Z"
          }
      }
* Get message list
  * Method: Get
  * Url: /message/{limit}/{offset}
  * Param: 10, 0
  * Response: {
    "message": "",
    "body": {
            "_id": "617d52f9ac41968e55221df6",
            "message": "test",
            "created_at": "2021-10-30T14:13:13.0121858Z",
            "updated_at": "2021-10-30T14:13:13.0121859Z"
        }
    }
* Delete message
    * Method: Delete
    * Url: /message/{id}
    * Param: 617d52f9ac41968e55221df6
    * Response: {
      "message": "message deleted",
      "body": 1
      }
      
To run API's on postman you can find the file "Test.postman_collection.json" and import into postman.
