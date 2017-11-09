package util

import (
	"fmt"
	"bytes"
	"io"
	"strconv"
	"os"
)


type Paper struct {
	Id			int
    Title 		string
	AuthorList  string
	Abstract 	string
}

type StrSet map[string]bool
//Convert string slice to StrSet
func NewStrSet(strs ...string) StrSet {
	ss := StrSet(make(map[string]bool))
	for _, str := range strs {
		ss.Put(str)
	}
	return ss
}

func (this StrSet) Put(str string)                { this[str] = true }
func (this StrSet) Del(str string)                { delete(this, str) }
func (this StrSet) Contains(str string) (ok bool) { _, ok = this[str]; return ok }
func (this StrSet) Merge(that StrSet) {
	for str := range that {
		this[str] = true
	}
}
func (this StrSet) Array() []string {
	ret := make([]string, 0, len(this))
	for str := range this {
		ret = append(ret, str)
	}
	return ret
}


//write to dot file
func WriteDot(pkgimports map[string]StrSet, writer io.Writer) (err error) {
	nodes := NewStrSet()
	edges := [][2]string{}
	for pkg, imps := range pkgimports {
		nodes.Put(pkg)
		for imp := range imps {
			nodes.Put(imp)
			edges = append(edges, [2]string{pkg, imp})
		}
	}
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString("digraph G {\n")
	for _, edge := range edges {
		buf.WriteString(fmt.Sprintf(`"%s"->"%s";`, edge[0], edge[1]))
		buf.WriteByte('\n')
	}
	for pkg, _ := range nodes {
		buf.WriteString(fmt.Sprintf(`"%s";`, pkg))
		buf.WriteByte('\n')
	}
	buf.WriteString("}\n")
	_, err = writer.Write(buf.Bytes())
	return
}

func Write2Dot(paperlist []Paper,writer io.Writer) (err error) {
	nodes := NewStrSet()
	buf := bytes.NewBuffer([]byte{})
	edges := [][2]string{}
	buf.WriteString("digraph G {\n")
		nodes.Put("USENIX")
		for _,p :=range paperlist {
			//edges = append(edges, {{"USE","s"}})
			//edges = append(edges, [2]string{})
			//b:=[][2]string{{p.title, p.abstract}}
			ID:=strconv.Itoa(p.Id)
			edges = append(edges, [][2]string{{"USENIX", ID}}...)
			edges = append(edges, [][2]string{{ID, p.Title}}...)
			edges = append(edges, [][2]string{{ID, p.AuthorList}}...)
			edges = append(edges, [][2]string{{ID, p.Abstract}}...)
			nodes.Put(ID)
		}
		for _, edge := range edges {
			buf.WriteString(fmt.Sprintf(`"%s"->"%s";`, edge[0], edge[1]))
			buf.WriteByte('\n')
		}
		for pid, _ := range nodes {
			buf.WriteString(fmt.Sprintf(`"%s";`, pid))
			buf.WriteByte('\n')
		}	
	buf.WriteString("}\n")
	_, err = writer.Write(buf.Bytes())
	return
}


func Write2Dotf(paperlist []Paper, dotf string) (err error) {
	nodes := NewStrSet()
	buf := bytes.NewBuffer([]byte{})
	edges := [][2]string{}
	buf.WriteString("digraph G {\n")
		nodes.Put("USENIX")
		for _,p :=range paperlist {
			//edges = append(edges, {{"USE","s"}})
			//edges = append(edges, [2]string{})
			//b:=[][2]string{{p.title, p.abstract}}
			ID:=strconv.Itoa(p.Id)
			edges = append(edges, [][2]string{{"USENIX", ID}}...)
			edges = append(edges, [][2]string{{ID, p.Title}}...)
			edges = append(edges, [][2]string{{ID, p.AuthorList}}...)
			edges = append(edges, [][2]string{{ID, p.Abstract}}...)
			nodes.Put(ID)
		}
		for _, edge := range edges {
			buf.WriteString(fmt.Sprintf(`"%s"->"%s";`, edge[0], edge[1]))
			buf.WriteByte('\n')
		}
		for pid, _ := range nodes {
			buf.WriteString(fmt.Sprintf(`"%s";`, pid))
			buf.WriteByte('\n')
		}	
	buf.WriteString("}\n")

	//f, err := os.Open(dotf)
	f, err := os.Create(dotf)
	defer f.Close()
	_,err = f.Write(buf.Bytes())
	//ioutil.WriteFile(dotf,buf,0644)
	//_, err = writer.Write(buf.Bytes())
	return
}