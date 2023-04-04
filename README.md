A restaurant keeps a log of (eater_id, foodmenu_id) for all the diners. The eater_id is a unique number for every diner and foodmenu_id is unique for every food item served on the menu. Write a program that reads this log file and returns the top 3 menu items consumed. If you find an eater_id with the same foodmenu_id more than once then show an error.

Expected output:

1.     Code that can handle the above problem statement

2.     Testcases (with example log files) that checks the possible conditions with unit tests.

3.     Code has to be submitted in your github repository (share the repo link).

4.     Containerize the application and host the image.


## How the service works?
- You can pass any file path as command line argument with **fpath** name.
- It will validate the file for the mentioned cases.

#### To build and Run the docker image:

- Use`docker build --tag diner_service  .` to build the image tag.
- Use `docker run diner_service` to run the docker image.
- You can pass other file names as well in the **fpath** flag
   example:
   `docker run dinner-test -fpath testdata/test_log_duplicate.csv`

Note : It uses the default file which is located in _app_ named **log.csv**