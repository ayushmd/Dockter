## Tasks of Master Node
- Pool every worker node
- Decide where the **Task** recides (Load balancer)
- Send build request to builder
- Connect with database to resolve where the **Task** is stored in which node (External Routing)
- Check if above threshold then **Task** preemtion

## What will Master recieve
- Request for task server
- Request from builder to deploy the build
- Response of pooling from worker nodes