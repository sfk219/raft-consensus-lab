// Package raftsim is an educational, synchronous Raft-inspired model, not a production implementation.
package raftsim
import ("errors"; "sort")
var ErrNoQuorum=errors.New("majority unavailable")
var ErrNoLeader=errors.New("no leader")
type Entry struct{ Term int; Index int; Key,Value string }
type Node struct{ ID string; Online bool; Log []Entry; Store map[string]string }
type Cluster struct{ Nodes map[string]*Node; Leader string; Term int }
func New(ids ...string)*Cluster{c:=&Cluster{Nodes:map[string]*Node{}};for _,id:=range ids{c.Nodes[id]=&Node{ID:id,Online:true,Store:map[string]string{}}};return c}
func(c *Cluster) quorum()bool{n:=0;for _,v:=range c.Nodes{if v.Online{n++}};return n>=len(c.Nodes)/2+1}
func(c *Cluster) Elect(id string)error{candidate,ok:=c.Nodes[id];if !ok||!candidate.Online{return ErrNoLeader};if !c.quorum(){return ErrNoQuorum};votes:=0;for _,n:=range c.Nodes{if n.Online&&len(n.Log)<=len(candidate.Log){votes++}};if votes<len(c.Nodes)/2+1{return ErrNoQuorum};c.Term++;c.Leader=id;for _,n:=range c.Nodes{if n.Online{c.sync(n,candidate)}};return nil}
func(c *Cluster) sync(dst,src *Node){dst.Log=append([]Entry(nil),src.Log...);dst.Store=map[string]string{};for k,v:=range src.Store{dst.Store[k]=v}}
func(c *Cluster) Put(k,v string)error{n,ok:=c.Nodes[c.Leader];if !ok||!n.Online{return ErrNoLeader};if !c.quorum(){return ErrNoQuorum};n.Log=append(n.Log,Entry{c.Term,len(n.Log)+1,k,v});n.Store[k]=v;for _,x:=range c.Nodes{if x.Online{c.sync(x,n)}};return nil}
func(c *Cluster) Stop(id string){if n,ok:=c.Nodes[id];ok{n.Online=false};if id==c.Leader{c.Leader=""}}
func(c *Cluster) Start(id string){if n,ok:=c.Nodes[id];ok{n.Online=true;if leader,exists:=c.Nodes[c.Leader];exists{c.sync(n,leader)}}}
func(c *Cluster) IDs()[]string{r:=[]string{};for k:=range c.Nodes{r=append(r,k)};sort.Strings(r);return r}
