Ansible Inventory
==============================
I use Ansible to maintain / update three networks. But I have no control over local DNS or DHCP. Therefore, I can not use names to refer to machines.

Since each machine receives DHCP addresses (and they can change) I need some way to generate dinamic inventory lists.

This tool is to dynamically generate the Ansible inventory based on the machines it finds in a range of IP addresses

Installation
-------------------
First install dependencies:

    $ go get github.com/utrescu/listIP

And build:

    $ go build ansibleinventory

Usage
-------------------
The program supports different optional parameters that can be seen with '-h'

    $ ./ansibleinventory -h
    Usage of ./ansibleinventory:
    -debug
            verbose output
    -i string
            Name of configuration time (default "conf.yaml")
    -input string
            Name of configuration time (default "conf.yaml")
    -p int
            Port to scan (default 22)
    -port int
            Port to scan (default 22)
    -parallel int
            Number of parallel connections (default 32)
    -t string
            Network timeout (default "1000ms")
    -timeout string
            Network timeout (default "1000ms")

The configuration is defined in a YAML file:

    # Comments
    groups:
    - name: servers
      networks:
        - 192.168.4.1
        - 192.168.4.2
        - 192.168.4.254
    - name: hosts
      networks:
        - 192.168.9.0/24
        - 192.168.10.0/24


where `name` is the host group and `networks` are a list of network addresses in the group in CIDR format.

Execution
---------------------------------
To generate an inventory of machines defined in `conf.yml` that have port 22 open.. I can execute:

    $ ./ansibleinventory -i conf.yaml
    [servers]
    192.168.4.2
    192.168.4.254

    [hosts]
    192.168.9.12
    192.168.9.25
    192.168.10.5
    192.168.10.23
    192.168.10.26
    192.168.10.35

The inventory can be redirected to a file with `-o` or `-output`

    $ ./ansibleinventory -i conf.yaml -o inventory
