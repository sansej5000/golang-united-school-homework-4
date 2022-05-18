import "testing"

func StringEmptyTest(t *testing. T){
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
