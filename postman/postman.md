# Beispiel für Development:
{
  "completed_at": "2023-02-20T04:59:12.000+00:00",
  "responsibility": "development",
  "description": "This is the Description.",
  "created_at": "2023-02-10T18:23:01.000+00:00",
  "reporter": {
    "firstname": "Bob",
    "role": "development",
    "surname": "Baumeister",
    "id": 8,
    "email": "bob.bau@example.com"
  },
  "id": 2,
  "assignee": {},
  "title": "This is the Title.",
  "priority": "low",
  "status": "created"
}

---
localhost:8081/todos
localhost:8081/todo/status/created
----------------------------------------
# Beispiel für Support:
{
  "completed_at": "2023-02-24T09:56:49.000+00:00",
  "responsibility": "support",
  "description": "This is the Description.",
  "created_at": "2023-02-23T14:50:07.000+00:00",
  "reporter": {
    "firstname": "Max",
    "role": "support",
    "surname": "Mustermann",
    "id": 9,
    "email": "max.mus@example.com"
  },
  "id": 3,
  "assigned": false,
  "title": "This is the Title.",
  "priority": "high",
  "status": "created"
}

---
localhost:8080/todos
localhost:8080/todo/responsibility/development/status/created
