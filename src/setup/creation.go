package setup
import (
        "fmt"
        "github.com/urfave/cli"
        )

func init(){
    //flags
    cli.HelpFlag=cli.BoolFlag{Name: "help"}
    cli.VersionFlag=cli.BoolFlag{Name: "print-version, V"}

    //printers
    //not yet thought up something for this
    //laying ground work
    /*cli.HelpPrinter=func(w io.Writer, templ string, data interface{}){
        fmt.Fprintf(w, "Not yet implemented\n")
    }*/
    cli.VersionPrinter=func(c *cli.Context){
        fmt.Fprintf(c.App.Writer, "version = %s\n", c.App.Version)
    }
}

/*
 * Will create the application on install with required name and definition.
 * Hoping that this will work out the box because there are some issues when installing.
 * 042518
 */
func Creat() *cli.App {
    app := cli.NewApp()
    app.Name = "Watering_Hole"
    app.Usage = "Connect to mastadon like a cool retro nerd~!"
    
    app.Flags=[]cli.Flag{
        //inside the StringFlag struct we have:
        //Name, Value, Usage,
        //should help out with new usage parameters like username, instance etc.
        cli.StringFlag{
            Name: "lang",
            Value: "english",
            Usage: "Application Language",
        },
    }

    app.Action = func(c *cli.Context)error{
        fmt.Println("hello world")
        //should print out the arguments
        fmt.Printf("Arguments: %q", c.Args().Get(0))
        return nil
    }
    return app
}
