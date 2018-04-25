package setup
import (
        "fmt"
        "github.com/urfave/cli"
        )

func Creat() *cli.App {
    app := cli.NewApp()
    app.Name = "Watering Hole"
    app.Usage = "Connect to mastadon like a cool retro nerd~!"
    app.Action = func(c *cli.Context)error{
        fmt.Println("hello world")
        return nil
    }
    return app
}
