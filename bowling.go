package main

import (
  "fmt"
  "math/rand"
  "strconv"
  "time"
)

func main() {
  play_frame(-1)
  play_frame(5)
  play_frame(10)
}

func take_turn(pinsToHit int) (int, string) {
  ballOne := bowl(10,pinsToHit)
  pinsDown := ballOne 
  ballTwo := bowl(10-ballOne,pinsToHit)
  var result string
  if ballOne == 10 {
    result = "strike"
    fmt.Print("X ")
  } else {
    pinsDown += ballTwo
    fmt.Print(num_to_s(ballOne))
    if ballTwo == 10-ballOne {
      fmt.Print("/ ")
      result = "spare"
    } else if ballTwo == 0 {
      fmt.Print("- ")
    } else {
      fmt.Print(num_to_s(ballTwo)+" ")
    }
  }
  return pinsDown, result
}

func score_turn(pinsDown int, previousResult string, result string) int {
  switch previousResult {
    case "strike":
      if result != "strike" {
        return 10 + pinsDown
      } else {
        return 20 + pinsDown
      }
    case "spare":
      return 10 + 5
    default:
      return pinsDown
  }
}

func bowl( remaining_pins, pinsToHit int) int {
  rand.Seed(time.Now().UnixNano())
  if pinsToHit == -1 {
    return rand.Intn(remaining_pins+1)
  } else {
    return pinsToHit
  }
}

func num_to_s(number int) string {
  return strconv.Itoa(number)
}

func play_frame(pinsToHit int) {
  i := 1
  var score int
  var previousResult string
  for i <= 10 {
    var pins int
    var result string 
    pins, result = take_turn(pinsToHit)
    score += score_turn(pins,previousResult,result)
    i = i + 1
    if i == 10 {
      if result == "strike" {
        pinsDown := bowl(10,pinsToHit)
                fmt.Print(pinsDown)
        if pinsDown == 10 {
          previousResult = "strike"
        }

        score += score_turn(pinsDown,previousResult,result)
      }
    }
    previousResult=result
  }
  fmt.Println(num_to_s(score))
}