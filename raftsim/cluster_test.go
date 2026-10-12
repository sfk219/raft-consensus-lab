package raftsim
import "testing"
func TestElectionReplication(t *testing.T){c:=New("a","b","c");if err:=c.Elect("a");err!=nil{t.Fatal(err)};if err:=c.Put("region","eu");err!=nil{t.Fatal(err)};for _,n:=range c.Nodes{if n.Store["region"]!="eu"{t.Fatal("replication failed")}}}
func TestQuorum(t *testing.T){c:=New("a","b","c");_ =c.Elect("a");c.Stop("b");c.Stop("c");if err:=c.Put("x","y");err!=ErrNoQuorum{t.Fatalf("expected quorum error, got %v",err)}}
func TestFailover(t *testing.T){c:=New("a","b","c");_ =c.Elect("a");_ =c.Put("v","1");c.Stop("a");if err:=c.Elect("b");err!=nil{t.Fatal(err)};_ =c.Put("v","2");c.Start("a");if c.Nodes["a"].Store["v"]!="2"{t.Fatal("failed recovery")}}
