import "tasting"

func StringEmptyTest(t *testing. T){
  str:=" "
  _,err:=StringSum(str)
  if err==nil{
    t.Errorf("result is wrong, the result should be 'input is empty'")
  }
}
