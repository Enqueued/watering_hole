package setup
import (
        "fmt"
        "github.com/urfave/cli"
        )

/*
 * Will create the application on install with required name and definition.
 * Hoping that this will work out the box because there are some issues when installing.
 * 042518
 */
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
