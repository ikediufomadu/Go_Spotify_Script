# Go Spotify Script
## Overview
This Go program leverages the Spotify API to interact with user playlists. It includes functionalities for authenticating with Spotify, web scraping music data, manipulating an authorized user's public and private playlists, and performing specific queries on the Spotify platform.
## Prerequisites
Before running the program, ensure you have the following:
- Go installed on your machine.
- Have an application that allows you to modify .env files (notepad, vscode).
- A Spotify Developer account to obtain your client ID and client secret. https://developer.spotify.com/dashboard
- You must have this link as your Redirect URI: http://localhost:8080/callback
- Check Web API and Web Playback SDK when creating your spotify app.

## Setup 

- Clone or download the repository
- Create a .env file in the root directory and add your Spotify client ID and client secret:
    -  CLIENT_ID = your_spotify_client_id
    -  CLIENT_SECRET = your_spotify_client_secret
- Define the target website in the main.go file: targetWebsite := "https://example.com".
- Replace "https://example.com" with the website you want to scrape music data from.
- Note that you will have to manually change the html elements the program will search for if you choose a new website.
    - You can change this in the web_handler.go file look for the webScrapper method.
- Pick which playlist you want to add music to. You can do that by changing the second parameter of the AddTrackToPlaylist() method in the search.go file.
  - The playlist name is NOT case-sensitive.
- Choose a query type by updating the GenreQuery method parameter in main.go: GenreQuery("rock")
  - Replace 'rock' with the desired query type you can check the supported types in query.go the WhichQuery method.
## Usage

1. Spotify User Authentication - The program will then start a local server and output a Spotify authentication URL. Paste the string into your browser to log in to your Spotify account and authorize the program.
2. Web Scraping - Then it will visit the site to collect music data from.
3. Search - It then searches for the track data it collected from the music website of your choice.
4. Add to Playlist - Finally, it will add the first result to a playlist of your choice then it will stop execution.

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
