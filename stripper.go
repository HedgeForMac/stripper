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

func replace(source string, nodes []string) (replaced string, err error) {

  replaced = source

  for i := 0; i < len(nodes); i++ {

    // Get node
    node := nodes[i]
    fmt.Printf("Replacing %s\n", node)

    // Define regex
    reg, err := regexp.Compile("(<" + node + ">[^<>]+</" + node + ">)")
    if err != nil {
      return "", err
    }
    replaced = reg.ReplaceAllString(replaced, "")

  }

  // Remove extra new lines
  reg, err := regexp.Compile("(\t+\n\t+\n)")
  if err != nil {
    return "", err
  }
  replaced = reg.ReplaceAllString(replaced, "")

  return replaced, nil
}

func main() {

  // Get command line arguments
  args := os.Args[1:]
  if len(args) < 2 {
    fmt.Println("Usage:\n./stripped path node [node]")
    os.Exit(2)
  }

  // Read file
  inputPath := args[0]
  dat, readErr := ioutil.ReadFile(inputPath)
  check(readErr)

  // Get nodes
  nodes := args[1:]

  // Find replace
  output, err := replace(string(dat), nodes)
  check(err)

  // Create output directory
  outputPath := "." + string(filepath.Separator) + "stripped"
  os.Mkdir(outputPath, 0777)

  // Write output
  byteOutput := []byte(output)
  _, filename := path.Split(inputPath)
  writeErr := ioutil.WriteFile(outputPath + string(filepath.Separator) + filename, byteOutput, 0644)
  check(writeErr)

  fmt.Println("All done")
}
