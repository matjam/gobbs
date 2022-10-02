# gobbs
BBS Software written in Go



# LUA Functions for the Terminal interface:

* input_char() - waits for a keypress, returns the input as a single char
* input_string(len int) - Will accept text input up to the specific length
                          and return it.
* send(s string) - sends a string to the connected user.
* parse_template(file string) - sends a given template to the connected user.
* get_users() - returns a list of all connected users.
* authenticate(username string, password string) - 