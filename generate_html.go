package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "Sairat"
	picURL := "https://m.media-amazon.com/images/M/MV5BMjBjNWYyY2UtOGNjZC00MTg4LWIwYWItYmZiNDI5MGUwNWRkXkEyXkFqcGdeQXVyNjI1NjA5NjE@._V1_.jpg"
	titleID := "tt5312232"
	year := "2016"
	description := "In interior Maharashtra, a fisherman's son and a local politician's daughter fall in love against the restrictions of caste hierarchy."
	streamingURL := "https://www.zee5.com/movies/details/sairat/0-0-movie_1921832123"
	
	color := ""
	streamingService := ""
	if strings.Contains(streamingURL, "netflix") {
		color = "Red"
		streamingService = "NETFLIX"
	} else if strings.Contains(streamingURL, "primevideo") {
		color = "DodgerBlue"
		streamingService = "prime video"
	} else if strings.Contains(streamingURL, "hotstar") {
		color = "LimeGreen"
		streamingService = "hotstar"
	} else if strings.Contains(streamingURL, "zee5") {
		color = "Black"
		streamingService = "ZEE5"
	}

	code := `
<h2>%s</h2>
<table style="border:none;">
  <tr>
    <td style="width:20%%; border:none; float:left;">
      <img src="%s" alt="%s">
    </td>
    <td style="width:80%%; border:none; float:left;">
      <span class="imdbRatingPlugin" data-user="ur100837743" data-title="%s" data-style="p1">
	<a href="https://www.imdb.com/title/%s/?ref_=plg_rt_1">
	  <img src="https://ia.media-imdb.com/images/G/01/imdb/plugins/rating/images/imdb_46x22.png" alt="%s (%s) on IMDb" />
	</a>
      </span>
      <script>(function(d,s,id){var js,stags=d.getElementsByTagName(s)[0];if(d.getElementById(id)){return;}js=d.createElement(s);js.id=id;js.src="https://ia.media-imdb.com/images/G/01/imdb/plugins/rating/js/rating.js";stags.parentNode.insertBefore(js,stags);})(document,"script","imdb-rating-api");
      </script> &nbsp;&nbsp;&nbsp;&nbsp;Year : %s
      <p>
        <i>
	%s
	</i>  
        <br><br>
	Available on <a href="%s" style="color:%s; font-weight:bold;">%s</a>        
      </p>  
    </td>
  </tr>
</table>
`

	fmt.Printf(code, name, picURL, name, titleID, titleID, name, year, year, description, streamingURL, color, streamingService)
}
