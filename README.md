
					██████╗ ██████╗  ██████╗ ██╗  ██╗ ██████╗██╗     ██╗
					██╔══██╗██╔══██╗██╔═══██╗╚██╗██╔╝██╔════╝██║     ██║
					██████╔╝██████╔╝██║   ██║ ╚███╔╝ ██║     ██║     ██║
					██╔═══╝ ██╔══██╗██║   ██║ ██╔██╗ ██║     ██║     ██║
					██║     ██║  ██║╚██████╔╝██╔╝ ██╗╚██████╗███████╗██║
					╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝╚══════╝╚═╝


Proxcli is a small CLI client for the Proxmox API that allows you to start, stop and list the VMs and LXC containers on your Proxmox server. It also supports the creation of an inventory file to be able to call the VMs by its name and not just by ID. You can create groups inside the inventory file, where you can group VMs with a similar purpose to start and stop them together.

## Configuration

Proxcli expect a **yaml** config file called **proxcli.yml** under the folder **$HOME/.proxcli** and the content needs to have the following structure:

```yaml
config:
  - node: <node-name>
    ip: '<node-ip>'
    security:
      user: <user>
      realm: pve
      tokenid: <token-id>
      token: <api-token>
```

Currently, it only supports authentication through an API token, so you will need to create one and give enough permissions to it and the user  to be able to perform the actions.

>  Note:  At the moment, you are limited to one node, but I am working to add support to manage multiple nodes.

## Inventory

Once you put your configuration information in the configuration file, you can create the inventory of your VMs and Lxc containers, either manually or using the command as shown below.

![](img/Pasted%20image%2020240512234049.png)

The command by default will collect the information of all the VMs and Lxc containers on the node and create a file called **inventory.yml** under the same folder as the configuration file (**$HOME/.proxcli**) with the following structure that you can use to create the file manually if preferred.

```yaml
vms:
    - name: <vm-name>
      id: <vm-id>
    - name: <vm-name>
      id: <vm-id>
lxc:
    - name: <lxc-name>
      id: <lxc-id>
    - name: <lxc-name>
      id: <lxc-id>
```

If you change any Vm or container info like the name, id or delete one, and want to update the information  you can just run the command again whit the source that you want to update using the **--source** flag  with one of the following values **vm**, **lxc** or **all**, and it will append another block of **vms** or **lxc** to the file with the most recent information, and you just delete the old one.


![](img/Pasted%20image%2020240512235405.png)


![](img/Pasted%20image%2020240512235432.png)

## Groups

Within the inventory file you can add a section called groups where you can specify the name of the group and the Vms that are part of the group. The section needs to have the following structure:

```yaml
groups:
  - name: <group-name>
    vms:
      - name: <vm-name>
        id: <vm-id>
      - name: <vm-name>
        id: <vm-id>
  - name: <group-name>
    vms:
      - name: <vm-name>
        id: <vm-id>
      - name: <vm-name>
        id: <vm-id>
```

## Commands


### Inventory
Create inventory file of VMs and Lxc containers
- ### --source
	- #### all (default)
	
		![](img/Pasted%20image%2020240513003336.png)

	- #### vms
	
		![](img/Pasted%20image%2020240512235432.png)

	- #### lxc

		![](img/Pasted%20image%2020240512235405.png)

### Vm
- #### get
	Display more detail info about the Vm, like cpu and memory.
	- by id
	
		![](img/Pasted%20image%2020240430164323.png)

		 
	- by name
	

		![](img/Pasted%20image%2020240430164419.png)


- #### list
	Display all VMs along with their current status. 
	 -  state (default "all")
	 
		![](img/Pasted%20image%2020240430182155.png)


	 - state running
	 
		![](img/Pasted%20image%2020240430182321.png)


	 - state stopped
	
		![](img/Pasted%20image%2020240430182423.png)
		
- #### start
	Start a Vm by specifying its ID or, if the inventory is configured, also by Name.
	 - by id 
	
		![](img/Pasted%20image%2020240430183040.png)

	 - by name

		![](img/Pasted%20image%2020240430183229.png)
		
- #### stop
	Stop a Vm by specifying its ID or, if the inventory is configured, also by Name.
	 - by id

		![](img/Pasted%20image%2020240430183426.png)
		
	 - by name
	
		![](img/Pasted%20image%2020240430183510.png)\
- #### group
	perform actions on a group if it is configured in the inventory file.
	 - list
	 
		List the groups and their members
	
		![](img/Pasted%20image%2020240430183745.png)
		
	 - get

		Display more detailed information about the VMs, which are part of the group.
	
		![](img/Pasted%20image%2020240430184055.png)
		
	 - start

		Start the VMs in a group, specifying the group name.
		
		![](img/Pasted%20image%2020240430184339.png)
		
	 - stop

		Stop the VMs in a group, specifying the group name.
				
		![](img/Pasted%20image%2020240430184209.png)

### lxc
- #### list 
	Display all containers along with their current status.
	 -  state (default "all")

		![](img/Pasted%20image%2020240512235920.png)	

	 - state running
	 
		![](img/Pasted%20image%2020240513000119.png)

     - state stopped

		![](img/Pasted%20image%2020240513000200.png)
- #### get
	Display more detail info about the Vm, like cpu and memory.
	- by id
	
		![](img/Pasted%20image%2020240513000351.png)

	- by name

		![](img/Pasted%20image%2020240513000712.png)

- ### start
	- by id
	
		![](img/Pasted%20image%2020240513002602.png)

	- by name

		![](img/Pasted%20image%2020240513002800.png)
	
- ### stop
	- by id
	
		![](img/Pasted%20image%2020240513002713.png)

	- by name

		![](img/Pasted%20image%2020240513002849.png)
### node
- #### get
	Display more detail info about the node, like cpu and memory and load average.
	
	![](img/Pasted%20image%2020240430185613.png)