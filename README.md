# docstract
golang package to extract attachments out of .msg files

```
import "github.com/jacthomp/DocStract"

func main(){
  file, _ := ioutil.ReadFile("x.msg")
  
  files, count, err := DocStract.Extract(file)
}
