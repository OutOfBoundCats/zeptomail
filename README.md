![Logo](data/zeptomail-go-logo.png)
# Golang SDK For Zoho Zepto Mail

This golang sdk provides way to send email using zoho [zepto mailing service](https://www.zoho.com/zeptomail/)


## Features

- Written in Pure Go
- Used Generics
- Returns all zepto api errors back


## Usage/Examples

### Send Template Email
```go
import (
    "zeptomail"
)


type WelcomeEmailData struct {
    ProductName string `json:"product name"`
    Product     string `json:"product"`
    Link        string `json:"link"`
    SupportID   string `json:"support id"`
    Brand       string `json:"brand"`
    Username    string `json:"username"`
}

func sendSimpleTemaplteEmail() {
        zeptoClient := zeptomail.New("MAIL_TOKEN", "MAIL_AGENT_ALIAS")
        mergeData := WelcomeEmailData{
        ProductName: "name",
        Product:     "data",
        Link:        "link",
        SupportID:   "12345",
        Brand:       "brand",
        Username:    "name",
    }
    var FromEmail, ToEmail, ReplyTo zmodels.EmailAddress
    FromEmail = zmodels.EmailAddress{
        Address: "",
        Name:    "",
    }
    ToEmail = zmodels.EmailAddress{
        Address: "",
        Name:    "",
    }
    ReplyTo = zmodels.EmailAddress{
        Address: "",
        Name:    "",
    }
    var data = zmodels.ZeptoTemplateEmail[WelcomeEmailData]{
        TemplateKey: "Key",
        From:        FromEmail,
        To:          []zmodels.SendEmailTo{zmodels.SendEmailTo{EmailAddress: ToEmail}},
        MergeInfo:   mergeData,
        ReplyTo:     []zmodels.EmailAddress{ReplyTo},
    }
    
    successResp, failureResp, err := zeptomail.SendTemplateEmail[WelcomeEmailData](&zeptoClient, data)
    
    if err != nil {
        //handle error
    }
    if failureResp != nil {
        //handle zepto error this can be due to any reason as mentioend in belwo link
        //https://www.zoho.com/zeptomail/help/api/error-codes.html
    }
    //success data as returned by api
    fmt.Println(successResp.Message)

}

```

### Send Simple Email
To send simple email with either text or html body hover on zeptomail.EmailData to get the list of all fields or refer to [link](https://www.zoho.com/zeptomail/help/api/email-templates.html)
```go
 
var data = zeptomail.EmailData{
	From:    FromEmail,
	To:      []zeptomail.SendEmailTo{zeptomail.SendEmailTo{EmailAddress: ToEmail}},
	Subject: "Subject",
    Textbody: "Hi There"
	ReplyTo: []zeptomail.EmailAddress{ReplyTo},
}
successResp, failureResp, err := zeptoMail.SendSimpleEmail(&zeptoClient, data)			

```

### Send Template batch Emails
To send simple email with either text or html body hover on zeptomail.EmailData to get the list of all fields or refer to [link](https://www.zoho.com/zeptomail/help/api/batch-email-templates.html)
```go
to1:=zmodels.BatchTemplateTo[WelcomeEmailData]{
ToEmailAddress: ToEmail
MergeInfo: WelcomeData //this can be any type
} 
var data = zmodels.TemplateEmailBatch{
    TemplateKey: "Key",
    From:        FromEmail,
    To:          []zmodels.BatchTemplateTo{to1, ......add additional to emails with different merge data},
    ReplyTo:     []zmodels.EmailAddress{ReplyTo},
}
successResp, failureResp, err := zeptoMail.SendTemplateBatchEmail(&zeptoClient, data)			

```

## Authors

- [@OutOfBoundCats](https://github.com/OutOfBoundCats)

## Support

For support raise an issue in this repository.

## License

[MIT](https://choosealicense.com/licenses/mit/)
