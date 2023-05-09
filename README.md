# MicroService architecture for law enforcement

The services in this repo represent those that will deliver software features for core policing and law enforcment domains.

The microservices belong in two groups:

1. Enterprise domains - used by the application domains and are more general in application
1. Applications domains - implement features and use cases that are specific to law enforcement

## Application Domains

The domains in this application are the major doamins specific to law enforcement:

1. Incident Response - recording reports of incidents and responding to them
1. Investigation - the management of investigations into criminal activity
1. Custody - the process of taking a person into custody and managing their welfare while in custody
1. Intelligence - gathering and organising information from various sources that assists in detecting or investigating crime
1. Legal Process - recording charges, the legal nature of offences, the preparation of a legal brief, court appearances, and recording sentences

Further domains may be added as they become clear. For example, a domain for counter terrorism may be added.

## Enterprise domains

Enterprise domains offer common features that may be used by several application domains.

The enterprise domains are:

1. Party: the identity of individual persons and organisations
2. Object: both tangible and intangible "things" that can be involved in incidents or may be used for evidence
3. Location: the geographical place, or address, or electronic space where an event occurs
4. Event: the occurrence at a given time that is of interest, involving people and/or organisations, objects, and locations
5. Activity: the things the organisation does in response to events, or to provide a service
6. Resource: facilities, assets, information, etc. that may be used in many business processes
7. Service: a business service that the organisation provides; not to be confused with a software service that might support or enable business services. For example, a protection service may be provided for VIPs, protected witnesses in a matter, informants, etc. A missing person service may be provided to find missing persons and to support their relatives and friends


## Law Enforcement Use Cases

The use cases here are organized by the application domain microservices that implement them.

1. Incident Response - responding to calls for assistance from the public
1. Custody - taking a person into custody (as a suspected offender, for questioning regarding an investigation, or as a vulnerable person)
1. Investigation Management - investigating criminal activity, usually arising from the response to incidents
1. Intelligence Gathering - gathering, organizing, and disseminating intelligence about crime or terrorism
1. Legal Process - charging and prosecuting an offender


# Jolt

Domain driven design and data architecture come together to speed your delivery of software in law enforcement.

This project provides a number of high-level bounded contexts to guide a micro-service architeture for just about any organisation or business software. It is based on a well defined business architecture that should be used by any commercial, government, or non-profit organisation.

The foundation principle for the business architecture is:
A business exists to deliver products or services to clients in their constituent market segments via a business model and strategy that delivers a desired return on investment. 

The decisions about what products and services to deliver and the business model(s) to use are driven by the organisation's strategy. 

An organisation's strategy is defined by its:

- Mission - why the organisation exists and where the organisation sees itself in the world
- Vision - where the organisation wants to go
- Goals and objectives - how the organisation will measure its progress towards the vision

The Business Model consists of the story and numbers that describe how the organisation will deliver its desired return on investment by making choices about the channels, level of quality, pricing, and market segments to target.

To deliver the services and products within the selected business models and strategy, the organisation needs to manage:

- Data
- Process
- Business rules
- Resources - human, technical, financial, and informational

To measure progress towards the goals and objectives, the organisation must gather information at the across the entire delivery chain. Worklaod indicators and performance measures are gathered and used to assist in decision-making and to report against the strategy. These decisions affect the projects and other initiatives that deliver the capabilities needed to support the evolving strategy.

An understanding of the influences exerted on the organisation, both internally and externally, will help miitgate against risks and maximise opportunities for improved product and service delivery.

A software architecture that aligns with this understanding of the behaviour of any business can be more resilient to change.

This project attempts to define the data domains to support this thinking. These domains can contribute to a microservice architecture and provide patterns for the development of software components within boudned contexts. 