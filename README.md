# goflip
A straight-forward and basic library for creating memes with the imgflip.com API

### Create ImgFlip client

    client := goflip.NewClient("your ImgFlip username", "your ImgFlip password")

### Create new meme through the ImgFlip API

    request := goflip.APIRequest{
        TemplateID: "61582",
        Text0: "",
        Text1: "",
        Boxes: []goflip.Textbox{
            goflip.Textbox{
                Text: "Oh, you thought you could make your own memes?",
                X: 0,
                Y: 0,
                Width: 550,
                Height: 125,
                Color: goflip.BoxDefaultColor,
                OutlineColor: goflip.BoxDefaultOutlineColor
            },
            goflip.Textbox{
                Text: "Well isn't that special",
                X: 0,
                Y: 420,
                Width: 550,
                Height: 125,
                Color: goflip.BoxDefaultColor,
                OutlineColor: goflip.BoxDefaultOutlineColor
            }
        }
    }

    response, err := client.SendRequest(&request)

    // Check for errors
    if err != nil && response.Error == "" && response.Success {

        memeURL := response.Data.PageURL

        imageUrl := response.Data.URL
    }
