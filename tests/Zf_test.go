package tests
import(
	"testing"
	"log"
	"changliang/fanshe"
	"changliang/zf"
)

func TestPrintname(t *testing.T) {
	log.Println(fanshe.Fangfaming(false))
	log.Println(zf.Zfs.Fangfaming(false))
}
