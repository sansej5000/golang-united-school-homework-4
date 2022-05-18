import "tasting"

func StringEmptyTest(t *testing. T){
  str:=" "
  _,err:=StringSum(str)
  if err!=nil{
    t.Errorf("Str is empty: %s", err)
  }
}
