import {
  "testing"
  "fmt"}

func TestEmptyInput(t *testing. T){
  str:=" "
  _,err:=StringSum(str)
  if err==nil{
    t.Errorf("result is wrong, the result should be 'input is empty'")
  }
}

func TestLetterInFirstOperand(t *testing. T){
  str:="2s+5"
  _,err:=StringSum(str)
  if err==nil{
    t.Errorf("result is wrong, the result should be 'strconv.Atoi: parsing '2s': invalid syntax'")
  }
}

func TestLetterInSecondOperand(t *testing. T){
  str:="2+5s"
  _,err:=StringSum(str)
  if err==nil{
    t.Errorf("result is wrong, the result should be 'strconv.Atoi: parsing '5s': invalid syntax'")
  }
}

func TestFirstOperandNegative(t *testing. T){
  str:="-2+5"
  res,_:=StringSum(str)
  if res!="3"{
    t.Errorf("Result = %d; want 3", res)
  }
}

func TestSecondOperandNegative(t *testing. T){
  str:="2-5"
  res,_:=StringSum(str)
  if res!="-3"{
    t.Errorf("Result = %d; want -3", res)
  }
}

func TestBothOperandNegative(t *testing. T){
  str:="-2-5"
  res,_:=StringSum(str)
  if res!="-7"{
    t.Errorf("Result = %d; want -7", res)
  }
}

func TestBothOperandPositive(t *testing. T){
  str:="2+5"
  res,_:=StringSum(str)
  if res!="7"{
    t.Errorf("Result = %d; want 7", res)
  }
}

func TestWithWhispace(t *testing. T){
  str:=" 2 + 5 "
  res,_:=StringSum(str)
  if res!="7"{
    t.Errorf("Result = %d; want 7", res)
  }
}

func TestThreeOperands(t *testing. T){
  str:="2+5-3 "
  _,err:=StringSum(str)
  if err==nil{
    t.Errorf("result is wrong, the result should be 'expecting two operands, but received more or less'")
  }
}

func TestOneOperands(t *testing. T){
  str:="-3 "
  _,err:=StringSum(str)
  if err==nil{
    t.Errorf("result is wrong, the result should be 'expecting two operands, but received more or less'")
  }
}

func TestStringSum(t *testing. T){
  str:="2+3"
  res,_:=StringSum(str)
  rst,_:=fmt.Printf("%T", res)
  if rst!="string6"{
    t.Errorf("result is wrong, the result should be 'string6'")
  }
}
