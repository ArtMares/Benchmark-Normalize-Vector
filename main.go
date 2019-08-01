package main

import (
    "flag"
    "fmt"
    "math"
    "math/rand"
    "os"
    "time"
)

type Vec3 [3]float64

func (v *Vec3) Mul(i interface{}) *Vec3 {
    switch b := i.(type) {
    case float64:
        v[0] *= b
        v[1] *= b
        v[2] *= b
    }
    return v
}

func (v *Vec3) Dev(i interface{}) *Vec3 {
    switch b := i.(type) {
    case float64:
        v[0] /= b
        v[1] /= b
        v[2] /= b
    }
    return v
}

func (v *Vec3) X() float64 {
    return v[0]
}

func (v *Vec3) Y() float64 {
    return v[1]
}

func (v *Vec3) Z() float64 {
    return v[2]
}

func (v *Vec3) NormD() {
    l := math.Sqrt(v[0] * v[0] + v[1] * v[1] + v[2] * v[2])
    v[0] /= l
    v[1] /= l
    v[2] /= l
}

func (v *Vec3) NormDF() {
    l := math.Sqrt(v[0] * v[0] + v[1] * v[1] + v[2] * v[2])
    v.Dev(l)
}

func (v *Vec3) NormDX() {
    l := math.Sqrt(v.X() * v.X() + v.Y() * v.Y() + v.Z() * v.Z())
    v[0] /= l
    v[1] /= l
    v[2] /= l
}

func (v *Vec3) NormDFX() {
    l := math.Sqrt(v.X() * v.X() + v.Y() * v.Y() + v.Z() * v.Z())
    v.Dev(l)
}

func (v *Vec3) NormDM() {
    l := v.S()
    v[0] /= l
    v[1] /= l
    v[2] /= l
}

func (v *Vec3) NormDFS() {
    l := v.S()
    v.Dev(l)
}

func (v *Vec3) NormDSX() {
    l := v.SX()
    v[0] /= l
    v[1] /= l
    v[2] /= l
}

func (v *Vec3) NormDFSX() {
    l := v.SX()
    v[0] /= l
    v[1] /= l
    v[2] /= l
}

func (v *Vec3) NormM() {
    l := 1 / math.Sqrt(v[0] * v[0] + v[1] * v[1] + v[2] * v[2])
    v[0] *= l
    v[1] *= l
    v[2] *= l
}

func (v *Vec3) NormMF() {
    l := 1 / math.Sqrt(v[0] * v[0] + v[1] * v[1] + v[2] * v[2])
    v.Mul(l)
}

func (v *Vec3) NormMX() {
    l := 1 / math.Sqrt(v.X() * v.X() + v.Y() * v.Y() + v.Z() * v.Z())
    v[0] *= l
    v[1] *= l
    v[2] *= l
}

func (v *Vec3) NormMFX() {
    l := 1 / math.Sqrt(v.X() * v.X() + v.Y() * v.Y() + v.Z() * v.Z())
    v.Mul(l)
}

func (v *Vec3) NormMS() {
    l := 1 / v.S()
    v[0] *= l
    v[1] *= l
    v[2] *= l
}

func (v *Vec3) NormMFS() {
    l := 1 / v.S()
    v.Mul(l)
}

func (v *Vec3) NormMSX() {
    l := 1 / v.SX()
    v[0] *= l
    v[1] *= l
    v[2] *= l
}

func (v *Vec3) NormMFSX() {
    l := 1 / v.SX()
    v.Mul(l)
}

func (v *Vec3) S() float64 {
    return math.Sqrt(v[0] * v[0] + v[1] * v[1] + v[2] * v[2])
}

func (v *Vec3) SX() float64 {
    return math.Sqrt(v.X() * v.X() + v.Y() * v.Y() + v.Z() * v.Z())
}

type Vec3s struct {
    X, Y, Z   float64
}

func (v *Vec3s) Mul(i interface{}) *Vec3s {
    switch b := i.(type) {
    case float64:
        v.X *= b
        v.Y *= b
        v.Z *= b
    }
    return v
}

func (v *Vec3s) Dev(i interface{}) *Vec3s {
    switch b := i.(type) {
    case float64:
        v.X /= b
        v.Y /= b
        v.Z /= b
    }
    return v
}

func (v *Vec3s) NormD() {
    l := math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
    v.X /= l
    v.Y /= l
    v.Z /= l
}
func (v *Vec3s) NormDF() {
    l := math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
    v.Dev(l)
}


func (v *Vec3s) NormDS() {
    l := v.M()
    v.X /= l
    v.Y /= l
    v.Z /= l
}

func (v *Vec3s) NormDFS() {
    l := v.M()
    v.Dev(l)
}

func (v *Vec3s) NormM() {
    l := 1 / math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
    v.X *= l
    v.Y *= l
    v.Z *= l
}

func (v *Vec3s) NormMF() {
    l := math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
    v.Mul(l)
}

func (v *Vec3s) NormMS() {
    l := 1 / v.M()
    v.X *= l
    v.Y *= l
    v.Z *= l
}

func (v *Vec3s) NormMFS() {
    l := v.M()
    v.Mul(l)
}

func (v *Vec3s) M() float64 {
    return math.Sqrt(v.X * v.X + v.Y * v.Y + v.Z * v.Z)
}

type Benchmark struct {
    Name    string
    Steps   []time.Duration
    Avg     time.Duration
    Min     time.Duration
    Max     time.Duration
}

func (b *Benchmark) String() string {
    res := fmt.Sprintf("Test: %s\n\tAvg: %s\n\tMin: %s\n\tMax: %s\n\tSteps:\n", b.Name, b.Avg, b.Min, b.Max)
    for i, t := range b.Steps {
        res += fmt.Sprintf("\t\t%d\t%s\n", i + 1, t)
    }
    return res
}

func calculate(b *Benchmark) {
    var total float64
    var min, max int64
    for _, t := range b.Steps {
        total += float64(t)
        if min == 0 {
            min = int64(t)
        } else if min > int64(t) {
            min = int64(t)
        }
        if max == 0 {
            max = int64(t)
        } else if max < int64(t) {
            max = int64(t)
        }
        b.Avg = time.Duration(total / float64(len(b.Steps)))
        b.Min = time.Duration(min)
        b.Max = time.Duration(max)
    }
}

func main() {
    steps := 10
    iterations := 10000000
    rand.NewSource(time.Now().Unix())

    benchName := flag.String("name", "", "Test name")
    flag.Parse()

    benchmarks := []*Benchmark{
        {
            Name: "Vec3.NormD",
        },
        {
            Name: "Vec3.NormDF",
        },
        {
            Name: "Vec3.NormDX",
        },
        {
            Name: "Vec3.NormDFX",
        },
        {
            Name: "Vec3.NormDS",
        },
        {
            Name: "Vec3.NormDFS",
        },
        {
            Name: "Vec3.NormDSX",
        },
        {
            Name: "Vec3.NormDFSX",
        },
        {
            Name: "Vec3.NormM",
        },
        {
            Name: "Vec3.NormMF",
        },
        {
            Name: "Vec3.NormMX",
        },
        {
            Name: "Vec3.NormMFX",
        },
        {
            Name: "Vec3.NormMS",
        },
        {
            Name: "Vec3.NormMFS",
        },
        {
            Name: "Vec3.NormMSX",
        },
        {
            Name: "Vec3.NormMFSX",
        },
        {
            Name: "Vec3s.NormD",
        },
        {
            Name: "Vec3s.NormDF",
        },
        {
            Name: "Vec3s.NormDS",
        },
        {
            Name: "Vec3s.NormDFS",
        },
        {
            Name: "Vec3s.NormM",
        },
        {
            Name: "Vec3s.NormMF",
        },
        {
            Name: "Vec3s.NormMS",
        },
        {
            Name: "Vec3s.NormMFS",
        },
    }

    if *benchName == "" {
        fmt.Print("Need argument \"-name\"\nAvailable values:\n")
        for _, b := range benchmarks {
            fmt.Printf("\t%s\n", b.Name)
        }
        os.Exit(0)
    }

    count := iterations / 4



    for _, b := range benchmarks {
        for i := 0; i < steps; i++ {
            t := time.Now()
            for j := 0; j < iterations; j++ {
                switch b.Name {
                case "Vec3.Norm()":
                    v := &Vec3{6, 5.3, 2.7}
                    v.Norm()
                case "Vec3.NormF()":
                    v := &Vec3{6, 5.3, 2.7}
                    v.NormF()
                case "Vec3.NormXYZ()":
                    v := &Vec3{6, 5.3, 2.7}
                    v.NormXYZ()
                case "Vec3.NormSqrMag()":
                    v := &Vec3{6, 5.3, 2.7}
                    v.NormSqrMag()
                case "Vec3.NormSqrMagXYZ()":
                    v := &Vec3{6, 5.3, 2.7}
                    v.NormSqrMagXYZ()
                case "Vec3.NormMul()" :
                    v := &Vec3{6, 5.3, 2.7}
                    v.NormMul()
                case "Vec3.NormMulXYZ()":
                    v := &Vec3{6, 5.3, 2.7}
                    v.NormMulXYZ()
                case "Vec3.NormMulSqrMag()":
                    v := &Vec3{6, 5.3, 2.7}
                    v.NormMulSqrMag()
                case "Vec3.NormMulSqrMagXYZ()":
                    v := &Vec3{6, 5.3, 2.7}
                    v.NormMulSqrMagXYZ()
                case "Vec3s.Norm()":
                    v := &Vec3s{6, 5.3, 2.7}
                    v.Norm()
                case "Vec3s.NormSqrMag()":
                    v := &Vec3s{6, 5.3, 2.7}
                    v.NormSqrMag()
                case "Vec3s.NormF()":
                    v := &Vec3s{6, 5.3, 2.7}
                    v.NormF()
                case "Vec3s.NormFSqrMag()":
                    v := &Vec3s{6, 5.3, 2.7}
                    v.NormFSqrMag()
                case "Vec3s.NormMul()":
                    v := &Vec3s{6, 5.3, 2.7}
                    v.NormMul()
                case "vec3s.NormMulSqrMag()":
                    v := &Vec3s{6, 5.3, 2.7}
                    v.NormMulSqrMag()
                case "Vec3s.NormFMul()":
                    v := &Vec3s{6, 5.3, 2.7}
                    v.NormFMul()
                case "vec3s.NormFMulSqrMag()":
                    v := &Vec3s{6, 5.3, 2.7}
                    v.NormFMulSqrMag()
                }
            }
            b.Steps = append(b.Steps, time.Since(t))
        }
        calculate(b)
        fmt.Println(b)
    }
}
