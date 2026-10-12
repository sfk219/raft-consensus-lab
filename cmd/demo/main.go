package main
import ("fmt";"github.com/sfk219/raft-consensus-lab/raftsim")
func main(){c:=raftsim.New("node-1","node-2","node-3");_ =c.Elect("node-1");_ =c.Put("release","v1");c.Stop("node-1");_ =c.Elect("node-2");_ =c.Put("release","v2");c.Start("node-1");for _,id:=range c.IDs(){fmt.Println(id,c.Nodes[id].Store)}}
