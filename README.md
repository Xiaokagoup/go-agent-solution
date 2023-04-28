# HelloWorldGoAgent

## Added by Jiean

One thing to note is that this is **a preliminary codebase**, which means that there are many areas for improvement, including naming conventions for variables that may be influenced by personal habits. I will listen to everyone's suggestions during the development process to improve the code. The code is uploaded now to give everyone an initial understanding and to be able to understand the relevant code if needed.

This is our agent code repository, and let me explain the project structure:

- /src/main.go: This is the entry point of the project, where an agent instance is created and the start function is called.
- /src/agent.go: This is the core business logic, which handles the heartbeat signal, receives and handles messages, and sends results and collected metrics to the backend. Each action is provided by a module in the tools folder.
- /src/agentHttp: This provides external APIs for easy calling by the backend when needed, and these interfaces are secured with a PSK key for security.
- /src/tools: This is our tool library, where each folder represents a module with different functions providing different features or actions.
- /src/docs: This folder is used to support automatic generation of swagger files, providing a visual interface for our APIs. (Run `cd src, ./swagger.sh`)
- /\*_/_\_test.go: This is the Go language's test folder.
