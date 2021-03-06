package codegen_c

import (
	"../../util"
	"fmt"
)

type ClassTable struct {
	table map[string]*ClassBinding
}

func ClassTable_new() *ClassTable {
	o := new(ClassTable)
	o.table = make(map[string]*ClassBinding)
	return o
}

func (this *ClassTable) init(current string, extends string) {
	this.table[current] = ClassBinding_new(extends)
}

func (this *ClassTable) initDecs(current string, decs []Dec) {
	cb := this.table[current]
	for _, dec := range decs {
		if d, ok := dec.(*DecSingle); ok {
			cb.putField(current, d.Tp, d.Name)
		} else {
			panic("impossible")
		}
	}
}

func (this *ClassTable) initMethod(current string, ret Type, args []Dec, mid string) {
	cb := this.table[current]
	cb.putMethod(current, ret, args, mid)
}

func (this *ClassTable) get(c string) *ClassBinding {
	return this.table[c]
}

func (this *ClassTable) inherit(name string) {
	cb := this.table[name]
	if cb.visited == true {
		return
	}
	if cb.extends == "" {
		cb.visited = true
		return
	}
	this.inherit(cb.extends)

	super := this.table[cb.extends]
	var new_fields []*Tuple
	for _, t := range super.fields {
		new_fields = append(new_fields, t)
	}
	for _, t := range cb.fields {
		override := false
		for _, t2 := range new_fields {
			if Tuple_equals(t, t2) == true {
				override = true
			}
		}
		if override == false {
			new_fields = append(new_fields, t)
		}
	}
	cb.updateFields(new_fields)

	//method override
	var new_methods []*Ftuple
	for _, m := range super.methods {
		new_methods = append(new_methods, m)
	}
	for _, m := range cb.methods {
		for index, t := range new_methods {
			if Ftuple_equals(m, t) == true {
				new_methods[index] = m
			}
		}
	}
	cb.updateMethods(new_methods)
	cb.visited = true
}

func (this *ClassTable) String() string {
	util.Todo()
	return ""
}

func (this *ClassTable) dump() {
	for name, c := range this.table {
		fmt.Printf(name)
		c.dump()
		fmt.Println("")
	}
}
