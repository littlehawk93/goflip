# goflip
A straight-forward and basic library for creating memes with the imgflip.com API

### Create ImgFlip client

    client := goflip.NewClient("your ImgFlip username", "your ImgFlip password")

### Create new meme through the ImgFlip API

    request := goflip.APIRequest{
		TemplateID: "61582",
		Text0:      "",
		Text1:      "",
		Boxes: []goflip.Textbox{
			goflip.Textbox{
				Text:         "Oh, you thought you could make your own memes?",
				X:            10,
				Y:            10,
				Width:        530,
				Height:       125,
				Color:        goflip.BoxDefaultColor,
				OutlineColor: goflip.BoxDefaultOutlineColor,
			},
			goflip.Textbox{
				Text:         "Well isn't that special",
				X:            10,
				Y:            410,
				Width:        530,
				Height:       125,
				Color:        goflip.BoxDefaultColor,
				OutlineColor: goflip.BoxDefaultOutlineColor,
			},
		},
	}

	response, err := client.SendRequest(&request)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if !response.Success || response.Error != "" {
		fmt.Fprintln(os.Stderr, response.Error)
		return
	}

	fmt.Println(response.Data.PageURL)
	fmt.Println(response.Data.URL)
