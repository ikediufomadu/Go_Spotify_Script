# Go Spotify Script
## _Overview_
This Go program leverages the Spotify API to interact with user playlists. It includes functionalities for authenticating with Spotify, web scraping music data, manipulating an authorized user's public and private playlists, and performing specific queries on the Spotify platform.
## Prerequisites
Before running the program, ensure you have the following:
- Go installed on your machine.
- A Spotify Developer account to obtain your client ID and client secret.
- You must have this as your Redirect URI http://localhost:8080/callback

## Setup

- Clone the repository: 
    - git clone https://github.com/ikediufomadu/Go_Spotify_Script.git
- Move to the folder location
    -  cd Go_Spotify_Script
-  Create a .env file in the root directory and add your Spotify client ID and client secret:
    -  CLIENT_ID = your_spotify_client_id
    -  CLIENT_SECRET = your_spotify_client_secret
## Usage
- Web Scraping
    - Define the target website in the main.go file:
    - targetWebsite := "https://example.com"
    - Replace "https://example.com" with the website you want to scrape music data from.
- Authentication
    - Run the program to start a local server for Spotify authentication. It is set to run on port 8080, you can change this if needed.
        - **If you want to change the starting port you will have to change the the redirect URI on the spotify dashboard. localhost:8080/callback and also in the redirectURI in the authenticator.go file**
    - Visit the provided authentication URL in your browser to log in to your Spotify account.
- Query Spotify
   - Choose a query type by updating the which query function parameter in main.go: 
    - Handler.WhichQuery(1)
    - Replace 1 with the desired query type (1 for song name query, 2 for artist name query, 3 for genre query).
- Run the Program
    - Run the program to scrape data, authenticate, and perform the specified query:

## Technicalities
- Playlist ID
    - As of 11/30/2023 you will have to create a new playlist, get the string from the CLI and then manually input it into the getPlaylist method in the playlist.go file.
- Changing server port
    - As stated before in the Authentication section of usage, if you want to change the  port number in authenticator.go file you will also need to change the callback URI's port number in the Spotify dashboard and also in the constant redirectURI in the authenticator.go file as well.

## License

This project is licensed under the MIT License.

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [dill]: <https://github.com/joemccann/dillinger>
   [git-repo-url]: <https://github.com/joemccann/dillinger.git>
   [john gruber]: <http://daringfireball.net>
   [df1]: <http://daringfireball.net/projects/markdown/>
   [markdown-it]: <https://github.com/markdown-it/markdown-it>
   [Ace Editor]: <http://ace.ajax.org>
   [node.js]: <http://nodejs.org>
   [Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
   [jQuery]: <http://jquery.com>
   [@tjholowaychuk]: <http://twitter.com/tjholowaychuk>
   [express]: <http://expressjs.com>
   [AngularJS]: <http://angularjs.org>
   [Gulp]: <http://gulpjs.com>

   [PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
   [PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
   [PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
   [PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
   [PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
   [PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>
