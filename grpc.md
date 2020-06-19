# How do we make apps talk to each other?

# How do we build API?

# What is REST?
- Architecutral design for HTTP. 

# Data format
Most commonly JSON

# Why HTTP/REST is good?
- Text based protocol, easy to understand
- Good tooling available 
- Maybe more    

# API lifecycle 
- Implementation
- Documentation
- Client library 

# Problems with HTTP/REST
- Writing client libaries is not fun.
    - Different languages
    - Naming convention/Types, etc

- Documentation can feel like a chore depending on framework. 

- Streaming is difficult

- RESTful doesn't always make sense

- We can do better than text based transport 

# gRPC 
- Define the services first. 

- Auto generate code.
    - Create client libraries
    - I've seen this with WSDL/SOAP
    - Server stubs 
        - Creating the service
        - Registering the service to the gRPC server

- Better data serialization
    - protobuf(pluggable)
- Streaming is easy and can be bi-directional


# What's the catch?
- Error handling
- No more REST, clients usually are HTTP/REST based for lots of frameworks. (More of an adoption issue)    

# Misc

grpcurl -plaintext localhost:9090 Welcome.World
grpcurl -plaintext -msg-template -emit-defaults -d '{"name": "CNCF"}' localhost:9090  Lookup.Find 






 