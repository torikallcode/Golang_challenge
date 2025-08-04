package beginner

import "fmt"

func TestLogic() {

	follower := []Wrapper{
		{
			StringListData: []StringData{
				{Value: "Torikal"},
				{Value: "Akbar"},
				{Value: "Leal"},
			},
		},
	}

	following := InstagramData{
		Relationship_following: []Wrapper{
			{
				StringListData: []StringData{
					{Value: "Torikal"},
					{Value: "Akbar"},
					{Value: "Leal"},
					{Value: "bukangrace"},
				},
			},
		},
	}

	data_follower := make(map[string]bool)
	for _, i := range follower {
		for _, j := range i.StringListData {
			data_follower[j.Value] = true
		}
	}

	data_following := make(map[string]bool)
	for _, k := range following.Relationship_following {
		for _, l := range k.StringListData {
			data_following[l.Value] = true
		}
	}

	index := 1
	for i := range data_following {
		if !data_follower[i] {
			fmt.Println(index, ".", i)
			index++
		}
	}

}
