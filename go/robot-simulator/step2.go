package robot

import "fmt"

type Action int

func StartRobot(com <- chan Command, act chan Action) {
    for c := range com {
        switch string(c) {
        case "R":
            fmt.Println("Right")
        case "L":
            fmt.Println("Left")
        case "A":
            fmt.Println("Advance")
        default:
            fmt.Println(c)
        }
    }
}

func Room(extent Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
    //for a := range act {
    //    switch a {
    //
    //    }
    //}
}
