Server keeps asking for password after I've copied my SSH Public Key to authorized_keys

I have an Ubuntu Server, running in a Cloud. I created a user (`git`). In the folder `/home/git`, I have created the `.ssh/` dir, and the `authorized_keys` file.

But, when I put my SSH Public Key in the `authorized_keys` file, the server continues to ask me the password.

What did I do wrong?



## Generate Key
You can generate an SSH key pair in Ubuntu using the following command:


`ssh-keygen -t rsa -b 4096 -C "your_email@example.com"`

Replace "your_email@example.com" with your actual email address.



## Upload public key to host

After generating the key pair, you can upload the public key to the server by adding it to the `~/.ssh/authorized_keys` file on the server. You can use the `ssh-copy-id` command to automate this process:


`ssh-copy-id user@hostname`

Replace `user` with your username on the server and `hostname` with the server's hostname or IP address. This command will copy your public key to the server and add it to the `authorized_keys` file.


If you want to specify a different public key, you can use the `-i` option with `ssh-copy-id`. For example:

```bash
ssh-copy-id -i /path/to/public_key user@hostname
```

```bash
 ssh 'vagrant@192.168.56.56' -i id_rsa_homestead
```

## scp copy file

```bash
scp source_file_path username@remote_host:destination_file_path
```