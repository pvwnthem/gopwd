# gopwd
### gopwd is a pass like terminal based password manager written in golang

## Initalizing a vault
#### run gopwd vault init (-p path to the vault (optional). Default - $HOME/.gopwd) (-n name of the vault (optional). Default - vault) 
EXAMPLE
```
gopwd vault init -p /home/'your username'/vaults -n main
```

#### this will initailize a vault at /home/'your username'/vaults/main

### Additional information about initalizing a vault
if you use the default options, all other commands can be carried out without flags or a config file, they will use the default location. HOWEVER, if you used a custom location it is recommended to setup the path and name of your vault in a config file. If you do not do this, in order to use your custom vault you will have to set your path flag and name flag for every command you run (using -p [path] and -n [name]). Having a config will automatically fill in the path and name inside the config to the command so you don't have to
To setup a config either run
```
gopwd config init -p "path to where your vault is NOT THE FULL PATH, JUST THE PATH TO THE DIR IT'S IN" -n "name of your vault (name of the actual vault folder)"
```
or run
```
gopwd config init 
```
and then set the config options manually with
```
gopwd config set "field" "value"
```
the field will be either "path" or "name"
### You can change this config anytime using the above command
## Commands
```
gopwd insert [service that the password is used for]
```
#### this will insert a password into your vault for the service you specify
this service name can be anything, including a website, or a username, or anything else you want to use it for. IMPORTANT: it can not contain spaces, if you want to use a space use a dash (-) instead, it can also not contain any special characters other than a dash (-) or an underscore (_)

#### the service name can be nested aswell, for example if you have multiple github passwords you could set one as github/personal and another as github/work. This is not required but can be useful for organization. If you do not want to use this feature, just use the service name as the service you want to use it for

```
gopwd generate [service that the password is used for] [length of password]
```
this command is the same as insert except it generates a password for you along with inserting that password into the vault. 
Some example service names and examples of what they look like in the vault: 
```
github 
- vault
 | - github
    | - password

github/personal
- vault
 | - github
    | - personal
       | - password

github/work and github/personal in the same vault
- vault
 | - github
    | - personal
       | - password
    | - work
       | - password
```
#### dont worry about the password file, that is just your encrypted password which will be automactically retrieved when you run the command "gopwd show [service name]"

```
gopwd show [service that the password is used for]
```
this command shows the password for the service that you provide.

```
gopwd rm [service that the password is used for]
```
this command removes the password and folder for the service that you provide.

```
gopwd edit [service that the password is used for]
```
this opens up a nano window and allows you to edit the password or add any metada you want such as an email or username if you use different ones for different accounts