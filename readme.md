# PawNotes

A dog notes site for tracking your pet's activities, health, and memorable moments.

## Table of Contents
- [PawNotes](#pawnotes)
  - [Table of Contents](#table-of-contents)
    - [Overview](#overview)
    - [Tech Stack](#tech-stack)
    - [Design notes](#design-notes)

---

### Overview
PawNotes is a web application designed for pet owners to document their dog(s) activities, health, and memorable moments

### Tech Stack
Built using:
- **Frontend**: Vue.js
- **Backend**: Go, Air(used for live reloading)
- **Database**: PostgreSQL
- **Other**: Docker


### Design notes
- Handler -> Service -> Repository w/ interfaces for easy DI and swapping out. I.e Service holds a reference to a dbRepo, which implements some interface
  - Thus I could use a struct that holds a test SQLite DB and use that to implement the interface as well, and inject that into Service. This works since service doesn't care about the underyling type. Benefits is decoupling the code

