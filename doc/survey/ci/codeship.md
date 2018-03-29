# Codeship

https://codeship.com/

codeship-services.yml

````yaml
lines (17 sloc)  399 Bytes
web-codeship-example-php:
  build: .
  depends_on:
    - postgres
  environment:
    DATABASE_URL: postgres://todoapp@postgres/todos
  cached: true
postgres:
  image: healthcheck/postgres:alpine
  environment:
    POSTGRES_USER: todoapp
    POSTGRES_DB: todos
codeship-heroku-deployment:
  image: codeship/heroku-deployment
  encrypted_env_file: deployment.env.encrypted
  volumes:
    - ./:/deploy
````

codeship-steps.yml

````yaml
- name: tests
  service: web-codeship-example-php
  command: phpunit
- name: deploy
  tag: master
  service: codeship-heroku-deployment
  command: codeship_heroku deploy /deploy php-laravel-todoapp
- name: migrate
  tag: master
  service: codeship-heroku-deployment
  command: heroku run --app php-laravel-todoapp -- php artisan migrate --no-interaction
````