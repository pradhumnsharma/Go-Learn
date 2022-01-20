## Introduction

In this folder, I have implemented the CRUD operation on Students and Teachers. I have learned about Struct, Nested Structs, Anonymous Structs, Slices, Http routing and controllers.

I have provided dummy json for students and teachers. I have created separate struct for Response type.


### How to use

* Start the application using `go run main.go`. Application will start on `http://localhost:6001`.
* Install Thunder client extension if using VS code or Postman for making api calls
* Select data from StudentsDummy.json for Students and TeachersDummy.json for Teachers

**Note:** For Students **RollNo** is unique identifier and for Teachers **Id** is unique identifier. When performing create operation these will generate randomly for first entry and then increment by 1 for rest entries.

Following operations are available for both tables

1. Create
2. Update
3. Get All
4. Get One
5. Delete All
6. Delete One


### What I Learned
1. How to create random numbers
2. How to use Structs and Nested structs
3. How to modify response keys as per UI requirements
4. How to hide a field from not showing on UI
5. How to use params for updating and deleting
6. How to use query params for  updating and deleting (Not implemented in this but other place)
7. How to convert string to number and number to string
8. Better understanding of pointer, reference and copy variable.



### `Note:` Please review the work. All suggestions related to improvements are highly welcome and appreciated.
