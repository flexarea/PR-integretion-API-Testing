Testing Trello API

-----------------------
Accomplishment checklist

1. **Trello Account and API Key:**
    - [x]  Create a Trello account 
    - [x]  Generate API Key
    - [x]  Generate an OAuth token
2. **Environment Setup:**
    - [x]  Create a `.env` file in project directory.
    - [x]  Add Trello API key and OAuth token to the `.env` file.
    - [x]  Install the `godotenv` package for loading environment variables in Go.
3. **Testing Authentication:**
    - [x]  Write a simple Go script to load the API key and token from the `.env` file.
    - [x]  Make a test API call to Trello (e.g., get a list of boards) to ensure authentication is working
-----------------------
**Note**
**docs on authorization**
https://developer.atlassian.com/cloud/trello/guides/rest-api/authorization/
-----------------------

**comment:**
    Testing on the first api call took quite a while as the trello Rest Api documentation is not that clear. Spent quite a while trying to figure out how to authorize, approve, and being able to write on user's trello account using provided endpoint

-----------------------
**Question and Issue:**
Unable and don't know how to use call the api with OAUTH token instead of (single) API generated token
-----------------------
4. **Reading Board:**
    - [x]  read Board's detail (e.g., name, description).
5. **Updating Board:**
    - [x]  update Board's detail (e.g., name, description).
    - [x]  Test the function by updating you created previously.
6. **lists**
    - [x]  read all the lists on the board
    - [x]  create a new list on the board
    - [x]  read a list
    - [x]  read all cards from a list
    - [x]  delete a card
7. **List Actions:**
8. **lists**
    - [x]  read information from card
    - [x]  Move card to targetList within same Board

-----------------------

Working on

    - [x]  refactoring code

