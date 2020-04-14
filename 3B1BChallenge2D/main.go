// https://youtu.be/OkmNXy7er84?list=WL&t=116
// Three points in a circle P1, P2 and P3 make a triangle.
// Whats the probability the triangle contains the center of the cicle?
package main

import (
    "fmt"
    "math"
    "math/rand"
)

var counter = 0.0
var inside  = 0.0
var outside = 0.0

func main() {
    for{
        // p = (x, y)
        var p1_x, p1_y float64

        // random  -1 < p_x < 1 TODO: How include -1 and 1?
        p1_x = rand.Float64() * math.Pow(-1.0, float64(rand.Intn(2)))
        // p_x^2 + p_y^2 = R^2
        p1_y = math.Pow(-1.0, float64(rand.Intn(2))) * math.Sqrt(1- math.Pow(p1_x, 2.0))

        // fmt.Println(math.Pow(p1_x, 2.0) + math.Pow(p1_y, 2.0))
        // ~ 1

        // Same...
        var p2_x, p2_y float64
        p2_x = rand.Float64() * math.Pow(-1.0, float64(rand.Intn(2)))
        p2_y = math.Pow(-1.0, float64(rand.Intn(2))) * math.Sqrt(1- math.Pow(p2_x, 2.0))
        var p3_x, p3_y float64
        p3_x = rand.Float64() * math.Pow(-1.0, float64(rand.Intn(2)))
        p3_y = math.Pow(-1.0, float64(rand.Intn(2))) * math.Sqrt(1- math.Pow(p3_x, 2.0))

        var dx, dy float64
        // l its the side for the triangle
        dx = p1_x-p2_x
        dy = p1_y-p2_y
        // l^2 = (dx)^2 + (dy)^2
        var l1 = math.Sqrt(math.Pow(dx, 2.0) + math.Pow(dy, 2.0))

        dx = p1_x-p3_x
        dy = p1_y-p3_y
        var l2 = math.Sqrt(math.Pow(dx, 2.0) + math.Pow(dy, 2.0))

        dx = p3_x-p2_x
        dy = p3_y-p2_y
        var l3 = math.Sqrt(math.Pow(dx, 2.0) + math.Pow(dy, 2.0))

        // cos law c^2 = a^2 + b^2 -2*a*b*cos(ĉ)
        var cos_theta1 = (math.Pow(l1, 2.0) - (math.Pow(l2, 2.0)+math.Pow(l3, 2.0)))/(-2.0*l2*l3)
        var cos_theta2 = (math.Pow(l2, 2.0) - (math.Pow(l1, 2.0)+math.Pow(l3, 2.0)))/(-2.0*l1*l3)
        var cos_theta3 = (math.Pow(l3, 2.0) - (math.Pow(l2, 2.0)+math.Pow(l1, 2.0)))/(-2.0*l2*l1)

        // If cosine of an interior angle is lower than zero
        // It's mean that angule is biggest then 90º and we have a obtusangle
        // And it's doesn't contains the circle
        if cos_theta1 < 0 || cos_theta2 < 0 || cos_theta3 < 0 {
            // obtusangle
            outside ++
        } else {
            // acutangle
            inside ++
        }
        counter ++

        fmt.Println(inside/counter, "\t", counter)
        // ~ 0.25
        // fmt.Println(inside/counter + outside/counter)
        // 1
    }
}
// TODO: Look for a draw package and plot every triangle inside of the circle.
// TODO: I think a better approach is with analytc geometryc instead of plane geometry.
// Must exist a theoreme to return in one form the angles of a triangle on the plane.
// Example: it's a knowledge that tan(alpha) = Dy/Dx => alpha = tan^-1(Dy/Dx)
// TODO: Made the challenge in 3D
