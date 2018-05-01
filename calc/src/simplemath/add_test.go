package simplemath

import "testing"

func TestAdd(t *testing.T){
    r := Add(1,2)
    if r != 3{
        t.Errorf("Add(1,2) failed.Got %d,expected 3.",r)
    }
}

func BenchmarkAdd1(b* testing.B){
    for i:=0; i < b.N; i++{
        Add(1,2)
    }
}

func BenchmarkAdd2(b *testing.B){
    b.StopTimer()
    b.StartTimer()

    for i:=0; i < b.N; i++{
        Add(1,2)
    }
}