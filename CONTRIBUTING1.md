# How to contribute


## Use Cases

### Format


### Granularity
Please do not contribute use cases that are at the "CRUD" level. Any at that level of granularity will be rejected.

A use case should offer business value to the user and as such should be more meaningful than "Add Person", or "Update Address". A 

Any attempt to contribute a use case for logging in will not even be looked at. 

Use cases should not be used to design data actions. They should describe a broad purpose for which the user wants to use the system, and may result in a number of functions, classes, method calls, or other actions by the system. For example, the use case for "Assess Person Welfare" in the custody domain involves the Custody Manager using a questionnaire and recording the responses to the questions, recording their own visual assessment, and reports from other officers. The user will also search for the custody record of the person being held, assuming they have already been identified and their details entered in the use case: *"Accept Person Into Custody"*. **There should not be a separate use case for searching the list of people in custody!** Depending on the solution design, the user may access several screens and several database transactions may result. However, the user is only concerend with achieving one business goal. 

The delivery of a use case may be done in slices, or increments. However, the use case is not a design of the software. Use cases represent the requirements and should be expressed in a way that is meaningful to the person who will use the system.

Domain models, ER diagrams, sequence diagrams, component diagrams, and other types of graphical representation will be used in this project to design the software. Please do not put design into the use cases. There are other places to do that in this project.




