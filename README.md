# tantan test
- golang
- echo
- postgresql
- go-pg
- restful

## config file
- postgresql ==>> src/config/db.json

## postgresql structure

- database: tt_users
- schema: test
- table: tt_users、relation

```
tt_users-# \d test.tt_users;
                                       Table "test.tt_users"
 Column |          Type          | Collation | Nullable |                  Default
--------+------------------------+-----------+----------+-------------------------------------------
 id     | bigint                 |           | not null | nextval('test.tt_users_id_seq'::regclass)
 name   | character varying(128) |           | not null | ''::character varying
Indexes:
    "tt_users_pk" PRIMARY KEY, btree (id)
    "tt_users_name_uindex" UNIQUE, btree (name)

```
```
tt_users-# \d test.relation;
                                Table "test.relation"
 Column |  Type   | Collation | Nullable |                  Default
--------+---------+-----------+----------+-------------------------------------------
 id     | integer |           | not null | nextval('test.relation_id_seq'::regclass)
 uid1   | bigint  |           | not null | 0
 uid2   | bigint  |           | not null | 0
 state  | integer |           | not null | 0
Indexes:
    "relation_pk" PRIMARY KEY, btree (id)
    "uid" UNIQUE, btree (uid1, uid2)

```

## Run
go run main.go

## API

```
Users: 

GET
 /users 
List all users 
Example: 
$curl -XGET "http://localhost:80/users" 
[ 
  { 
    "id": "21341231231", 
    "name": "Bob" ,
    "type": "user" 
  }, 
  { 
    "id": "31231242322", 
    "name": "Samantha" ,
    "type": "user" 
  } 
] 


POST
 /users 
Create a user 
allowed fields: 
  name = string 
Example: 
$curl -XPOST -d '{"name":"Alice"}' "http://localhost:80/users" 
{ 
  "id": "11231244213", 
  "name": "Alice" ,
  "type": "user" 
}


Relationships: 
GET
 /users/:user_id/relationships 
List a users all relationships 
Example: 
$curl -XGET "http://localhost:80/users/11231244213/relationships" 
[ 
  { 
    "user_id": "222333444", 
    "state": "liked" ,
    "type": "relationship" 
  }, 
  { 
    "user_id": "333222444", 
    "state": "matched" ,
    "type": "relationship" 
  }, 
  { 
    "user_id": "444333222", 
    "state": "disliked" ,
    "type": "relationship" 
  }
] 


PUT
 /users/:user_id/relationships/:other_user_id 
Create/update relationship state to another user. 
allowed fields: 
   state = "liked"|"disliked" 
If two users have "liked" each other, then the state of the relationship is "matched" 
Example: 
$curl -XPUT -d '{"state":"liked"}' 
"http://localhost:80/users/11231244213/relationships/21341231231" 
{ 
  "user_id": "21341231231", 
  "state": "liked" ,
  "type": "relationship" 
} 
$curl -XPUT -d '{"state":"liked"}' 
"http://localhost:80/users/21341231231/relationships/11231244213" 
{ 
  "user_id": "11231244213", 
  "state": "matched" ,
  "type": "relationship" 
} 
$curl -XPUT -d '{"state":"disliked"}' 
"http://localhost:80/users/21341231231/relationships/11231244213" 
{ 
  "user_id": "11231244213", 
  "state": "disliked" ,
  "type": "relationship" 
} 

```

