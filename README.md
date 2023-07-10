# gopwd - A Terminal-Based Password Manager in Golang

## Introduction
gopwd is a command-line password manager written in Golang. It provides a secure and convenient way to store and manage your passwords.

## Initializing a Vault
To initialize a vault, use the following command:

```
gopwd vault init [-p <path>] [-n <name>]
```

- `-p` (optional): Specify the path to the vault (default: `$HOME/.gopwd`).
- `-n` (optional): Specify the name of the vault (default: "vault").

**Example:**

```
gopwd vault init -p /home/'your username'/vaults -n main
```

This command initializes a vault at `/home/'your username'/vaults/main`.

### Additional Information on Initializing a Vault
If you use the default options, all other commands can be executed without specifying flags or a config file. They will automatically use the default vault location. However, if you have used a custom location, it is recommended to set up the path and name of your vault in a config file. Without a config file, you will need to specify the path and name flags for every command you run (`-p [path]` and `-n [name]`).

To set up a config file, you have two options:

1. Run the following command:

   ```
   gopwd config init -p "path to the directory where your vault is (not the full path)" -n "name of your vault (name of the actual vault folder)"
   ```

2. Run the following command to initialize the config file:

   ```
   gopwd config init
   ```

   Then, manually set the config options using:

   ```
   gopwd config set <field> <value>
   ```

   The `<field>` can be either "path" or "name".

You can change the config settings anytime using the above commands.

## Commands

### Inserting a Password

To insert a password into your vault for a specific service, use the following command:

```
gopwd insert <service>
```

- `<service>`: Specify the service for which the password is used. The service name can be anything, including a website, username, or any other identifier. **Important:** The service name cannot contain spaces. Use a dash (-) instead. Special characters other than a dash or underscore are also not allowed.

You can nest the service names to organize your passwords. For example, if you have multiple GitHub passwords, you can set one as `github/personal` and another as `github/work`. Nesting is optional and is useful for organizational purposes. If you don't want to use this feature, simply use the service name as the service identifier.

### Generating a Password

To generate a password for a specific service and insert it into the vault, use the following command:

```
gopwd generate <service> <length>
```

- `<service>`: Specify the service for which the password is used.
- `<length>`: Specify the length of the generated password.

Example service names and their representation in the vault:

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

**Note:** The password file is the encrypted version of your password, which will be automatically retrieved when you run the command `gopwd show <service>`.

### Showing a Password

To view the password for a specific service, use the following command:

```
gopwd show <service>
```

- `<service>`: Specify the service for which you want to view the password.

### Removing a Password

To remove a password and its associated folder for a specific service, use the following command:

```
gopwd rm <service>
```

- `<service>`: Specify the service for which you want to remove the password.

### Editing a Password

To edit a password or add metadata such as an email or username, use the following command:

```
gopwd edit <service>
```

- `<service>`: Specify the service for which you want to edit the password.

This command opens a Nano text editor window, allowing you to modify the password or add any desired metadata.

Feel free to reach out if you have any further questions!