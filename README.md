
        ██████╗ ██████╗  ██████╗ ██╗  ██╗ ██████╗██╗     ██╗
        ██╔══██╗██╔══██╗██╔═══██╗╚██╗██╔╝██╔════╝██║     ██║
        ██████╔╝██████╔╝██║   ██║ ╚███╔╝ ██║     ██║     ██║
        ██╔═══╝ ██╔══██╗██║   ██║ ██╔██╗ ██║     ██║     ██║
        ██║     ██║  ██║╚██████╔╝██╔╝ ██╗╚██████╗███████╗██║
        ╚═╝     ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝ ╚═════╝╚══════╝╚═╝


Proxcli is a small CLI client for the Proxmox API that allows you to start, stop and list the VMs on your Proxmox server. It also supports the creation of an inventory file to be able to call the VMs by its name and not just by ID. You can create groups inside the inventory file, where you can group VMs with a similar purpose to start and stop them together.

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

Once you put your configuration information in the configuration file, you can create the inventory of your VMs, either manually or using the command as shown below.

![](img/Pasted%20image%2020240430142718.png)

The command will collect the information of all the VMs on the node and create a file called **inventory.yml** under the same folder as the configuration file (**$HOME/.proxcli**) with the following structure that you can use to create the file manually.

```yaml
vms:
    - name: <vm-name>
      id: <vm-id>
    - name: <vm-name>
      id: <vm-id>
    - name: <vm-name>
      id: <vm-id>
    - name: <vm-name>
      id: <vm-id>
```

If you change any Vm info like the name, id or delete one, you can just run the command again, and it will append another block of **vms** to the file with the most recent information, and you just delete the old one.

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

### Vm
- #### inventory
	Create or update inventory file.
	
	![](img/Pasted%20image%2020240430142718.png)


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
		
### node
- #### get
	Display more detail info about the node, like cpu and memory and load average.
	
	![](img/Pasted%20image%2020240430185613.png)