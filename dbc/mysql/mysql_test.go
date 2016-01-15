package mysql

import (
	"fmt"
	"github.com/eaciit/dbox"
	"github.com/eaciit/toolkit"
	//"reflect"
	"testing"
	"time"
	"strconv"
)

type User struct {
	Id      string
	Name    string
	Tanggal time.Time
	Umur    int
}

func prepareConnection() (dbox.IConnection, error) {
	ci := &dbox.ConnectionInfo{"localhost:3306", "test", "root", "", nil}
	c, e := dbox.NewConnection("mysql", ci)
	if e != nil {
		return nil, e
	}

	e = c.Connect()
	if e != nil {
		return nil, e
	}
	return c, nil
}

func TestConnect(t *testing.T) {
	c, e := prepareConnection()
	if e != nil {
		t.Errorf("Unable to connect: %s \n", e.Error())
		fmt.Println(e)
	} else {
		// fmt.Println(c)
	}
	defer c.Close()
}

// func TestFilter(t *testing.T) {
// 	fb := dbox.NewFilterBuilder(new(FilterBuilder))
// 	fb.AddFilter(dbox.And(
// 		dbox.Eq("_id", "33"),dbox.Eq("_id", "34"),dbox.Eq("_id", "35")))
// 	// fb.AddFilter(dbox.Eq("_id", "33"), dbox.Eq("_id", "35"))
// 	b, e := fb.Build()
// 	if e != nil {
// 		t.Errorf("Error %s", e.Error())
// 	} else {
// 		fmt.Printf("Result:\n%v\n", toolkit.JsonString(b))
// 	}

// }

func TestSelect(t *testing.T) {
	c, e := prepareConnection()

	if e != nil {
		t.Errorf("Unable to connect %s \n", e.Error())
	}
	defer c.Close()

	//csr, e := c.NewQuery().Select().From("tes").Where(dbox.Eq("id", "3")).Cursor(nil)
	csr, e := c.NewQuery().Select("id", "name").From("tes").Cursor(nil)

	if e != nil {
		t.Errorf("Cursor pre error: %s \n", e.Error())
		return
	}
	if csr == nil {
		t.Errorf("Cursor not initialized")
		return
	}
	defer csr.Close()

	results := make([]map[string]interface{}, 0)

	err := csr.Fetch(&results, 0, false)
	if err != nil {
		t.Errorf("Unable to fetch all: %s \n", err.Error())
	} else {
		fmt.Println("=========================")
		fmt.Println("Select with NO filter")
		fmt.Println("=========================")

		for _, val := range results {
			fmt.Printf("Fetch N OK. Result: %v \n",
				toolkit.JsonString(val))
		}
	}

	e = csr.ResetFetch()
	if e != nil {
		t.Errorf("Unable to reset fetch: %s \n", e.Error())
	}
}

// func TestProcedure(t *testing.T) {
// 	c, _ := prepareConnection()
// 	defer c.Close()

// 	csr, e := c.NewQuery().Command("procedure", toolkit.M{}.Set("name", "spSomething").Set("parms", toolkit.M{}.Set("@name", "EACIIT"))).Cursor(nil)
// 	if csr == nil {
// 		t.Errorf("Cursor not initialized")
// 		return
// 	}
// 	defer csr.Close()

// 	ds, e := csr.Fetch(nil, 0, false)
// 	if e != nil {
// 		t.Errorf("Unable to fetch: %s \n", e.Error())
// 	} else {
// 		fmt.Printf("Fetch OK. Result: %v \n",
// 			toolkit.JsonString(ds.Data))
// 	}
// }

func TestSelectFilter(t *testing.T) {
	c, e := prepareConnection()
	if e != nil {
		t.Errorf("Unable to connect %s \n", e.Error())
		return
	}
	defer c.Close()

	csr, e := c.NewQuery().
		Select("id", "name", "umur").
		From("tes").
		//Where(dbox.Eq("name", "Bourne")).
		//Where(dbox.Neq("name", "Bourne")).
		//Where(dbox.Gt("umur", 25)).
		//Where(dbox.Gte("umur", 25)).
		//Where(dbox.Lt("umur", 25)).
		//Where(dbox.Lte("umur", 25)).
		//Where(dbox.In("name", "vidal", "bourne")).
		//Where(dbox.In("umur", "25", "30")).
		Where(dbox.Nin("umur", "25", "30")).
		//Where(dbox.In("tanggal", "2016-01-12 14:35:54", "2016-01-12 14:36:15")).
		//Where(dbox.And(dbox.Gt("umur", 25), dbox.Eq("name", "Roy"))).
		Cursor(nil)
	if e != nil {
		t.Errorf("Cursor pre error: %s \n", e.Error())
		return
	}
	if csr == nil {
		t.Errorf("Cursor not initialized")
		return
	}
	defer csr.Close()

	results := make([]map[string]interface{}, 0)
	err := csr.Fetch(&results, 0, false)
	if err != nil {
		t.Errorf("Unable to fetch: %s \n", err.Error())
	} else {
		fmt.Println("======================")
		fmt.Println("Select with FILTER")
		fmt.Println("======================")
		for _, val := range results {
			fmt.Printf("Fetch N OK. Result: %v \n",
				toolkit.JsonString(val))
		}
	}
}

// func TestSelectAggregate(t *testing.T) {
// 	c, e := prepareConnection()
// 	if e != nil {
// 		t.Errorf("Unable to connect %s \n", e.Error())
// 	}
// 	defer c.Close()

// 	fb := c.Fb()
// 	csr, e := c.NewQuery().
// 		//Select("_id", "email").
// 		//Where(c.Fb().Eq("email", "arief@eaciit.com")).
// 		Aggr(dbox.AggSum, 1, "Count").
// 		Aggr(dbox.AggSum, 1, "Avg").
// 		From("appusers").
// 		Group("").
// 		Cursor(nil)
// 	if e != nil {
// 		t.Errorf("Cursor pre error: %s \n", e.Error())
// 		return
// 	}
// 	if csr == nil {
// 		t.Errorf("Cursor not initialized")
// 		return
// 	}
// 	defer csr.Close()

// 	//rets := []toolkit.M{}

// 	ds, e := csr.Fetch(nil, 0, false)
// 	if e != nil {
// 		t.Errorf("Unable to fetch: %s \n", e.Error())
// 	} else {
// 		fmt.Printf("Fetch OK. Result: %v \n",
// 			toolkit.JsonString(ds.Data[0]))

// 	}
// }

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).UTC())
	return []byte(stamp), nil
}

func TestCRUD(t *testing.T) {
	//t.Skip()
	c, e := prepareConnection()
	if e != nil {
		t.Errorf("Unable to connect %s \n", e.Error())
		return
	}
	defer c.Close()

	//===============================INSERT==============================

	q := c.NewQuery().SetConfig("multiexec", true).From("tes").Save()

	// dataInsert := User{}
	// dataInsert.Id = fmt.Sprintf("6")
	// dataInsert.Name = fmt.Sprintf("Barkley")
	// //dataInsert.Tanggal = JSONTime(time.Now())
	// dataInsert.Tanggal = time.Now()
	// dataInsert.Umur = 21

	// e = q.Exec(toolkit.M{
	// 	"data": dataInsert,
	// })
	// if e != nil {
	// 	t.Errorf("Unable to save: %s \n", e.Error())
	// }

	//===============================INSERT MANY==============================

	// nama := [] string{"Barkley", "Vidal", "Arnautovic", "Agger", "Wijnaldum", "Ighalo", "Mahrez"}
	// dataInsert := User{}

	// for i, val := range nama{
		
	// 	dataInsert.Id = strconv.Itoa (i+1)
	// 	dataInsert.Name = fmt.Sprintf(val)
	// 	dataInsert.Tanggal = time.Now()
	// 	dataInsert.Umur = i+20
	// 	e = q.Exec(toolkit.M{
	// 	"data": dataInsert,
	// 	})
	// 	if e != nil {
	// 		t.Errorf("Unable to save: %s \n", e.Error())
	// 	}
	// }

	//===============================UPDATE==============================

	// data := User{}
	// data.Id = "7"
	// data.Name = "Oscar"
	// data.Tanggal = time.Now()
	// data.Umur = 24
	// e = c.NewQuery().From("tes").Where(dbox.Eq("id", "7")).Update().Exec(toolkit.M{"data": data})
	// if e != nil {
	// 	t.Errorf("Unable to update: %s \n", e.Error())
	// }

	//===============================DELETE==============================
	// e = c.NewQuery().From("tes").Where(dbox.And(dbox.Eq("id", "6"),dbox.Eq("name", "Barkley"))).Delete().Exec(nil)
	// if e != nil {
	// 	t.Errorf("Unablet to delete table %s\n", e.Error())
	// 	return
	// }


	//===============================CLEAR ALL TABLE DATA==============================

	// e = c.NewQuery().From("tes").Delete().Exec(nil)
	// if e != nil {
	// 	t.Errorf("Unablet to clear table %s\n", e.Error())
	// 	return
	// }

}