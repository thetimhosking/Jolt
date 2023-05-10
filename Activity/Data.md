# Activity Data 


## Activity

### Activity external interface

```json
    GUID string,
    desc string
```

### Activity resource

```json
"activity": {
    "description": "An instance of a managed business process or workflow, consisting of a number of tasks",
    "fields": [
        { "name": "id", "type": "guuid" },
        { "name": "begin-date", "type": "date"}
    ]

}
    
    end-date date,
    name string,
    ref-number string,
    type-code string,
    priority-type string,
    complexity-level string,
    reason-code string,
    total-time-target int,
    total-time-taken int,
    security-classification string,
    summary string,
    resolution-description
```


