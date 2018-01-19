# Design draft

- contains UI for both normal user and BenchHub admin (since it can be self hosted)
- [ ] 

## directory layout

m for module, c for component

- auth (m)
  - register (c)
  - login (c)
  - oauth ?
- dashboard (m)
- admin (m)

````bash
ng generate module auth --routing
cd src/app/auth
ng generate component auth --module auth --flat --inline-style --inline-template
ng generate component login --module auth
ng generate component register --module auth
````
