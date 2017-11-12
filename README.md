Ansible Inventory
==============================
I use Ansible to maintain / update three networks. But I have no control over local DNS or DHCP. Therefore, I can not use names to refer to machines.

Since each machine receives DHCP addresses (and therefore they change) I needed some way to generate inventory lists.

This tool is used to generate the Ansible inventory dynamically

Installation
-------------------

    $ github.com/utrescu/listIP
    $ go build ansibleinventory

Usage
-------------------
The program supports different parameters that can be seen with '-h'

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
    -t string
            Network timeout (default "1000ms")
    -timeout string
            Network timeout (default "1000ms")

The configuration is defined in a YAML file:

    # Coments
    networks:
    - name: servers
        address:
        - 192.168.4.0/24
    - name: hosts
        address:
        - 192.168.9.0/24
        - 192.168.10.0/24


where `name` is the host group and `address` are a list of addresses for a group.

Execution
---------------------------------
If I need to generate a inventory of machines with port 22 open using the configuration in `conf.yaml`. I can execute:

    $ ./ansibleinventori -i conf.yaml
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


