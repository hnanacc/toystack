* With the introduction of virtual machines, the operating systems were decoupled from the physical machine. The broker being hypevisor.
* But the applications are still coupled to the operating systems.
* With the introduction of docker, the application are decoupled from the operating system.
* In context of orchestrator, a container is like a process to an operating system. An OS manages processes and orchestrator manages container.
* **Task** is some process, a restful server, haproxy, or anything else. At minimum we need three things to describe a task: 1. resources(cpu, disk, mem), 2. what to do when crashed (restart policy), 3. the image to use to instantiated the task.
* **Job** is a group of task. A complete set of microservices and load balances. We need two things to describe a Job, 1. tasks to run and their properties(num), 2. type of job, optionally, 3. which datacenter to run in.
* **Scheduler** decides on which node(some machine) each task runs. Formally, 1. it determines which nodes can run the task. 2. score each of these nodes, 3. pick the node based on the best score.
* **Worker** runs on a task given to it and collects statistics about itself and the task. It can be several instances on a single machine or several machines.
* **Manager** is the brain of the orchestrator, it gets info from everywhere(user, scheduler, worker) and instructs the workers what to do.
* **Cluster** is the logical grouping of all of the above components. It could be deployed across several machines or on a single machine.
* **CLI** is the user primary interface we use to talk to the manager. Start/stop tasks, start/stop workers and manager and get statistics.
* 
