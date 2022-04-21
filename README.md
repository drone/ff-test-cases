# Server SDK - Evaluator test cases

Integration tests and fixtures for FF server side SDK

## Introduction to v2

- Better separation of flag and segment data from tests
- tests are more simple now

## Create the flag file

flag files are now in flags directory and the name of file should be exact the same as flag feature field:

bool_off_flag.json

```json
{
  "defaultServe": { "distribution": null, "variation": "true" },
  "environment": "demoenv",
  "feature": "bool_off_flag",
  "kind": "boolean",
  "offVariation": "false",
  "prerequisites": [],
  "project": "FeatureFlagsDemo",
  "rules": [],
  "state": "off",
  "variationToTargetMap": null,
  "variations": [
    {
      "description": null,
      "identifier": "true",
      "name": "Enable Page",
      "value": "true"
    },
    {
      "description": null,
      "identifier": "false",
      "name": "Disable Page",
      "value": "false"
    }
  ]
}
```

## Create the segment file

segment files are now in segments directory and the name of file should be exact the same as segment identifier field:

```json
{
  "identifier": "flagteam",
  "name": "FlagTeam",
  "environment": "demoenv",
  "included": [
    {
      "identifier": "harness",
      "name": "Harness"
    }
  ],
  "excluded": [],
  "rules": [],
  "version": 1
}
```

## Target data

targets can be created in targets directory and only mandatory fields are identifier and name

```json
{
  "identifier": "harness",
  "name": "Harness",
  "attributes": { // optional fields
    "email": "enver.bisevac@harness.io"
  }
}
```

## Test files

test files are much more simplified now and can reuse any of above entities

```json
{
  "flag": "demo_feature",
  "expected": {
    "_no_target": true,
    "harness": false,
    "enver": true
  }
}
```

## Validation of entity files

In progress