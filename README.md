# Welcome

Your mission, should you choose to accept it, involves these tasks.

## 1. Get the application up and running

1. install [go](https://golang.org)
2. install [revel](http://revel.github.io) (tested with 0.12.0)
3. install a recent version of [postgresql](http://www.postgresql.org). create a db user and database according to these specs `user=evard_bugs password=EvardBugs? dbname=evard_bugs`
4. launch the app and go to http://localhost:9000

## 2. Fix message submit
Right now the application will crash when you press Submit. Fix this.

## 3. Fix the missing page title
The page title is currently empty. Make sure it is set correctly from the PageLoad() method.

## 4. Add tests
Add automatic tests that verify that the two bugs (#2 and #3) above have been fixed.
