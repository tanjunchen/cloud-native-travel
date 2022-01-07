namespace go Sample

struct User {
    1:required i32 id;
    2:required string name;
    3:required string avatar;
    4:required string address;
    5:required string mobile;
}

struct UserList {
    1:required list<User> userList;
    2:required i32 page;
    3:required i32 limit;
}