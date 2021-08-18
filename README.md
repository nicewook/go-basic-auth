# Basic user authentication with password

## Overview

Let's make a HTTP API server which will sign up user and sign in user. It will accept new user with username/password and
hashed password using bcrypt, then save them to `sqlite` server. It will also accept username/password and check if the 
username is saved and password is correct.

## In detail

### Bcrypt

`Bcrypt` is a hash algorithm for password hashing. There are many other algorithms for this usage, and some of them are
much secure then `bcrypt`. Though, you can say `bycrypt` is the most common hash algorithm for the general usage. 

One of the major characteristic of the password hash algorithm is that it is slow enough against bruth forch attack.

### How to save, and how to compare

First, We make hashed password with bcrypt. then save it to DB. Hashed password have its own `salt` (You can find more 
info about `salt` by googling)

Second, If someone try to sign in, then we are making hash with the same `salt` which the saved (hashed password) already made of. 

## Blog posting in Korean 

- Blog post: https://jusths.tistory.com/228


## Reference

- Link: https://www.sohamkamani.com/golang/password-authentication-and-storage/
- GitHub Repo: https://github.com/sohamkamani/go-password-auth-example
