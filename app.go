package regist

import (
	"context"
	"fmt"
	"github.com/UltimaQuazar/go-mastodon"
	"log"
)

func reg_app() {
    app, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
        Server: "https://toot.cafe", ClientName: "client-name",
        Scopes: "read write follow", Website: "https://github.com/UltimaQuazar/go-mastodon",
    })
    if err != nil{
        log.Fatal(err)
    }
    fmt.Printf("client-id: %s\n", app.ClientID)
    fmt.Printf("client-secret: %s\n", app.ClientSecret)
}
