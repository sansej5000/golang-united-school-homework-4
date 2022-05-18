import {
  "testing"
  "fmt"}

func EmptyInputTest(t *testing. T){
  str:=" "
  _,err:=StringSum(str)
  if err==nil{
    t.Errorf("result is wrong, the result should be 'input is empty'")
  }
}

func LetterInFirstOperandTest(t *testing. T){
  str:="2s+5"
  _,err:=StringSum(str)
  if err==nil{
    t.Errorf("result is wrong, the result should be 'strconv.Atoi: parsing '2s': invalid syntax'")
  }
}

func LetterInSecondOperandTest(t *testing. T){
  str:="2+5s"
  _,err:=StringSum(str)
  if err==nil{
    t.Errorf("result is wrong, the result should be 'strconv.Atoi: parsing '5s': invalid syntax'")
  }
}

func FirstOperandNegativeTest(t *testing. T){
  str:="-2+5"
  res,_:=StringSum(str)
  if res!="3"{
    t.Errorf("result is wrong, the result should be '3'")
  }
}

func SecondOperandNegativeTest(t *testing. T){
  str:="2-5"
  res,_:=StringSum(str)
  if res!="-3"{
    t.Errorf("result is wrong, the result should be '-3'")
  }
}

func BothOperandNegativeTest(t *testing. T){
  str:="-2-5"
  res,_:=StringSum(str)
  if res!="-7"{
    t.Errorf("result is wrong, the result should be '-7'")
  }
}

func BothOperandPositiveTest(t *testing. T){
  str:="2+5"
  res,_:=StringSum(str)
  if res!="7"{
    t.Errorf("result is wrong, the result should be '7'")
  }
}

func WithWhispaceTest(t *testing. T){
  str:=" 2 + 5 "
  res,_:=StringSum(str)
  if res!="7"{
    t.Errorf("result is wrong, the result should be '7'")
  }
}

func ThreeOperandsTest(t *testing. T){
  str:="2+5-3 "
  _,err:=StringSum(str)
  if err==nil{
    t.Errorf("result is wrong, the result should be 'expecting two operands, but received more or less'")
  }
}

func OneOperandsTest(t *testing. T){
  str:="-3 "
  _,err:=StringSum(str)
  if err==nil{
    t.Errorf("result is wrong, the result should be 'expecting two operands, but received more or less'")
  }
}

func StringSumTest(t *testing. T){
  str:="2+3"
  res,_:=StringSum(str)
  rst,_:=fmt.Printf("%T", res)
  if rst!="string6"{
    t.Errorf("result is wrong, the result should be 'string6'")
  }
}