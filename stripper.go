package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "regexp"
  "path"
  "path/filepath"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func replace(source string) (replaced string) {

  // Define regex
  reg, err := regexp.Compile("(<pproTicks[a-zA-Z]+>[0-9]+</pproTicks[a-zA-Z]+>)")
  check(err)

  //found := reg.FindAllString(source, -1)
  //fmt.Println(found)

  replaced = reg.ReplaceAllString(source, "")
  return replaced
}

func main() {

  // Get command line arguments
  args := os.Args[1:]
  if len(args) < 1 {
    fmt.Println("You need to add the path to a file to find/replace in")
    os.Exit(2)
  }

  // Read file
  inputPath := args[0]
  dat, readErr := ioutil.ReadFile(inputPath)
  check(readErr)

  // Find replace
  output := replace(string(dat))
  fmt.Println(output)

  // Create output directory
  outputPath := "." + string(filepath.Separator) + "stripped"
  os.Mkdir(outputPath, 0777)

  // Write output
  byteOutput := []byte(output)
  _, filename := path.Split(inputPath)
  writeErr := ioutil.WriteFile(outputPath + string(filepath.Separator) + filename, byteOutput, 0644)
  check(writeErr)
}
