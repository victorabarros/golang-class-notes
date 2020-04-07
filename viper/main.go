package main

import (
    "os"
    "fmt"
    "github.com/spf13/viper"
)

func main() {
    // Before `AutomaticEnv` Viper did not check envs
    fmt.Println("\nViper (\"LOGNAME\")\t", viper.GetString("LOGNAME"))
    fmt.Println("OS    (\"LOGNAME\")\t", os.Getenv("LOGNAME"))

    viper.AutomaticEnv()
    // After `AutomaticEnv`
    fmt.Println("\nViper (\"LOGNAME\")\t", viper.GetString("LOGNAME"))
    fmt.Println("OS    (\"LOGNAME\")\t", os.Getenv("LOGNAME"))


    // Read .env file
    viper.SetConfigType("env") // REQUIRED if the config file does not have the extension in the name
    viper.SetConfigName(".env") // name of config file (without extension)
    viper.AddConfigPath(".")               // optionally look for config in the working directory
    err := viper.ReadInConfig() // Find and read the config file
    if err != nil { // Handle errors reading the config file
        fmt.Println(err)
    }
    fmt.Println("\nViper     (\"FOO\")\t", viper.GetString("FOO"))
}
