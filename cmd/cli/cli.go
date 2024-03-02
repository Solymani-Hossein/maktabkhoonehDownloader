package cli

import (
	"os"
	"web-scrapper-go/internal/entity"

	"github.com/spf13/cobra"
)

var arguments = entity.Args{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "maktabkhoonehdl-cli",
	Short: "season link and define quality",
	Long:  `long description`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() entity.Args {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	return arguments
}

func init() {

	//url := "https://maktabkhooneh.org/course/برنامه-نویسی-پیشرفته-جاوا-mk242"
	url := "https://maktabkhooneh.org/course/آموزش-تحلیل-بدافزار-مقدماتی-mk933/"
	rootCmd.Flags().StringVarP(&arguments.Url, "url", "u", url, "The url to Scrapping.")

	rootCmd.Flags().IntVarP((*int)(&arguments.Quality), "quality", "q", int(entity.High), "The quality of video.")

	//if err := rootCmd.MarkFlagRequired("url"); err != nil {
	//	fmt.Println(err)
	//}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
