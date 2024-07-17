https://juejin.cn/post/6984210017398292487

Cookie is the session id in client
Session is the session in server

# cookie-session 
Server holds the id to identify the cookie is valid or not

Just like , guest has a id card number, and server also have the id card record, so server can judge the id card number is authenticated

# jwt token
Server do not need hold the session id, session id is included in the token, named sub
server check the id  by the sign is correct or not.

Just like only guest have id card number, but the id card number is signed, only the server can check if the signed is correct or not.