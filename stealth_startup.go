package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

)

const (
	ConsumerKey    = "ConsumerKey"
	ConsumerSecret = "ConsumerSecret"
)



type AutoGenerated struct {
	Statuses []struct {
		CreatedAt string `json:"created_at"`
		ID int64 `json:"id"`
		IDStr string `json:"id_str"`
		Text string `json:"text"`
		Truncated bool `json:"truncated"`
		Entities struct {
			Hashtags []struct {
				Text string `json:"text"`
				Indices []int `json:"indices"`
			} `json:"hashtags"`
			Symbols []interface{} `json:"symbols"`
			UserMentions []interface{} `json:"user_mentions"`
			Urls []struct {
				URL string `json:"url"`
				ExpandedURL string `json:"expanded_url"`
				DisplayURL string `json:"display_url"`
				Indices []int `json:"indices"`
			} `json:"urls"`
		} `json:"entities"`
		Metadata struct {
			IsoLanguageCode string `json:"iso_language_code"`
			ResultType string `json:"result_type"`
		} `json:"metadata"`
		Source string `json:"source"`
		InReplyToStatusID interface{} `json:"in_reply_to_status_id"`
		InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
		InReplyToUserID interface{} `json:"in_reply_to_user_id"`
		InReplyToUserIDStr interface{} `json:"in_reply_to_user_id_str"`
		InReplyToScreenName interface{} `json:"in_reply_to_screen_name"`
		User struct {
			ID int `json:"id"`
			IDStr string `json:"id_str"`
			Name string `json:"name"`
			ScreenName string `json:"screen_name"`
			Location string `json:"location"`
			Description string `json:"description"`
			URL interface{} `json:"url"`
			Entities struct {
				Description struct {
					Urls []interface{} `json:"urls"`
				} `json:"description"`
			} `json:"entities"`
			Protected bool `json:"protected"`
			FollowersCount int `json:"followers_count"`
			FriendsCount int `json:"friends_count"`
			ListedCount int `json:"listed_count"`
			CreatedAt string `json:"created_at"`
			FavouritesCount int `json:"favourites_count"`
			UtcOffset int `json:"utc_offset"`
			TimeZone string `json:"time_zone"`
			GeoEnabled bool `json:"geo_enabled"`
			Verified bool `json:"verified"`
			StatusesCount int `json:"statuses_count"`
			Lang string `json:"lang"`
			ContributorsEnabled bool `json:"contributors_enabled"`
			IsTranslator bool `json:"is_translator"`
			IsTranslationEnabled bool `json:"is_translation_enabled"`
			ProfileBackgroundColor string `json:"profile_background_color"`
			ProfileBackgroundImageURL string `json:"profile_background_image_url"`
			ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
			ProfileBackgroundTile bool `json:"profile_background_tile"`
			ProfileImageURL string `json:"profile_image_url"`
			ProfileImageURLHTTPS string `json:"profile_image_url_https"`
			ProfileBannerURL string `json:"profile_banner_url"`
			ProfileLinkColor string `json:"profile_link_color"`
			ProfileSidebarBorderColor string `json:"profile_sidebar_border_color"`
			ProfileSidebarFillColor string `json:"profile_sidebar_fill_color"`
			ProfileTextColor string `json:"profile_text_color"`
			ProfileUseBackgroundImage bool `json:"profile_use_background_image"`
			HasExtendedProfile bool `json:"has_extended_profile"`
			DefaultProfile bool `json:"default_profile"`
			DefaultProfileImage bool `json:"default_profile_image"`
			Following interface{} `json:"following"`
			FollowRequestSent interface{} `json:"follow_request_sent"`
			Notifications interface{} `json:"notifications"`
			TranslatorType string `json:"translator_type"`
		} `json:"user"`
		Geo interface{} `json:"geo"`
		Coordinates interface{} `json:"coordinates"`
		Place interface{} `json:"place"`
		Contributors interface{} `json:"contributors"`
		IsQuoteStatus bool `json:"is_quote_status"`
		RetweetCount int `json:"retweet_count"`
		FavoriteCount int `json:"favorite_count"`
		Favorited bool `json:"favorited"`
		Retweeted bool `json:"retweeted"`
		PossiblySensitive bool `json:"possibly_sensitive,omitempty"`
		Lang string `json:"lang"`
		RetweetedStatus struct {
			CreatedAt string `json:"created_at"`
			ID int64 `json:"id"`
			IDStr string `json:"id_str"`
			Text string `json:"text"`
			Truncated bool `json:"truncated"`
			Entities struct {
				Hashtags []struct {
					Text string `json:"text"`
					Indices []int `json:"indices"`
				} `json:"hashtags"`
				Symbols []interface{} `json:"symbols"`
				UserMentions []struct {
					ScreenName string `json:"screen_name"`
					Name string `json:"name"`
					ID int64 `json:"id"`
					IDStr string `json:"id_str"`
					Indices []int `json:"indices"`
				} `json:"user_mentions"`
				Urls []interface{} `json:"urls"`
			} `json:"entities"`
			Metadata struct {
				IsoLanguageCode string `json:"iso_language_code"`
				ResultType string `json:"result_type"`
			} `json:"metadata"`
			Source string `json:"source"`
			InReplyToStatusID interface{} `json:"in_reply_to_status_id"`
			InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
			InReplyToUserID interface{} `json:"in_reply_to_user_id"`
			InReplyToUserIDStr interface{} `json:"in_reply_to_user_id_str"`
			InReplyToScreenName interface{} `json:"in_reply_to_screen_name"`
			User struct {
				ID int64 `json:"id"`
				IDStr string `json:"id_str"`
				Name string `json:"name"`
				ScreenName string `json:"screen_name"`
				Location string `json:"location"`
				Description string `json:"description"`
				URL interface{} `json:"url"`
				Entities struct {
					Description struct {
						Urls []interface{} `json:"urls"`
					} `json:"description"`
				} `json:"entities"`
				Protected bool `json:"protected"`
				FollowersCount int `json:"followers_count"`
				FriendsCount int `json:"friends_count"`
				ListedCount int `json:"listed_count"`
				CreatedAt string `json:"created_at"`
				FavouritesCount int `json:"favourites_count"`
				UtcOffset interface{} `json:"utc_offset"`
				TimeZone interface{} `json:"time_zone"`
				GeoEnabled bool `json:"geo_enabled"`
				Verified bool `json:"verified"`
				StatusesCount int `json:"statuses_count"`
				Lang string `json:"lang"`
				ContributorsEnabled bool `json:"contributors_enabled"`
				IsTranslator bool `json:"is_translator"`
				IsTranslationEnabled bool `json:"is_translation_enabled"`
				ProfileBackgroundColor string `json:"profile_background_color"`
				ProfileBackgroundImageURL string `json:"profile_background_image_url"`
				ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
				ProfileBackgroundTile bool `json:"profile_background_tile"`
				ProfileImageURL string `json:"profile_image_url"`
				ProfileImageURLHTTPS string `json:"profile_image_url_https"`
				ProfileBannerURL string `json:"profile_banner_url"`
				ProfileLinkColor string `json:"profile_link_color"`
				ProfileSidebarBorderColor string `json:"profile_sidebar_border_color"`
				ProfileSidebarFillColor string `json:"profile_sidebar_fill_color"`
				ProfileTextColor string `json:"profile_text_color"`
				ProfileUseBackgroundImage bool `json:"profile_use_background_image"`
				HasExtendedProfile bool `json:"has_extended_profile"`
				DefaultProfile bool `json:"default_profile"`
				DefaultProfileImage bool `json:"default_profile_image"`
				Following interface{} `json:"following"`
				FollowRequestSent interface{} `json:"follow_request_sent"`
				Notifications interface{} `json:"notifications"`
				TranslatorType string `json:"translator_type"`
			} `json:"user"`
			Geo interface{} `json:"geo"`
			Coordinates interface{} `json:"coordinates"`
			Place struct {
				ID string `json:"id"`
				URL string `json:"url"`
				PlaceType string `json:"place_type"`
				Name string `json:"name"`
				FullName string `json:"full_name"`
				CountryCode string `json:"country_code"`
				Country string `json:"country"`
				ContainedWithin []interface{} `json:"contained_within"`
				BoundingBox struct {
					Type string `json:"type"`
					Coordinates []struct {
						Num0 []float64 `json:"0"`
						Num1 []float64 `json:"1"`
						Num2 []float64 `json:"2"`
						Num3 []float64 `json:"3"`
					} `json:"coordinates"`
				} `json:"bounding_box"`
				Attributes struct {
				} `json:"attributes"`
			} `json:"place"`
			Contributors interface{} `json:"contributors"`
			IsQuoteStatus bool `json:"is_quote_status"`
			RetweetCount int `json:"retweet_count"`
			FavoriteCount int `json:"favorite_count"`
			Favorited bool `json:"favorited"`
			Retweeted bool `json:"retweeted"`
			Lang string `json:"lang"`
		} `json:"retweeted_status,omitempty"`
	} `json:"statuses"`
	SearchMetadata struct {
		CompletedIn float64 `json:"completed_in"`
		MaxID int64 `json:"max_id"`
		MaxIDStr string `json:"max_id_str"`
		NextResults string `json:"next_results"`
		Query string `json:"query"`
		RefreshURL string `json:"refresh_url"`
		Count int `json:"count"`
		SinceID int `json:"since_id"`
		SinceIDStr string `json:"since_id_str"`
	} `json:"search_metadata"`
}




type WatsonResponse struct {
	DocumentTone struct {
		ToneCategories []struct {
			Tones []struct {
				Score float64 `json:"score"`
				ToneID string `json:"tone_id"`
				ToneName string `json:"tone_name"`
			} `json:"tones"`
			CategoryID string `json:"category_id"`
			CategoryName string `json:"category_name"`
		} `json:"tone_categories"`
	} `json:"document_tone"`
	SentencesTone []struct {
		SentenceID int `json:"sentence_id"`
		Text string `json:"text"`
		InputFrom int `json:"input_from"`
		InputTo int `json:"input_to"`
		ToneCategories []struct {
			Tones []struct {
				Score float64 `json:"score"`
				ToneID string `json:"tone_id"`
				ToneName string `json:"tone_name"`
			} `json:"tones"`
			CategoryID string `json:"category_id"`
			CategoryName string `json:"category_name"`
		} `json:"tone_categories"`
	} `json:"sentences_tone"`
}



var finalGlobal string
	
func main() {

var buffer bytes.Buffer

	client := &http.Client{}
	//Step 1: Encode consumer key and secret
	encodedKeySecret := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s",
		url.QueryEscape(ConsumerKey),
		url.QueryEscape(ConsumerSecret))))

	//Step 2: Obtain a bearer token
	//The body of the request must be grant_type=client_credentials
	reqBody := bytes.NewBuffer([]byte(`grant_type=client_credentials`))
	//The request must be a HTTP POST request


	// trytest, err := ioutil.ReadAll(reqBody)
	// if err != nil {
	// 	log.Fatal(err, trytest)
	// }

	// fmt.Println(string(trytest))



	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", reqBody)
	 

	if err != nil {
		log.Fatal(err, client, req)
	}
	//The request must include an Authorization header formatted as
	//Basic <base64 encoded value from step 1>.
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedKeySecret))
	//The request must include a Content-Type header with
	//the value of application/x-www-form-urlencoded;charset=UTF-8.
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Content-Length", strconv.Itoa(reqBody.Len()))

	//Issue the request and get the bearer token from the JSON you get back
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err, resp)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err, respBody)
	}

	type BearerToken struct {
		AccessToken string `json:"access_token"`
	}
	var b BearerToken
	json.Unmarshal(respBody, &b)

	// fmt.Println(string(respBody))

	

	//choose your API endpoint that supports application only auth context
	//and create a request object with that
	twitterEndPoint := "https://api.twitter.com/1.1/search/tweets.json?q=%23goa&count=290"
	req, err = http.NewRequest("GET", twitterEndPoint, nil)
	if err != nil {
		//log.Fatal(err)
	}

	//Step 3: Authenticate API requests with the bearer token
	//include an Authorization header formatted as
	//Bearer <bearer token value from step 2>
	req.Header.Add("Authorization",
		fmt.Sprintf("Bearer %s", b.AccessToken))

	//Issue the request and get the JSON API response
	resp, err = client.Do(req)
	if err != nil {
		//log.Fatal(err, resp)
	}
	
	// fmt.Println("test: ",resp);

	defer resp.Body.Close()//will execute at the end of the block
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Auto AutoGenerated
	json.Unmarshal(respBody, &Auto)


	
	for i := 0; i <  len(Auto.Statuses); i++ {
	
	fmt.Println("Tweet: ",Auto.Statuses[i].Text)
	textTemp := Auto.Statuses[i].Text
	buffer.WriteString("Tweet : ")
	 buffer.WriteString(textTemp)
	 buffer.WriteString(" \n ")
	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// fmt.Println(string(respBody))



	type SendDataToWatson struct {
        Text string `json:"text"`
    }

    nonJSONData := SendDataToWatson{
        Text:textTemp,
    }


    sendDataInJsonFormat, err:=json.Marshal(nonJSONData);
    if err != nil {
      //converting to json error
    }

    urlDataInStringJSONFrom :=bytes.NewReader(sendDataInJsonFormat);
// request http api
    req2, err := http.NewRequest("POST", "https://gateway.watsonplatform.net/tone-analyzer/api/v3/tone?version=2016-05-19", urlDataInStringJSONFrom);

    if err != nil {
    	fmt.Println(err)
    }

    req2.SetBasicAuth("IBM Watson key ", "IBM Watson password");
    req2.Header.Set("Content-Type","application/json");

    client2 := http.Client{}
    res, err := client2.Do(req2)
    if err != nil {
    	fmt.Println(err)
    }

   // fmt.Println("StatusCode:", res.StatusCode)

// read body
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
    	fmt.Println(err)
    }

    var Watson WatsonResponse
	json.Unmarshal(body, &Watson)

	for j := 0; j <  len(Watson.DocumentTone.ToneCategories); j++ {
		buffer.WriteString(" \n ")
		 // fmt.Println(Watson.DocumentTone.ToneCategories[j].Tones[0].Score)
		 // fmt.Println(Watson.DocumentTone.ToneCategories[j].Tones[0].ToneName)
		fmt.Println(Watson.DocumentTone.ToneCategories[j].CategoryName)
		buffer.WriteString(Watson.DocumentTone.ToneCategories[j].CategoryName)
		buffer.WriteString(" \n ")
		for k := 0; k <  len(Watson.DocumentTone.ToneCategories[j].Tones); k++ {

			fmt.Println(Watson.DocumentTone.ToneCategories[j].Tones[k].ToneName , Watson.DocumentTone.ToneCategories[j].Tones[k].Score)

			buffer.WriteString(Watson.DocumentTone.ToneCategories[j].Tones[k].ToneName)
			buffer.WriteString(" : ")
			//toneSore = strconv.FormatFloat(Watson.DocumentTone.ToneCategories[j].Tones[0].Score, 'f', -1, 64)
			buffer.WriteString(strconv.FormatFloat(Watson.DocumentTone.ToneCategories[j].Tones[0].Score, 'f', -1, 64))
			buffer.WriteString(" \n ")
 			// fmt.Println(Watson.DocumentTone.ToneCategories[j].Tones[k].ToneName)
		}

	}
buffer.WriteString(" \n ")
	    //fmt.Println("Body: %s\n", string(body))
}

//fmt.Println(buffer.String())
finalString := buffer.String()
finalGlobal = finalString
// fmt.Println(finalGlobal)

  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "The reponse:\n %s", finalGlobal)
}