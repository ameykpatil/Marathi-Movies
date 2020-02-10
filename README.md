# Marathi Movies

## Idea
This project has been started to tell the movie enthusiast about the amazing movies made in Marathi language.
It will serve with the basic information of the movie such as rating, very short summary, streaming service it is available on etc.

## Execution
This is simple static project using the github pages & it's default themes.
It converts markdown file into an html. Even if it requires markdown the html has been preferred so that it can be migrated easily in future.

## How to add a new movie?
1. Go to this [link](https://play.golang.org/p/RgOiEBm5rrR). It holds the code snippet to generate html code.  
If due to some reason the link is broken, then go to [playground](https://play.golang.org) & copy the code in [generate_html.go](generate_html.go). 
1. Replace the values with the details corressponding to this new movie.  
1. Click Run.  
1. Copy the html in the output.  
1. Add it in index.md. 
1. Commit & Check the page after couple of minutes.     

## What are the necessary details of the movie & how to get them?
- `name` : Name of the movie  
- `picURL` : Link of the picture which will be displayed for this movie. You can head to movie's imdb page & open an image of the movie & copy the image address.  
- `titleID` : When you head to movie's imdb page, check the url for title-id starting from `tt` e.g. `tt5312232`   
- `year` : Year when the movie was released. Check movie's imdb page.  
- `description` : Short description of the movie. Check movie's imdb page.  
- `streamingURL` : Streaming service link where this movie is available. You need to search on different streaming platforms for this.  
