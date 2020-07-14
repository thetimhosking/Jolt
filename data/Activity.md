# Activity

An activity includes the actions, processes, jobs, and tasks the organisation does across its value chain. 

The main goal of these activities is to deliver products and services to clients in their constituent market segments. However, some activities are "backroom" activities that are necessary to manage the resources needed to deliver products or services. Or, they may be activities needed to meet compliance with regulations, such as, auditing of financial records, etc.

Some activities are long-lasting, as in case management, which is a main sub-type of activity.

Activities begin with some triggering event and have a defined completion. Events can also be triggered by activities.

Activities can be measured for their time and cost as well as their success in meeting agreed standards. They can be monitored while being performed. They can be evaluated after completion.

Activities have participants: the parties in their given roles who perform or are otherwise involved in the activity.

Activities occur at given locations, including electronic locations.

```mermaid
erDiagram
    Activity }o..|{ Activity : includes
    Activity }|..|{ Participation : has
    Activity }|..o{ Cost : incurs
    Activity }|..o{ Risk : "has or mitigates"
    Activity }|..o{ Outcome : achieves
    Activity }|..o{ Issue : addresses
    Activity }o..o{ Event : "triggers or is triggered by"
    Activity }o..o{ Evaluation : "have success measured by"
    Participation }o..|{ PartyRole : of
    Participation }o..o{ Location : at
    Participation }o..o{ Resource : of
```

```mermaid
classDiagram
    Activity <.. Resource Management : 
    Activity <.. Case Management 
    
```
